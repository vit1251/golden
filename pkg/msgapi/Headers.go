package msgapi

import (
	"github.com/vit1251/golden/pkg/utils"
//	"time"
)

type Header struct {
	ID string                // Message ID (i.e. index value provided by engine to repetable search in base)
	UMSGID uint32            // Message ID in UMSGID value
	From string              // From
	FromAddr string
	To string                // To
	ToAddr string
	Subject string           // Subject
	DateWritten string       // Message date written
	DateArrived string       // Message date arrival
	Hash string              // MD5 digest of message body
	Body []byte              // Message body
}

func (self Header) GetContent() string {
	return utils.MakeString(self.Body[:])
}
