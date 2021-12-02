package action

import (
	"fmt"
	"github.com/gorilla/mux"
	commonfunc "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/site/widgets"
	"log"
	"net/http"
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

	indexTable := widgets.NewTableWidget().
		SetClass("file-area-index-table")

	indexTable.AddRow(widgets.NewTableRowWidget().
		SetClass("file-area-index-header").
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Name"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Summary"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Date"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Action"))))

	for _, f := range files {

		newDate := commonfunc.MakeHumanTime(f.GetTime())

		actions := widgets.NewVBoxWidget()

		actions.Add(widgets.NewLinkWidget().
			SetContent("View").
			SetClass("btn").
			SetLink(fmt.Sprintf("/file/%s/tic/%s/view", f.GetArea(), f.GetFile())))

		actions.Add(widgets.NewLinkWidget().
			SetContent("Remove").
			SetClass("btn").
			SetLink(fmt.Sprintf("/file/%s/tic/%s/remove", f.GetArea(), f.GetFile())))

		log.Printf("file = %+v", f)
		indexTable.AddRow(widgets.NewTableRowWidget().
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(f.GetFile()))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(f.GetDesc()))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(newDate))).
			AddCell(widgets.NewTableCellWidget().SetWidget(actions)))
	}

	containerVBox.Add(indexTable)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
