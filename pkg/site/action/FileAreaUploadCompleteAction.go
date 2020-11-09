package action

import (
	"bufio"
	"fmt"
	"github.com/gorilla/mux"
	cmn "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/tracker"
	"hash/crc32"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

type FileAreaUploadCompleteAction struct {
	Action
}

func NewFileAreaUploadCompleteAction() *FileAreaUploadCompleteAction {
	return new(FileAreaUploadCompleteAction)
}

func (self FileAreaUploadCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	fileMapper := mapperManager.GetFileMapper()
	configMapper := mapperManager.GetConfigMapper()

	passwd, _ := configMapper.Get("main", "Password")
	myAddr, _ := configMapper.Get("main", "Address")

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	/* Get file area */
	area, err1 := fileMapper.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	/* ... */
	var maxMemory int64 = 128 * 1024 * 1024
	r.ParseMultipartForm(maxMemory)

	// in your case file would be fileupload
	stream, header, err2 := r.FormFile("file")
	if err2 != nil {
		panic(err2)
	}
	defer stream.Close()

	/* Description */
	desc := r.PostForm.Get("desc")
	to := r.PostForm.Get("desc")
	ldesc := r.PostForm.Get("ldesc")

	//
	log.Printf("FileAreaUploadCompleteAction: filename = %+v", header.Filename)

	// Copy the file data to my buffer
	outboundDirectory := cmn.GetOutboundDirectory()
	tmpFile := path.Join(outboundDirectory, header.Filename)
	writeStream, err3 := os.Create(tmpFile)
	if err3 != nil {
		panic(err3)
	}
	cacheWriter := bufio.NewWriter(writeStream)
	defer func () {
		cacheWriter.Flush()
		writeStream.Close()
	}()

	/* Copy */
	crcWriter := crc32.NewIEEE()

	outStreams := io.MultiWriter(cacheWriter, crcWriter)

	size, err4 := io.Copy(outStreams, stream)
	if err4 != nil {
		panic(err4)
	}

	crc := crcWriter.Sum32()
	crcValue := fmt.Sprintf("%08X", crc)

	/* Create TIC description */
	ticBuilder := tracker.NewTicBuilder()

	ticBuilder.SetArea(area.GetName())
	ticBuilder.SetOrigin(myAddr)
	ticBuilder.SetFrom(myAddr)
	ticBuilder.SetFile(header.Filename)
	ticBuilder.SetDesc(desc)
	ticBuilder.SetLDesc(ldesc)
	ticBuilder.SetSize(size)
	ticBuilder.SetPw(passwd)
	ticBuilder.SetCrc(crcValue)
	ticBuilder.SetTo(to)
	ticBuilder.SetDate(time.Now())
	ticBuilder.AddSeenby(myAddr)

	newTicPath := fmt.Sprintf("%s %d %s %s/%s %s (%s)",
		myAddr,
		time.Now().Unix(), time.Now().Format("Mon Nov 09 09:03:02 2020 UTC"),
		"GoldenPoint", cmn.GetPlatform(), cmn.GetVersion(), cmn.GetReleaseDate())
	ticBuilder.AddPath(newTicPath)


	/* Save TIC on disk */
	newName := cmn.MakeTickName()
	newPath := path.Join(outboundDirectory, newName)

	newContent := ticBuilder.Build()
	writer, err5 := os.Create(newPath)
	if err5 != nil {
		panic(err5)
	}
	cacheWriter2 := bufio.NewWriter(writer)
	defer func() {
		cacheWriter2.Flush()
		writer.Close()
	}()
	cacheWriter2.WriteString(newContent)

	/* Redirect */
	newLocation := fmt.Sprintf("/file/%s", area.GetName())
	http.Redirect(w, r, newLocation, 303)

}
