package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/site/utils"
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
	log.Printf("My name is %s", myName)

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
	container.AddWidget(containerVBox)

	/* Context actions */
	amw := widgets.NewActionMenuWidget().
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/echo/%s/message/compose", newArea.GetName())).
			SetIcon("icofont-edit").
			SetLabel("Compose")).
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/echo/%s/mark", newArea.GetName())).
			SetIcon("icofont-mark-as-read").
			SetLabel("Mark as read")).
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/echo/%s/update", newArea.GetName())).
			SetIcon("icofont-update").
			SetLabel("Settings"))

	containerVBox.Add(amw)

	indexTable := widgets.NewDivWidget().
		SetClass("echo-msg-index-table").
		SetStyle("width: 100%")

	/* Render message index column names */
	//rowWidget := self.renderHeader()
	//indexTable.AddWidget(rowWidget)

	for _, msg := range msgHeaders {

		/* Check message sender and receiver in Twit names */
		if self.checkSenderInTwit(msg, twitNames) {
			continue
		}

		/* Render message row */
		msgRow := self.renderRow(&msg, myName)
		indexTable.AddWidget(msgRow)

	}
	containerVBox.Add(indexTable)

	/* Render server response */
	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self *EchoMsgIndexAction) checkSenderInTwit(msg msg.Message, twitNames []mapper.Twit) bool {
	for _, t := range twitNames {
		var twitName string = t.GetName()
		if strings.EqualFold(twitName, msg.From) || strings.EqualFold(twitName, msg.To) {
			return true
		}
	}
	return false
}

func (self *EchoMsgIndexAction) renderHeader() widgets.IWidget {

	rowWidget := widgets.NewDivWidget()
	rowWidget.SetStyle("display: flex")
	rowWidget.SetStyle("direction: column")

	userpicWidget := widgets.NewDivWidget().
		SetWidth("30px").
		SetHeight("30px").
		SetContent("VS")
	rowWidget.AddWidget(userpicWidget)

	sourceWidget := widgets.NewDivWidget().
		SetWidth("190px").
		SetHeight("38px").
		SetContent("Vitold Sedyshev")
	rowWidget.AddWidget(sourceWidget)

	subjectWidget := widgets.NewDivWidget().
		SetStyle("min-width: 350px").
		SetHeight("38px").
		SetStyle("flex-grow: 1").
		SetContent("Пора готовить к новому году")

	rowWidget.AddWidget(subjectWidget)

	dateWidget := widgets.NewDivWidget().
		SetHeight("38px").
		SetWidth("160px").
		SetContent("2021-12-23")
	rowWidget.AddWidget(dateWidget)

	return rowWidget

}

func (self *EchoMsgIndexAction) renderRow(m *msg.Message, myName string) widgets.IWidget {

	/* Make message row container */
	rowWidget := widgets.NewDivWidget().
		SetStyle("display: flex").
		SetStyle("direction: column").
		SetStyle("align-items: center")

	var classNames []string
	classNames = append(classNames, "echo-msg-index-item")
	if m.ViewCount == 0 {
		classNames = append(classNames, "echo-msg-index-item-new")
	}
	if strings.EqualFold(m.From, myName) || strings.EqualFold(m.To, myName) {
		classNames = append(classNames, "echo-msg-index-item-own")
	}
	rowWidget.SetClass(strings.Join(classNames, " "))

	/* Render user pic */
	nameTitle := utils.TextHelper_makeNameTitle(m.From)
	nameColor := utils.TextHelper_makeColorByName(m.From)
	userpicWidget := widgets.NewDivWidget().
		SetWidth("30px").
		SetHeight("30px").
		SetStyle("flex-shrink: 0").
		//SetStyle("border: 1px solid yellow").
		SetStyle(fmt.Sprintf("background-color: %s", nameColor)).
		SetStyle("border-radius: 50%").
		SetContent(nameTitle)
	rowWidget.AddWidget(userpicWidget)

	/* Render sender name */
	sourceWidget := widgets.NewDivWidget().
		SetWidth("190px").
		SetHeight("38px").
		SetStyle("flex-shrink: 0").
		SetStyle("white-space: nowrap").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		//SetStyle("border: 1px solid green").
		SetContent(m.From)
	rowWidget.AddWidget(sourceWidget)
	// TODO - add `m.To` under m.From ....

	subjectWidget := widgets.NewDivWidget().
		SetStyle("min-width: 350px").
		SetHeight("38px").
		SetStyle("flex-grow: 1").
		SetStyle("white-space: nowrap").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		//SetStyle("border: 1px solid red").
		SetContent(m.Subject)

	rowWidget.AddWidget(subjectWidget)

	msgDate := utils.DateHelper_renderDate(m.DateWritten)
	dateWidget := widgets.NewDivWidget().
		SetHeight("38px").
		SetWidth("160px").
		SetStyle("flex-shrink: 0").
		//SetStyle("border: 1px solid blue").
		SetContent(msgDate)
	rowWidget.AddWidget(dateWidget)

	/* Link container */
	navigateAddress := fmt.Sprintf("/echo/%s/message/%s/view", m.Area, m.Hash)
	navigateItem := widgets.NewLinkWidget().
		SetLink(navigateAddress).
		AddWidget(rowWidget)

	return navigateItem

}
