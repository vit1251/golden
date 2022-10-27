package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/internal/i18n"
	utils2 "github.com/vit1251/golden/internal/site/utils"
	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/internal/um"
	"github.com/vit1251/golden/pkg/config"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/msg"
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

func (self *EchoMsgIndexAction) renderActions(newArea *mapper.Area) widgets2.IWidget {

	urlManager := um.RestoreUrlManager(self.GetRegistry())

	var mainLanguage string = i18n.GetDefaultLanguage()

	actionBar := widgets2.NewActionMenuWidget()

	/* Compose */
	composeAddr := urlManager.CreateUrl("/echo/{area_index}/message/compose").
		SetParam("area_index", newArea.GetAreaIndex()).
		Build()
	composeTitle := i18n.GetText(mainLanguage, "EchoMsgIndexAction", "action-compose-button")
	actionBar.Add(widgets2.NewMenuAction().
		SetLink(composeAddr).
		SetIcon("icofont-edit").
		SetClass("mr-2").
		SetLabel(composeTitle))

	/* Tree */
	treeAddr := urlManager.CreateUrl("/echo/{area_index}/tree").
		SetParam("area_index", newArea.GetAreaIndex()).
		Build()
	treeTitle := i18n.GetText(mainLanguage, "EchoMsgIndexAction", "action-tree-button")
	actionBar.Add(widgets2.NewMenuAction().
		SetLink(treeAddr).
		SetIcon("icon-tree").
		SetClass("mr-2").
		SetLabel(treeTitle))

	/* Mark as read */
	markAsReadAddr := urlManager.CreateUrl("/echo/{area_index}/mark").
		SetParam("area_index", newArea.GetAreaIndex()).
		Build()
	markAsReadTitle := i18n.GetText(mainLanguage, "EchoMsgIndexAction", "action-mark-as-read-button")
	actionBar.Add(widgets2.NewMenuAction().
		SetLink(markAsReadAddr).
		SetIcon("icofont-mark-as-read").
		SetClass("mr-2").
		SetLabel(markAsReadTitle))

	/* Settings */
	settingsAddr := urlManager.CreateUrl("/echo/{area_index}/update").
		SetParam("area_index", newArea.GetAreaIndex()).
		Build()
	settingsTitle := i18n.GetText(mainLanguage, "EchoMsgIndexAction", "action-settings-button")
	actionBar.Add(widgets2.NewMenuAction().
		SetLink(settingsAddr).
		SetIcon("icofont-update").
		SetClass("mr-2").
		SetLabel(settingsTitle))

	return actionBar

}

