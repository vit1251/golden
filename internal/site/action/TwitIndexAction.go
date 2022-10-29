package action

import (
	"fmt"
	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/pkg/mapper"
	"net/http"
)

type TwitIndexAction struct {
	Action
}

func NewTwitIndexAction() *TwitIndexAction {
	return new(TwitIndexAction)
}

func (self TwitIndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
	twitMapper := mapperManager.GetTwitMapper()

	/* Restore twits */
	twitNames, err1 := twitMapper.GetTwitNames()
	if err1 != nil {
		status := fmt.Sprintf("Fail in GetTwitNames on twitMapper: err = %+v", err1)
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
		SetClass("twit-index-table")

	twitTable.AddRow(widgets2.NewTableRowWidget().
		SetClass("twit-index-header").
		AddCell(widgets2.NewTableCellWidget().SetWidget(widgets2.NewTextWidgetWithText("Name"))).
		AddCell(widgets2.NewTableCellWidget().SetWidget(widgets2.NewTextWidgetWithText("Action"))))

	for _, t := range twitNames {

		twitTable.AddRow(widgets2.NewTableRowWidget().
			AddCell(widgets2.NewTableCellWidget().SetWidget(widgets2.NewTextWidgetWithText(t.GetName()))).
			AddCell(widgets2.NewTableCellWidget().SetWidget(widgets2.NewLinkWidget().
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
