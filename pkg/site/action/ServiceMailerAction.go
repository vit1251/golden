package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/site/utils"
	"github.com/vit1251/golden/pkg/site/widgets"
	"log"
	"net/http"
	"strings"
	"time"
)

type ServiceMailerAction struct {
	Action
}

func NewServiceMailerAction() *ServiceMailerAction {
	sa := new(ServiceMailerAction)
	return sa
}

func (self *ServiceMailerAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	statMailerMapper := mapperManager.GetStatMailerMapper()

	/* Get statistics */
	sessions, err1 := statMailerMapper.GetMailerSummary()
	if err1 != nil {
		response := fmt.Sprintf("Fail GetStat on statMapper: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Create statistics */

	/* Render */
	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets.NewVBoxWidget()

	/* Context actions */
	amw := widgets.NewActionMenuWidget().
		Add(widgets.NewMenuAction().
			SetLink(fmt.Sprintf("/service/mailer/event")).
			SetIcon("icofont-update").
			SetLabel("Run"))

	containerVBox.Add(amw)

	container.AddWidget(containerVBox)

	vBox.Add(container)

	for _, s := range sessions {
		log.Printf("session = %#v", s)
		row := self.renderRow(&s)
		containerVBox.Add(row)
	}

	/* Render result */
	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self *ServiceMailerAction) renderRow(s *mapper.StatMailer) widgets.IWidget {

	/* Make message row container */
	rowWidget := widgets.NewDivWidget().
		SetStyle("display: flex").
		SetStyle("direction: column").
		SetStyle("align-items: center")

	var classNames []string
	classNames = append(classNames, "netmail-index-item")
	rowWidget.SetClass(strings.Join(classNames, " "))

	/* Render sender name */
	sessionStartTime := time.UnixMilli(s.SessionStart)
	sessionStart := utils.DateHelper_renderDate(sessionStartTime)
	startMailerWidget := widgets.NewDivWidget().
		SetWidth("190px").
		SetHeight("38px").
		SetStyle("flex-shrink: 0").
		SetStyle("white-space: nowrap").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		//SetStyle("border: 1px solid green").
		SetContent(sessionStart)
	rowWidget.AddWidget(startMailerWidget)

	/* Render sender name */

	sessionDuration := utils.DateHelper_renderDurationInMilli(s.SessionStop - s.SessionStart)
	stopMailerWidget := widgets.NewDivWidget().
		SetWidth("190px").
		SetHeight("38px").
		SetStyle("flex-shrink: 0").
		SetStyle("white-space: nowrap").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		//SetStyle("border: 1px solid green").
		SetContent(sessionDuration)
	rowWidget.AddWidget(stopMailerWidget)

	/* Render sender name */
	summaryWidget := widgets.NewDivWidget().
		SetWidth("320px").
		SetHeight("38px").
		SetStyle("flex-shrink: 0").
		SetStyle("white-space: nowrap").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		//SetStyle("border: 1px solid green").
		SetContent(s.Status)
	rowWidget.AddWidget(summaryWidget)

	return rowWidget

}
