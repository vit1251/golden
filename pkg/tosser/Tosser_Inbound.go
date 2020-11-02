package tosser

import (
	"bufio"
	"fmt"
	"github.com/vit1251/golden/pkg/echomail"
	"github.com/vit1251/golden/pkg/fidotime"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/mailer/cache"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/netmail"
	"github.com/vit1251/golden/pkg/packet"
	"github.com/vit1251/golden/pkg/tosser/arcmail"
	"io"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

func (self *Tosser) processNewMessage(pktMessage *TosserPacketMessage) error {

	msgContentParser := msg.NewMessageContentParser()
	msgBody, err1 := msgContentParser.Parse(pktMessage.Body)
	if err1 != nil {
		return err1
	}

	/* Process message */
	if msgBody.IsArea() {
		return self.processNewEchoMessage(pktMessage.Header, msgBody)
	} else {
		return self.processNewDirectMessage(pktMessage.Header, msgBody)
	}

	return nil

}

func (self *Tosser) processNewDirectMessage(msgHeader *packet.PacketMessageHeader, msgBody *msg.MessageContent) error {

	charsetManager := self.restoreCharsetManager()
	netmailManager := self.restoreNetmailManager()

	var msgID string
	var msgHash string
	var msgCharset string = "CP866"
	var msgTime time.Time = time.Now()

	/* FST-40001 - Parse NETMAIL source address */
	for _, k := range msgBody.GetKludges() {
		if k.Name == "INTL" {

			value := strings.Trim(k.Value, " ")
			parts := strings.Split(value, " ")
			if len(parts) == 2 {

				orig := parts[0]
				dest := parts[1]

				origAddr := packet.NewNetAddr()
				if err := origAddr.SetAddr(orig); err == nil {
					msgHeader.OrigAddr.Zone = origAddr.Zone
					msgHeader.OrigAddr.Net = origAddr.Net
					msgHeader.OrigAddr.Node = origAddr.Node
				} else {
					log.Printf("INTL parse addr error: err = %+v", err)
				}

				destAddr := packet.NewNetAddr()
				if err := destAddr.SetAddr(dest); err == nil {
					msgHeader.DestAddr.Zone = destAddr.Zone
					msgHeader.DestAddr.Net = destAddr.Net
					msgHeader.DestAddr.Node = destAddr.Node
				} else {
					log.Printf("INTL parse addr error: err = %+v", err)
				}

				log.Printf("NETMAIL kludge INTL: orig = %+v dest = %+v", orig, dest)

			}

		} else if k.Name == "TOPT" {
			if newPoint, err := strconv.ParseUint(k.Value, 10, 16); err == nil {
				msgHeader.DestAddr.Point = uint16(newPoint)
			} else {
				log.Printf("Parse TOPT klude error: value = %q", k.Value)
			}
		} else if k.Name == "FMPT" {
			if newPoint, err := strconv.ParseUint(k.Value, 10, 16); err == nil {
				msgHeader.OrigAddr.Point = uint16(newPoint)
			} else {
				log.Printf("Parse FMPT klude error: value = %q", k.Value)
			}
		} else if k.Name == "CHRS" {

			value := strings.Trim(k.Value, " ")

			parts := strings.Split(value, " ")
			if len(parts) == 2 {
				msgCharset = parts[0]
			}

		} else if k.Name == "TZUTC" {

			value := strings.Trim(k.Value, " ")

			zParser := fidotime.NewTimeZoneParser()
			msgZone, err10 := zParser.Parse(value)

			if err10 != nil {
				return err10
			}
			log.Printf("Message zone: zone = %+v", msgZone)

			msgTime, err11 := msgHeader.Time.CreateTime(msgZone)
			if err11 != nil {
				return err11
			}

			log.Printf("fidoTime = %+v zone = %+v msgTime = %+v", msgHeader.Time, msgZone, msgTime)

		} else if k.Name == "MSGID" {

			value := strings.Trim(k.Value, " ")
			msgID = value
			log.Printf("msgid = %s", value)

			parts := strings.Split(value, " ")
			if len(parts) == 2 {
				//source := parts[0]
				msgHash = parts[1]
			}

		} else if k.Name == "REPLY" {

			// TODO - ...

		} else {
			log.Printf("Unknown kludge: name = %+v value = %+v", k.Name, k.Value)
		}
	}

	log.Printf("NETMAIL message %q -> %q", msgHeader.OrigAddr, msgHeader.DestAddr)

	/* Decode headers */
	newSubject, err1 := charsetManager.DecodeMessageBody(msgHeader.Subject, msgCharset)
	if err1 != nil {
		return err1
	}
	newFrom, err2 := charsetManager.DecodeMessageBody(msgHeader.FromUserName, msgCharset)
	if err2 != nil {
		return err2
	}
	newTo, err3 := charsetManager.DecodeMessageBody(msgHeader.ToUserName, msgCharset)
	if err3 != nil {
		return err3
	}

	/* Populate message */
	newMsg := netmail.NewNetmailMessage()

	newMsg.SetFrom(string(newFrom))
	newMsg.SetTo(string(newTo))
	newMsg.SetSubject(string(newSubject))
	newMsg.SetTime(msgTime)
	newMsg.SetMsgID(msgID)
	newMsg.SetHash(msgHash)
	//newMsg.SetPacket(msgBody.RAW)

	/* Decode message body */
	msgContent := msgBody.GetContent()
	newBody, err9 := charsetManager.DecodeMessageBody(msgContent, msgCharset)
	if err9 != nil {
		return err9
	}

	newMsg.SetContent(newBody)

	/* Write message */
	err := netmailManager.Write(newMsg)
	log.Printf("err = %v", err)

	return nil

}

func (self *Tosser) processNewEchoMessage(msgHeader *packet.PacketMessageHeader, msgBody *msg.MessageContent) error {

	areaManager := self.restoreAreaManager()
	messageManager := self.restoreMessageManager()
	statManager := self.restoreStatManager()
	charsetManager := self.restoreCharsetManager()

	log.Printf("Process ECHOMAIL message: %q -> %q", msgHeader.OrigAddr, msgHeader.DestAddr)

	areaName := msgBody.GetArea()

	/* Auto create area */
	a := echomail.NewArea()
	a.SetName(areaName)
	areaManager.Register(a)

	/* No message encoding */
	var msgID string
	var msgHash string // TODO - make message hash ...
	var msgTime time.Time = time.Now()
	var newSubject string = string(msgHeader.Subject)
	var newFrom string = string(msgHeader.FromUserName)
	var newTo string = string(msgHeader.ToUserName)
	var msgCharset string = "CP866"

	/* Process message kludges */
	for _, k := range msgBody.GetKludges() {
		if k.Name == "CHRS" {

			value := strings.Trim(k.Value, " ")

			parts := strings.Split(value, " ")
			if len(parts) == 2 {
				msgCharset = parts[0]
			}

		} else if k.Name == "MSGID" {

			value := strings.Trim(k.Value, " ")
			msgID = value

			parts := strings.Split(value, " ")
			if len(parts) == 2 {
				//source := msgidParts[0]
				msgHash = parts[1]
			} else {
				log.Printf("Problem with MSGID kludge: value = %s", value)
			}

		} else if k.Name == "REPLY" {

			value := strings.Trim(k.Value, " ")
			reply := strings.Trim(value, " ")
			log.Printf("reply = %s", reply)
			// TODO - save reply chains ...

		} else if k.Name == "TZUTC" {

			value := strings.Trim(k.Value, " ")

			zParser := fidotime.NewTimeZoneParser()
			msgZone, err10 := zParser.Parse(value)
			if err10 != nil {
				return err10
			}
			log.Printf("Message zone: zone = %+v", msgZone)

			msgTimePtr, err11 := msgHeader.Time.CreateTime(msgZone)
			if err11 != nil {
				return err11
			}
			msgTime = *msgTimePtr // TODO - I check this is legal ... buy require clarification ...
			log.Printf("fidoTime = %+v timezone = %+v msgTime = %+v", msgHeader.Time, msgZone, msgTime)

		} else {
			log.Printf("Unknown kludge: name = %s value = %s", k.Name, k.Value)
		}
	}

	msgContent := msgBody.GetContent()

	/* Decode routine */
	newBody, err1 := charsetManager.DecodeMessageBody(msgContent, msgCharset)
	if err1 != nil {
		return err1
	}
	newSubject, err2 := charsetManager.DecodeMessageBody(msgHeader.Subject, msgCharset)
	if err2 != nil {
		return err2
	}
	newFrom, err3 := charsetManager.DecodeMessageBody(msgHeader.FromUserName, msgCharset)
	if err3 != nil{
		return err3
	}
	newTo, err4 := charsetManager.DecodeMessageBody(msgHeader.ToUserName, msgCharset)
	if err4 != nil {
		return err4
	}

	/* Check duplicate message */
	exists, err9 := messageManager.IsMessageExistsByHash(areaName, msgHash)
	if err9 != nil {
		return err9
	}
	if exists {
		log.Printf("Message %s already exists", msgHash)
		statManager.RegisterDupe()
		return nil
	}


	/* Debug message */
	newFromAddr := fmt.Sprintf("\"%s\" <%s>", newFrom, msgHeader.OrigAddr)
	newToAddr := fmt.Sprintf("\"%s\" <%s>", newTo, msgHeader.DestAddr)
	log.Printf("Process ECHOMAIL message:\n\tFrom: %s\n\tTo: %s", newFromAddr, newToAddr)

	/* Create Message */
	newMsg := msg.NewMessage()
	newMsg.SetMsgID(msgID)
	newMsg.SetMsgHash(msgHash)
	newMsg.SetArea(areaName)
	newMsg.SetFrom(newFrom)
	newMsg.SetTo(newTo)
	newMsg.SetSubject(newSubject)
	newMsg.SetTime(msgTime)
	newMsg.SetContent(newBody)
	//newMsg.SetPacket(msgBody.RAW)

	/* Save message in storage */
	if err := messageManager.Write(*newMsg); err != nil {
		log.Printf("Fail on Write in MessageManager: err = %+v", err)
	}

	/* Update counter */
	statManager.RegisterInMessage()

	return nil
}

func (self *Tosser) ProcessPacket(name string) error {

	/* Open stream */
	stream, err1 := os.Open(name)
	if err1 != nil {
		return err1
	}
	defer stream.Close()

	/* Cache open stream */
	cacheStream := bufio.NewReader(stream)

	/* Parse packet */
	pr := packet.NewPacketReader(cacheStream)

	/* Parse packet header */
	pktHeader, err4 := pr.ReadPacketHeader()
	if err4 != nil {
		return err4
	}
	log.Printf("pktHeader = %+v", pktHeader)

	/* Validate packet */
	// TODO - check password and destination address ...

	/* Parse packet messages */
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

		/* Create message */
		msgTosser := NewTosserPacketMessage()
		msgTosser.Header = msgHeader
		msgTosser.Body = rawBody

		/* Process message */
		err7 := self.processNewMessage(msgTosser)
		if err7 != nil {
			log.Printf("Tosser: ProcessPacket: err = %+v", err7)
		}

		/* Update message count */
		msgCount += 1

	}

	/* Show summary */
	log.Printf("Toss area message: %d", msgCount)

	return nil
}

