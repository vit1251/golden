package mailer

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

type CommandFrame struct {
	CommandID CommandID     /* Command ID  */
	Body []byte
}

type DataFrame struct {
	Body []byte
}

type Frame struct {
	Command bool
	DataFrame
	CommandFrame
}

type Mailer struct {
	protocolState     ProtocolState          /* Protocol state      */
	sessionSetupState SessionSetupStageState /* Session setup state */
	transferState     FileTransferStageState /* Mailer state         */
	rxState           RxState                /* */
	txState           TxState                /* */
	inDataFrames      chan Frame             /* RX processing Frame routine */
	outDataFrames     chan Frame             /* TX processing Frame routine */
	conn              net.Conn               /* Network address */
	reader            *bufio.Reader          /* Network address */
	writer            *bufio.Writer          /* Network address */
	sessionComplete   chan bool              /* Network address */
	addr              string                 /* Network address */
	secret            string                 /* Secret password */
	ServerAddr        string
	inboundDirectory  string                 /* */
	outboundDirectory string
	respAuthorization string
	outStream         *os.File
	writeSize         int
	size              int
	connComplete      chan int
	recvUnix          int
	recvName		  string
}

func NewMailer() (*Mailer) {
	m := new(Mailer)
	m.inDataFrames = make(chan Frame)
	m.outDataFrames = make(chan Frame)
	m.connComplete = make(chan int)
	return m
}

func (self *Mailer) closeSession() {
    self.conn.Close()
}

func (self *Mailer) writeTrafic(mail int, data int) {
	raw := fmt.Sprintf("TRF %d %d", mail, data)
	self.writeComment(raw)
}

func (self *Mailer) SetAddr(addr string) {
	self.addr = addr
}

func (self *Mailer) SetSecret(secret string) {
	self.secret = secret
}

func (self *Mailer) Start() {
	/* Setup start state */
	self.SetProtocolState(SessionSetupState)
	self.SetSessionSetupState(SessionSetupConnInitState)
	/* Process state */
	go self.run()
}

func (self *Mailer) processState() {
	if (self.protocolState == SessionSetupState) {
		self.processSessionSetupState()
	} else if (self.protocolState == FileTransferState) {
		self.processFileTransferState()
	} else {
		log.Panicf("Wrong mailer state: state = %v", self.protocolState)
	}
}

func (self *Mailer) run() {
	for {
		log.Printf("Process state: protocolState = %v", self.protocolState.String())
		self.processState()
	}
}

func (self *Mailer) Wait() {
	<- self.connComplete
}

func (self *Mailer) SetProtocolState(state ProtocolState) {
	log.Printf("Set protocol state: %s", state.String())
	self.protocolState = state
}

func (self *Mailer) SetSessionSetupState(state SessionSetupStageState) {
	log.Printf("Set session setup state: %s", state.String())
	self.sessionSetupState = state
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
