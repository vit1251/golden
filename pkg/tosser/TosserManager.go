package tosser

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/vit1251/golden/internal/common"
	"github.com/vit1251/golden/internal/utils"
	"github.com/vit1251/golden/pkg/charset"
	"github.com/vit1251/golden/pkg/config"
	"github.com/vit1251/golden/pkg/eventbus"
	"github.com/vit1251/golden/pkg/fidotime"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/packet"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/internal/tmpl"
	"hash/crc32"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"
)

type TosserManager struct {
	registry.Service
	event chan bool
}

func NewTosserManager(registry *registry.Container) *TosserManager {

	tm := new(TosserManager)
	tm.SetRegistry(registry)

	tm.event = make(chan bool)

	/* Step 2. Register in EventBus */
	eventBus := eventbus.RestoreEventBus(registry)
	eventBus.Register(tm)

	return tm
}

func (self *TosserManager) HandleEvent(event string) {
	log.Printf("Tosser event receive")
	if event == "Tosser" {
		if self.event != nil {
			self.event <- true
		}
	}
}

func (self *TosserManager) Start() {
	go self.run()
}

func (self *TosserManager) processTosser() {
	newTosser := NewTosser(self.GetRegistry())
	newTosser.Toss()
}

func (self *TosserManager) run() {
	log.Printf(" * Tosser service start")
	var procIteration int
	for range self.event {
		log.Printf(" * Tosser start (%d)", procIteration)
		self.processTosser()
		log.Printf(" * Tosser complete (%d)", procIteration)
		procIteration += 1
	}
	log.Printf(" * Tosser service stop")
}

func (self *TosserManager) makeTimeZone() string {
	newTime := time.Now()
	_, offset := newTime.Zone()
	var sign string = "+"
	if offset < 0 {
		offset = -offset
		sign = "-"
	}
	ZHour, Zmin := offset/3600, offset%3600
	var newZone string
	if sign == "+" {
		newZone = fmt.Sprintf("%02d%02d", ZHour, Zmin)
	} else if sign == "-" {
		newZone = fmt.Sprintf("-%02d%02d", ZHour, Zmin)
	}
	log.Printf("zone = %s", newZone)
	return newZone
}

// / ORIGIN LENGTH 79 http://ftsc.org/docs/fsc-0068.001
func (self *TosserManager) prepareOrigin(Origin string) string {

	/* Set empty origin */
	result := ""

	/* Process origins notebook */
	if strings.HasPrefix(Origin, "@") {
		/* Remove @ */
		newPath := strings.TrimPrefix(Origin, "@")
		/* Reading notebook content */
		content, err := ioutil.ReadFile(newPath)
		if err == nil {
			rows := bytes.Split(content, []byte("\r"))
			rand.Seed(time.Now().Unix())
			n := rand.Intn(len(rows))
			oneLine := rows[n]
			newOneLine := bytes.Trim(oneLine, " \t\n\r")
			result = string(newOneLine)
		}
	} else {
		/* Set Origin (i.e. Origin without @ prefix) */
		result = Origin
	}

	/* Processing with origin in context UTF-8 rues */
	originRunes := []rune(result)
	if len(originRunes) >= 79 {
		result = string(originRunes[:79])
	}

	return result
}

