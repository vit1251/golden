package mailer

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/vit1251/golden/pkg/mailer/auth"
	"github.com/vit1251/golden/pkg/mailer/cache"
	stream2 "github.com/vit1251/golden/pkg/mailer/stream"
	"github.com/vit1251/golden/pkg/mailer/util"
	"github.com/vit1251/golden/pkg/setup"
	"log"
	"os"
	"sync"
	"time"
)

type Mailer struct {
	activeState IMailerState /* Mailer state                */

	rxState RxState /*                             */
	txState TxState /*                             */

	stream *stream2.MailerStream

	reader *bufio.Reader /* Network address             */
	writer *bufio.Writer /* Network address             */

	wait sync.WaitGroup /* Add wait */

	addr       string /* Network address             */
	secret     string /* Secret password             */
	ServerAddr string /*  */

	inboundDirectory  string /*   */
	outboundDirectory string /*   */

	respAuthorization string /*   */

	recvStream *os.File /* Stream using in Rx routines */
	sendStream *os.File /* Stream using in Tx routines */

	readSize  int64 /* Size incoming download */
	writeSize int64 /* Size outgoing upload   */

	connComplete chan int /*  */
	recvUnix     int      /*  */

	sendName *cache.FileEntry   /* Upload entry     */
	recvName *cache.FileEntry    /* Download entry   */

	workInbound  string /*  */
	workOutbound string /*  */

	work         string

	SetupManager *setup.ConfigManager /*   */

	InFileCount  int
	OutFileCount int

	//
	workPath   string
	systemName string
	userName   string
	location   string

	outboundQueue []cache.FileEntry
	inboundQueue  []cache.FileEntry

	queue          *util.TheQueue       /* TheQueue      */

	pendingFiles    util.Directory

}

func NewMailer(sm *setup.ConfigManager) *Mailer {
	m := new(Mailer)

	m.connComplete = make(chan int)
	m.SetupManager = sm
	m.queue = util.NewTheQueue()

	return m
}

func (self *Mailer) writeTrafic(mail int, data int) {
	raw := fmt.Sprintf("TRF %d %d", mail, data)
	self.stream.WriteComment(raw)
}

func (self *Mailer) SetTempOutbound(workOutbound string) {
	self.workOutbound = workOutbound
}

func (self *Mailer) SetAddr(addr string) {
	self.addr = addr
}

func (self *Mailer) SetSecret(secret string) {
	self.secret = secret
}

func (self *Mailer) Start() {

	/* Start state */
	self.activeState = NewMailerStateStart()

	/* Add wait */
	self.wait.Add(1)

	/* Play! */
	go self.run()

}

func (self *Mailer) IsTransmitting() bool {
	return self.sendName != nil
}

func (self *Mailer) IsReceiving() bool {
	return self.recvName != nil
}

func (self *Mailer) run() {

	mailerStart := time.Now()
	log.Printf("Start mailer routine")
	for {
		log.Printf("mailer: process state %s", self.activeState)
		newState := self.activeState.Process(self)
		log.Printf("mailer: chage state: %s -> %s", self.activeState, newState)
		self.activeState = newState
		/* Stop processing when done */
		if newState == nil {
			log.Printf("mailer: No more processing")
			break
		}
	}
	log.Printf("Stop mailer routine")
	elapsed := time.Since(mailerStart)
	log.Printf("Mailer session: %+v", elapsed)

	/* Close connection */
	self.wait.Done()

}

func (self *Mailer) Wait() {
	self.wait.Wait()
}

func (self *Mailer) SetServerAddr(addr string) {
	self.ServerAddr = addr
}

func (self *Mailer) SetInboundDirectory(inb string) {
	self.inboundDirectory = inb
}

func (self *Mailer) SetOutboundDirectory(outb string) {
	self.outboundDirectory = outb
}

func (self *Mailer) SetTempInbound(workInbound string) {
	self.workInbound = workInbound
}

func (self *Mailer) SetTemp(work string) {
	self.work = work
}

func (self *Mailer) GetWorkOutbound() string {
	return self.workOutbound
}

func (self *Mailer) GetAddr() string {
	return self.addr
}

func (self *Mailer) GetSystemName() string {
	return self.systemName
}

func (self *Mailer) GetUserName() string {
	return self.userName
}

func (self *Mailer) GetLocation() string {
	return self.location
}

func (self *Mailer) SetLocation(location string) {
	self.location = location
}

func (self *Mailer) SetUserName(name string) {
	self.userName = name
}

func (self *Mailer) SetStationName(name string) {
	self.systemName = name
}

func (self *Mailer) AddOutbound(path cache.FileEntry) {
	self.outboundQueue = append(self.outboundQueue, path)
}

func (self *Mailer) createAuthorization(chData []byte) (string) {
	a := auth.NewAuthorizer()
	a.SetChallengeData(string(chData))
	a.SetSecret(self.secret)
	responseDigest, err := a.CalculateDigest()
	if err != nil {
		panic(err)
	}
	password := fmt.Sprintf("%s-%s-%s", "CRAM", "MD5", responseDigest)
	return password
}

func (self *Mailer) processNulOptFrame(rawOptions []byte)  {

	log.Printf("Remote server option: %s", rawOptions)

	/* Split options */
	options := bytes.Fields(rawOptions)
	for _, option := range options {
		if bytes.HasPrefix(option, []byte("CRAM-")) {
			parts := bytes.SplitN(option, []byte("-"), 3)
			authScheme := parts[1]
			if bytes.Equal(authScheme, []byte("MD5")) {
				authDigest := parts[2]
				log.Printf("Use %s as digest", authDigest)
				self.respAuthorization = self.createAuthorization(authDigest)
			} else {
				log.Panicf("Wrong mechanism: authScheme = %s", authScheme)
			}
		}
	}

}

func (self *Mailer) processNulFrame(nextFrame stream2.Frame) {

	packet := nextFrame.CommandFrame.Body
	values := bytes.SplitN(packet, []byte(" "), 2)

	log.Printf("Remote side M_NUL with values: values = %+v", values)

	if len(values) == 2 {

		key := values[0]
		value := values[1]

		if bytes.Equal(key, []byte("OPT")) {
			self.processNulOptFrame(value)
		}

	} else {
		log.Printf("Remote side M_NUL parse error")
	}

}

func (self *Mailer) GetWork() string {
	return self.work
}
