package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/site/widgets"
	"log"
	"net/http"
)

type EchoMsgRemoveAction struct {
	Action
}

func NewEchoMsgRemoveAction() *EchoMsgRemoveAction {
	ra := new(EchoMsgRemoveAction)
	return ra
}

func (self *EchoMsgRemoveAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	urlManager := self.restoreUrlManager()
	mapperManager := self.restoreMapperManager()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	echoMapper := mapperManager.GetEchoMapper()

	//
	vars := mux.Vars(r)
	areaIndex := vars["echoname"]
	log.Printf("areaIndex = %v", areaIndex)

	//
	area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	//
	msgHash := vars["msgid"]
	var areaName string = area.GetName()
	msg, err2 := echoMapper.GetMessageByHash(areaName, msgHash)
	if err2 != nil {
		response := fmt.Sprintf("Fail on GetMessageByHash on echoMapper: err = %+v", err2)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

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

	//<h1>Delete message?</h1>
	headerWidget := widgets.NewHeaderWidget().
		SetTitle("Delete message?")
	containerVBox.Add(headerWidget)

	//
	removeCompleteAddr := urlManager.CreateUrl("/echo/{area_index}/message/{message_index}/remove/complete").
		SetParam("area_index", area.GetAreaIndex()).
		SetParam("message_index", msg.Hash).
		Build()
	formWidget := widgets.NewFormWidget().
		SetMethod("POST").
		SetAction(removeCompleteAddr)
	formVBox := widgets.NewVBoxWidget()
	formWidget.SetWidget(formVBox)
	containerVBox.Add(formWidget)

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