func (self *TosserManager) makePacketEchoMessage(em *EchoMessage) (string, error) {

	configManager := config.RestoreConfigManager(self.GetRegistry())
	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	charsetManager := charset.RestoreCharsetManager(self.GetRegistry())

	newConfig := configManager.GetConfig()

	//
	tempOutbound := commonfunc.GetTempOutboundDirectory()
	packetName := commonfunc.MakePacketName()
	newPacketName := path.Join(tempOutbound, packetName)

	stream, err0 := os.Create(newPacketName)
	if err0 != nil {
		return "", err0
	}
	defer stream.Close()

	cacheStream := bufio.NewWriter(stream)

	/* Open outbound packet */
	pw, err1 := packet.NewPacketWriter(cacheStream)
	if err1 != nil {
		return "", err1
	}
	defer func() {
		cacheStream.Flush()
		stream.Close()
	}()

	/* Restore area */
	area, err2 := echoAreaMapper.GetAreaByName(em.AreaName)
	if err2 != nil {
		return "", err2
	}
	log.Printf("area = %+v", area)
	msgCharset := area.GetCharset()

	/* Write packet header */
	pktHeader := packet.NewPacketHeader()
	pktHeader.SetDate(time.Now())
	pktHeader.SetOrigAddr(newConfig.Main.Address)
	pktHeader.SetDestAddr(newConfig.Main.Link)
	pktHeader.SetPassword(newConfig.Main.Password)

	if err := pw.WritePacketHeader(pktHeader); err != nil {
		return "", err
	}

	/* Prepare origin */
	clearOrigin := self.prepareOrigin(newConfig.Main.Origin)

	/* Prepare new message */
	t := tmpl.NewTemplate()
	newTearLine, _ := t.Render(newConfig.Main.TearLine)
	newOrigin, _ := t.Render(clearOrigin)
	newTID, _ := t.Render("Golden/{GOLDEN_PLATFORM} {GOLDEN_VERSION} {GOLDEN_RELEASE_DATE} ({GOLDEN_RELEASE_HASH})")

	/* Append */
	em.body += packet.CR
	em.body += fmt.Sprintf("--- %s", newTearLine) + packet.CR
	em.body += fmt.Sprintf(" * Origin: %s (%s)", newOrigin, newConfig.Main.Address) + packet.CR

	/* Encode message headers */
	newSubject, err1 := charsetManager.EncodeMessageBody(em.Subject, msgCharset)
	if err1 != nil {
		return "", err1
	}
	newTo, err2 := charsetManager.EncodeMessageBody(em.To, msgCharset)
	if err2 != nil {
		return "", err2
	}
	newFrom, err3 := charsetManager.EncodeMessageBody(newConfig.Main.RealName, msgCharset)
	if err3 != nil {
		return "", err3
	}
	newBody, err4 := charsetManager.EncodeMessageBody(em.GetBody(), msgCharset)
	if err4 != nil {
		return "", err3
	}

	/* Prepare packet message */
	packedMessage := packet.NewPackedMessage()
	packedMessage.OrigAddr.SetAddr(newConfig.Main.Address)
	packedMessage.DestAddr.SetAddr(newConfig.Main.Link)
	packedMessage.SetToUserName(newTo)
	packedMessage.SetFromUserName(newFrom)
	packedMessage.SetSubject(newSubject)
	var now *fidotime.FidoDate = fidotime.NewFidoDate()
	now.SetNow()
	packedMessage.SetTime(now)

	newZone := self.makeTimeZone()

	/* Write message body */
	msgBody := packet.NewMessageBody()
	//
	msgBody.SetArea(em.AreaName)
	//
	msgBody.AddKludge(packet.Kludge{
		Name:  "TZUTC",
		Value: newZone,
		Raw:   []byte(fmt.Sprintf("\x01TZUTC: %s", newZone)),
	})
	chrsKludge := self.makeChrsKludgeByCharsetName(area.GetCharset())
	msgBody.AddKludge(packet.Kludge{
		Name:  "CHRS",
		Value: chrsKludge,
		Raw:   []byte(fmt.Sprintf("\x01CHRS: %s", chrsKludge)),
	})
	msgIdValue := fmt.Sprintf("%s %s", newConfig.Main.Address, makeCRC32(newBody))
	msgBody.AddKludge(packet.Kludge{
		Name:  "MSGID",
		Value: msgIdValue,
		Raw:   []byte(fmt.Sprintf("\x01MSGID: %s", msgIdValue)),
	})
	var uuidValue string = utils.IndexHelper_makeUUID()
	msgBody.AddKludge(packet.Kludge{
		Name:  "UUID",
		Value: uuidValue,
		Raw:   []byte(fmt.Sprintf("\x01UUID: %s", uuidValue)),
	})
	msgBody.AddKludge(packet.Kludge{
		Name:  "TID",
		Value: newTID,
		Raw:   []byte(fmt.Sprintf("\x01TID: %s", newTID)),
	})

	if em.Reply != "" {
		msgBody.AddKludge(packet.Kludge{
			Name:  "REPLY",
			Value: em.Reply,
			Raw:   []byte(fmt.Sprintf("\x01REPLY: %s", em.Reply)),
		})
	}

	msgBody.SetContent(newBody)

	packedMessage.SetText(msgBody.Bytes())

	/* Write packed message */
	err5 := pw.WritePackedMessage(packedMessage)
	if err5 != nil {
		return "", err5
	}

	/* Write packed end */
	err6 := pw.WritePacketEnd()
	if err6 != nil {
		return "", err6
	}

	return packetName, nil

}