func (self *EchoMsgIndexAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	configManager := config.RestoreConfigManager(self.GetRegistry())
	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	echoMapper := mapperManager.GetEchoMapper()
	twitMapper := mapperManager.GetTwitMapper()

	newConfig := configManager.GetConfig()

	/* Parse URL parameters */
	vars := mux.Vars(r)
	areaIndex := vars["echoname"]
	log.Printf("areaIndex = %v", areaIndex)

	newArea, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName where areaIndex is %s: err = %+v", areaIndex, err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("area = %+v", newArea)

	/* Get message headers */
	var areaName string = newArea.GetName()
	msgHeaders, err2 := echoMapper.GetMessageHeaders(areaName)
	if err2 != nil {
		response := fmt.Sprintf("Fail on GetMessageHeaders where areaName is %s: err = %+v", areaName, err2)
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

	bw := widgets2.NewBaseWidget()

	vBox := widgets2.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets2.NewDivWidget()
	container.SetClass("container")
	vBox.Add(container)

	containerVBox := widgets2.NewVBoxWidget()
	container.AddWidget(containerVBox)

	/* Context actions */
	actionsBar := self.renderActions(newArea)
	containerVBox.Add(actionsBar)

	indexTable := widgets2.NewDivWidget().
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
		msgRow := self.renderRow(newArea, &msg, newConfig.Main.RealName)
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

func (self *EchoMsgIndexAction) renderRow(area *mapper.Area, m *msg.Message, myName string) widgets2.IWidget {

	urlManager := um.RestoreUrlManager(self.GetRegistry())

	/* Make message row container */
	rowWidget := widgets2.NewDivWidget().
		SetStyle("display: flex").
		SetStyle("direction: column").
		SetStyle("align-items: center")

	var classNames []string
	classNames = append(classNames, "echo-msg-index-item")
	if m.ViewCount == 0 {
		classNames = append(classNames, "echo-msg-index-item-new")
	}
	rowWidget.SetClass(strings.Join(classNames, " "))

	/* Render user pic */
	nameTitle := utils2.TextHelper_makeNameTitle(m.From)
	nameColor := utils2.TextHelper_makeColorByName(m.From)
	userpicWidget := widgets2.NewDivWidget().
		SetWidth("30px").
		SetHeight("30px").
		SetStyle("flex-shrink: 0").
		SetStyle("margin-right: 8px").
		SetStyle("display: flex").
		SetStyle("align-items: center").
		SetStyle("justify-content: center").
		//SetStyle("border: 1px solid yellow").
		SetStyle(fmt.Sprintf("background-color: %s", nameColor)).
		SetStyle("border-radius: 50%").
		SetContent(nameTitle)
	rowWidget.AddWidget(userpicWidget)

	/* Render sender name */
	sourceWidget := widgets2.NewDivWidget().
		SetWidth("190px").
		SetHeight("38px").
		SetStyle("flex-shrink: 0").
		SetStyle("white-space: nowrap").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		SetStyle("display: flex").
		SetStyle("flex-direction: column").
		SetStyle("align-items: flex-start").
		SetStyle("justify-content: center").
		SetContent(m.From)
	rowWidget.AddWidget(sourceWidget)
	// TODO - add `m.To` under m.From ....

	/* Render NEW point */
	var newPointContent string = ""
	if m.IsNew() {
		newPointContent = "•"
	}
	newPointWidget := widgets2.NewDivWidget().
		SetWidth("20px").
		SetHeight("38px").
		SetStyle("flex-shrink: 0").
		SetStyle("white-space: nowrap").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		SetStyle("display: flex").
		SetStyle("flex-direction: column").
		SetStyle("align-items: center").
		SetStyle("justify-content: center").
		//SetStyle("border: 1px solid green").
		SetStyle("color: yellow").
		SetContent(newPointContent)
	rowWidget.AddWidget(newPointWidget)

	/* Render subject */
	subjectWidget := widgets2.NewDivWidget().
		SetStyle("min-width: 350px").
		SetHeight("38px").
		SetStyle("flex-grow: 1").
		SetStyle("white-space: nowrap").
		SetStyle("display: flex").
		SetStyle("flex-direction: column").
		SetStyle("align-items: flex-start").
		SetStyle("justify-content: center").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		//SetStyle("border: 1px solid red").
		SetContent(m.Subject)

	rowWidget.AddWidget(subjectWidget)

	/* Render date and time */
	msgDate := utils2.DateHelper_renderDate(m.DateWritten)
	dateWidget := widgets2.NewDivWidget().
		SetHeight("38px").
		SetWidth("180px").
		SetStyle("flex-shrink: 0").
		SetStyle("white-space: nowrap").
		SetStyle("display: flex").
		SetStyle("flex-direction: column").
		SetStyle("align-items: flex-end").
		SetStyle("justify-content: center").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		//SetStyle("border: 1px solid blue").
		SetContent(msgDate)
	rowWidget.AddWidget(dateWidget)

	/* Link container */
	navigateAddress := urlManager.CreateUrl("/echo/{area_index}/message/{message_hash}/view").
		SetParam("area_index", area.GetAreaIndex()).
		SetParam("message_hash", m.Hash).
		Build()

	navigateItem := widgets2.NewLinkWidget().
		SetLink(navigateAddress).
		AddWidget(rowWidget)

	return navigateItem

}
