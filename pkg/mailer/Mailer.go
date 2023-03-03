package mailer

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/vit1251/golden/pkg/mailer/auth"
	stream2 "github.com/vit1251/golden/pkg/mailer/stream"
	"github.com/vit1251/golden/pkg/mailer/util"
	"github.com/vit1251/golden/pkg/queue"
	"github.com/vit1251/golden/pkg/registry"
	"log"
	"os"
	"strings"
	"sync"
	"time"
	"reflect"
	"runtime"
)

type Mailer struct {
	registry *registry.Container /* ???                         */

	activeState mailerStateFn /* Mailer state                */

	rxState RxState /* RX FSM                      */
	txState TxState /* TX FSM                      */

	stream *stream2.MailerStream /* ???                         */

	reader *bufio.Reader /* RX network stream           */
	writer *bufio.Writer /* TX network stream           */

	connectionCount int            /* Active session count        */
	wait            sync.WaitGroup /* Sync                        */

	addr         string  /* Network address             */
	secret       string  /* Secret password             */
	ServerAddr   string  /* Server IPv4 or FQDN address */

	inboundDirectory  string /* ???                         */
	outboundDirectory string /* ???                         */

	respAuthorization string /* ???                         */

	recvStream *os.File /* Stream using in Rx routines */
	sendStream *os.File /* Stream using in Tx routines */

	readSize  int64 /* Size incoming download      */
	writeSize int64 /* Size outgoing upload        */

	connComplete chan int /* ???                         */
	recvUnix     int      /* ???                         */

	sendName *queue.FileEntry /* Upload entry                */
	recvName *queue.FileEntry /* Download entry              */

	workInbound  string /* ???                         */
	workOutbound string /* ???                         */

	work string /* ???                         */

	InFileCount  int
	OutFileCount int

	workPath   string
	systemName string
	userName   string
	location   string

	rxRoutineResult ReceiveRoutineResult
	txRoutineResult TransmitRoutineResult

	outboundQueue []queue.FileEntry
	inboundQueue  []queue.FileEntry

	queue *util.TheQueue /* TheQueue      */

	pendingFiles util.Directory
	chunk        []byte
	report       *MailerReport /* Mailer report             */
}

func NewMailer(r *registry.Container) *Mailer {
	m := new(Mailer)

	m.connComplete = make(chan int)
	m.registry = r
	m.queue = util.NewTheQueue()
	m.connectionCount = 0

	return m
}

func (self *Mailer) GetReport() *MailerReport {
	return self.report
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

func (self *Mailer) Start() (error, *MailerReport) {

	if self.connectionCount > 0 {
		return fmt.Errorf("fido session alredy in progress"), nil
	}

	/* Initialize new report */
	self.report = NewMailerReport()

	/* Add wait */
	self.wait.Add(1)

	/* Play! */
	go self.run()

	return nil, self.report
}

func (self *Mailer) IsTransmitting() bool {
	return self.sendName != nil
}

func (self *Mailer) IsReceiving() bool {
	return self.recvName != nil
}

func makeFuncName(i interface{}) string {
	wideName := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	parts := strings.Split(wideName, ".")
	ourName := parts[len(parts) - 1]
	return ourName
}

func (self *Mailer) run() {

	self.report.SetSessionStart(time.Now())

	/* Reset active state */
	self.activeState = mailerStateStart

	/* Start processing */
	log.Printf("Start mailer routine")
	for {
		log.Printf("mailer: state = %s", makeFuncName(self.activeState))
		newState := self.activeState(self)
		log.Printf("mailer: change state: %s -> %s", makeFuncName(self.activeState), makeFuncName(newState))
		self.activeState = newState
		/* Stop processing when done */
		if newState == nil {
			log.Printf("mailer: Reach Exit state")
			break
		}
	}
	log.Printf("Stop mailer routine")

	/* Update report */
	self.report.SetSessionStop(time.Now())

	/* Close connection */
	self.wait.Done()

	/* Remove counter */
	self.connectionCount = self.connectionCount - 1

}

func (self *Mailer) Wait() *MailerReport {

	/* Wait session complete */
	self.wait.Wait()

	/* Dump report */
	self.report.Dump()

	return self.report
}

func (self *Mailer) SetServerAddr(addr string) {

	if !strings.Contains(addr, ":") {
		defaultPort := 24554
		addr = fmt.Sprintf("%s:%d", addr, defaultPort)
	}

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

func (self *Mailer) AddOutbound(path queue.FileEntry) {
	self.outboundQueue = append(self.outboundQueue, path)
}

func (self *Mailer) createAuthorization(chData []byte) string {
	a := auth.NewAuthorizer()
	a.SetChallengeData(chData)
	a.SetSecret([]byte(self.secret))
	responseDigest, err := a.CalculateDigest()
	if err != nil {
		panic(err)
	}
	password := fmt.Sprintf("%s-%s-%s", "CRAM", "MD5", responseDigest)
	return password
}

func (self *Mailer) processNulOptFrame(rawOptions []byte) {

	log.Printf("Mailer: Remote server option: %s", rawOptions)

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

	if len(values) == 2 {

		key := values[0]
		value := values[1]

		log.Printf("Mailer: Remote side M_NUL packet: name = %s value = %s", key, value)

		if bytes.Equal(key, []byte("OPT")) {
			self.processNulOptFrame(value)
		}

	} else {
		log.Printf("Mailer: Remote side M_NUL parse error: packet = %s", packet)
	}

}

func (self *Mailer) GetWork() string {
	return self.work
}

func (self *Mailer) IsReceiveName(name string) bool {
	if self.recvName != nil {
		if name == self.recvName.Name {
			return true
		}
	}
	return false
}

func (self *Mailer) IsTransmitName(name string) bool {
	if self.sendName != nil {
		if name == self.sendName.Name {
			return true
		}
	}
	return false
}
