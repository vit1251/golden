package tosser

import (
	"bufio"
	cmn "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/fidotime"
	"github.com/vit1251/golden/pkg/mailer/cache"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/msg"
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

func (self *Tosser) processNewMessage(pkt *TosserPacket) error {

	//packetHeader := packet.GetHeader()
	packedMessage := pkt.GetMessage()

	msgBodyParser := packet.NewMessageBodyParser()
	msgBody, err1 := msgBodyParser.Parse(packedMessage.Text)
	if err1 != nil {
		return err1
	}

	/* Process message */
	if msgBody.IsArea() {
		return self.processNewEchoMessage(packedMessage, msgBody)
	} else {
		return self.processNewDirectMessage(packedMessage, msgBody)
	}

	return nil

}

func (self *Tosser) processNewDirectMessage(msgHeader *packet.PackedMessage, msgBody *packet.MessageBody) error {

	charsetManager := self.restoreCharsetManager()
	mapperManager := self.restoreMapperManager()
	netmailMapper := mapperManager.GetNetmailMapper()

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

				/* Parse INTL "dest" "orig" */
				destAddres := parts[0]
				origAddres := parts[1]

				origAddr := packet.NewNetAddr()
				if err := origAddr.SetAddr(origAddres); err == nil {
					msgHeader.OrigAddr.Zone = origAddr.Zone
					msgHeader.OrigAddr.Net = origAddr.Net
					msgHeader.OrigAddr.Node = origAddr.Node
				} else {
					log.Printf("INTL parse addr error: err = %+v", err)
				}

				destAddr := packet.NewNetAddr()
				if err := destAddr.SetAddr(destAddres); err == nil {
					msgHeader.DestAddr.Zone = destAddr.Zone
					msgHeader.DestAddr.Net = destAddr.Net
					msgHeader.DestAddr.Node = destAddr.Node
				} else {
					log.Printf("INTL parse addr error: err = %+v", err)
				}

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
	newMsg := mapper.NewNetmailMsg()

	newMsg.SetFrom(newFrom)
	newMsg.SetTo(newTo)
	newMsg.SetSubject(newSubject)
	newMsg.SetTime(msgTime)
	newMsg.SetMsgID(msgID)
	newMsg.SetHash(msgHash)
	newMsg.SetOrigAddr(msgHeader.OrigAddr.String())
	newMsg.SetDestAddr(msgHeader.DestAddr.String())

	/* Decode message body */
	msgContent := msgBody.GetContent()
	newBody, err9 := charsetManager.DecodeMessageBody(msgContent, msgCharset)
	if err9 != nil {
		return err9
	}

	newMsg.SetContent(newBody)

	/* Write message */
	err := netmailMapper.Write(newMsg)
	log.Printf("err = %v", err)

	return nil

}

func (self *Tosser) processNewEchoMessage(msgHeader *packet.PackedMessage, msgBody *packet.MessageBody) error {

	mapperManager := self.restoreMapperManager()
	echoMapper := mapperManager.GetEchoMapper()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	statMapper := mapperManager.GetStatMapper()
	charsetManager := self.restoreCharsetManager()

	log.Printf("Process ECHOMAIL message: %q -> %q", msgHeader.OrigAddr, msgHeader.DestAddr)

	areaName := msgBody.GetArea()

	/* Auto create area */
	a := mapper.NewArea()
	a.SetName(areaName)
	echoAreaMapper.Register(a)

	/* No message encoding */
	var msgID string
	var reply string
	var msgHash string // TODO - make message hash ...
	var msgTime time.Time
	var newSubject string = string(msgHeader.Subject)
	var newFrom string = string(msgHeader.FromUserName)
	var newTo string = string(msgHeader.ToUserName)
	var msgCharset string = "CP866"
	var noTimeZone bool = true

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

			reply = value
			log.Printf("reply = %s", reply)

		} else if k.Name == "TZUTC" {

			/* Check filezone exist */
			noTimeZone = false

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
			msgTime = *msgTimePtr // NOTE - I check this is legal variant copy date struct ...
			log.Printf("fidoTime = %+v timezone = %+v msgTime = %+v", msgHeader.Time, msgZone, msgTime)

		} else {
			log.Printf("Unknown kludge: name = %s value = %s", k.Name, k.Value)
		}
	}

	/* Failsafe time */
	if noTimeZone {
		if msgTimePtr, err := msgHeader.Time.CreateTime(time.Local); err == nil {
			msgTime = *msgTimePtr
		}
	}

	/**/
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
	exists, err5 := echoMapper.IsMessageExistsByHash(areaName, msgHash)
	if err5 != nil {
		return err5
	}
	if exists {
		log.Printf("Message %s already exists", msgHash)
		statMapper.RegisterDupe()
		return nil
	}

	newOrig := msgBody.GetOrigin()

	/* Parse origin */
	originParser := NewOriginParser()
	originAddr := originParser.Parse(newOrig)
	newOriginAddr := string(originAddr)
	log.Printf("Origin: addr = %s", originAddr)


	/* Debug message */
	log.Printf("Process ECHOMAIL message:\n\tFrom: %s\n\tTo: %s", newFrom, newTo)

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
	newMsg.SetPacket(msgBody.GetPacket())
	newMsg.SetReply(reply)
	newMsg.SetFromAddr(newOriginAddr)

	/* Save message in storage */
	err6 := echoMapper.Write(*newMsg)
	if err6 != nil {
		log.Printf("Fail on Write in echoMapper: err = %+v", err6)
	}

	/* Update counter */
	statMapper.RegisterInMessage()

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
	log.Printf("Tosser: ProcessPacket: name = %s date = %+v", name, pktHeader.GetDate())
	log.Printf("Tosser: pktHeader = %+v", pktHeader)

	/* Validate packet */
	// TODO - check password and destination address ...

	/* Parse packet messages */
	var msgCount int = 0
	for {

		/* Read message header */
		packedMessage, err5 := pr.ReadPackedMessage()
		if err5 == io.EOF {
			break
		}
		if err5 != nil {
			return err5
		}
		log.Printf("packetMessage = %+v", packedMessage)

		/* Create message */
		msgTosser := NewTosserPacket()
		msgTosser.SetHeader(pktHeader)
		msgTosser.SetMessage(packedMessage)

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

func (self *Tosser) processNetmail(item cache.FileEntry) error {

	mapperManager := self.restoreMapperManager()
	statMapper := mapperManager.GetStatMapper()

	statMapper.RegisterInPacket()

	err1 := self.ProcessPacket(item.AbsolutePath)
	if err1 != nil {
		return err1
	}

	/* Construct new path */
	inbTempPath := cmn.GetTempInboundDirectory()
	newArcPath := path.Join(inbTempPath, item.Name)

	/* Move in area */
	log.Printf("Move %s -> %s", item.AbsolutePath, newArcPath)
	if err2 := os.Rename(item.AbsolutePath, newArcPath);err2 != nil {
		log.Printf("Fail on Rename: err = %+v", err2)
	}

	return nil
}

func (self *Tosser) processARCmail(item cache.FileEntry) error {

	mapperManager := self.restoreMapperManager()
	statMapper := mapperManager.GetStatMapper()

	newInbTempDir := cmn.GetTempInboundDirectory()

	/* Unpack */
	packets, err1 := archmail.Unpack(item.AbsolutePath, newInbTempDir)
	if err1 != nil {
		return err1
	}

	log.Printf("-- ARCmail provide FTN packets: packets = %+v", packets)

	for _, p := range packets {

		log.Printf("-- Process FTN packet start: packet = %+v", p)

		/* Register packet */
		if err := statMapper.RegisterInPacket(); err != nil {
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
	if err3 := os.Rename(item.AbsolutePath, newArcPath); err3 != nil {
		log.Printf("Fail on Rename: err = %+v", err3)
	}

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
		} else {
			log.Printf("Tosser: Unknoen packet: name = %s", item.Name)
		}
	}

	return nil

}
