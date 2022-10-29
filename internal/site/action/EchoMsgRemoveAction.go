package action

import (
	"fmt"
	"github.com/gorilla/mux"
	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/internal/um"
	"github.com/vit1251/golden/pkg/mapper"
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

	urlManager := um.RestoreUrlManager(self.GetRegistry())
	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
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

	//<h1>Delete message?</h1>
	headerWidget := widgets2.NewHeaderWidget().
		SetTitle("Delete message?")
	containerVBox.Add(headerWidget)

	//
	removeCompleteAddr := urlManager.CreateUrl("/echo/{area_index}/message/{message_index}/remove/complete").
		SetParam("area_index", area.GetAreaIndex()).
		SetParam("message_index", msg.Hash).
		Build()
	formWidget := widgets2.NewFormWidget().
		SetMethod("POST").
		SetAction(removeCompleteAddr)
	formVBox := widgets2.NewVBoxWidget()
	formWidget.SetWidget(formVBox)
	containerVBox.Add(formWidget)

	qustionWidget := widgets2.NewDivWidget().
		SetContent(fmt.Sprintf("A you sure to remove '%s' message?", msg.Subject))
	formVBox.Add(qustionWidget)

	buttonWidget := widgets2.NewFormButtonWidget().
		SetTitle("Remove")
	formVBox.Add(buttonWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
