package squish

import (
	"fmt"
	"github.com/vit1251/golden/pkg/utils"
	"os"
	"io"
	"bufio"
	"encoding/binary"
	"log"
	"errors"
	"golang.org/x/text/encoding/charmap"
	"encoding/hex"
	"crypto/md5"
)

type SquishMessage struct {
}

type SquishMessageBaseHeader struct {
	Size uint16            /* sizeof(SqshBaseT)                 */
	Reserved1 uint16       /* Reserved                          */
	MessageNumber uint32   /* Number of messages                */
	HighestMsg uint32      /* Highest msg in area               */
	Protmsgs uint32        /* Skip killing first x msgs in area */
	Highwatermark uint32   /* Relno (not Tagno) of HWM          */
	Nextmsgno uint32       /* Next message number to use        */
	Name [80]byte          /* Base name of SquishFile           */
	BeginFrame uint32      /* Offset of the first frame         */
	LastFrame uint32       // Offset to last frame in file
	FirstFreeFrame uint32  // Offset of first FREE frame in file
	LastFreeFrame uint32   // Offset of last free frame in file
	EndFrame uint32        // Pointer to end of file
	MaxMsgs uint32         // Max # of msgs to keep in area
	DaysToKeep uint16      // Max age of msgs in area (for packing util)
	FrameSize uint16       // sizeof(SqshFrmT)
	Reserved2 [124]byte    // Reserved by Squish for future use
}

const SQFRAMEID = 0xAFAE4453

type FrameType uint16

const (
    FrameTypeNormal FrameType = 0
    FrameTypeFree   FrameType = 1
    FrameTypeLZSS   FrameType = 2
    FrameTypeUpdate FrameType = 3
)

type SquishMessageBaseFrameHeader struct {
	ID uint32             // Must always equal SQFRAMEID
	NextFrame uint32      // Offset of next frame
	PrevFrame uint32      // Offset of previous frame
	FrameLength uint32    // Length of this frame
	MsgLength uint32      // Length of data in frame (hdr+ctl+txt)
	ControlLength uint32  // Length of control info
	FrameType FrameType   // Frm type (normal or free)
	Reserved uint16       // Reserved
}

const MAX_REPLY = 9

type UMSGID uint32

func (self UMSGID) GetMessageID() string {
	var result string = ""
	var value uint32 = uint32(self)
	result = fmt.Sprintf("%08x", value)
	return result
}

type SquishMessageBaseMessageHeader struct {
	Attr        uint32                   // Message attributes
	From        [36]byte                 // From
	To          [36]byte                 // To
	Subject     [72]byte                 // Subject
	Orig        NetAddress               // 
	Dest        NetAddress               //
	DateWritten FidoDate                 // When user wrote the msg (UTC)
	DateArrived FidoDate                 // When msg arrived on-line (UTC)
	UtcOffset   int16                    // Offset from UTC of message writer, in minutes.
	ReplyTo     UMSGID
	Replies     [MAX_REPLY]UMSGID
	UMSGID      UMSGID
	DateFTSC    [20]byte
}

type SquishMessageBase struct {
	messageCount   int
	Msgs         []SquishMessage
}

func (self *SquishMessageBase) readMessageBaseHeader(reader *bufio.Reader) (error) {
	var header SquishMessageBaseHeader
	size := binary.Size(header)
	log.Printf("size = %d", size)
	err := binary.Read(reader, binary.LittleEndian, &header)
	if err != nil {
		return err
	}
	var msg string
	msg += "SqusihMessageBase:\n"
	msg += fmt.Sprintf(" message_count = %d\n", header.MessageNumber)
	log.Printf(msg)
	//
	self.messageCount = int(header.MessageNumber)
	//
	return nil
}

func (self *SquishMessageBase) readMessageFrameHeader(reader *bufio.Reader) (*SquishMessageBaseFrameHeader, error) {
	//
	frameHeader := new(SquishMessageBaseFrameHeader)
	size := binary.Size(frameHeader)
	log.Printf("size = %d", size)
	err1 := binary.Read(reader, binary.LittleEndian, frameHeader)
	if err1 != nil {
		return nil, err1
	}
	if frameHeader.ID != SQFRAMEID {
		return nil, errors.New("Squish message base is corrup.")
	}
	//
	{
		var msg string
		msg = "SqusihMessageBaseFrameHeader\n"
		msg += fmt.Sprintf(".. FrameType = %d\n", frameHeader.FrameType)
		msg += fmt.Sprintf(".. NextFrame = %d\n", frameHeader.NextFrame)
		msg += fmt.Sprintf(".. PrevFrame = %d\n", frameHeader.PrevFrame)
		msg += fmt.Sprintf(".. FrameLength = %d\n", frameHeader.FrameLength)
		msg += fmt.Sprintf(".. MsgLength = %d\n", frameHeader.MsgLength)
		msg += fmt.Sprintf(".. ControlLength = %d\n", frameHeader.ControlLength)
		log.Printf(msg)
	}
	return frameHeader, nil
}

