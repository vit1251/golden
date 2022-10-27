package action

import (
	"archive/zip"
	"fmt"
	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/internal/um"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
)

type FileEchoAreaViewAction struct {
	Action
}

func NewFileEchoAreaViewAction() *FileEchoAreaViewAction {
	return new(FileEchoAreaViewAction)
}

func (self *FileEchoAreaViewAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	urlManager := um.RestoreUrlManager(self.GetRegistry())
	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
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
	var areaName string = area.GetName()
	err3 := fileMapper.ViewFileByIndexName(areaName, indexName)
	if err3 != nil {
		response := fmt.Sprintf("Fail on ViewFileByIndexName on fileMapper")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Render page */
	bw := widgets2.NewBaseWidget()

	vBox := widgets2.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets2.NewDivWidget()
	container.SetClass("container")

	vBox.Add(container)

	containerVBox := widgets2.NewVBoxWidget()

	container.AddWidget(containerVBox)

	/* Context actions */
	actionBar := self.renderActions(area, indexName)
	containerVBox.Add(actionBar)

	/* Show preview widget */
	var origPath string = file.GetAbsolutePath()
	var origName = file.GetOrigName()

	/* Common attributes */
	containerVBox.Add(self.renderAttribute("Original TIC filename", origName))
	containerVBox.Add(self.renderAttribute("Disk path", origPath))
	fi, err := os.Stat(origPath)
	if err == nil {
		var storageSize string = fmt.Sprintf("%d", fi.Size()/1024)
		containerVBox.Add(self.renderAttribute("Disk size, KB", storageSize))
	}

	/* Extension base attributes */
	if IsImage(origName) {

		/* Render preview */
		imageAddr := urlManager.CreateUrl("/file/{farea_name}/tic/{tic_index}/download").
			SetParam("farea_name", areaName).
			SetParam("tic_index", indexName).
			Build()

		imageWidget := widgets2.NewImageWidget()
		imageWidget.SetSource(imageAddr)
		imageWidget.SetClass("preview")
		containerVBox.Add(imageWidget)

		/* Render image attributes */
		// TODO - image size and etc...

	} else if IsZipArchive(origName) {

		/* Render archive attributes */
		// TODO - add more attributes ...

		/* Processing archive */
		reader, err1 := zip.OpenReader(origPath)
		if err1 == nil {

			/* Render archive comment */
			divBox1 := widgets2.NewPreWidget()
			var comment string = reader.Comment
			comment = strings.Replace(comment, "\r\n", "\n", -1) // CRLF -> NL
			divBox1.SetContent(comment)
			containerVBox.Add(divBox1)

			/* Render archive index */
			var out string
			for _, f := range reader.File {
				out += fmt.Sprintf("%s - %s (%d byte)<br />", f.Name, f.Comment, f.UncompressedSize64)
			}
			divBox2 := widgets2.NewDivWidget()
			divBox2.SetContent(out)
			containerVBox.Add(divBox2)

		} else {
			// TODO - process broken ZIP arcive ...
		}

	}

	/* Done */
	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self *FileEchoAreaViewAction) renderActions(area *mapper.FileArea, newFile string) widgets2.IWidget {

	actionMenu := widgets2.NewActionMenuWidget()

	actionMenu.Add(widgets2.NewMenuAction().
		SetLink(fmt.Sprintf("/file/%s/tic/%s/download", area.GetName(), newFile)).
		SetIcon("icofont-edit").
		SetClass("mr-2").
		SetLabel("Download"))

	actionMenu.Add(widgets2.NewMenuAction().
		SetLink(fmt.Sprintf("/file/%s/tic/%s/remove", area.GetName(), newFile)).
		SetIcon("icofont-edit").
		SetClass("mr-2").
		SetLabel("Remove"))

	return actionMenu
}

func (self *FileEchoAreaViewAction) renderAttribute(name string, value string) widgets2.IWidget {

	output := fmt.Sprintf("%s = %+v", name, value)

	divBox := widgets2.NewDivWidget()
	divBox.SetContent(output)

	return divBox
}

func IsZipArchive(filename string) bool {
	upperName := strings.ToUpper(filename)
	return strings.HasSuffix(upperName, ".ZIP")
}

func IsImage(filename string) bool {
	upperName := strings.ToUpper(filename)
	return strings.HasSuffix(upperName, ".GIF") || strings.HasSuffix(upperName, ".JPG") || strings.HasSuffix(upperName, ".PNG")
}
