package mailer

import (
	"fmt"
	cmn "github.com/vit1251/golden/pkg/common"
	"time"
)

type MailerStateWaitConn struct {
	MailerState
}

func NewMailerWaitConn() *MailerStateWaitConn {
	msth := new(MailerStateWaitConn)
	return msth
}

func (self *MailerStateWaitConn) String() string {
	return "MailerStateWaitConn"
}

func (self *MailerStateWaitConn) makeSystemTime() string {
	now := time.Now().Format(time.RFC822)
	return now
}

func (self *MailerStateWaitConn) makeOperationSystemName() string {
	return cmn.GetPlatform()
}

func (self *MailerStateWaitConn) makeVersionString() string {
	appName := "GoldenMailer"
	appVersion := cmn.GetVersion()
	protocolVersion := "binkp/1.0"
	return fmt.Sprintf("%s/%s %s", appName, appVersion, protocolVersion)
}

func (self *MailerStateWaitConn) processWelcome(mailer *Mailer) {

	/* Send M_NUL frames with system info (optional) */
	if username := mailer.GetUserName(); username != "" {
		mailer.stream.WriteInfo("ZYZ", username)
	}
	if stationName := mailer.GetSystemName(); stationName != "" {
		mailer.stream.WriteInfo("SYS", stationName)
	}
	if location := mailer.GetLocation(); location != "" {
		mailer.stream.WriteInfo("LOC", location)
	}
	mailer.stream.WriteInfo("TIME", self.makeSystemTime())
	mailer.stream.WriteInfo("OS", self.makeOperationSystemName())
	mailer.stream.WriteInfo("VER", self.makeVersionString())

	/* Send M_ADR frame with system address */
	mailer.stream.WriteAddress(mailer.GetAddr())

}

func (self *MailerStateWaitConn) Process(mailer *Mailer) IMailerState {

	select {
	case <-mailer.stream.InFrameReady: // TODO - replace on connection ready channel ...
		self.processWelcome(mailer)
		return NewMailerStateAdditionalStep()

	case <-time.After(15 * time.Second):
		return NewMailerStateEnd()
	}

}
