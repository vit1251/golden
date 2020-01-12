package tosser

import (
	"log"
	"path"
	"io"
	"io/ioutil"
	"github.com/vit1251/golden/pkg/packet"
	"github.com/vit1251/golden/pkg/msg"
	"strings"
	"errors"
)

func (self *Tosser) ProcessPacket(name string) (error) {

	messageManager := msg.NewMessageManager()

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

		/* Determine msgid */
		msgid := msgBody.GetKludge("MSGID", "")
		msgid = strings.Trim(msgid, " ")
		msgidParts := strings.Split(msgid, " ")
		var msgHash string
		if len(msgidParts) == 2 {
			//source := msgidParts[0]
			msgHash = msgidParts[1]
		}

//		/* Determine reply */
//		reply := msgBody.GetKludge("REPLY", "")
//		reply = strings.Trim(reply, " ")
//		log.Printf("reply = %s", reply)
//		replyParts := strings.Split(reply, " ")
//		var replyHash string
//		if len(replyParts) == 2 {
//			//source := replyParts[0]
//			replyHash = replyParts[1]
//		}

		/* Check unique item */
		exists, err9 := messageManager.IsMessageExistsByHash(areaName, msgHash)
		if err9 != nil {
			return err9
		}
		if exists {

			log.Printf("Message %s already exists", msgHash)

		} else {

			/* Create msgapi.Message */
			newMsg := new(msg.Message)
			newMsg.SetMsgID(msgHash)
			newMsg.SetArea(areaName)
			newMsg.SetFrom(msgHeader.FromUserName)
			newMsg.SetTo(msgHeader.ToUserName)
			newMsg.SetSubject(msgHeader.Subject)
			newMsg.SetTime(msgHeader.Time)

			newMsg.SetContent(newBody)

			/* Store message */
			err := messageManager.Write(newMsg)
			log.Printf("err = %v", err)

			/* Update counter */
			msgCount += 1
		}

	}

	/* Show summary */
	log.Printf("Toss area message: %d", msgCount)

	return nil
}

func (self *Tosser) SearchArcmail() {

	items, err := ioutil.ReadDir(self.inboundDirectory)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range items {

		absPath := path.Join(self.inboundDirectory, item.Name())

		mode := item.Mode()
		if (mode.IsRegular()) {

			if IsNetmail(item.Name()) {

				log.Printf("Found netmail packet %s. Skip.", item.Name() )

			} else if IsArchmail(item.Name()) {

				log.Printf("Found archmail packet %s. Process.", item.Name() )

				log.Printf("Process %s", absPath)

				packets, err := Unpack(absPath, self.workInboundDirectory)
				if err != nil {
					log.Fatal(err)
				}

				log.Printf("Process FTN packets %+v", packets)

				for _, packet := range packets {
					self.ProcessPacket(packet)
				}

			} else {

				log.Printf("Found unknown packet %s. Skip.", item.Name() )

			}
		}
	}

}

func (self *Tosser) ProcessInbound() (error) {
	self.SearchArcmail()
	//ProcessPacket("testdata/5de3695e.pkt")
	return nil
}
