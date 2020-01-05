package main

import (
	"github.com/vit1251/golden/pkg/packet"
	"github.com/satori/go.uuid"
	"time"
	"fmt"
)

func (self *Application) ProcessOutbound() (error) {

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
	msgHeader.SetSubject("Golden Tosser (generator)")
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

	/* Write message body */
	msgBody := packet.NewMessageBody()
	//
	msgBody.SetArea("HOBBIT.TEST")
	//
	msgBody.AddKludge("TZUTC", "0300")
	msgBody.AddKludge("CHRS", "CP866 2")
	msgBody.AddKludge("UUID", fmt.Sprintf("%s", u1))
	msgBody.AddKludge("TID", "golden/lnx 1.2.0 2020-01-04")
	//
	msgBody.SetBody([]byte("Hello,\r\n\r\nI am Golden tosser outbound message generator stub message. I am uses common message to check packet delivery.\r\n\r\nFeature: Direct\r\n\r\n * Origin: Yo Adrian, I Did It! (c) Rocky II"))
	//
	if err5 := pw.WriteMessage(msgBody); err5 != nil {
		return err5
	}

	return nil
}