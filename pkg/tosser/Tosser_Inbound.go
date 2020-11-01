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

func (self *Tosser) parseCharsetKludge(charsetKludge string) (string) {

	var result string = "CP866"

	/* Trim spaces */
	charsetKludge = strings.Trim(charsetKludge, " ")

	/* Split */
	charsetParts := strings.Split(charsetKludge, " ")
	charsetPartCount := len(charsetParts)
	if charsetPartCount == 2 {
		result = charsetParts[0]
	}

	log.Printf("Message charset: %+v", result)

	return result

}

func (self *Tosser) processNewMessage(msg *TosserPacketMessage) error {

	msgParser := packet.NewMessageBodyParser()
	msgBody, err8 := msgParser.Parse(msg.Body)
	if err8 != nil {
		return err8
	}

	/* Determine area */
	areaName := msgBody.GetArea()
	if areaName == "" {
		return self.processNewDirectMessage(msg.Header, msgBody)
	} else {
		return self.processNewEchoMessage(msg.Header, msgBody)
	}

	return nil
}

func (self *Tosser) processNewDirectMessage(msgHeader *packet.PacketMessageHeader, msgBody *packet.MessageBody) error {

	charsetManager := self.restoreCharsetManager()
	netmailManager := self.restoreNetmailManager()

	/* FST-40001 - Parse NETMAIL source address */
	for _, k := range msgBody.GetKludges() {
		if k.Name == "INTL" {
			// TODO - parse ...
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
		}
	}

	log.Printf("NETMAIL message %q -> %q", msgHeader.OrigAddr, msgHeader.DestAddr)

	/* Decode headers */
	newSubject, err1 := charsetManager.Decode(msgHeader.Subject)
	if err1 != nil {
		return err1
	}
	newFrom, err2 := charsetManager.Decode(msgHeader.FromUserName)
	if err2 != nil {
		return err2
	}
	newTo, err3 := charsetManager.Decode(msgHeader.ToUserName)
	if err3 != nil {
		return err3
	}

	/* Populate message */
	newMsg := netmail.NewNetmailMessage()

	newMsg.SetFrom(string(newFrom))
	newMsg.SetTo(string(newTo))
	newMsg.SetSubject(string(newSubject))

	// TODO - populate message ...

	/* Decode message */
	charsetKludge := msgBody.GetKludge("CHRS", "CP866 2")
	charset := self.parseCharsetKludge(charsetKludge)

	/* Determine msgid */
	msgid := msgBody.GetKludge("MSGID", "")
	msgid = strings.Trim(msgid, " ")
	log.Printf("msgid = %s", msgid)
	msgidParts := strings.Split(msgid, " ")
	var msgHash string
	if len(msgidParts) == 2 {
		//source := msgidParts[0]
		msgHash = msgidParts[1]
	}

	/* Determine zone */
	zone := msgBody.GetKludge("TZUTC", " 0300")
	log.Printf("zone = %+v", zone)
	zone = strings.Trim(zone, " \t")
	zParser := fidotime.NewTimeZoneParser()
	msgZone, err10 := zParser.Parse(zone)
	if err10 != nil {
		return err10
	}
	log.Printf("Message zone: zone = %+v", msgZone)

	log.Printf("orig: msgTime = %+v", msgHeader.Time)
	msgTime, err11 := msgHeader.Time.CreateTime(msgZone)
	if err11 != nil {
		return err11
	}
	log.Printf("msgTime = %+v", msgTime)
	log.Printf("msgTime (Local) = %+v", msgTime.Local())

	newMsg.SetTime(msgTime)

	newMsg.SetMsgID(msgid)
	newMsg.SetHash(msgHash)
	newMsg.SetPacket(msgBody.RAW)

	/* Decode message body */
	newBody, err9 := charsetManager.DecodeMessageBody(msgBody.RAW, charset)
	if err9 != nil {
		return err9
	}

	newMsg.SetContent(newBody)

	/* Write message */
	err := netmailManager.Write(newMsg)
	log.Printf("err = %v", err)

	return nil

}

func (self *Tosser) processNewEchoMessage(msgHeader *packet.PacketMessageHeader, msgBody *packet.MessageBody) error {

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
	var newBody string = string(msgBody.RAW)
	var newSubject string = string(msgHeader.Subject)
	var newFrom string = string(msgHeader.FromUserName)
	var newTo string = string(msgHeader.ToUserName)

	/* Process message kludges */
	for _, k := range msgBody.GetKludges() {
		if k.Name == "CHRS" {

			charset := self.parseCharsetKludge(k.Value)

			var err1 error
			var err2 error
			var err3 error
			var err4 error

			/* Decode body */
			newBody, err1 = charsetManager.DecodeMessageBody(msgBody.RAW, charset)
			if err1 != nil {
				return err1
			}

			/* Decode headers */
			newSubject, err2 = charsetManager.DecodeMessageBody(msgHeader.Subject, charset)
			if err2 != nil {
				return err2
			}

			newFrom, err3 = charsetManager.DecodeMessageBody(msgHeader.FromUserName, charset)
			if err3 != nil{
				return err3
			}

			newTo, err4 = charsetManager.DecodeMessageBody(msgHeader.ToUserName, charset)
			if err4 != nil {
				return err4
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
	newMsg.SetPacket(msgBody.RAW)

	/* Save message in storage */
	if err := messageManager.Write(newMsg); err != nil {
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
