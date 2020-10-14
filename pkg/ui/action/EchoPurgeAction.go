package action

import (
	"fmt"
	"github.com/gorilla/mux"
	msg2 "github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/ui/widgets"
	"log"
	"net/http"
)

type EchoPurgeAction struct {
	Action
}

func NewEchoPurgeAction() *EchoPurgeAction {
	ra := new(EchoPurgeAction)
	return ra
}

func (self *EchoPurgeAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var areaManager *msg2.AreaManager
	var messageManager *msg2.MessageManager
	self.Container.Invoke(func(am *msg2.AreaManager, mm *msg2.MessageManager) {
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

	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets.NewMainMenuWidget()
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")
	vBox.Add(container)

	containerVBox := widgets.NewVBoxWidget()
	container.SetWidget(containerVBox)

	//<h1>Delete message?</h1>
	headerWidget := widgets.NewHeaderWidget().
		SetTitle("Purge area?")
	containerVBox.Add(headerWidget)

	//
	formWidget := widgets.NewFormWidget().
		SetMethod("POST").
		SetAction(fmt.Sprintf("/echo/%s/purge/complete", area.Name()))
	formVBox := widgets.NewVBoxWidget()
	formWidget.SetWidget(formVBox)
	containerVBox.Add(formWidget)

	qustionWidget := widgets.NewDivWidget().
		SetContent(fmt.Sprintf("A you sure to remove '%s' area?", area.Name()))
	formVBox.Add(qustionWidget)

	buttonWidget := widgets.NewFormButtonWidget().
		SetTitle("Purge")
	formVBox.Add(buttonWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}