package tosser

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/packet"
	"github.com/vit1251/golden/pkg/setup"
	"hash/crc32"
	"log"
	"path"
	"time"
)

type TosserManager struct {
	SetupManager *setup.SetupManager
}

type NetmailMessage struct {
	Subject string
	To string
	From string
	Body string
}

type EchoMessage struct {
	Subject string
	To string
	From string
	Body string
	AreaName string
}

func (m *EchoMessage) SetSubject(subj string) {
	m.Subject = subj
}

func NewTosserManager(sm *setup.SetupManager) *TosserManager {
	tm := new(TosserManager)
	tm.SetupManager = sm
	return tm
}

func (self *TosserManager) WriteEchoMessage(em *EchoMessage) error {

	/* Create packet name */
	outb, err1 := self.SetupManager.Get("main", "Outbound", ".")
	if err1 != nil {
		return err1
	}

	name := path.Join(outb, "compose.pkt")

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
	msgHeader.SetToUserName(em.To)
	msgHeader.SetFromUserName("Vitold Sedyshev")
	msgHeader.SetSubject(em.Subject)
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

	msgContent.SetCharset("CP866")

	msgContent.AddLine(em.Body)
	msgContent.AddLine("")
	msgContent.AddLine("--- Golden/LNX 1.2.8 2020-02-14 11:10:20 MSK (master)")
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
	msgBody.SetArea(em.AreaName)
	//
	msgBody.AddKludge("TZUTC", "0300")
	//msgBody.AddKludge("CHRS", "UTF-8 4")
	msgBody.AddKludge("CHRS", "CP866 2")
	msgBody.AddKludge("MSGID", fmt.Sprintf("%s %08x", "2:5023/24.3752", hs))
	msgBody.AddKludge("UUID", fmt.Sprintf("%s", u1))
	msgBody.AddKludge("TID", "golden/lnx 1.2.8 2020-01-05 20:41 (master)")
	//
	msgBody.SetRaw(rawMsg)
	//
	if err5 := pw.WriteMessage(msgBody); err5 != nil {
		return err5
	}

	return nil
}

func (self *TosserManager) WriteNetmailMessage(em *NetmailMessage) error {

	/* Create packet name */
	name := "new_netmail.pkt"

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
	msgHeader.SetToUserName(em.To)
	msgHeader.SetFromUserName("Vitold Sedyshev")
	msgHeader.SetSubject(em.Subject)
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
	msgContent.AddLine(em.Body)
	msgContent.AddLine("")
	msgContent.AddLine("--- Golden/LNX 1.2.8 2020-01-05 18:29:20 MSK (master)")
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
	msgBody.AddKludge("TZUTC", "0300")
	msgBody.AddKludge("CHRS", "CP866 2")
	msgBody.AddKludge("MSGID", fmt.Sprintf("%s %08x", "2:5023/24.3752", hs))
	msgBody.AddKludge("UUID", fmt.Sprintf("%s", u1))
	msgBody.AddKludge("TID", "golden/lnx 1.2.1 2020-01-05 20:41 (master)")
	msgBody.AddKludge("TOPT", "0")
	msgBody.AddKludge("FMPT", "3752")

	//
	msgBody.SetRaw(rawMsg)

	//
	if err5 := pw.WriteMessage(msgBody); err5 != nil {
		return err5
	}

	return nil
}

func (self *TosserManager) NewEchoMessage() *EchoMessage {
	em := new(EchoMessage)
	return em
}

func (self *TosserManager) NewNetmailMessage() *NetmailMessage {
	nm := new(NetmailMessage)
	return nm
}