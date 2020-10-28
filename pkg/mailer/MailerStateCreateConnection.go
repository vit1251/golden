package mailer

import (
	"bufio"
	"log"
	"net"
	"time"
)

type MailerStateCreateConnection struct {
	MailerState
}

func NewMailerStateCreateConnection() *MailerStateCreateConnection {
	msc := new(MailerStateCreateConnection)
	return msc
}

func (self *MailerStateCreateConnection) String() string {
	return "MailerStateCreateConnection"
}

func (self *MailerStateCreateConnection) Process(mailer *Mailer) IMailerState {

	conn, err1 := net.DialTimeout("tcp", mailer.ServerAddr, time.Millisecond*1200)
	if err1 != nil {
		log.Printf("Fail on create network connection: err = %+v", err1)
		return NewMailerStateCloseConnection()
	}

	mailer.conn = conn

	/* Create reader and writer */
	mailer.reader = bufio.NewReader(conn)
	mailer.writer = bufio.NewWriter(conn)

	/* Register wait */
	mailer.wait.Add(2)

	/* Start frame processing */
	go mailer.processRX()
	go mailer.processTX()

	return NewMailerStateTxHello()
}
