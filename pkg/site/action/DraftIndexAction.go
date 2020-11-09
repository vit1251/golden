package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/site/widgets"
	"net/http"
)

type DraftIndexAction struct {
	Action
}

func NewDraftIndexAction() *DraftIndexAction {
	return new(DraftIndexAction)
}

func (self DraftIndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	draftMapper := mapperManager.GetDraftMapper()

	/* Restore draft index */
	drafts, err1 := draftMapper.GetDraftMessages(mapper.DraftStateActive)
	if err1 != nil {
		status := fmt.Sprintf("Fail in GetDraftMessages on draftMapper: err = %+v", err1)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

	/* Render base wiew */
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

	/* Add custom param field */
	twitTable := widgets.NewTableWidget().
			SetClass("draft-index-table")

	twitTable.AddRow(widgets.NewTableRowWidget().
		SetClass("draft-index-header").
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Name"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Action"))))

	for _, d := range drafts {

		twitTable.AddRow(widgets.NewTableRowWidget().
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(d.GetSubject()))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewLinkWidget().
				SetContent("Edit").
				SetClass("btn").
				SetLink(fmt.Sprintf("/draft/%s/edit", d.GetUUID())))))

	}

	containerVBox.Add(twitTable)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
