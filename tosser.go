package main

import (
	"log"
	"path"
	"io"
	"io/ioutil"
	"github.com/vit1251/golden/pkg/packet"
	"github.com/vit1251/golden/pkg/msgapi/sqlite"
	"github.com/vit1251/golden/pkg/tosser"
	"path/filepath"
	"time"
)

type Application struct {
}

func IsNetmail(name string) bool {
	var ext string = filepath.Ext(name)
	return ext == ".pkt"
}

func IsArchmail(name string) bool {
	var ext string = filepath.Ext(name)
	return ext != ".pkt"
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
	pktHeader.OrigAddr.SetAddr("2:5030/1592.11")
	pktHeader.DestAddr.SetAddr("2:5030/1592.0")
	if err2 := pw.WritePacketHeader(pktHeader); err2 != nil {
		return err2
	}

	/* Prepare packet message */
	msgHeader := packet.NewPacketMessageHeader()
	msgHeader.OrigAddr.SetAddr("2:5030/1592.11")
	msgHeader.DestAddr.SetAddr("2:5030/1592.0")
	msgHeader.SetAttribute("Direct")
	msgHeader.SetToUserName("All")
	msgHeader.SetFromUserName("Golden Tosser")
	msgHeader.SetSubject("Tosser statistics")
	var now time.Time = time.Now()
	msgHeader.SetTime(&now)
	if err3 := pw.WriteMessageHeader(msgHeader); err3 != nil {
		return err3
	}

	/* Write message body */
	msgBody := packet.NewMessageBody()
	msgBody.SetKludge([]byte("TZUTC"), []byte("0300"))
	msgBody.SetKludge([]byte("AREA"), []byte("HOBBIT.TEST"))
	msgBody.SetKludge([]byte("CHARSET"), []byte("CP866 2"))
	msgBody.SetBody([]byte("Hello, no statistics template exists"))
	if err4 := pw.WriteMessage(msgBody); err4 != nil {
		return err4
	}

	return nil
}

func (self *Application) Run() {
	self.ProcessInbound()
	self.ProcessOutbound()
}

func Tosser() {
	app := new(Application)
	app.Run()
}

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
		var areaName string = "NETMAIL"
		if area, ok := msgBody.Kludges["AREA"]; ok {
			areaName = area
		}

		/* Decode message */
		/* TODO - search charmap in message kludge */
//		newBody, err9 := packet.DecodeText(msgBody.RAW)
//		if err9 != nil {
//			return err9
//		}

		/* Determine dupe */
		// TODO - add checking dupe message ...

		/* Create msgapi.Message */
		newMsg := new(sqlite.Message)
		newMsg.SetArea(areaName)
		newMsg.SetFrom(msgHeader.FromUserName)
		newMsg.SetTo(msgHeader.ToUserName)
		newMsg.SetSubject(msgHeader.Subject)
		newMsg.SetTime(msgHeader.Time)

		newMsg.SetContent(msgBody.Body)

		/* Store message */
		mBaseWriter.Write(newMsg)

		/* Update counter */
		msgCount += 1
	}

	/* Show summary */
	log.Printf("Toss area message: %d", msgCount)

	return nil
}