func (self *Tosser) processNetmail(item *cache.FileEntry) error {

	configManager := self.restoreConfigManager()
	statManager := self.restoreStatManager()

	inbTempPath, _ := configManager.Get("main", "TempInbound")

	statManager.RegisterInPacket()

	err1 := self.ProcessPacket(item.AbsolutePath)
	if err1 != nil {
		return err1
	}

	/* Construct new path */
	newArcPath := path.Join(inbTempPath, item.Name)

	/* Move in area*/
	log.Printf("Move %s -> %s", item.AbsolutePath, newArcPath)
	os.Rename(item.AbsolutePath, newArcPath)


	return nil
}

func (self *Tosser) processARCmail(item *cache.FileEntry) error {

	configManager := self.restoreConfigManager()
	statManager := self.restoreStatManager()

	newInbTempDir, _ := configManager.Get("main", "TempInbound")

	/* Unpack */
	packets, err1 := archmail.Unpack(item.AbsolutePath, newInbTempDir)
	if err1 != nil {
		return err1
	}

	log.Printf("-- ARCmail provide FTN packets: packets = %+v", packets)

	for _, p := range packets {

		log.Printf("-- Process FTN packet start: packet = %+v", p)

		/* Register packet */
		if err := statManager.RegisterInPacket(); err != nil {
			log.Printf("Fail on RegisterInPacket: err = %+v", err)
		}

		/* Process PKT network packet */
		if err := self.ProcessPacket(p); err != nil {
			log.Printf("error durng parse package: err = %+v", err)
		}

		log.Printf("-- Process FTN packet complete: packet = %+v", p)

	}

	/* Construct new path */
	newArcPath := path.Join(newInbTempDir, item.Name)

	/* Move in area*/
	log.Printf("Move %s -> %s", item.AbsolutePath, newArcPath)
	os.Rename(item.AbsolutePath, newArcPath)

	return nil

}

