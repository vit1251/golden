package mailer

import (
	"fmt"
	"github.com/vit1251/golden/pkg/mailer/stream"
)

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

	err1 := s.OpenSession(mailer.ServerAddr)
	if err1 != nil {
		mailer.report.SetStatus(fmt.Sprintf("Unable to open session: err = %+v", err1))
		return NewMailerStateEnd()
	}

	mailer.stream = s

	return NewMailerWaitConn()

}
