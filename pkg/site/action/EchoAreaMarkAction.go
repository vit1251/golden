package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/site/widgets"
	"github.com/vit1251/golden/pkg/um"
	"log"
	"net/http"
)

type EchoAreaMarkAction struct {
	Action
}

func NewEchoAreaMarkAction() *EchoAreaMarkAction {
	ra := new(EchoAreaMarkAction)
	return ra
}

func (self *EchoAreaMarkAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	urlManager := um.RestoreUrlManager(self.GetRegistry())
	mapperManager := mapper.RestoreMapperManager(self.registry)
	echoAreaMapper := mapperManager.GetEchoAreaMapper()

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
		SetTitle("Purge area?")
	containerVBox.Add(headerWidget)

	//
	markCompleteAddr := urlManager.CreateUrl("/echo/{area_index}/mark/complete").
		SetParam("area_index", area.GetAreaIndex()).
		Build()
	formWidget := widgets.NewFormWidget().
		SetMethod("POST").
		SetAction(markCompleteAddr)
	formVBox := widgets.NewVBoxWidget()
	formWidget.SetWidget(formVBox)
	containerVBox.Add(formWidget)

	qustionWidget := widgets.NewDivWidget().
		SetContent(fmt.Sprintf("A you sure to mark all read in '%s' area?", area.GetName()))
	formVBox.Add(qustionWidget)

	buttonWidget := widgets.NewFormButtonWidget().
		SetTitle("Mark")
	formVBox.Add(buttonWidget)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}
