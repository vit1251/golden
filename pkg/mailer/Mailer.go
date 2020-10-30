package mailer

import (
	"bufio"
	"fmt"
	"github.com/vit1251/golden/pkg/mailer/cache"
	stream2 "github.com/vit1251/golden/pkg/mailer/stream"
	"github.com/vit1251/golden/pkg/setup"
	"log"
	"os"
	"sync"
	"time"
)

type Mailer struct {
	activeState IMailerState /* Mailer state                */

	transferState FileTransferStageState /* Mailer state                */

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

}

func NewMailer(sm *setup.ConfigManager) *Mailer {
	m := new(Mailer)

	m.connComplete = make(chan int)
	m.SetupManager = sm

	return m
}

func (self *Mailer) closeSession() {
    self.stream.CloseSession()
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

func (self *Mailer) SetFileTransferState(state FileTransferStageState) {
	self.transferState = state
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
