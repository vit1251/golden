package action

import (
	"archive/zip"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/site/widgets"
)

type FileEchoAreaViewAction struct {
	Action
}

func NewFileEchoAreaViewAction() *FileEchoAreaViewAction {
	return new(FileEchoAreaViewAction)
}

func (self *FileEchoAreaViewAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	fileAreaMapper := mapperManager.GetFileAreaMapper()
	fileMapper := mapperManager.GetFileMapper()

	/* Parse URL parameters */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)
	indexName := vars["file"]
	log.Printf("indexName = %v", indexName)

	/* Get message area */
	area, err1 := fileAreaMapper.GetAreaByName(echoTag)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName on fileMapper")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("area = %+v", area)

	/**/
	file, err2 := fileMapper.GetFileByIndexName(echoTag, indexName)
	if err2 != nil {
		response := fmt.Sprintf("Fail on GetFileByFileName on fileMapper")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Update view counter */
	err3 := fileMapper.ViewFileByIndexName(area.GetName(), indexName)
	if err3 != nil {
		response := fmt.Sprintf("Fail on ViewFileByIndexName on fileMapper")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Render page */
	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")

	vBox.Add(container)

	containerVBox := widgets.NewVBoxWidget()

	container.AddWidget(containerVBox)

	/* Context actions */
	actionBar := self.renderActions(area, indexName)
	containerVBox.Add(actionBar)

	/* TODO - show meta here ... */
	if IsImage(file.GetFile()) {
		imageURL := fmt.Sprintf("/file/%s/tic/%s/download", area.GetName(), indexName)
		imageWidget := widgets.NewImageWidget()
		imageWidget.SetSource(imageURL)
		imageWidget.SetClass("preview")
		containerVBox.Add(imageWidget)
	}

	if IsZipArchive(file.GetFile()) {
		path := file.GetAbsolutePath()
		//
		divBox := widgets.NewDivWidget()
		divBox.SetContent(path)
		containerVBox.Add(divBox)
		//
		reader, err1 := zip.OpenReader(path)
		if err1 == nil {
			divBox1 := widgets.NewPreWidget()
			var comment string = reader.Comment
			comment = strings.Replace(comment, "\r\n", "\n", -1) // CRLF -> NL
			divBox1.SetContent(comment)
			containerVBox.Add(divBox1)
			//
			var out string
			for _, f := range reader.File {
				out += fmt.Sprintf("%s - %s (%d byte)<br />", f.Name, f.Comment, f.UncompressedSize64)
			}
			//
			divBox2 := widgets.NewDivWidget()
			divBox2.SetContent(out)
			containerVBox.Add(divBox2)
		}

	}

	/* Done */
	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self *FileEchoAreaViewAction) renderActions(area *mapper.FileArea, newFile string) widgets.IWidget {

	actionMenu := widgets.NewActionMenuWidget()

	actionMenu.Add(widgets.NewMenuAction().
		SetLink(fmt.Sprintf("/file/%s/tic/%s/download", area.GetName(), newFile)).
		SetIcon("icofont-edit").
		SetClass("mr-2").
		SetLabel("Download"))

	actionMenu.Add(widgets.NewMenuAction().
		SetLink(fmt.Sprintf("/file/%s/tic/%s/remove", area.GetName(), newFile)).
		SetIcon("icofont-edit").
		SetClass("mr-2").
		SetLabel("Remove"))

	return actionMenu
}

func IsZipArchive(filename string) bool {
	upperName := strings.ToUpper(filename)
	return strings.HasSuffix(upperName, ".ZIP")
}

func IsImage(filename string) bool {
	upperName := strings.ToUpper(filename)
	return strings.HasSuffix(upperName, ".GIF") || strings.HasSuffix(upperName, ".JPG") || strings.HasSuffix(upperName, ".PNG")
}
