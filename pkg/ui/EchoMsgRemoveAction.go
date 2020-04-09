package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	area2 "github.com/vit1251/golden/pkg/area"
	msg2 "github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/ui/widgets"
	"log"
	"net/http"
)

type EchoMsgRemoveAction struct {
	Action
}

func NewEchoMsgRemoveAction() *EchoMsgRemoveAction {
	ra:=new(EchoMsgRemoveAction)
	return ra
}

func (self *EchoMsgRemoveAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var areaManager *area2.AreaManager
	var messageManager *msg2.MessageManager
	self.Container.Invoke(func(am *area2.AreaManager, mm *msg2.MessageManager) {
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
	msg, err2 := messageManager.GetMessageByHash(echoTag, msgHash)
	if err2 != nil {
		response := fmt.Sprintf("Fail on GetMessageByHash on MessageManager: err = %+v", err2)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets.NewMainMenuWidget()
	vBox.Add(mmw)

	//<h1>Delete message?</h1>
	headerWidget := widgets.NewHeaderWidget().
		SetTitle("Delete message?")
	vBox.Add(headerWidget)

	//
	formWidget := widgets.NewFormWidget().
		SetMethod("POST").
		SetAction(fmt.Sprintf("/echo/%s/message/%s/remove/complete", area.Name(), msg.Hash))
	formVBox := widgets.NewVBoxWidget()
	formWidget.SetWidget(formVBox)
	vBox.Add(formWidget)

	qustionWidget := widgets.NewDivWidget().
		SetContent(fmt.Sprintf("A you sure to remove '%s' message?", msg.Subject))
	formVBox.Add(qustionWidget)

	buttonWidget := widgets.NewFormButtonWidget().
		SetTitle("Remove")
	formVBox.Add(buttonWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
