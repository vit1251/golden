package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/netmail"
	"github.com/vit1251/golden/pkg/site/widgets"
	"log"
	"net/http"
)

type NetmailReplyAction struct {
	Action
}

func NewNetmailReplyAction() *NetmailReplyAction {
	return new(NetmailReplyAction)
}

func (self *NetmailReplyAction) preprocessMessage(origMsg *netmail.NetmailMessage) string {
	cmap := msg.NewMessageAuthorParser()
	ma, _ := cmap.Parse(origMsg.From)

	/* Make reply content */
	mtp := msg.NewMessageTextProcessor()
	mtp.Prepare(origMsg.Content)
	newContent := mtp.Content()
	log.Printf("reply: orig = %+v", newContent)

	/* Message replay transform */
	mrt := msg.NewMessageReplyTransformer()
	mrt.SetAuthor(ma.QuoteName)
	newContent2 := mrt.Transform(newContent)

	return newContent2
}

func (self *NetmailReplyAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	netmailManager := self.restoreNetmailManager()
	//configManager := self.restoreConfigManager()

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

	msgFrom := origMsg.From
	msgFromAddr := fmt.Sprintf("%s", origMsg.OrigAddr)
	newBody := self.preprocessMessage(origMsg)

	/* Compact header*/
	sc:= msg.NewSubjectCompactor()
	newSubject := sc.Compact(origMsg.Subject)

	/* Render */
	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()

	mmw := self.makeMenu()
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
		Add(widgets.NewFormInputWidget().SetTitle("ToName").SetName("to").SetPlaceholder("Enter FidoNet destination name").SetValue(msgFrom)).
		Add(widgets.NewFormInputWidget().SetTitle("ToAddr").SetName("to_addr").SetPlaceholder("Enter FidoNet destination address Zone:Net/Node.Point").SetValue(msgFromAddr)).
		Add(widgets.NewFormInputWidget().SetTitle("Subject").SetName("subject").SetPlaceholder("Enter message body").SetValue(newSubject)).
		Add(widgets.NewFormTextWidget().SetName("body").SetValue(newBody)).
		Add(widgets.NewFormButtonWidget().SetType("submit").SetTitle("Send")))

	section.SetWidget(composeForm)
	containerVBox.Add(section)

	bw.SetWidget(vBox)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
	}

}