func (self *TosserManager) WriteEchoMessage(em *EchoMessage) error {

	inbound := commonfunc.GetInboundDirectory()
	outbound := commonfunc.GetOutboundDirectory()
	tempOutbound := commonfunc.GetTempOutboundDirectory()

	packetName, err1 := self.makePacketEchoMessage(em)
	if err1 != nil {
		return err1
	}
	tempPacketPath := path.Join(tempOutbound, packetName)

	// Copy to Inbound and Outbound
	log.Printf("Publish packet: name = %+v", tempPacketPath)
	if err := self.PushPacket(tempPacketPath, path.Join(inbound, packetName)); err != nil {
		log.Printf("Fail on copy Inbound: err = %+v", err)
	}
	if err := self.PushPacket(tempPacketPath, path.Join(outbound, packetName)); err != nil {
		log.Printf("Fail on copy Outbound: err = %+v", err)
	}

	return nil
}

func (self *TosserManager) PushPacket(src string, dst string) error {

	log.Printf("Publish packet: %s -> %s", src, dst)

	source, err1 := os.Open(src)
	if err1 != nil {
		return err1
	}
	defer source.Close()

	destination, err2 := os.Create(dst)
	if err2 != nil {
		return err2
	}
	defer destination.Close()

	nBytes, err3 := io.Copy(destination, source)
	log.Printf("Copy %+v", nBytes)

	return err3
}

