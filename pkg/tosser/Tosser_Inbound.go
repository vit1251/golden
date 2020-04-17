package tosser

import (
	"errors"
	"github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/fidotime"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/mailer"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/netmail"
	"github.com/vit1251/golden/pkg/packet"
	"io"
	"log"
	"os"
	"path"
	"strings"
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

func (self *Tosser) decodeMessageBody(msgBody []byte, charset string) (string, error) {

	var result string

	if charset == "CP866" {

		if unicodeBody, err1 := self.CharsetManager.DecodeText(msgBody); err1 == nil {
			result = string(unicodeBody)
		} else {
			return result, err1
		}

	} else if charset == "UTF-8" {

		result = string(msgBody)

	} else if charset == "LATIN-1" {

		result = string(msgBody)

	} else {

		return result, errors.New("wrong charset on message")

	}

	return result, nil
}

func (self *Tosser) processNewMessage(msgHeader *packet.PacketMessageHeader, rawBody []byte) error {

	msgParser := packet.NewMessageBodyParser()
	msgBody, err8 := msgParser.Parse(rawBody)
	if err8 != nil {
		return err8
	}

	/* Determine area */
	areaName := msgBody.GetArea()
	if areaName == "" {
		return self.processNewDirectMessage(msgHeader, msgBody)
	} else {
		return self.processNewEchoMessage(msgHeader, msgBody)
	}

	return nil
}

func (self *Tosser) processNewDirectMessage(msgHeader *packet.PacketMessageHeader, msgBody *packet.MessageBody) error {

	log.Printf("direct: msg = %+v", msgHeader)

	/* Decode headers */
	newSubject, err1 := self.CharsetManager.DecodeText(msgHeader.Subject)
	if err1 != nil {
		return err1
	}
	newFrom, err2 := self.CharsetManager.DecodeText(msgHeader.FromUserName)
	if err2 != nil {
		return err2
	}
	newTo, err3 := self.CharsetManager.DecodeText(msgHeader.ToUserName)
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

	/* Decode message body */
	newBody, err9 := self.decodeMessageBody(msgBody.RAW, charset)
	if err9 != nil {
		return err9
	}

	newMsg.SetContent(newBody)

	/* Write message */
	err := self.NetmailManager.Write(newMsg)
	log.Printf("err = %v", err)

	return nil

}

func (self *Tosser) processNewEchoMessage(msgHeader *packet.PacketMessageHeader, msgBody *packet.MessageBody) error {

	areaName := msgBody.GetArea()

	/* Auto create area */
	a := area.NewArea()
	a.SetName(areaName)
	self.AreaManager.Register(a)

	/* Decode message */
	charsetKludge := msgBody.GetKludge("CHRS", "CP866 2")
	charset := self.parseCharsetKludge(charsetKludge)

	/* Decode message body */
	newBody, err9 := self.decodeMessageBody(msgBody.RAW, charset)
	if err9 != nil {
		return err9
	}

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

	/* Determine reply */
	reply := msgBody.GetKludge("REPLY", "")
	reply = strings.Trim(reply, " ")
	log.Printf("reply = %s", reply)

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

	/* Check unique item */
	exists, err9 := self.MessageManager.IsMessageExistsByHash(areaName, msgHash)
	if err9 != nil {
		return err9
	}
	if exists {
		log.Printf("Message %s already exists", msgHash)
		self.StatManager.RegisterDupe()
		return nil
	}

	/* Decode headers */
	newSubject, err1 := self.CharsetManager.DecodeText(msgHeader.Subject)
	if err1 != nil {
		return err1
	}
	newFrom, err2 := self.CharsetManager.DecodeText(msgHeader.FromUserName)
	if err2 != nil {
		return err2
	}
	newTo, err3 := self.CharsetManager.DecodeText(msgHeader.ToUserName)
	if err3 != nil {
		return err3
	}

	/* Create msgapi.Message */
	newMsg := new(msg.Message)
	newMsg.SetMsgID(msgid)
	newMsg.SetMsgHash(msgHash)
	newMsg.SetArea(areaName)
	newMsg.SetFrom(string(newFrom))
	newMsg.SetTo(string(newTo))
	newMsg.SetSubject(string(newSubject))
	newMsg.SetTime(msgTime)

	newMsg.SetContent(newBody)

	/* Store message */
	err := self.MessageManager.Write(newMsg)
	log.Printf("err = %v", err)

	/* Update counter */
	self.StatManager.RegisterInMessage()

	return nil
}

func (self *Tosser) ProcessPacket(name string) error {

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
		self.processNewMessage(msgHeader, rawBody)

		/* Update message count */
		msgCount += 1

	}

	/* Show summary */
	log.Printf("Toss area message: %d", msgCount)

	return nil
}

