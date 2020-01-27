package mailer

type MailerOutbound struct {
}

func (self *MailerOutbound) TransmitFile(name string) {
	// TODO - make write
}

func NewMailerOutbound() (*MailerOutbound) {
	result := new(MailerOutbound)
	return result
}

