package main

import (
	"log"
	"path"
//	"strings"
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
	outPack, err1 := packet.NewPacketWriter(name)
	if err1 != nil {
		return err1
	}
	defer outPack.Close()

	/* Write packet header */
	pktHeader := packet.NewPacketHeader()
	pktHeader.OrigAddr.SetAddr("2:5030/1592.11")
	pktHeader.DestAddr.SetAddr("2:5030/1592.0")
	outPack.WriteHeader(pktHeader)

	/* Prepare packet message */
	pktMessage := packet.NewPacketMessage()
	pktMessage.OrigAddr.SetAddr("2:5030/1592.11")
	pktMessage.DestAddr.SetAddr("2:5030/1592.0")
	pktMessage.SetAttribute("Direct")
	pktMessage.SetToUserName("All")
	pktMessage.SetFromUserName("Golden Tosser")
	pktMessage.SetSubject("Tosser statistics")
	pktMessage.SetText("Hello, no statistics template exists")
	var now time.Time = time.Now()
	pktMessage.SetTime(&now)
	outPack.WriteMessage(pktMessage)

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

	packet, err3 := packet.NewPacketReader(name)
	if err3 != nil {
		return err3
	}
	defer packet.Close()

	/* Process message */
	var msgCount int = 0
	for msg := range packet.Scan() {
		/* Create msgapi.Message */
		newMsg := new(sqlite.Message)
		newMsg.SetArea(msg.Area)
		newMsg.SetFrom(msg.From)
		newMsg.SetTo(msg.To)
		newMsg.SetSubject(msg.Subject)
		newMsg.SetContent(msg.Content)
		newMsg.SetUnixTime(msg.UnixTime)
		/* Store message */
		mBaseWriter.Write(newMsg)
		msgCount += 1
	}

	/* Show summary */
	log.Printf("mBase write count: %d", msgCount)

	return nil
}
