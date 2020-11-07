package action

import (
	"fmt"
	"github.com/gorilla/mux"
	commonfunc "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/site/widgets"
	"log"
	"net/http"
)

type FileAreaViewAction struct {
	Action
}

func NewFileAreaViewAction() *FileAreaViewAction {
	fa := new(FileAreaViewAction)
	return fa
}


func (self *FileAreaViewAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fileManager := self.restoreFileManager()

	/* Parse URL parameters */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	/* Get message area */
	area, err1 := fileManager.GetAreaByName(echoTag)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName on FileManager")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	if area == nil {
		response := fmt.Sprintf("Fail on GetAreaByName on FileManager with area name %s", echoTag)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("area = %+v", area)

	files, err2 := fileManager.GetFileHeaders(echoTag)
	if err2 != nil {
		response := fmt.Sprintf("Fail on GetFileHeaders on FileManager")
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
	container.SetWidget(containerVBox)
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
		SetClass("table")

	indexTable.AddRow(widgets.NewTableRowWidget().
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("GetName"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Summary"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("GetAge"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Action"))))

	for _, f := range files {

		newDate := commonfunc.MakeHumanTime(f.GetTime())

		log.Printf("file = %+v", f)
		indexTable.AddRow(widgets.NewTableRowWidget().
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(f.GetFile()))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(f.GetDesc()))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(newDate))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewLinkWidget().
				SetContent("Download"). // /file/BOOK-DOP/tic/KORSAV01.RAR/view
				SetLink(fmt.Sprintf("/file/%s/tic/%s/view", f.GetArea(), f.GetFile())))))
	}

	containerVBox.Add(indexTable)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
