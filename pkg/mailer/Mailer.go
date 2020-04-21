package mailer

import (
	"bufio"
	"fmt"
	"github.com/vit1251/golden/pkg/setup"
	"log"
	"net"
	"os"
	"time"
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
	activeState       IMailerState           /* Mailer state                */
	sessionSetupState SessionSetupStageState /* Session setup state         */
	transferState     FileTransferStageState /* Mailer state                */
	rxState           RxState                /*                             */
	txState           TxState                /*                             */
	inDataFrames      chan Frame             /* RX processing Frame routine */
	outDataFrames     chan Frame             /* TX processing Frame routine */
	conn              net.Conn               /* Network address             */
	reader            *bufio.Reader          /* Network address             */
	writer            *bufio.Writer          /* Network address             */
	sessionComplete   chan bool              /* Network address             */
	addr              string                 /* Network address             */
	secret            string                 /* Secret password             */
	ServerAddr        string                 /*  */
	inboundDirectory  string                 /*   */
	outboundDirectory string                 /*   */
	respAuthorization string                 /*   */
	outStream         *os.File               /*    */
	writeSize         int                    /*    */
	size              int                    /*  */
	connComplete      chan int               /*  */
	recvUnix          int                    /*  */
	recvName          string                 /*   */
	TempOutbound      string                 /*  */
	SetupManager      *setup.ConfigManager   /*   */
	InFileCount       int
	OutFileCount      int
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

	/* Start state */
	self.activeState = NewMailerStateCreateConnection()

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
	close(self.connComplete)
}

func (self *Mailer) Wait() {
	<- self.connComplete
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
	return "1.2.11"
}
