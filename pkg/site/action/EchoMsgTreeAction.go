package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/site/widgets"
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

	mapperManager := self.restoreMapperManager()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	echoMapper := mapperManager.GetEchoMapper()

	/* Parse URL parameters */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	newArea, err1 := echoAreaMapper.GetAreaByName(echoTag)
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetAreaByName where echoTag is %s: err = %+v", echoTag, err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("area = %+v", newArea)

	/* Get message headers */
	msgHeaders, err2 := echoMapper.GetMessageHeaders(echoTag)
	if err2 != nil {
		response := fmt.Sprintf("Fail on GetMessageHeaders where echoTag is %s: err = %+v", echoTag, err2)
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
	container.SetWidget(containerVBox)

	/* Create node tree */
	tree := msg.NewMessageTree()
	for _, m := range msgHeaders {
		tree.RegisterMessage(m)
	}

	/* Render tree */
	log.Printf("Tree = %+v", tree.GetRoot())
	nodeTree := self.renderTree(*tree.GetRoot())

	containerVBox.Add(nodeTree)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self EchoMsgTreeAction) renderTree(root msg.MessageNode) widgets.IWidget {

	list := widgets.NewListWidget()

	for _, node := range root.Items {

		newMsg := node.GetValue()

		newSubject := fmt.Sprintf("%s", newMsg.Subject)

		newLink := widgets.NewLinkWidget().
			SetContent(newSubject).
			SetLink(fmt.Sprintf("/echo/%s/message/%s/view", newMsg.Area, newMsg.Hash))

		if len(node.Items) == 0 {

			list.AddItem(newLink)

		} else {

			vbox := widgets.NewVBoxWidget()

			vbox.Add(newLink)
			vbox.Add(self.renderTree(node))

			list.AddItem(vbox)

		}

	}

	return list
}
