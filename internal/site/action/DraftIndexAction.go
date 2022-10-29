package action

import (
	"fmt"
	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/pkg/mapper"
	"net/http"
)

type DraftIndexAction struct {
	Action
}

func NewDraftIndexAction() *DraftIndexAction {
	return new(DraftIndexAction)
}

func (self DraftIndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.registry)
	draftMapper := mapperManager.GetDraftMapper()

	/* Restore draft index */
	drafts, err1 := draftMapper.GetDraftMessages(mapper.DraftStateActive)
	if err1 != nil {
		status := fmt.Sprintf("Fail in GetDraftMessages on draftMapper: err = %+v", err1)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

	/* Render base wiew */
	bw := widgets2.NewBaseWidget()

	vBox := widgets2.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets2.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets2.NewVBoxWidget()

	container.AddWidget(containerVBox)

	vBox.Add(container)

	/* Add custom param field */
	twitTable := widgets2.NewTableWidget().
		SetClass("draft-index-table")

	twitTable.AddRow(widgets2.NewTableRowWidget().
		SetClass("draft-index-header").
		AddCell(widgets2.NewTableCellWidget().SetWidget(widgets2.NewTextWidgetWithText("Name"))).
		AddCell(widgets2.NewTableCellWidget().SetWidget(widgets2.NewTextWidgetWithText("Action"))))

	for _, d := range drafts {

		msgSubject := d.GetSubject()
		if msgSubject == "" {
			msgSubject = "-EMPTY-"
		}

		twitTable.AddRow(widgets2.NewTableRowWidget().
			AddCell(widgets2.NewTableCellWidget().SetWidget(widgets2.NewTextWidgetWithText(msgSubject))).
			AddCell(widgets2.NewTableCellWidget().SetWidget(widgets2.NewLinkWidget().
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
