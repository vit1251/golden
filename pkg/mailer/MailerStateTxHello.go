package mailer

import (
	cmn "github.com/vit1251/golden/pkg/common"
)

type MailerStateTxHello struct {
	MailerState
	IMailerState
}

func NewMailerStateTxHello() *MailerStateTxHello {
	msth := new(MailerStateTxHello)
	return msth
}

func (self *MailerStateTxHello) String() string {
	return "MailerStateTxHello"
}

func (self *MailerStateTxHello) Process(mailer *Mailer) IMailerState {

	/* System name */
	systemName := mailer.GetSystemName()
	if systemName != "" {
		mailer.WriteInfo("SYS", systemName)
	}

	/* User name */
	username := mailer.GetUserName()
	if username != "" {
		mailer.WriteInfo("ZYZ", username)
	}

	location := mailer.GetLocation()
	if location != "" {
		mailer.WriteInfo("LOC", location)
	}

	mailer.WriteInfo("NDL", "115200,TCP,BINKP")
	mailer.WriteInfo("TIME", cmn.GetTime())
	mailer.WriteInfo("OS", cmn.GetPlatform())
	mailer.WriteVersion()
	addr := mailer.GetAddr()
	mailer.WriteAddress(addr)

	return NewMailerStateRxHello()
}