func (self *SquishMessageBase) readMessageLetterHeader(reader *bufio.Reader) (*SquishMessageBaseMessageHeader, error) {
	//
	msgHeader := new (SquishMessageBaseMessageHeader)
	size := binary.Size(msgHeader)
	log.Printf("size = %d", size)
	err2 := binary.Read(reader, binary.LittleEndian, msgHeader)
	if err2 != nil {
		return nil, err2
	}
	//
	{
		var msg string
		msg = "SqusihMessageBaseMessageHeader\n"
		msg += fmt.Sprintf(".. UMSGID = %X\n", msgHeader.UMSGID)
		msg += fmt.Sprintf(".. From = %s\n", string(msgHeader.From[:]))
		msg += fmt.Sprintf(".. To = %s\n", string(msgHeader.To[:]))
		msg += fmt.Sprintf(".. Subject = %s\n", string(msgHeader.Subject[:]))
		log.Printf(msg)
	}
	//
	return msgHeader, nil
}

type Message struct {
	Control []byte
	Body    []byte
}

func (self *SquishMessageBase) readMessageLetterBody(reader *bufio.Reader, frame *SquishMessageBaseFrameHeader) (*Message, error) {
	//
	msgSize := int(frame.FrameLength) - 238 // TODO - replace 238 on sizeof(MessageHeader)
	controlSize := int(frame.ControlLength)
	//
	bodySize := msgSize - controlSize
	// Read control
	log.Printf("controlSize = %d", controlSize)
	control := make([]byte, controlSize)
	_, err := io.ReadFull(reader, control)
	if err != nil {
		return nil, err
	}
	log.Printf("%s", hex.Dump(control))
	// Read body
	log.Printf("bodySize = %d", bodySize)
	body := make([]byte, bodySize)
	_, err1 := io.ReadFull(reader, body)
	if err1 != nil {
		return nil, err1
	}
	log.Printf("%s", hex.Dump(body))
	//
	res := new(Message)
	res.Control = control
	res.Body = body
	//
	return res, nil
}

func (self *SquishMessageBase) readMessage(reader *bufio.Reader) (*Header, error) {
	//
	frameHeader, err1 := self.readMessageFrameHeader(reader)
	if err1 != nil {
		return nil, err1
	}
	if frameHeader.FrameType == FrameTypeNormal {

		msgHeader, err2 := self.readMessageLetterHeader(reader)
		if err2 != nil {
			return nil, err2
		}
		//
		decoder1 := charmap.CodePage866.NewDecoder()
		newSubject := make([]byte, 512)
		_, _, err3 := decoder1.Transform(newSubject, msgHeader.Subject[:], false)
		if err3 != nil {
			panic(err3)
		}
		//
		msg, err4 := self.readMessageLetterBody( reader, frameHeader )
		if err4 != nil {
			panic(err4)
		}
		//
		decoder2 := charmap.CodePage866.NewDecoder()
		newBody := make([]byte, 2 * len(msg.Body) + 512) // TODO - additional buffer ...
		_, _, err5 := decoder2.Transform(newBody, msg.Body[:], false)
		if err5 != nil {
			panic(err5)
		}
		//
		res := new(Header)
		res.UMSGID = uint32(msgHeader.UMSGID)
		res.ID = msgHeader.UMSGID.GetMessageID()
		res.From = utils.MakeString(msgHeader.From[:])
		res.FromAddr = msgHeader.Orig.GetAddr()
		res.To = utils.MakeString(msgHeader.To[:])
		res.ToAddr = msgHeader.Dest.GetAddr()
		res.Subject = utils.MakeString(newSubject)
		res.DateWritten = msgHeader.DateWritten.GetDateTime()
		res.DateArrived = msgHeader.DateArrived.GetDateTime()
		res.Hash = fmt.Sprintf("%32x", md5.Sum(msg.Body[:]))
		res.Body = newBody
//		res.Control = msg.Control // TODO - parse kludge here
		//
		return res, nil

	} else {

		log.Printf("Unknown squish frame ( FrameType = %v ). Skip.", frameHeader.FrameType)
		var skipSize int = int(frameHeader.FrameLength)
		reader.Discard(skipSize)

	}
	//
	return nil, nil
}

func (self *SquishMessageBase) ReadBase(filename string) ([]*Header, error) {
	messageBaseFileName := fmt.Sprintf("%s.sqd", filename)
	stream, err := os.Open(messageBaseFileName)
	if err != nil {
		return nil, err
	}
	defer stream.Close()
	//
	var res []*Header
	reader := bufio.NewReader(stream)
	self.readMessageBaseHeader(reader)
	for {
		resp, err1 := self.readMessage(reader)
		if err1 == io.EOF {
			break
		} else if err1 != nil {
			panic(err1)
		}
		if resp != nil {
			res = append(res, resp)
		} else {
			log.Printf("No message return at readMessage.")
		}
	}
	//
	return res, nil
}

func (self *SquishMessageBase) ReadMessage(filename string, UMSGID uint32) (*Header, error) {
	messageBaseFileName := fmt.Sprintf("%s.sqd", filename)
	stream, err := os.Open(messageBaseFileName)
	if err != nil {
		return nil, err
	}
	defer stream.Close()
	//
	reader := bufio.NewReader(stream)
	self.readMessageBaseHeader(reader)
	for {
		log.Printf(" --- Read message ---")
		resp, err1 := self.readMessage(reader)
		if err1 == io.EOF {
			break
		} else if err1 != nil {
			panic(err1)
		}
		if resp != nil {
			log.Printf("Compare: %X <-> %X", resp.UMSGID, UMSGID)
			if resp.UMSGID == UMSGID {
				return resp, nil
			}
		} else {
			log.Printf("No message return at readMessage.")
		}
	}
	//
	return nil, errors.New("No message exists.")
}

func (self *SquishMessageBase) GetMessageCount() int {
	return self.messageCount
}
