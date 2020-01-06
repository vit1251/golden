package packet

import (
	"log"
	"bytes"
	"io"
)

type ParserState int

const (
	MBP_STATE_KLUDGE ParserState = 1
	MBP_STATE_BODY   ParserState = 2
)

type KludgeState int

const (
	MBPK_STATE_OPTIONAL KludgeState = 0
	MBPK_STATE_NAME     KludgeState = 1
	MBPK_STATE_VALUE    KludgeState = 2
)

type MessageBodyParser struct {
	state         ParserState      /* Parser state */
	kludgeState   KludgeState
	kludgeName  []byte
	kludgeValue []byte
	message     []byte
	newLine       bool
	lineCount     int
	result       *MessageBody
}

func NewMessageBodyParser() (*MessageBodyParser, error) {
	mp := new(MessageBodyParser)
	mp.result = NewMessageBody()
	mp.state = MBP_STATE_KLUDGE
	mp.kludgeState = MBPK_STATE_NAME
	return mp, nil
}

func (self *MessageBodyParser) processKludge() {
	//
	log.Printf("Meet kludge: name = %v value = %v", self.kludgeName, self.kludgeValue)
	//
	var name string = string(self.kludgeName)
	var value string = string(self.kludgeValue)
	//
	if name == "AREA" {
		self.result.SetArea(value)
	} else {
		self.result.AddKludge(name, value)
	}
	//
	self.kludgeName = nil // make([]byte, 0)
	self.kludgeValue = nil // make([]byte, 0)
	//
}

func (self *MessageBodyParser) processKludgeByte(value byte) {
	if self.kludgeState == MBPK_STATE_OPTIONAL {
		if value == '\x01' {
			self.kludgeState = MBPK_STATE_NAME
		} else {
			/* Unread byte */
			self.message = append(self.message, value)
			/* Debug message */
			self.state = MBP_STATE_BODY
		}
	} else if self.kludgeState == MBPK_STATE_NAME {
		if value == '\x0A' || value == '\x0D' {
			self.processKludge()
			self.kludgeState = MBPK_STATE_OPTIONAL
		} else if value == ':' || value == ' ' {
			self.kludgeState = MBPK_STATE_VALUE
		} else {
			self.kludgeName = append(self.kludgeName, value)
		}
	} else if self.kludgeState == MBPK_STATE_VALUE {
		if value == '\x0A' || value == '\x0D' {
			/* Reset kludge cache */
			self.processKludge()
			/* Set new state */
			self.kludgeState = MBPK_STATE_OPTIONAL
		} else {
			self.kludgeValue = append(self.kludgeValue, value)
		}
	}
}

func (self *MessageBodyParser) processMessageByte(value byte) {
	self.message = append(self.message, value)
}

func (self *MessageBodyParser) Parse(msg []byte) (*MessageBody, error) {

	stream := bytes.NewBuffer(msg)

	for {
		value, err1 := stream.ReadByte()
		if err1 != nil {
			if err1 == io.EOF {
				log.Printf("msg = %q", self.message)
				self.result.SetRaw(self.message)
				break
			} else {
				log.Fatal(err1)
			}
		}
		if self.state == MBP_STATE_KLUDGE {
			self.processKludgeByte(value)
		} else if self.state == MBP_STATE_BODY {
			self.processMessageByte(value)
		} else {
		    // eror
		}
	}

	/* Message body */
	return self.result, nil
}
