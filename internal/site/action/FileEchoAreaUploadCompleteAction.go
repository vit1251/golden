package action

import (
	"bufio"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/internal/common"
	"github.com/vit1251/golden/pkg/charset"
	"github.com/vit1251/golden/pkg/config"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/tracker"
	"hash/crc32"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"time"
)

type FileEchoAreaUploadCompleteAction struct {
	Action
}

func NewFileEchoAreaUploadCompleteAction() *FileEchoAreaUploadCompleteAction {
	return new(FileEchoAreaUploadCompleteAction)
}

func (self FileEchoAreaUploadCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	configManager := config.RestoreConfigManager(self.GetRegistry())
	charsetManager := charset.RestoreCharsetManager(self.GetRegistry())
	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
	fileAreaMapper := mapperManager.GetFileAreaMapper()
	//fileMapper := mapperManager.GetFileMapper()

	newConfig := configManager.GetConfig()

	passwd := newConfig.Main.Password
	myAddr := newConfig.Main.Address

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	/* Get file area */
	area, err1 := fileAreaMapper.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)
	var areaCharset string = area.GetCharset()

	/* ... */
	var maxMemory int64 = 128 * 1024 * 1024

	err2 := r.ParseMultipartForm(maxMemory)
	if err2 != nil {
		panic(err2)
	}

	// in your case file would be fileupload
	stream, header, err3 := r.FormFile("file")
	if err3 != nil {
		panic(err3)
	}
	defer stream.Close()

	/* Description */
	desc := r.PostForm.Get("desc")
	to := r.PostForm.Get("to")
	ldesc := r.PostForm.Get("ldesc")

	//
	log.Printf("FileEchoAreaUploadCompleteAction: filename = %+v", header.Filename)

	// Copy the file data to my buffer
	outboundDirectory := commonfunc.GetOutboundDirectory()
	tmpFile := path.Join(outboundDirectory, header.Filename)
	writeStream, err6 := os.Create(tmpFile)
	if err6 != nil {
		panic(err6)
	}
	cacheWriter := bufio.NewWriter(writeStream)
	defer func() {
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
		"GoldenPoint", commonfunc.GetPlatform(), commonfunc.GetVersion(), commonfunc.GetReleaseDate())
	ticBuilder.AddPath(newTicPath)

	/* Save TIC on disk */
	newName := commonfunc.MakeTickName()
	newPath := path.Join(outboundDirectory, newName)

	content := ticBuilder.Build()

	/* Change charset */
	newContent, err4 := charsetManager.EncodeMessageBody(content, areaCharset)
	if err4 != nil {
		panic(err4)
	}

	writer, err5 := os.Create(newPath)
	if err5 != nil {
		panic(err5)
	}
	cacheWriter2 := bufio.NewWriter(writer)
	defer func() {
		cacheWriter2.Flush()
		writer.Close()
	}()
	cacheWriter2.Write(newContent)

	/* Redirect */
	newLocation := fmt.Sprintf("/file/%s", area.GetName())
	http.Redirect(w, r, newLocation, 303)

}
