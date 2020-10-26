package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/netmail"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/ui/widgets"
	"net/http"
)

type NetmailReplyAction struct {
	Action
}

func NewNetmailReplyAction() *NetmailReplyAction {
	return new(NetmailReplyAction)
}

func (self *NetmailReplyAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var netmailManager *netmail.NetmailManager
	var configManager *setup.ConfigManager
	self.Container.Invoke(func(nm *netmail.NetmailManager, cm *setup.ConfigManager) {
		netmailManager = nm
		configManager = cm
	})

	vars := mux.Vars(r)

	/* Recover message */
	msgHash := vars["msgid"]

	origMsg, err3 := netmailManager.GetMessageByHash(msgHash)
	if err3 != nil {
		response := fmt.Sprintf("Fail on GetMessageByHash")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	if origMsg == nil {
		response := fmt.Sprintf("Fail on GetMessageByHash")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	content := origMsg.GetContent()

	msgFrom := origMsg.From
	msgFromAddr := "" // origMsg.FromAddr
	msgSubject := origMsg.Subject

	newContent := content

	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()

	mmw := widgets.NewMainMenuWidget()
	vBox.Add(mmw)

	container := widgets.NewDivWidget().SetClass("container")
	vBox.Add(container)

	containerVBox := widgets.NewVBoxWidget()
	container.SetWidget(containerVBox)

	section := widgets.NewSectionWidget().
		SetTitle("Reply NETMAIL message")

	composeForm := widgets.NewFormWidget().
		SetAction(fmt.Sprintf("/netmail/%s/reply/complete", msgHash)).
		SetMethod("POST")

	composeForm.SetWidget(widgets.NewVBoxWidget().
		Add(widgets.NewFormInputWidget().SetTitle("ToName").SetName("to").SetPlaceholder("Vitold Sedyshev").SetValue(msgFrom)).
		Add(widgets.NewFormInputWidget().SetTitle("ToAddr").SetName("to_addr").SetPlaceholder("2:5023/24.3752").SetValue(msgFromAddr)).
		Add(widgets.NewFormInputWidget().SetTitle("Subject").SetName("subject").SetPlaceholder("RE: Hello, world!").SetValue(msgSubject)).
		Add(widgets.NewFormTextWidget().SetName("body").SetValue(newContent)).
		Add(widgets.NewFormButtonWidget().SetType("submit").SetTitle("Send")))

	section.SetWidget(composeForm)
	containerVBox.Add(section)

	bw.SetWidget(vBox)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
	}

}
