package msg

import (
	"fmt"
	"testing"
)

func TestMessageTree_RegisterMessage(t *testing.T) {

	msgTree := NewMessageTree()
	msgTree.RegisterMessage(Message{MsgID: "A", Reply: ""})
	msgTree.RegisterMessage(Message{MsgID: "B", Reply: "A"})
	msgTree.RegisterMessage(Message{MsgID: "C", Reply: "B"})
	msgTree.RegisterMessage(Message{MsgID: "D", Reply: "C"})

	root := msgTree.GetRoot()
	renderTree(*root, 0)

}

func renderTree(root MessageNode, level int) {
	fmt.Printf("[%d] Node: value = %+v orphan = %+v childrens = %d\n", level, root.value, root.orphan, len(root.Items))
	for _, node := range root.Items {
		renderTree(*node, level+1)
	}
}
