package stream

import (
	"bufio"
	"log"
	"net"
	"sync"
	"time"
)

type MailerStream struct {
	conn net.Conn

	reader *bufio.Reader
	writer *bufio.Writer

	wait sync.WaitGroup

	InDataFrames  chan Frame
	OutDataFrames chan Frame

	readReady     bool
	writeReady    bool

}

func NewMailerStream() *MailerStream {
	stream := new(MailerStream)
	return stream
}

func (self *MailerStream) IsReadReady() bool {
	return self.readReady
}

func (self *MailerStream) IsWriteReady() bool {
	return self.writeReady
}

func (self *MailerStream) OpenSession(system string) error {

	conn, err1 := net.DialTimeout("tcp", system, time.Millisecond*1000)
	if err1 != nil {
		log.Printf("MailerStream: Fail to open and connect socket: err = %+v", err1)
		return err1
	}

	log.Printf("MailerStream: Socket open")

	self.InDataFrames = make(chan Frame)
	self.OutDataFrames = make(chan Frame)

	self.conn = conn

	/* Create reader and writer */
	self.reader = bufio.NewReader(conn)
	self.writer = bufio.NewWriter(conn)

	/* Initial reading value */
	self.readReady = false
	self.writeReady = true

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
	self.OutDataFrames <- newFrame
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
	self.OutDataFrames <- newFrame

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
	self.OutDataFrames <- newFrame

	return nil
}

func (self *MailerStream) CloseSession() error {

	log.Printf("MailerSession: Closing session.")

	/* Close stream */
	close(self.OutDataFrames)

	log.Printf("MailerSession: Wait socket reader and writer")
	self.wait.Wait()

	log.Printf("MailerStream: Socket close.")
	return self.conn.Close()

}

func (self *MailerStream) RemReadReady() {
	self.readReady = false
}

func (self *MailerStream) GetFrame() Frame {
	nextFrame := <- self.InDataFrames
	self.RemReadReady()
	return nextFrame
}
