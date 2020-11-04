package msg

import "log"

type MessageNode struct {
	value  *Message
	Items  []MessageNode
	orphan bool
}

func (self MessageNode) GetValue() *Message {
	return self.value
}

func NewMessageNode(msg *Message) *MessageNode {
	newNode := new(MessageNode)
	newNode.value = msg
	newNode.orphan = true
	return newNode
}

type MessageTree struct {
	index   map[string]*MessageNode
	order   []*MessageNode
}

func NewMessageTree() *MessageTree {
	tree := new(MessageTree)
	tree.index = make(map[string]*MessageNode)
	return tree
}

func (self MessageTree) GetRoot() *MessageNode {
	root := NewMessageNode(nil)
	for _, node := range self.order {
		if node.value.Reply == "" || node.orphan {
			root.AddNode(*node)
		}
	}
	return root
}

func (self *MessageNode) SearchById(msgid string) *MessageNode {
	//for _, n := range self.GetItems() {
		//n.SearchById()
	//}
	return nil
}

func (self *MessageNode) AddMessage(m Message) {
	node := NewMessageNode(&m)
	self.Items = append(self.Items, *node)
}

func (self *MessageNode) AddNode(node MessageNode) {
	self.Items = append(self.Items, node)
}

func (self *MessageTree) RegisterMessage(m Message) {
	var msgID string = m.MsgID
	node := NewMessageNode(&m)
	self.index[msgID] = node
	self.order = append(self.order, node)
	self.Compact()
}

func (self *MessageTree) Compact() {
	var keys []string
	for k, _ := range self.index {
		keys = append(keys, k)
	}
	log.Printf("keys = %+v", keys)
	for _, k := range keys {
		//
		node := self.index[k]
		if node.orphan {
			m := node.value
			reply := m.Reply
			//
			if parent, ok := self.index[reply]; ok {
				log.Printf("Attach %+v -> %+v", k, reply)
				parent.AddNode(*node)
				node.orphan = false
			}
		}
	}
}
