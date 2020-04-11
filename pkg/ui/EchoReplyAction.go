package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	area2 "github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/ui/widgets"
	"log"
	"net/http"
)

type EchoReplyAction struct {
	Action
}

func NewEchoReplyAction() *EchoReplyAction {
	ra := new(EchoReplyAction)
	return ra
}

func (self *EchoReplyAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var areaManager *area2.AreaManager
	var messageManager *msg.MessageManager
	self.Container.Invoke(func(am *area2.AreaManager, mm *msg.MessageManager) {
		areaManager = am
		messageManager = mm
	})

	//
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	//
	area, err1 := areaManager.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	//
	msgHash := vars["msgid"]
	origMsg, err3 := messageManager.GetMessageByHash(echoTag, msgHash)
	if err3 != nil {
		response := fmt.Sprintf("Fail on GetMessageByHash")
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	/* Detect sender */
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
	log.Printf("reply: reply = %+v", newContent2)

	//    <form method="post" action="/echo/{{ .Area.Name }}/message/{{ .Msg.Hash }}/reply/complete">
	//        <div><input class="input" type="text" name="to" value="{{ .Msg.From }}">
	//        <div><input class="input" type="text" value="{{ .Msg.Subject }}" name="subject">
	//        <textarea class="input input-area" name="body">{{ .Content }}</textarea>
	//        <button type="submit" name="action" value="send">Send</button>
	//    </form>

	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets.NewMainMenuWidget()
	vBox.Add(mmw)

	formVBox := widgets.NewVBoxWidget()

	formWidget := widgets.NewFormWidget()
	formWidget.
		SetMethod("POST").
		SetAction(fmt.Sprintf("/echo/%s/message/%s/reply/complete", area.Name(), origMsg.Hash)).
		SetWidget(formVBox)

	formVBox.Add(widgets.NewFormInputWidget().SetTitle("TO").SetName("to").SetValue(origMsg.From))
	formVBox.Add(widgets.NewFormInputWidget().SetTitle("SUBJ").SetName("subject").SetValue(fmt.Sprintf("RE: %s", origMsg.Subject)))
	formVBox.Add(widgets.NewFormTextWidget().SetName("body").SetValue(newContent2))
	formVBox.Add(widgets.NewFormButtonWidget().SetTitle("Compose").SetType("submit"))

	vBox.Add(formWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}


}
