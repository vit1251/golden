package tosser

import (
	"bufio"
	"bytes"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/vit1251/golden/pkg/charset"
	cmn "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/eventbus"
	"github.com/vit1251/golden/pkg/fidotime"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/packet"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/tmpl"
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
	event          chan bool
	registry      *registry.Container
}

func NewTosserManager(registry *registry.Container) *TosserManager {

	tm := new(TosserManager)
	tm.registry = registry

	tm.event = make(chan bool)

	eventBus := tm.restoreEventBus()
	eventBus.Register(tm)

	return tm
}

func (self *TosserManager) HandleEvent(event string ) {
	log.Printf("Tosser event receive")
	if event == "Tosser" {
		if self.event != nil {
			self.event <- true
		}
	}
}

func (self *TosserManager) restoreEventBus() *eventbus.EventBus {
	managerPtr := self.registry.Get("EventBus")
	if manager, ok := managerPtr.(*eventbus.EventBus); ok {
		return manager
	} else {
		panic("no event bus")
	}
}

func (self *TosserManager) Start() {
	go self.run()
}

func (self *TosserManager) processTosser() {
	newTosser := NewTosser(self.registry)
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
	ZHour, Zmin := offset / 3600, offset % 3600
	var newZone string
	if sign == "+" {
		newZone = fmt.Sprintf("%02d%02d", ZHour, Zmin)
	} else if sign == "-" {
		newZone = fmt.Sprintf("-%02d%02d", ZHour, Zmin)
	}
	log.Printf("zone = %s", newZone)
	return newZone
}

/// ORIGIN LENGTH 79 http://ftsc.org/docs/fsc-0068.001
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

	return  result
}

func (self *TosserManager) makePacketEchoMessage(em *EchoMessage) (string, error) {

	mapperManager := self.restoreMapperManager()
	configMapper := mapperManager.GetConfigMapper()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	charsetManager := self.restoreCharsetManager()

	/* Create packet name */
	pktPassword, _ := configMapper.Get("main", "Password")

	//
	tempOutbound := cmn.GetTempOutboundDirectory()
	packetName := cmn.MakePacketName()
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

	/* Ask source address */
	myAddr, _ := configMapper.Get("main", "Address")
	bossAddr, _ := configMapper.Get("main", "Link")
	realName, _ := configMapper.Get("main", "RealName")
	TearLine, _ := configMapper.Get("main", "TearLine")
	Origin, _ := configMapper.Get("main", "Origin")

	/* Restore area */
	area, err2 := echoAreaMapper.GetAreaByName(em.AreaName)
	if err2 != nil {
		return "", err2
	}
	log.Printf("area = %+v", area)
	msgCharset := area.Charset
	if msgCharset == "" {
		msgCharset = "CP866"
	}

	/* Write packet header */
	pktHeader := packet.NewPacketHeader()
	pktHeader.SetOrigAddr(myAddr)
	pktHeader.SetDestAddr(bossAddr)
	pktHeader.SetPassword(pktPassword)

	if err := pw.WritePacketHeader(pktHeader); err != nil {
		return "", err
	}

	/* Prepare origin */
	Origin = self.prepareOrigin(Origin)

	/* Prepare new message */
	t := tmpl.NewTemplate()
	newTearLine, _ := t.Render(TearLine)
	newOrigin, _ := t.Render(Origin)
	newTID, _ := t.Render("Golden/{GOLDEN_PLATFORM} {GOLDEN_VERSION} {GOLDEN_RELEASE_DATE} ({GOLDEN_RELEASE_HASH})")

	/* Append */
	em.body += msg.CR
	em.body += fmt.Sprintf("--- %s", newTearLine) + msg.CR
	em.body += fmt.Sprintf(" * Origin: %s (%s)", newOrigin, myAddr) + msg.CR

	/* Encode message headers */
	newSubject, err1 := charsetManager.EncodeMessageBody([]rune(em.Subject), msgCharset)
	if err1 != nil {
		return "", err1
	}
	newTo, err2 := charsetManager.EncodeMessageBody([]rune(em.To), msgCharset)
	if err2 != nil {
		return "", err2
	}
	newFrom, err3 := charsetManager.EncodeMessageBody([]rune(realName), msgCharset)
	if err3 != nil {
		return "", err3
	}
	newBody, err4 := charsetManager.EncodeMessageBody([]rune(em.GetBody()), msgCharset)
	if err4 != nil {
		return "", err3
	}

	/* Prepare packet message */
	msgHeader := packet.NewPacketMessageHeader()
	msgHeader.OrigAddr.SetAddr(myAddr)
	msgHeader.DestAddr.SetAddr(bossAddr)
	msgHeader.SetToUserName(newTo)
	msgHeader.SetFromUserName(newFrom)
	msgHeader.SetSubject(newSubject)
	var now *fidotime.FidoDate = fidotime.NewFidoDate()
	now.SetNow()
	msgHeader.SetTime(now)

	if err := pw.WriteMessageHeader(msgHeader); err != nil {
		return "", err
	}

	newZone := self.makeTimeZone()

	/* Write message body */
	msgBody := packet.NewMessageBody()
	//
	msgBody.SetArea(em.AreaName)
	//
	msgBody.AddKludge(packet.Kludge{
		Name: "TZUTC",
		Value: newZone,
		Raw: []byte(fmt.Sprintf("\x01TZUTC %s", newZone)),
	})
	chrsKludge := self.makeChrsKludgeByCharsetName(area.Charset)
	msgBody.AddKludge(packet.Kludge{
		Name: "CHRS",
		Value: chrsKludge,
		Raw: []byte(fmt.Sprintf("\x01CHRS: %s", chrsKludge)),
	})
	msgIdValue := fmt.Sprintf("%s %s", myAddr, makeCRC32(newBody))
	msgBody.AddKludge(packet.Kludge{
		Name: "MSGID",
		Value: msgIdValue,
		Raw: []byte(fmt.Sprintf("\x01MSGID: %s", msgIdValue)),
	})
	uuidValue := fmt.Sprintf("%s", makeUUID())
	msgBody.AddKludge(packet.Kludge{
		Name: "UUID",
		Value: uuidValue,
		Raw: []byte(fmt.Sprintf("\x01UUID: %s", uuidValue)),
	})
	msgBody.AddKludge(packet.Kludge{
		Name: "TID",
		Value: newTID,
		Raw: []byte(fmt.Sprintf("\x01TID: %s", newTID)),
	})

	if em.Reply != "" {
		msgBody.AddKludge(packet.Kludge{
			Name: "REPLY",
			Value: em.Reply,
			Raw: []byte(fmt.Sprintf("\x01REPLY: %s", em.Reply)),
		})
	}

	msgBody.SetRaw(newBody)

	if err5 := pw.WriteMessage(msgBody); err5 != nil {
		return "", err5
	}

	/* Write complete bytes */
	if err := pw.WritePacketEnd(); err != nil {
		return "", err
	}

	return packetName, nil

}

