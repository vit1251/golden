package mailer

import "log"

type MailerOutbound struct {
}

type Item struct {
	Name string
	AbsolutePath string
//	Type
}

func (self *MailerOutbound) TransmitFile(filename string) {
	log.Printf("Schedule to transmit %s", filename)
}

func NewMailerOutbound() *MailerOutbound {
	result := new(MailerOutbound)
	return result
}

func (self *MailerOutbound) GetItems() []*Item {
	return nil
}