func (self *Tosser) processTICmail(item *cache.FileEntry) (error) {

	charsetManager := self.restoreCharsetManager()
	fileManager := self.restoreFileManager()
	configManager := self.restoreConfigManager()
	statManager := self.restoreStatManager()

	/* Parse */
	newTicParser := file.NewTicParser(charsetManager)
	tic, err1 := newTicParser.ParseFile(item.AbsolutePath)
	if err1 != nil {
		return err1
	}
	log.Printf("tic = %+v", tic)

	areaName := tic.GetArea()

	/* Search area */
	fa, err1 := fileManager.GetAreaByName(areaName)
	if err1 != nil {
		return err1
	}

	/* Prepare area directory */
	boxBasePath, _ := configManager.Get("main", "FileBox")
	inboxBasePath, _ := configManager.Get("main", "Inbound")

	areaLocation := path.Join(boxBasePath, areaName)
	os.MkdirAll(areaLocation, 0755)

	/* Create area */
	if fa == nil {
		/* Prepare area */
		newFa := file.NewFileArea()
		newFa.SetName(areaName)
		newFa.Path = areaLocation
		/* Create area */
		if err := fileManager.CreateFileArea(newFa); err != nil {
			log.Printf("Fail CreateFileArea on FileManager: area = %s err = %+v", areaName, err)
			return err
		}
	}

	/* Create new path */
	inboxTicLocation := path.Join(inboxBasePath, tic.File)
	areaFileLocation := path.Join(areaLocation, tic.File)
	log.Printf("inboxTicLocation = %s areaFileLocation = %s", inboxTicLocation, areaFileLocation)

	/* Move */
	os.Rename(inboxTicLocation, areaFileLocation)

	/* Register file */
	fileManager.RegisterFile(tic)

	/* Register status */
	statManager.RegisterInFile(tic.File)

	/* Move TIC */
	areaTicLocation := path.Join(areaLocation, item.Name)
	log.Printf("areaTicLocation = %s", areaTicLocation)
	os.Rename(item.AbsolutePath, areaTicLocation)

	return nil
}

func (self *Tosser) ProcessInbound() error {

	log.Printf("ProcessInbound")

	/* New mailer inbound */
	mi := cache.NewMailerInbound(self.registry)

	/* Scan inbound */
	items, err2 := mi.Scan()
	if err2 != nil {
		return err2
	}
	log.Printf("items = %+v", items)

	for _, item := range items {
		if item.Type == cache.TypeNetmail {
			log.Printf("Tosser: Netmail packet: name = %s", item.Name)
			self.processNetmail(item)
		} else if item.Type == cache.TypeARCmail {
			log.Printf("Tosser: ARCmail packet: name = %s", item.Name)
			self.processARCmail(item)
		} else if item.Type == cache.TypeTICmail {
			log.Printf("Tosser: TIC packet: name = %s", item.Name)
			self.processTICmail(item)
		} else {
			log.Printf("Tosser: Unknoen packet: name = %s", item.Name)
		}
	}

	return nil

}
