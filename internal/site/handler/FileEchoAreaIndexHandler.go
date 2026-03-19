package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/vit1251/golden/internal/site/utils"
	widgets2 "github.com/vit1251/golden/internal/site/widgets"

	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
)

type FileEchoAreaIndexHandler struct {
	registry *registry.Container
}

func NewFileEchoAreaIndexHandler(registry *registry.Container) *FileEchoAreaIndexHandler {
	return &FileEchoAreaIndexHandler{
		registry: registry,
	}
}

func (self *FileEchoAreaIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.registry)
	fileAreaMapper := mapperManager.GetFileAreaMapper()
	fileMapper := mapperManager.GetFileMapper()

	/* Parse URL parameters */
	var areaIndex string = r.PathValue("echoname")
	log.Printf("echoTag = %v", areaIndex)

	/* Get message area */
	area, err1 := fileAreaMapper.GetAreaByName(areaIndex)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName on fileMapper")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	if area == nil {
		response := fmt.Sprintf("Fail on GetAreaByName on fileMapper with area name %s", areaIndex)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("area = %+v", area)

	files, err2 := fileMapper.GetFileHeaders(areaIndex)
	if err2 != nil {
		response := fmt.Sprintf("Fail on GetFileHeaders on fileMapper: %+v", err2)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("files = %+v", files)

	bw := widgets2.NewBaseWidget()

	vBox := widgets2.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets2.NewMainMenuWidget()
	vBox.Add(mmw)

	container := widgets2.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets2.NewVBoxWidget()
	container.AddWidget(containerVBox)
	vBox.Add(container)

	/* Context handlers */
	actionBar := self.renderHandlers(area)
	containerVBox.Add(actionBar)

	indexTable := widgets2.NewDivWidget().
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

func (self *FileEchoAreaIndexHandler) renderRow(area *mapper.FileArea, file *mapper.File) widgets2.IWidget {

	mapperManager := mapper.RestoreMapperManager(self.registry)
	fileMapper := mapperManager.GetFileMapper()

	/* Make message row container */
	rowWidget := widgets2.NewDivWidget().
		SetStyle("display: flex").
		SetStyle("direction: column").
		SetStyle("align-items: center")

	var classNames []string
	classNames = append(classNames, "file-area-index-item")

	/* Check exists */
	var areaName string = area.GetName()
	var indexName string = file.GetFile()
	path := fileMapper.GetFileAbsolutePath(areaName, indexName)
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
	nameWidget := widgets2.NewDivWidget().
		SetWidth("190px").
		SetHeight("38px").
		SetStyle("flex-shrink: 0").
		SetStyle("white-space: nowrap").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		//SetStyle("border: 1px solid green").
		SetContent(file.GetOrigName())
	rowWidget.AddWidget(nameWidget)

	/* Render NEW point */
	var newPointContent string = ""
	if file.IsNew() {
		newPointContent = "•"
	}
	newPointWidget := widgets2.NewDivWidget().
		SetWidth("20px").
		SetHeight("38px").
		SetStyle("flex-shrink: 0").
		SetStyle("white-space: nowrap").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		//SetStyle("border: 1px solid green").
		SetStyle("color: yellow").
		SetContent(newPointContent)
	rowWidget.AddWidget(newPointWidget)

	/* Render summary */
	summaryWidget := widgets2.NewDivWidget().
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

	counterWidget := widgets2.NewDivWidget().
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

	navigateItem := widgets2.NewLinkWidget().
		SetLink(navigateAddress).
		AddWidget(rowWidget)

	return navigateItem

}

func (self *FileEchoAreaIndexHandler) checkExists(path string) bool {

	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return false

}

func (self *FileEchoAreaIndexHandler) renderHandlers(area *mapper.FileArea) widgets2.IWidget {

	actionBar := widgets2.NewActionMenuWidget()

	/* Upload */
	actionBar.Add(widgets2.NewMenuAction().
		SetLink(fmt.Sprintf("/file/%s/upload", area.GetName())).
		SetIcon("icofont-edit").
		SetClass("mr-2").
		SetLabel("Upload"))

	/* Settings */
	actionBar.Add(widgets2.NewMenuAction().
		SetLink(fmt.Sprintf("/file/%s/update", area.GetName())).
		SetIcon("icofont-update").
		SetClass("mr-2").
		SetLabel("Settings"))

	return actionBar

}