func (self *TosserManager) WriteNetmailMessage(nm *NetmailMessage) error {

	configManager := config.RestoreConfigManager(self.GetRegistry())
	charsetManager := charset.RestoreCharsetManager(self.GetRegistry())

	newConfig := configManager.GetConfig()

	From := newConfig.Main.Address
	FromName := newConfig.Main.RealName
	TearLine := newConfig.Main.TearLine
	Origin := newConfig.Main.Origin
	pktPassword := newConfig.Main.Password
	origin := newConfig.Main.Origin
	charset := newConfig.Netmail.Charset

	cleanOrigin := self.prepareOrigin(origin)
	Origin = cleanOrigin

	/* Create packet name */
	pktName := commonfunc.MakePacketName()
	outboundDirectory := commonfunc.GetOutboundDirectory()
	name := path.Join(outboundDirectory, pktName)
	log.Printf("Write Netmail packet %s", name)

	/* Create write stream */
	stream, err0 := os.Create(name)
	if err0 != nil {
		return err0
	}

	cacheStream := bufio.NewWriter(stream)

	defer func() {
		cacheStream.Flush()
		stream.Close()
	}()

	/* Open outbound packet */
	pw, err1 := packet.NewPacketWriter(cacheStream)
	if err1 != nil {
		return err1
	}

	/* Write packet header */
	pktHeader := packet.NewPacketHeader()
	pktHeader.SetDate(time.Now())
	pktHeader.SetOrigAddr(From)
	pktHeader.SetDestAddr(nm.ToAddr)
	pktHeader.SetPassword(pktPassword)

	if err := pw.WritePacketHeader(pktHeader); err != nil {
		return err
	}

	/* Prepare Origin */
	t := tmpl.NewTemplate()
	newTearLine, _ := t.Render(TearLine)
	newOrigin, _ := t.Render(Origin)
	newTID, _ := t.Render("Golden/{GOLDEN_PLATFORM} {GOLDEN_VERSION} {GOLDEN_RELEASE_DATE} ({GOLDEN_RELEASE_HASH})")

	nm.body += packet.CR
	nm.body += fmt.Sprintf("--- %s", newTearLine) + packet.CR
	nm.body += fmt.Sprintf(" * Origin: %s (%s)", newOrigin, From) + packet.CR

	/* Encode message */
	var msgCharset string = charset // TODO - check valid charset here ...
	newSubject, err1 := charsetManager.EncodeMessageBody(nm.Subject, msgCharset)
	if err1 != nil {
		return err1
	}
	newTo, err2 := charsetManager.EncodeMessageBody(nm.To, msgCharset)
	if err2 != nil {
		return err2
	}
	newFrom, err3 := charsetManager.EncodeMessageBody(FromName, msgCharset)
	if err3 != nil {
		return err3
	}
	newBody, err4 := charsetManager.EncodeMessageBody(nm.GetBody(), msgCharset)
	if err4 != nil {
		return err4
	}

	/* Prepare packet message */
	packedMessage := packet.NewPackedMessage()
	packedMessage.OrigAddr.SetAddr(From)
	packedMessage.DestAddr.SetAddr(nm.ToAddr)
	packedMessage.SetToUserName(newTo)
	packedMessage.SetFromUserName(newFrom)
	packedMessage.SetSubject(newSubject)
	packedMessage.SetAttribute(packet.MsgAttrPrivate)

	msgTime := fidotime.NewFidoDate()
	msgTime.SetNow()

	packedMessage.SetTime(msgTime)

	/* Write message body */
	msgBody := packet.NewMessageBody()

	/* Cross network NETMAIL */
	origAddr := fmt.Sprintf("%d:%d/%d", packedMessage.OrigAddr.Zone, packedMessage.OrigAddr.Net, packedMessage.OrigAddr.Node)
	destAddr := fmt.Sprintf("%d:%d/%d", packedMessage.DestAddr.Zone, packedMessage.DestAddr.Net, packedMessage.DestAddr.Node)

	log.Printf("Direct message: %+v -> %+v", origAddr, destAddr)

	newZone := self.makeTimeZone()

	/* Control paragraph write section */
	msgBody.AddKludge(packet.Kludge{
		Name:  "TZUTC",
		Value: newZone,
		Raw:   []byte(fmt.Sprintf("\x01TZUTC: %s", newZone)),
	})
	intlKludge := fmt.Sprintf("%s %s", destAddr, origAddr)
	msgBody.AddKludge(packet.Kludge{
		Name:  "INTL",
		Value: intlKludge,
		Raw:   []byte(fmt.Sprintf("\x01INTL %s", intlKludge)),
	})
	if packedMessage.OrigAddr.Point != 0 {
		fmptKludge := fmt.Sprintf("%d", packedMessage.OrigAddr.Point)
		msgBody.AddKludge(packet.Kludge{
			Name:  "FMPT",
			Value: fmptKludge,
			Raw:   []byte(fmt.Sprintf("\x01FMPT %s", fmptKludge)),
		})
	}
	if packedMessage.DestAddr.Point != 0 {
		toptKludge := fmt.Sprintf("%d", packedMessage.DestAddr.Point)
		msgBody.AddKludge(packet.Kludge{
			Name:  "TOPT",
			Value: toptKludge,
			Raw:   []byte(fmt.Sprintf("\x01TOPT %s", toptKludge)),
		})
	}
	chrsKludge := self.makeChrsKludgeByCharsetName(msgCharset)
	msgBody.AddKludge(packet.Kludge{
		Name:  "CHRS",
		Value: chrsKludge,
		Raw:   []byte(fmt.Sprintf("\x01CHRS: %s", chrsKludge)),
	})
	msgIdKludge := fmt.Sprintf("%s %s", From, makeCRC32(newBody))
	msgBody.AddKludge(packet.Kludge{
		Name:  "MSGID",
		Value: msgIdKludge,
		Raw:   []byte(fmt.Sprintf("\x01MSGID: %s", msgIdKludge)),
	})
	var uuidKludge string = utils.IndexHelper_makeUUID()
	msgBody.AddKludge(packet.Kludge{
		Name:  "UUID",
		Value: uuidKludge,
		Raw:   []byte(fmt.Sprintf("\x01UUID: %s", uuidKludge)),
	})
	msgBody.AddKludge(packet.Kludge{
		Name:  "TID",
		Value: newTID,
		Raw:   []byte(fmt.Sprintf("\x01TID: %s", newTID)),
	})

	/* Set message body */
	msgBody.SetContent(newBody)

	packedMessage.SetText(msgBody.Bytes())

	/* Write message in packet */
	err5 := pw.WritePackedMessage(packedMessage)
	if err5 != nil {
		return err5
	}

	/* Write complete bytes */
	err6 := pw.WritePacketEnd()
	if err6 != nil {
		return err6
	}

	return nil
}

func (self *TosserManager) makeChrsKludgeByCharsetName(charset string) string {
	if charset == "UTF-8" {
		return "UTF-8 4"
	} else if charset == "CP866" {
		return "CP866 2"
	} else {
		return "CP866 2"
	}
}

func makeCRC32(rawMsg []byte) string {
	h := crc32.NewIEEE()
	h.Write(rawMsg)
	hs := h.Sum32()
	return fmt.Sprintf("%08X", hs)
}
