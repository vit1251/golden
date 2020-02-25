package tosser

import (
	"errors"
	"github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/mailer"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/packet"
	"github.com/vit1251/golden/pkg/timezone"
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

		if unicodeBody, err1 := packet.DecodeText(msgBody); err1 == nil {
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

		/* Search area */
		// TODO - a, _ := self.AreaManager.GetAreaByName(areaName)

		/* Auto create area */
		a := area.NewArea()
		a.Name = strings.ToUpper(areaName)
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
		zParser := timezone.NewTimeZoneParser()
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
		exists, err9 := self.MessageManager.IsMessageExistsByHash(areaName, msgHash)
		if err9 != nil {
			return err9
		}
		if exists {

			log.Printf("Message %s already exists", msgHash)
			self.StatManager.RegisterDupe()

		} else {

			/* Create msgapi.Message */
			newMsg := new(msg.Message)
			newMsg.SetMsgID(msgHash)
			newMsg.SetArea(areaName)
			newMsg.SetFrom(msgHeader.FromUserName)
			newMsg.SetTo(msgHeader.ToUserName)
			newMsg.SetSubject(msgHeader.Subject)
			newMsg.SetTime(msgTime)

			newMsg.SetContent(newBody)

			/* Store message */
			err := self.MessageManager.Write(newMsg)
			log.Printf("err = %v", err)

			/* Update counter */
			self.StatManager.RegisterMessage()
			msgCount += 1
		}

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

	self.StatManager.RegisterPacket(item.AbsolutePath)

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
	self.StatManager.RegisterARCmail(item.AbsolutePath)

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

		/**/
		err2 := self.StatManager.RegisterPacket(p)
		log.Printf("error durng report stat package: err = %+v", err2)

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
	newTicParser := file.NewTicParser()
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
	self.StatManager.RegisterFile(tic.File)

	/* Remove TIC descriptor */
	areaTicLocation := path.Join(areaLocation, item.Name)
	os.Rename(item.AbsolutePath, areaTicLocation)

	return nil
}

func (self *Tosser) ProcessInbound() error {

	/* New mailer inbound */
	mi := mailer.NewMailerInbound(self.SetupManager)


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
			log.Printf(" - Found TIC packet %s. Process.", item.Name)
			self.processTICmail(item)
		} else {
			log.Printf(" - Found other packet %s. Skip.", item.Name)
		}
	}

	return nil

}
