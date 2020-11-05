package action

import (
	"bufio"
	"fmt"
	"github.com/gorilla/mux"
	cmn "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/file"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

type FileAreaUploadCompleteAction struct {
	Action
}

func NewFileAreaUploadCompleteAction() *FileAreaUploadCompleteAction {
	return new(FileAreaUploadCompleteAction)
}

func (self FileAreaUploadCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fileManager := self.restoreFileManager()
	configManager := self.restoreConfigManager()

	outb, _ := configManager.Get("main", "Outbound")
	passwd, _ := configManager.Get("main", "Password")
	from, _ := configManager.Get("main", "Address")
	to, _ := configManager.Get("main", "Link")

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	/* Get file area */
	area, err1 := fileManager.GetAreaByName(echoTag)
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

	//
	log.Printf("FileAreaUploadCompleteAction: filename = %+v", header.Filename)

	// Copy the file data to my buffer

	tmpFile := path.Join(outb, header.Filename)
	writeStream, err3 := os.Create(tmpFile)
	if err3 != nil {
		panic(err3)
	}
	cacheWriter := bufio.NewWriter(writeStream)
	defer func () {
		cacheWriter.Flush()
		writeStream.Close()
	}()

	size, err4 := io.Copy(cacheWriter, stream)
	if err4 != nil {
		panic(err4)
	}

	/* Create TIC description */
	ticFile := file.NewTicFile()

	/* Area RU.GOLDEN */
	ticFile.AddLine(fmt.Sprintf("Area %s", area.GetName()))

	/* From 2:5023/24.3752 */
	ticFile.AddLine(fmt.Sprintf("From %s", from))

	/* To 2:5023/24 */
	ticFile.AddLine(fmt.Sprintf("To %s", to))

	/* File GoldenPoint-20200423.zip */
	ticFile.AddLine(fmt.Sprintf("File %s", header.Filename))

	/* Desc Golden Point - Night - 2020-04-23 */
	ticFile.AddLine(fmt.Sprintf("Desc %s", desc))

	/* Size 0 */
	ticFile.AddLine(fmt.Sprintf("Size %d", size))

	/* Pw ****** */
	ticFile.AddLine(fmt.Sprintf("Pw %s", passwd))

	/* Save TIC on disnk */
	newName := cmn.MakeTickName()
	newPath := path.Join(outb, newName)
	ticFile.Save(newPath)

	/* Redirect */
	newLocation := fmt.Sprintf("/file/%s", area.GetName())
	http.Redirect(w, r, newLocation, 303)

}
