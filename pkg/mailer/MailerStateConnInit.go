package mailer

import (
	"fmt"
	"github.com/vit1251/golden/pkg/mailer/stream"
)

func mailerStateConnInit(mailer *Mailer) mailerStateFn {
	s := stream.NewMailerStream()
	err1 := s.OpenSession(mailer.ServerAddr)
	if err1 != nil {
		mailer.report.SetStatus(fmt.Sprintf("Unable to open session: err = %+v", err1))
		return mailerStateEnd
	}
	mailer.stream = s
	return mailerStateWaitConn
}
