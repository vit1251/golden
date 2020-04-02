package mailer

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/vit1251/golden/pkg/setup"
	"io"
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
	inboundDirectory  string /* */
	outboundDirectory string
	respAuthorization string
	outStream         *os.File
	writeSize         int
	size              int
	connComplete      chan int
	recvUnix          int
	recvName          string
	TempOutbound      string
	SetupManager     *setup.ConfigManager
}

func NewMailer(sm *setup.ConfigManager) (*Mailer) {
	m := new(Mailer)
	m.inDataFrames = make(chan Frame)
	m.outDataFrames = make(chan Frame)
	m.connComplete = make(chan int)
	m.SetupManager = sm
	return m
}

func (self *Mailer) closeSession() {
    self.conn.Close()
}

func (self *Mailer) writeTrafic(mail int, data int) {
	raw := fmt.Sprintf("TRF %d %d", mail, data)
	self.writeComment(raw)
}

func (self *Mailer) SetTempOutbound(TempOutbound string) {
	self.TempOutbound = TempOutbound
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

func (self *Mailer) processState() error {
	if self.protocolState == SessionSetupState {
		return self.processSessionSetupState()
	} else if self.protocolState == FileTransferState {
		return self.processFileTransferState()
	}
	return errors.New("wrong mailer state")
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

func (self *Mailer) Transmit(i Item) error {

	/* Open stream */
	stream, err1 := os.Open(i.AbsolutePath)
	if err1 != nil {
		return err1
	}
	defer stream.Close()

	/* Some status */
	streamInfo, err2 := stream.Stat()
	if err2 != nil {
		return err2
	}

	/* Transmit header */
	// p0018ea8.WE0 39678 1579714843 0
	streamSize := streamInfo.Size()
	streamTime := streamInfo.ModTime().Unix()
	fileStat := fmt.Sprintf("%s %d %d %d", i.Name, streamSize, streamTime, 0)
	log.Printf("TX %s", fileStat)
	self.writeHeader(fileStat)

	/* Transmit chunk */
	var outSize int = int(streamSize)
	for {

		/* Calculate transmit chunk */
		chunkSize := Min(outSize, 4096)
		chunk := make([]byte, chunkSize)

		/* Read */
		_, err3 := io.ReadFull(stream, chunk)
		if err3 != nil {
			return err3
		}

		/* Transmit chunk */
		self.writeData(chunk)

		/* Update TX size */
		outSize -= chunkSize
		if outSize == 0 {
			log.Printf("Transmit complete!")
			break
		}
	}

	/* Check error */

	return nil
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (self *Mailer) writeData(chunk []byte) error {
	log.Printf("TX data chunk %d", len(chunk))

	/* Transmit data frame */
	newFrame := Frame{
		Command: false,
		DataFrame: DataFrame{
			Body: chunk,
		},
	}
	self.outDataFrames <- newFrame

	return nil
}

func (self *Mailer) writeHeader(stat string) error {
	log.Printf("TX stream header %s", stat)

	newFrame := Frame{
		Command: true,
		CommandFrame: CommandFrame{
			CommandID: M_FILE,
			Body: []byte(stat),
		},
	}
	self.outDataFrames <- newFrame

	return nil
}

func (self *Mailer) GetVersion() string {
	return "1.2.10"
}
