package mailer

import (
	"log"
	"net"
	"bufio"
	"time"
	"fmt"
)

type Mailer struct {
	conn             net.Conn
	reader          *bufio.Reader
	writer          *bufio.Writer
	sessionComplete  chan bool
}

func NewMailer() (*Mailer) {
	m := new(Mailer)
	return m
}

func (self *Mailer) closeSession() {
    self.conn.Close()
}

func (self *Mailer) processRX() {
	for {
		if packet, err1 := self.processIncomingPacket(); err1 != nil {
			log.Printf("err = %v", err1)
			break
		} else {
			log.Printf("packet = %q", packet)
			self.processPacket(packet)
		}
	}
}

func (self *Mailer) openSession(host string) (error) {

	if conn, err := net.DialTimeout("tcp", host, time.Millisecond*1200); err != nil {
	} else {
		self.conn = conn

		self.reader = bufio.NewReader(conn)
		self.writer = bufio.NewWriter(conn)

		go self.processRX()
//		go self.processTX()

	}

	return nil
}

func (self *Mailer) writeTrafic(mail int, data int) {
	raw := fmt.Sprintf("TRF %d %d", mail, data)
	self.writeComment(raw)
}

func (self *Mailer) Check() {

	err := self.openSession("f24.n5023.z2.binkp.net:24554")
//	err := self.openSession("127.0.0.1:24554")
	if err != nil {
		log.Fatal(err)
	}
	defer self.closeSession()

	/* Start processing */
	self.writeInfo("SYS", "Vitold Station")
	self.writeInfo("ZYZ", "Vitold Sedyshev")
	self.writeInfo("LOC", "Saint-Petersburg, Russia")
//	self.writeComment("SYS Vitold Station")
//	self.writeComment("ZYZ Vitold Sedyshev")
//	self.writeComment("LOC Saint-Petersburg, Russia")
	self.writeComment("NDL 115200,TCP,BINKP")
	self.writeComment("TIME Sat, 16  Nov 2019 04:29:35 +0300")
	self.writeComment("OS Debian/10.0.0")
	self.writeComment("SIP vit1251@sipnet.ru")
	self.writeComment("XMPP vit1251@jabber.org")
	self.writeComment("VER GoldenMailer/1.0.0/Linux binkp/1.0")
	self.writeAddress("2:5023/24.3752@fidonet")
	self.writeComment("OPT DEBUG")
	self.writeTrafic(10234, 0)

	<- self.sessionComplete

}
