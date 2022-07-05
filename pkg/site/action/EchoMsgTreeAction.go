package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/site/widgets"
	"github.com/vit1251/golden/pkg/um"
	"log"
	"net/http"
)

type EchoMsgTreeAction struct {
	Action
}

func NewEchoMsgTreeAction() *EchoMsgTreeAction {
	return new(EchoMsgTreeAction)
}

func (self EchoMsgTreeAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	echoMapper := mapperManager.GetEchoMapper()

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

	// Views

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

	/* Create node tree */
	tree := msg.NewMessageTree()
	for _, m := range msgHeaders {
		tree.RegisterMessage(m)
	}

	/* Render tree */
	log.Printf("Tree = %+v", tree.GetRoot())
	nodeTree := self.renderTree(newArea, *tree.GetRoot())

	containerVBox.Add(nodeTree)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self EchoMsgTreeAction) renderTree(area *mapper.Area, root msg.MessageNode) widgets.IWidget {

	urlManager := um.RestoreUrlManager(self.GetRegistry())

	list := widgets.NewListWidget()

	for _, node := range root.Items {

		newMsg := node.GetValue()

		newSubject := fmt.Sprintf("%s", newMsg.Subject)

		msgAddr := urlManager.CreateUrl("/echo/{area_index}/message/{message_index}/view").
			SetParam("area_index", area.GetAreaIndex()).
			SetParam("message_index", newMsg.Hash).
			Build()
		newLink := widgets.NewLinkWidget().
			SetContent(newSubject).
			SetLink(msgAddr)

		if len(node.Items) == 0 {

			list.AddItem(newLink)

		} else {

			vbox := widgets.NewVBoxWidget()

			vbox.Add(newLink)
			vbox.Add(self.renderTree(area, *node))

			list.AddItem(vbox)

		}

	}

	return list
}
