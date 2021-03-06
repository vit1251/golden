package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/site/widgets"
	"log"
	"net/http"
	"strings"
)

type EchoMsgIndexAction struct {
	Action
}

func NewEchoMsgIndexAction() *EchoMsgIndexAction {
	ea := new(EchoMsgIndexAction)
	return ea
}

func (self *EchoMsgIndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	echoMapper := mapperManager.GetEchoMapper()
	twitMapper := mapperManager.GetTwitMapper()
	configMapper := mapperManager.GetConfigMapper()

	/* My name */
	myName, _ := configMapper.Get("main", "RealName")

	/* Parse URL parameters */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	newArea, err1 := echoAreaMapper.GetAreaByName(echoTag)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName where echoTag is %s: err = %+v", echoTag, err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("area = %+v", newArea)

	/* Get message headers */
	msgHeaders, err2 := echoMapper.GetMessageHeaders(echoTag)
	if err2 != nil {
		response := fmt.Sprintf("Fail on GetMessageHeaders where echoTag is %s: err = %+v", echoTag, err2)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("msgHeaders = %+v", msgHeaders)
	for _, msg := range msgHeaders {
		log.Printf("msg = %+v", msg)
	}

	/* Get twits */
	twitNames, err3 := twitMapper.GetTwitNames()
	if err3 != nil {
		response := fmt.Sprintf("Fail on GetTwitNames on twitMapper: err = %+v", err3)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	// Views

	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")
	vBox.Add(container)

	containerVBox := widgets.NewVBoxWidget()
	container.SetWidget(containerVBox)

	/* Context actions */
	amw := widgets.NewActionMenuWidget().
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/echo/%s/message/compose", newArea.GetName())).
			SetIcon("icofont-edit").
			SetLabel("Compose")).
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/echo/%s/update", newArea.GetName())).
			SetIcon("icofont-update").
			SetLabel("Settings"))

	containerVBox.Add(amw)

	indexTable := widgets.NewTableWidget().
		SetClass("echo-msg-index-table")

	indexTable.AddRow(widgets.NewTableRowWidget().
		SetClass("echo-msg-index-header").
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("From"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("To"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Subject"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Date"))).
		AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText("Action"))))

	for _, msg := range msgHeaders {
		log.Printf("msg = %+v", msg)

		var allowView bool = true
		for _, t := range twitNames {
			var twitName string = t.GetName()
			if strings.EqualFold(twitName, msg.From) || strings.EqualFold(twitName, msg.To) {
				allowView = false
			}
		}


		actions := widgets.NewVBoxWidget()

		if allowView {
			actions.Add(
				widgets.NewLinkWidget().
					SetContent("View").
					SetClass("btn").
					SetLink(fmt.Sprintf("/echo/%s/message/%s/view", msg.Area, msg.Hash)))
		} else {
			actions.Add(
				widgets.NewLinkWidget().
					SetContent("Remove").
					SetClass("btn").
					SetLink(fmt.Sprintf("/echo/%s/message/%s/remove", msg.Area, msg.Hash)))
		}

		row := widgets.NewTableRowWidget().
			SetClass("echo-msg-index-item").
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(msg.From))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(msg.To))).
			AddCell(widgets.NewTableCellWidget().SetWidget(widgets.NewTextWidgetWithText(msg.Subject))).
			AddCell(widgets.NewTableCellWidget().SetClass("echo-msg-index-date").SetWidget(widgets.NewTextWidgetWithText(msg.GetAge()))).
			AddCell(widgets.NewTableCellWidget().SetWidget(actions))

		var classes []string

		if msg.ViewCount == 0 {
			classes = append(classes, "echo-msg-index-item-new")
		}

		if strings.EqualFold(msg.From, myName) || strings.EqualFold(msg.To, myName) {
			classes = append(classes, "echo-msg-index-item-own")
		}

		row.SetClass(strings.Join(classes, " "))

		indexTable.AddRow(row)
	}

	containerVBox.Add(indexTable)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
