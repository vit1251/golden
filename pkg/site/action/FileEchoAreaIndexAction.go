package action

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/site/utils"
	"github.com/vit1251/golden/pkg/site/widgets"
	"log"
	"net/http"
	"os"
	"strings"
)

type FileEchoAreaIndexAction struct {
	Action
}

func NewFileEchoAreaIndexAction() *FileEchoAreaIndexAction {
	fa := new(FileEchoAreaIndexAction)
	return fa
}

func (self *FileEchoAreaIndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	fileAreaMapper := mapperManager.GetFileAreaMapper()
	fileMapper := mapperManager.GetFileMapper()

	/* Parse URL parameters */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	/* Get message area */
	area, err1 := fileAreaMapper.GetAreaByName(echoTag)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName on fileMapper")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	if area == nil {
		response := fmt.Sprintf("Fail on GetAreaByName on fileMapper with area name %s", echoTag)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("area = %+v", area)

	files, err2 := fileMapper.GetFileHeaders(echoTag)
	if err2 != nil {
		response := fmt.Sprintf("Fail on GetFileHeaders on fileMapper")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("files = %+v", files)

	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets.NewVBoxWidget()
	container.AddWidget(containerVBox)
	vBox.Add(container)

	/* Context actions */
	amw := widgets.NewActionMenuWidget().
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/file/%s/upload", area.GetName())).
			SetIcon("icofont-edit").
			SetLabel("Upload")).
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/file/%s/update", area.GetName())).
			SetIcon("icofont-update").
			SetLabel("Settings"))

	containerVBox.Add(amw)

	indexTable := widgets.NewDivWidget().
		SetClass("file-area-index-table")

	for _, f := range files {

		itemRow := self.renderRow(area, &f)
		indexTable.AddWidget(itemRow)

	}

	containerVBox.Add(indexTable)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self *FileEchoAreaIndexAction) renderRow(area *mapper.FileArea, file *mapper.File) widgets.IWidget {

	mapperManager := self.restoreMapperManager()
	fileMapper := mapperManager.GetFileMapper()

	/* Make message row container */
	rowWidget := widgets.NewDivWidget().
		SetStyle("display: flex").
		SetStyle("direction: column").
		SetStyle("align-items: center")

	var classNames []string
	classNames = append(classNames, "file-area-index-item")

	/* Check exists */
	var areaName string = area.GetName()
	var fileName string = file.GetFile()
	path := fileMapper.GetFileAbsolutePath(areaName, fileName)
	log.Printf("Check %s", path)

	viewCount := file.GetViewCount()
	if viewCount == 0 {
		classNames = append(classNames, "file-area-index-item-new")
	}
	if !self.checkExists(path) {
		classNames = append(classNames, "file-area-index-item-missing")
	}
	rowWidget.SetClass(strings.Join(classNames, " "))

	/* Render area name */
	nameWidget := widgets.NewDivWidget().
		SetWidth("190px").
		SetHeight("38px").
		SetStyle("flex-shrink: 0").
		SetStyle("white-space: nowrap").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		//SetStyle("border: 1px solid green").
		SetContent(file.GetFile())
	rowWidget.AddWidget(nameWidget)

	/* Render summary */
	summaryWidget := widgets.NewDivWidget().
		SetStyle("min-width: 350px").
		SetHeight("38px").
		SetStyle("flex-grow: 1").
		SetStyle("white-space: nowrap").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		//SetStyle("border: 1px solid red").
		SetContent(file.GetDesc())
	rowWidget.AddWidget(summaryWidget)

	/* Render counter widget */
	dateText := utils.DateHelper_renderDate(file.GetTime())

	counterWidget := widgets.NewDivWidget().
		SetHeight("38px").
		SetWidth("180px").
		SetStyle("white-space: nowrap").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		SetStyle("flex-shrink: 0").
		//SetStyle("border: 1px solid blue").
		SetContent(dateText)
	rowWidget.AddWidget(counterWidget)

	/* Link container */
	navigateAddress := fmt.Sprintf("/file/%s/tic/%s/view", file.GetArea(), file.GetFile())

	navigateItem := widgets.NewLinkWidget().
		SetLink(navigateAddress).
		AddWidget(rowWidget)

	return navigateItem

}

func (self *FileEchoAreaIndexAction) checkExists(path string) bool {

	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return false

}
