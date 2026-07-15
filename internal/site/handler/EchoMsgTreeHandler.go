package handler

import (
    "net/http"

    "github.com/vit1251/golden/internal/site/views"
    "github.com/vit1251/golden/pkg/mapper"
    "github.com/vit1251/golden/pkg/msg"
    "github.com/vit1251/golden/pkg/registry"
)

type EchoMsgTreeHandler struct {
    registry *registry.Container
}

func NewEchoMsgTreeHandler(registry *registry.Container) *EchoMsgTreeHandler {
    return &EchoMsgTreeHandler{
	registry: registry,
    }
}

func mapTreeNodes(nodes []*msg.MessageNode, areaIndex string) []views.TreeNode {
    var result []views.TreeNode
    for _, node := range nodes {
        m := node.GetValue()
        n := views.TreeNode{
            Subject: m.Subject,
            URL:     "/echo/" + areaIndex + "/message/" + m.Hash + "/view",
            Items:   mapTreeNodes(node.Items, areaIndex),
        }
        result = append(result, n)
    }
    return result
}

func (h *EchoMsgTreeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

    mapperManager := mapper.RestoreMapperManager(h.registry)
    echoAreaMapper := mapperManager.GetEchoAreaMapper()
    echoMapper := mapperManager.GetEchoMapper()

    /* Parse URL parameters */
    var areaIndex string = r.PathValue("echoname")

    newArea, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
    if err1 != nil {
	http.Error(w, "Fail on GetAreaByAreaIndex", http.StatusInternalServerError)
	return
    }

    /* Get message headers */
    var areaName string = newArea.GetName()
    msgHeaders, err2 := echoMapper.GetMessageHeaders(areaName)
    if err2 != nil {
    	http.Error(w, "Fail on GetMessageHeaders", http.StatusInternalServerError)
    	return
    }

    tree := msg.NewMessageTree()
    for _, m := range msgHeaders {
	tree.RegisterMessage(m)
    }

    data := views.EchoMsgTreeData{
	Actions: []views.ToolbarAction{
    	    {Label: "Back", URL: "/echo/" + areaIndex, Icon: "arrow-left"},
	},
	AreaName: areaName,
	Tree:     mapTreeNodes(tree.GetRoot().Items, areaIndex),
    }
    err := views.Page(areaName, views.EchoMsgTreeView(data)).Render(w)
    if err != nil {
	http.Error(w, err.Error(), http.StatusInternalServerError)
    }

}
