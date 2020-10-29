package stream

import (
	"bufio"
	"log"
	"net"
	"sync"
	"time"
)

type MailerStream struct {

	conn          net.Conn

	reader        *bufio.Reader
	writer        *bufio.Writer

	wait          sync.WaitGroup

	InDataFrames  chan Frame
	OutDataFrames chan Frame

}

func NewMailerStream() *MailerStream {
	stream := new(MailerStream)
	return stream
}

func (self *MailerStream) OpenSession(system string) error {

	conn, err1 := net.DialTimeout("tcp", system, time.Millisecond*1000)
	if err1 != nil {
		log.Printf("MailerStream: Fail to start session: err = %+v", err1)
		return err1
	}

	self.InDataFrames = make(chan Frame)
	self.OutDataFrames = make(chan Frame)

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

	log.Printf("MailerStream: Close")
	return self.conn.Close()

}
