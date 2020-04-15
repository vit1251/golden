package mailer

import (
	"fmt"
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

	//mailer.writeInfo("SYS", "Vitold Station")
	//mailer.writeInfo("ZYZ", "Vitold Sedyshev")
	//mailer.writeInfo("LOC", "Saint-Petersburg, Russia")
	mailer.writeInfo("NDL", "115200,TCP,BINKP")
	mailer.writeInfo("TIME", mailer.GetTime())
	mailer.writeInfo("OS", mailer.GetPlatform())
	appName := "GoldenMailer"
	appVersion := mailer.GetVersion()
	mailer.writeInfo("VER", fmt.Sprintf("%s/%s", appName, appVersion))
	mailer.writeAddress(mailer.addr)

	return NewMailerStateRxHello()
}
