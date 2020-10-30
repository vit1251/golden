package mailer

import "github.com/vit1251/golden/pkg/mailer/stream"

type MailerStateConnInit struct {
	MailerState
}

func NewMailerStateConnInit() *MailerStateConnInit {
	msc := new(MailerStateConnInit)
	return msc
}

func (self *MailerStateConnInit) String() string {
	return "MailerStateConnInit"
}

func (self *MailerStateConnInit) Process(mailer *Mailer) IMailerState {
	s := stream.NewMailerStream()
	if err1 := s.OpenSession(mailer.ServerAddr); err1 == nil {
		mailer.stream = s
		return NewMailerWaitConn()
	} else {
		return NewMailerStateEnd()
	}
}
