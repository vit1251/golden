package tosser

import (
	"log"
	"io"
	"github.com/vit1251/golden/pkg/packet"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/mailer"
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

func (self *Tosser) processNetmail(item *mailer.MailerInboundRec) (error) {
	return nil
}

func (self *Tosser) processARCmail(item *mailer.MailerInboundRec) (error) {

	packets, err1 := Unpack(item.AbsolutePath, self.workInboundDirectory)
	if err1 != nil {
		return err1
	}

	for _, packet := range packets {
		log.Printf("-- Process FTN packets %+v", packets)
		err2 := self.ProcessPacket(packet)
		log.Printf("error durng parse package: err = %+v", err2)
	}

	return nil

}

func (self *Tosser) processTICmail(item *mailer.MailerInboundRec) (error) {
	return nil
}

func (self *Tosser) ProcessInbound() (error) {

	/* New mailer inbound */
	mi := mailer.NewMailerInbound()
//	if err1 != nil {
//		return err1
//	}
	mi.SetInboundDirectory(self.inboundDirectory)

	/* Scan inbound */
	items, err2 := mi.Scan()
	if err2 != nil {
		return err2
	}

	for _, item := range items {
		if item.Type == mailer.TypeNetmail {
			log.Printf(" - Found Netmail packet %s. Skip.", item.Name)
			self.processNetmail(item)
		} else if item.Type == mailer.TypeARCmail {
			log.Printf(" - Found ARCmail packet %s. Process.", item.Name)
			self.processARCmail(item)
		} else if item.Type == mailer.TypeTICmail {
			log.Printf(" - Found TIC packet %s. Skip.", item.Name)
			self.processTICmail(item)
		} else {
			log.Printf(" - Found other packet %s. Skip.", item.Name)
		}
	}

	return nil

}
