package mailer

import (
	"fmt"
	cmn "github.com/vit1251/golden/pkg/common"
	"time"
)

type MailerStateWaitConnAction struct {
	MailerState
	IMailerState
}

func NewMailerWaitConnAction() *MailerStateWaitConnAction {
	msth := new(MailerStateWaitConnAction)
	return msth
}

func (self *MailerStateWaitConnAction) String() string {
	return "MailerStateWaitConnAction"
}

func (self *MailerStateWaitConnAction) makeSystemTime() string {
	now := time.Now().Format(time.RFC822)
	return now
}

func (self *MailerStateWaitConnAction) makeOperationSystemName() string {
	return cmn.GetPlatform()
}

func (self *MailerStateWaitConnAction) makeVersionString() string {
	appName := "GoldenMailer"
	appVersion := cmn.GetVersion()
	protocolVersion := "binkp/1.0"
	return fmt.Sprintf("%s/%s %s", appName, appVersion, protocolVersion)
}

func (self *MailerStateWaitConnAction) Process(mailer *Mailer) IMailerState {

	/* Send M_NUL frames with system info (optional) */
	if username := mailer.GetUserName(); username != "" {
		mailer.stream.WriteInfo("ZYZ", username)
	}
	if location := mailer.GetLocation(); location != "" {
		mailer.stream.WriteInfo("LOC", location)
	}
	mailer.stream.WriteInfo("TIME", self.makeSystemTime())
	mailer.stream.WriteInfo("OS", self.makeOperationSystemName())
	mailer.stream.WriteInfo("VER", self.makeVersionString())

	/* Send M_ADR frame with system address */
	mailer.stream.WriteAddress(mailer.GetAddr())

	return NewMailerStateAdditionalStepAction()

}
