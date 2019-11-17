package msgapi

import (
	"os"
	"fmt"
	"log"
)

type FidoMessage struct {
}

type FidoMessageBase struct {
	messageCount   int
	Msgs         []FidoMessage
}

func (self *FidoMessageBase) ReadBase(dirname string) ([]*Header, error) {
	for i := 1; i < 9999; i++ {
		filename := fmt.Sprintf("%s/%d.msg", dirname, i)
		log.Printf("Check message with name %s", filename)
		if !self.messageExists(filename) {
			log.Printf("No more message exists. Stop.")
			break
		}
		log.Printf("Message %s exists.", filename)
		self.messageCount += 1
	}
	return nil, nil
}

func (self *FidoMessageBase) GetMessageCount() int {
	return self.messageCount
}

func (self *FidoMessageBase) messageExists(filename string) bool {
    info, err := os.Stat(filename)
    if os.IsNotExist(err) {
        return false
    }
    return !info.IsDir()
}
