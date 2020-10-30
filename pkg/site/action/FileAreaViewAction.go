package action

import (
	"fmt"
	"github.com/gorilla/mux"
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

	/* Context actions */
	amw := widgets.NewActionMenuWidget().
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/file/%s/compose", area.Name())).
			SetIcon("icofont-edit").
			SetLabel("Compose")).
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/file/%s/update", area.Name())).
			SetIcon("icofont-update").
			SetLabel("Settings"))

	vBox.Add(amw)

	indexTable := widgets.NewTableWidget().
		SetClass("table")

	indexTable.AddRow(widgets.NewTableRowWidget().
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Name"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Summary"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Age"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Action"))))

	for _, f := range files {
		log.Printf("file = %+v", f)
		indexTable.AddRow(widgets.NewTableRowWidget().
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(f.File))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(f.Desc))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(f.Age()))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewLinkWidget().
				SetContent("Download"). // /file/BOOK-DOP/tic/KORSAV01.RAR/view
				SetLink(fmt.Sprintf("/file/%s/tic/%s/view", f.Area, f.File)))))
	}

	vBox.Add(indexTable)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
