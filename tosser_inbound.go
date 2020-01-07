package main

import (
	"log"
	"path"
	"io"
	"io/ioutil"
	"github.com/vit1251/golden/pkg/packet"
	"github.com/vit1251/golden/pkg/msgapi/sqlite"
	"github.com/vit1251/golden/pkg/tosser"
	"strings"
	"errors"
)

func ProcessPacket(name string) (error) {

	mBase, err1 := sqlite.NewMessageBase()
	if err1 != nil {
		return err1
	}
	mBaseWriter, err2 := sqlite.NewMessageBaseWriter(mBase)
	if err2 != nil {
		return err2
	}

	/* Start new packet reader */
	pr, err3 := packet.NewPacketReader(name)
	if err3 != nil {
		return err3
	}
	defer pr.Close()

	/* Read packet header */
	pktHeader, err4 := pr.ReadPacketHeader()
	if err4 != nil {
		return err4
	}
	log.Printf("pktHeader = %+v", pktHeader)

	/* Read messages */
	var msgCount int = 0
	for {

		/* Read message header */
		msgHeader, err5 := pr.ReadMessageHeader()
		if err5 == io.EOF {
			break
		}
		if err5 != nil {
			return err5
		}
		log.Printf("msgHeader = %+v", msgHeader)

		/* Read message body */
		rawBody, err6 := pr.ReadMessage()
		if err6 != nil {
			return err6
		}

		/* Process message */
		msgParser, err7 := packet.NewMessageBodyParser()
		if err7 != nil {
			return err7
		}
		msgBody, err8 := msgParser.Parse(rawBody)
		if err8 != nil {
			return err8
		}

		/* Determine area */
		var areaName string = msgBody.GetArea()
		if areaName == "" {
			areaName = "NETMAIL"
		}

		/* Decode message */
		charset := msgBody.GetKludge("CHRS", "CP866 2")
		charset = strings.Trim(charset, " ")
		log.Printf("charset = %+v", charset)
		var newBody string
		if charset == "CP866 2" {
			if unicodeBody, err9 := packet.DecodeText(msgBody.RAW); err9 == nil {
				newBody = string(unicodeBody)
			} else {
				return err9
			}
		} else if charset == "UTF-8 4"{
			newBody = string(msgBody.RAW)
		} else {
			return errors.New("Fail charset on message")
		}

		/* Determine message hash */
		msgid := msgBody.GetKludge("MSGID", "")
		msgid = strings.Trim(msgid, " ")
		msgidParts := strings.Split(msgid, " ")
		var msgHash string
		if len(msgidParts) == 2 {
			//source := msgidParts[0]
			msgHash = msgidParts[1]
		}

		/* Create msgapi.Message */
		newMsg := new(sqlite.Message)
		newMsg.SetMsgID(msgHash)
		newMsg.SetArea(areaName)
		newMsg.SetFrom(msgHeader.FromUserName)
		newMsg.SetTo(msgHeader.ToUserName)
		newMsg.SetSubject(msgHeader.Subject)
		newMsg.SetTime(msgHeader.Time)

		newMsg.SetContent(newBody)

		/* Store message */
		mBaseWriter.Write(newMsg)

		/* Update counter */
		msgCount += 1
	}

	/* Show summary */
	log.Printf("Toss area message: %d", msgCount)

	return nil
}

func SearchArcmail() {

	baseDir := "/var/spool/ftn/inb"
	workDirectory := "/var/spool/ftn/tmp.inb"

	items, err := ioutil.ReadDir(baseDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {
		mode := item.Mode()
		if (mode.IsRegular()) {
			if IsNetmail(item.Name()) {

			} else if IsArchmail(item.Name()) {
				absPath := path.Join(baseDir, item.Name())
				log.Printf("Process %s", absPath)
				packets, err := tosser.Unpack(absPath, workDirectory)
				if err != nil {
					log.Fatal(err)
				}
				for _, packet := range packets {
					ProcessPacket(packet)
				}
				log.Printf("Packets %s", packets)
			}
		}
	}

}

func (self *Application) ProcessInbound() (error) {
	SearchArcmail()
	//ProcessPacket("testdata/5de3695e.pkt")
	return nil
}