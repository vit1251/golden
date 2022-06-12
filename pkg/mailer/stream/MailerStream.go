package stream

import (
	"context"
	"bufio"
	"log"
	"net"
	"sync"
	"time"
)

type MailerStream struct {

	dialer net.Dialer
	dialerCancel context.CancelFunc

	conn net.Conn

	reader *bufio.Reader
	writer *bufio.Writer

	wait sync.WaitGroup

	InFrameReady  chan interface{}
	OutFrameReady chan interface{}

	InFrame  chan Frame
	OutFrame chan Frame

}

func NewMailerStream() *MailerStream {
	stream := new(MailerStream)
	return stream
}

func (self *MailerStream) OpenSession(remoteSystem string) error {

	/* Step 1. Create context */
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	self.dialerCancel = cancel

	/* Step 2. Start new TCP connection */
	conn, err1 := self.dialer.DialContext(ctx, "tcp", remoteSystem)
	if err1 != nil {
		log.Printf("MailerStream: Fail to open and connect socket: err = %+v", err1)
		return err1
	}
	log.Printf("MailerStream: Socket open")

	/* Step 3. Initialize stream channels */
	self.InFrameReady = make(chan interface{})
	self.OutFrameReady = make(chan interface{})

	self.InFrame = make(chan Frame)
	self.OutFrame = make(chan Frame)

	self.conn = conn

	/* Create reader and writer */
	self.reader = bufio.NewReader(conn)
	self.writer = bufio.NewWriter(conn)

	/* Register wait */
	self.wait.Add(2)

	/* Start frame processing */
	go self.processRX()
	go self.processTX()

	return nil
}

func (self *MailerStream) WriteCommandPacket(commandID CommandID, msgBody []byte) error {
	log.Printf("writeCommandPacket: commandID = %q body = %s", commandID, msgBody)
	newFrame := Frame{
		Command: true,
		CommandFrame: CommandFrame{
			CommandID: commandID,
			Body: msgBody,
		},
	}
	self.OutFrame <- newFrame
	return nil
}

func (self *MailerStream) WriteData(chunk []byte) error {

	log.Printf("MailerStream: TX stream: Write data frame: size = %d", len(chunk))

	/* Transmit data frame */
	newFrame := Frame{
		Command: false,
		DataFrame: DataFrame{
			Body: chunk,
		},
	}
	self.OutFrame <- newFrame

	return nil
}

func (self *MailerStream) WriteHeader(stat string) error {

	log.Printf("MailerStream: Uploading: stat = %s", stat)

	newFrame := Frame{
		Command: true,
		CommandFrame: CommandFrame{
			CommandID: M_FILE,
			Body:      []byte(stat),
		},
	}
	self.OutFrame <- newFrame

	return nil
}

func (self *MailerStream) CloseSession() error {

	log.Printf("MailerSession: Closing session.")

	/* Close RX stream */
	self.conn.SetReadDeadline(time.Now().Add(1 * time.Second))

	/* Close */
	close(self.OutFrame)

	log.Printf("MailerSession: Wait socket reader and writer")
	self.wait.Wait()

	log.Printf("MailerStream: Socket close.")
	return self.conn.Close()

}
