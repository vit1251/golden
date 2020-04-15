package mailer

import (
	"log"
	"os"
	"path"
)

type MailerStateTransfer struct {
	MailerState
}

func NewMailerStateTransfer() *MailerStateTransfer {
	mst := new(MailerStateTransfer)
	return mst
}

func (self *MailerStateTransfer) String() string {
	return "MailerStateTransfer"
}

func (self *MailerStateTransfer) Process(mailer *Mailer) IMailerState {

	mo := NewMailerOutbound(mailer.SetupManager)
	items, err2 := mo.GetItems()
	if err2 != nil {
		panic(err2)
	}

	for _, item := range items {

		/* Transmit packet */
		err1 := mailer.Transmit(*item)
		if err1 != nil {
			log.Printf("Unable transmit %s file: err = %+v", item.Name, err1)
			break
		}

		/* Complete routine */
		newName := path.Join(mailer.TempOutbound, item.Name)
		os.Rename(item.AbsolutePath, newName)

	}

	log.Printf("Sent complete!")

	return NewMailerStateCloseConnection()
}