func (self *TosserManager) WriteEchoMessage(em *EchoMessage) error {

	inbound := cmn.GetInboundDirectory()
	outbound := cmn.GetOutboundDirectory()
	tempOutbound := cmn.GetTempOutboundDirectory()

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

	mapperManager := self.restoreMapperManager()
	configMapper := mapperManager.GetConfigMapper()
	charsetManager := self.restoreCharsetManager()

	var From string
	var FromName string
	var TearLine string
	var Origin string

	/* Create packet name */
	From, _ = configMapper.Get("main", "Address")
	FromName, _ = configMapper.Get("main", "RealName")
	pktPassword, _ := configMapper.Get("main", "Password")
	TearLine, _ = configMapper.Get("main", "TearLine")

	origin, _ := configMapper.Get("main", "Origin")
	origin1 := self.prepareOrigin(origin)
	Origin = origin1

	/* Create packet name */
	pktName := cmn.MakePacketName()
	outboundDirectory := cmn.GetOutboundDirectory()
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
	msgCharset := "CP866"
	newSubject, err1 := charsetManager.EncodeMessageBody([]rune(nm.Subject), msgCharset)
	if err1 != nil {
		return err1
	}
	newTo, err2 := charsetManager.EncodeMessageBody([]rune(nm.To), msgCharset)
	if err2 != nil {
		return err2
	}
	newFrom, err3 := charsetManager.EncodeMessageBody([]rune(FromName), msgCharset)
	if err3 != nil {
		return err3
	}
	newBody, err4 := charsetManager.EncodeMessageBody([]rune(nm.GetBody()), msgCharset)
	if err4 != nil {
		return err4
	}

	/* Prepare packet message */
	msgHeader := packet.NewPacketMessageHeader()
	msgHeader.OrigAddr.SetAddr(From)
	msgHeader.DestAddr.SetAddr(nm.ToAddr)
	msgHeader.SetToUserName(newTo)
	msgHeader.SetFromUserName(newFrom)
	msgHeader.SetSubject(newSubject)
	msgHeader.SetAttribute(packet.MsgAttrPrivate)

	msgTime := fidotime.NewFidoDate()
	msgTime.SetNow()

	msgHeader.SetTime(msgTime)

	if err := pw.WriteMessageHeader(msgHeader); err != nil {
		return err
	}

	/* Write message body */
	msgBody := packet.NewMessageBody()

	/* Cross network NETMAIL */
	destinationAddress := fmt.Sprintf("%d:%d/%d", msgHeader.DestAddr.Zone, msgHeader.DestAddr.Net, msgHeader.DestAddr.Node)
	originAddress := fmt.Sprintf("%d:%d/%d", msgHeader.OrigAddr.Zone, msgHeader.OrigAddr.Net,  msgHeader.OrigAddr.Node)

	/* Control paragraph write section */
	intlKludge := fmt.Sprintf("%s %s", destinationAddress, originAddress)
	msgBody.AddKludge(packet.Kludge{
		Name: "INTL",
		Value: intlKludge,
		Raw: []byte(fmt.Sprintf("\x01INTL %s", intlKludge)),
	})
	fmptKludge := fmt.Sprintf("%d", msgHeader.OrigAddr.Point)
	msgBody.AddKludge(packet.Kludge{
		Name: "FMPT",
		Value: fmptKludge,
		Raw: []byte(fmt.Sprintf("\x01FMPT %s", fmptKludge)),
	})
	toptKludge := fmt.Sprintf("%d", msgHeader.DestAddr.Point)
	msgBody.AddKludge(packet.Kludge{
		Name: "TOPT",
		Value: toptKludge,
		Raw: []byte(fmt.Sprintf("\x01TOPT %s", toptKludge)),
	})
	chrsKludge := self.makeChrsKludgeByCharsetName(msgCharset)
	msgBody.AddKludge(packet.Kludge{
		Name: "CHRS",
		Value: chrsKludge,
		Raw: []byte(fmt.Sprintf("\x01CHRS: %s", chrsKludge)),
	})
	msgIdKludge := fmt.Sprintf("%s %s", From, makeCRC32(newBody))
	msgBody.AddKludge(packet.Kludge{
		Name: "MSGID",
		Value: msgIdKludge,
		Raw: []byte(fmt.Sprintf("\x01MSGID: %s", msgIdKludge)),
	})
	uuidKludge := fmt.Sprintf("%s", makeUUID())
	msgBody.AddKludge(packet.Kludge{
		Name: "UUID",
		Value: uuidKludge,
		Raw: []byte(fmt.Sprintf("\x01UUID: %s", uuidKludge)),
	})
	msgBody.AddKludge(packet.Kludge{
		Name: "TID",
		Value: newTID,
		Raw: []byte(fmt.Sprintf("\x01TID: %s", newTID)),
	})

	/* Set message body */
	msgBody.SetRaw(newBody)

	/* Write message in packet */
	if err := pw.WriteMessage(msgBody); err != nil {
		return err
	}

	/* Write complete bytes */
	if err := pw.WritePacketEnd(); err != nil {
		return err
	}

	return nil
}

func (self *TosserManager) restoreCharsetManager() *charset.CharsetManager {
	managerPtr := self.registry.Get("CharsetManager")
	if manager, ok := managerPtr.(*charset.CharsetManager); ok {
		return manager
	} else {
		panic("no charset manager")
	}
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

func (self TosserManager) restoreMapperManager() *mapper.MapperManager {
	managerPtr := self.registry.Get("MapperManager")
	if manager, ok := managerPtr.(*mapper.MapperManager); ok {
		return manager
	} else {
		panic("no mapper manager")
	}
}

func makeUUID() string {
	u1 := uuid.NewV4()
	//	u1, err4 := uuid.NewV4()
	//	if err4 != nil {
	//		return err4
	//	}
	return u1.String()
}

func makeCRC32(rawMsg []byte) string {
	h := crc32.NewIEEE()
	h.Write(rawMsg)
	hs := h.Sum32()
	return fmt.Sprintf("%08X", hs)
}