func (self *Tosser) processNetmail(item *mailer.MailerInboundRec) error {

	/**/
	inbTempPath, err0 := self.SetupManager.Get("main", "TempInbound", "")
	if err0 != nil {
		panic(err0)
	}

	self.StatManager.RegisterInPacket()

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

func (self *Tosser) processARCmail(item *mailer.MailerInboundRec) error {

	/**/
	newInbTempDir, err0 := self.SetupManager.Get("main", "TempInbound", "")
	if err0 != nil {
		panic(err0)
	}

	/* Unpack */
	packets, err1 := Unpack(item.AbsolutePath, newInbTempDir)
	if err1 != nil {
		return err1
	}

	for _, p := range packets {
		log.Printf("-- Process FTN packets %+v", packets)

		/* Register packet */
		if err := self.StatManager.RegisterInPacket(); err != nil {
			log.Printf("Fail on RegisterInPacket: err = %+v", err)
		}

		/**/
		err3 := self.ProcessPacket(p)
		log.Printf("error durng parse package: err = %+v", err3)
	}

	newInbTempDir, err3 := self.SetupManager.Get("main", "TempInbound", "")
	if err3 != nil {
		panic(err3)
	}

	/* Construct new path */
	newArcPath := path.Join(newInbTempDir, item.Name)

	/* Move in area*/
	log.Printf("Move %s -> %s", item.AbsolutePath, newArcPath)
	os.Rename(item.AbsolutePath, newArcPath)

	return nil

}

func (self *Tosser) processTICmail(item *mailer.MailerInboundRec) (error) {

	/* Parse */
	newTicParser := file.NewTicParser(self.CharsetManager)
	tic, err1 := newTicParser.ParseFile(item.AbsolutePath)
	if err1 != nil {
		return err1
	}
	log.Printf("tic = %+v", tic)

	/* Search area */
	fa, err1 := self.FileManager.GetAreaByName(tic.Area)
	if err1 != nil {
		panic(err1)
	}

	/* Prepare area directory */
	boxBasePath, err2 := self.SetupManager.Get("main", "FileBox", "")
	if err2 != nil {
		panic(err2)
	}
	inboxBasePath, err3 := self.SetupManager.Get("main", "Inbound", "")
	if err3 != nil {
		panic(err3)
	}
	areaLocation := path.Join(boxBasePath, tic.Area)
	os.MkdirAll(areaLocation, 0755)

	/* Create area */
	if fa == nil {
		/* Prepare area */
		newFa := file.NewFileArea()
		newFa.Name = tic.Area
		newFa.Path = areaLocation
		/* Create area */
		if err := self.FileManager.CreateArea(newFa); err != nil {
			log.Printf("Unable to create area %s: err = %+v", tic.Area, err)
		}
	}

	inboxTicLocation := path.Join(inboxBasePath, tic.File)

	areaFileLocation := path.Join(areaLocation, tic.File)
	log.Printf("Area path %+v", areaLocation)

	/* Move in area*/
	os.Rename(inboxTicLocation, areaFileLocation)

	/* Register file */
	self.FileManager.RegisterFile(tic)

	/* Register status */
	self.StatManager.RegisterInFile(tic.File)

	/* Remove TIC descriptor */
	areaTicLocation := path.Join(areaLocation, item.Name)
	os.Rename(item.AbsolutePath, areaTicLocation)

	return nil
}

func (self *Tosser) ProcessInbound() error {

	log.Printf("ProcessInbound")

	/* New mailer inbound */
	mi := mailer.NewMailerInbound(self.SetupManager)

	/* Scan inbound */
	items, err2 := mi.Scan()
	if err2 != nil {
		return err2
	}
	log.Printf("items = %+v", items)

	for _, item := range items {
		if item.Type == mailer.TypeNetmail {
			log.Printf(" - Found Netmail packet %s. Skip.", item.Name)
			self.processNetmail(item)
		} else if item.Type == mailer.TypeARCmail {
			log.Printf(" - Found ARCmail packet %s. Process.", item.Name)
			self.processARCmail(item)
		} else if item.Type == mailer.TypeTICmail {
			log.Printf(" - Found TIC packet %s. Process.", item.Name)
			self.processTICmail(item)
		} else {
			log.Printf(" - Found other packet %s. Skip.", item.Name)
		}
	}

	return nil

}
