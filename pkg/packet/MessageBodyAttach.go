package packet

import (
	"bytes"
	"fmt"
	"github.com/vit1251/golden/pkg/uue"
)

type MessageBodyAttach struct {
	permission string       /* Permission          */
	name       string       /* Name                */
	uue        []string     /* UUE rows            */
	data       bytes.Buffer /* Attachment data     */
	decoder    *uue.Decoder /* UUE decoder         */
	encoder    *uue.Encoder /* UUE encoder         */
}

func NewMessageBodyAttach() *MessageBodyAttach {
	msgBody := new(MessageBodyAttach)
	msgBody.decoder = uue.NewDecoder(&msgBody.data)
	msgBody.encoder = uue.NewEncoder(&msgBody.data)
	return msgBody
}

func (self *MessageBodyAttach) Write(row []byte) error {
	self.decoder.Decode(row)
	return nil
}

func (self *MessageBodyAttach) Read() ([]byte, error) {
	// TODO - encode not yet implement ...
	return nil, fmt.Errorf("not yet implemented")
}

func (self *MessageBodyAttach) SetPermission(permission string) {
	self.permission = permission
}

func (self *MessageBodyAttach) SetName(name string) {
	self.name = name
}

func (self *MessageBodyAttach) Len() int {
	return self.data.Len()
}

func (self *MessageBodyAttach) SetData(data []byte) {
	self.data.Reset()
	self.data.Write(data)
}

func (self *MessageBodyAttach) GetData() []byte {
	return self.data.Bytes()
}

func (self *MessageBodyAttach) GetName() string {
	return self.name
}
