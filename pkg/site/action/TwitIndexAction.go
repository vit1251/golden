package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/site/widgets"
	"net/http"
)

type TwitIndexAction struct {
	Action
}

func NewTwitIndexAction() *TwitIndexAction {
	return new(TwitIndexAction)
}

func (self TwitIndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	twitMapper := mapperManager.GetTwitMapper()

	/* Restore twits */
	twitNames, err1 := twitMapper.GetTwitNames()
	if err1 != nil {
		status := fmt.Sprintf("Fail in GetTwitNames on twitMapper: err = %+v", err1)
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
			SetClass("twit-index-table")

	twitTable.AddRow(widgets.NewTableRowWidget().
		SetClass("twit-index-header").
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Name"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Action"))))

	for _, t := range twitNames {

		twitTable.AddRow(widgets.NewTableRowWidget().
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(t.GetName()))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewLinkWidget().
				SetContent("Remove").
				SetClass("btn").
				SetLink(fmt.Sprintf("/twit/%s/remove", t.GetId())))))

	}

	containerVBox.Add(twitTable)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
