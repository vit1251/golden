package tosser

import (
	"github.com/vit1251/golden/pkg/packet"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/satori/go.uuid"
	"hash/crc32"
	"time"
	"fmt"
	"log"
)

func (self *Tosser) ProcessOutbound() (error) {
	return nil
}

func (self *Tosser) ProcessOutboundNew() (error) {

	/* Create packet name */
	name := "out.pkt"

	/* Open outbound packet */
	pw, err1 := packet.NewPacketWriter(name)
	if err1 != nil {
		return err1
	}
	defer pw.Close()

	/* Write packet header */
	pktHeader := packet.NewPacketHeader()
	pktHeader.OrigAddr.SetAddr("2:5023/24.3752")
	pktHeader.DestAddr.SetAddr("2:5023/24")
	if err2 := pw.WritePacketHeader(pktHeader); err2 != nil {
		return err2
	}

	/* Prepare packet message */
	msgHeader := packet.NewPacketMessageHeader()
	msgHeader.OrigAddr.SetAddr("2:5023/24.3752")
	msgHeader.DestAddr.SetAddr("2:5023/24.0")
	msgHeader.SetAttribute("Direct")
	msgHeader.SetToUserName("All")
	msgHeader.SetFromUserName("Vitold Sedyshev")
	msgHeader.SetSubject("Golden Point - Test 2")
	var now time.Time = time.Now()
	msgHeader.SetTime(&now)
	if err3 := pw.WriteMessageHeader(msgHeader); err3 != nil {
		return err3
	}

	/* Message UUID */
	u1 := uuid.NewV4()
//	u1, err4 := uuid.NewV4()
//	if err4 != nil {
//		return err4
//	}

	/* Construct message content */
	msgContent := msg.NewMessageContent()
	msgContent.AddLine("Добрый день,")
	msgContent.AddLine("")
	msgContent.AddLine("Это третье сообщение в UTF-8 и оно нужно быол что бы протестировать Ваши тоссеры, читалки и другой софт,")
	msgContent.AddLine("но если Вам удалось его прочитать, то я очень рад буду узнать о Вашем стеке технологий.")
	msgContent.AddLine("")
	msgContent.AddLine("Пожалуйста отпишитесь в ответ на это сообщение.")
	msgContent.AddLine("")
	msgContent.AddLine("Спасибо.")
	msgContent.AddLine("--- Golden/LNX 1.2.0 2020-01-05 18:29:20 MSK (master)")
	msgContent.AddLine(" * Origin: Yo Adrian, I Did It! (c) Rocky II (2:5023/24.3752)")
	rawMsg := msgContent.Pack()

	/* Calculate checksumm */
	h := crc32.NewIEEE()
	h.Write(rawMsg)
	hs := h.Sum32()
	log.Printf("crc32 = %+v", hs)

	/* Write message body */
	msgBody := packet.NewMessageBody()
	//
	msgBody.SetArea("UTF8.FTN.MESSAGING")
	//
	msgBody.AddKludge("TZUTC", "0300")
	msgBody.AddKludge("CHRS", "UTF-8 4")
	msgBody.AddKludge("MSGID", fmt.Sprintf("%s %08x", "2:5023/24.3752", hs))
	msgBody.AddKludge("UUID", fmt.Sprintf("%s", u1))
	msgBody.AddKludge("TID", "golden/lnx 1.2.1 2020-01-05 20:41 (master)")
	//
	msgBody.SetRaw(rawMsg)
	//
	if err5 := pw.WriteMessage(msgBody); err5 != nil {
		return err5
	}

	return nil
}
