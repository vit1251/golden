package tosser

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/vit1251/golden/pkg/charset"
	cmn "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/fidotime"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/packet"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/tmpl"
	"go.uber.org/dig"
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
	Container *dig.Container
	SetupManager *setup.ConfigManager
	MessageManager *msg.MessageManager
	CharsetManager *charset.CharsetManager
	event chan bool
}

func NewTosserManager(c *dig.Container) *TosserManager {
	tm := new(TosserManager)
	tm.Container = c
	//
	c.Invoke(func(cm *charset.CharsetManager, sm *setup.ConfigManager, mm *msg.MessageManager) {
		tm.CharsetManager = cm
		tm.SetupManager = sm
		tm.MessageManager = mm
	})
	//
	tm.checkDirectories()
	//
	tm.event = make(chan bool)
	go tm.run()
	//
	return tm
}

func (self *TosserManager) checkDirectory(cacheSection string) {

	cacheDirectory, _ := self.SetupManager.Get("main", cacheSection)
	if cacheDirectory == "" {
		log.Printf("Wrong directory: section = %+v", cacheSection)
		storageDirectory := cmn.GetStorageDirectory()
		cacheDirectory = path.Join(storageDirectory, "Fido", cacheSection)
		log.Printf("Construct new directory: section = %+v cacheDirectory = %+v", cacheSection, cacheDirectory)
		self.SetupManager.Set("main", cacheSection, cacheDirectory)
	}
	if _, err1 := os.Stat(cacheDirectory); err1 != nil {
		log.Printf("Directory check: name = %v - ERR", cacheSection)
		if os.IsNotExist(err1) {
			log.Printf("Initial create directory: path = %+v", cacheDirectory)
			os.MkdirAll(cacheDirectory, os.ModeDir|0755)
		} else {
			log.Fatal("Initernal error: err = %+v", err1)
		}
	} else {
		log.Printf("Directory check: name = %v - OK", cacheSection)
	}

}

func (self *TosserManager) checkDirectories() {

	/* Check mailer directory */
	self.checkDirectory("Inbound")
	self.checkDirectory("Outbound")
	self.checkDirectory("TempInbound")
	self.checkDirectory("TempOutbound")

	/* Check FileBox directory */
	self.checkDirectory("FileBox")

}

func (self *TosserManager) Start() {
	self.event <- true
}

func (self *TosserManager) processTosser() {
	newTosser := NewTosser(self.Container)
	newTosser.Toss()
}

func (self *TosserManager) run() {
	log.Printf(" * Tosser service start")
	var procIteration int
	tick := time.NewTicker(1 * time.Minute)
	for alive := true; alive; {
		select {
			case <-self.event:
			case <-tick.C:
				procIteration += 1
				log.Printf(" * Tosser start (%d)", procIteration)
				self.processTosser()
				log.Printf(" * Tosser complete (%d)", procIteration)
		}
	}
	log.Printf(" * Tosser service stop")
}

func (self *TosserManager) makePacketName() string {
	now := time.Now()
	unixTime := now.Unix()
	log.Printf("unixTime: dec = %d hex = %x", unixTime, unixTime)
	pktName := fmt.Sprintf("%08x.pkt", unixTime)
	log.Printf("pktName: name = %s", pktName)
	return pktName
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
	var result string = " -- No origins exist -- "
	/* Check Origin is external */
	if strings.HasPrefix(Origin, "@") {
		newPath := strings.TrimPrefix(Origin, "@")
		content, err := ioutil.ReadFile(newPath)
		if err == nil {
			newContent := string(content)
			rows := strings.Split(newContent, "\n")
			rand.Seed(time.Now().Unix())
			n := rand.Intn(len(rows))
			oneLine := rows[n]
			newOneLine := strings.Trim(oneLine, " \t\n\r")
			result = newOneLine
		}
	} else {
		result = Origin
	}
	if len(result) > 78 {
		result = result[:78]
	}
	return  result
}

func (self *TosserManager) makePacketEchoMessage(em *EchoMessage) (string, error) {

	/* Create packet name */
	tempOutbound, _ := self.SetupManager.Get("main", "TempOutbound")
	password, _ := self.SetupManager.Get("main", "Password")

	packetName := self.makePacketName()
	tempPacketPath := path.Join(tempOutbound, packetName)

	/* Open outbound packet */
	pw, err1 := packet.NewPacketWriter(tempPacketPath)
	if err1 != nil {
		return "", err1
	}
	defer pw.Close()

	/* Ask source address */
	myAddr, _ := self.SetupManager.Get("main", "Address")
	bossAddr, _ := self.SetupManager.Get("main", "Link")
	realName, _ := self.SetupManager.Get("main", "RealName")
	TearLine, _ := self.SetupManager.Get("main", "TearLine")
	Origin, _ := self.SetupManager.Get("main", "Origin")

	/* Write packet header */
	pktHeader := packet.NewPacketHeader()
	pktHeader.OrigAddr.SetAddr(myAddr)
	pktHeader.DestAddr.SetAddr(bossAddr)
	pktHeader.SetPassword(password)

	if err := pw.WritePacketHeader(pktHeader); err != nil {
		return "", err
	}

	/* Encode message headers */
	newSubject, err1 := self.CharsetManager.Encode([]rune(em.Subject))
	if err1 != nil {
		return "", err1
	}
	newTo, err2 := self.CharsetManager.Encode([]rune(em.To))
	if err2 != nil {
		return "", err2
	}
	newFrom, err3 := self.CharsetManager.Encode([]rune(realName))
	if err3 != nil {
		return "", err3
	}

	/* Prepare packet message */
	msgHeader := packet.NewPacketMessageHeader()
	msgHeader.OrigAddr.SetAddr(myAddr)
	msgHeader.DestAddr.SetAddr(bossAddr)
	msgHeader.SetAttribute(packet.PacketAttrDirect)
	msgHeader.SetToUserName(newTo)
	msgHeader.SetFromUserName(newFrom)
	msgHeader.SetSubject(newSubject)
	var now *fidotime.FidoDate = fidotime.NewFidoDate()
	now.SetNow()
	msgHeader.SetTime(now)

	if err := pw.WriteMessageHeader(msgHeader); err != nil {
		return "", err
	}

	/* Message UUID */
	u1 := uuid.NewV4()
	//	u1, err4 := uuid.NewV4()
	//	if err4 != nil {
	//		return err4
	//	}

	/* Prepare origin */
	Origin = self.prepareOrigin(Origin)

	/* Prepare new message */
	t := tmpl.NewTemplate()
	newTearLine, _ := t.Render(TearLine)
	newOrigin, _ := t.Render(Origin)
	newTID, _ := t.Render("Golden/{GOLDEN_PLATFORM} {GOLDEN_VERSION} {GOLDEN_RELEASE_DATE} ({GOLDEN_RELEASE_HASH})")

	/* Construct message content */
	msgContent := self.MessageManager.NewMessageContent()

	msgContent.SetCharset("CP866")

	msgContent.AddLine(em.GetBody())
	msgContent.AddLine("")
	msgContent.AddLine(fmt.Sprintf("--- %s", newTearLine))
	msgContent.AddLine(fmt.Sprintf(" * Origin: %s (%s)", newOrigin, myAddr))

	rawMsg := msgContent.Pack()

	/* Calculate checksumm */
	h := crc32.NewIEEE()
	h.Write(rawMsg)
	hs := h.Sum32()
	log.Printf("crc32 = %+v", hs)

	newZone := self.makeTimeZone()

	/* Write message body */
	msgBody := packet.NewMessageBody()
	//
	msgBody.SetArea(em.AreaName)
	//
	msgBody.AddKludge("TZUTC", newZone)
	//msgBody.AddKludge("CHRS", "UTF-8 4")
	msgBody.AddKludge("CHRS", "CP866 2")
	msgBody.AddKludge("MSGID", fmt.Sprintf("%s %08x", myAddr, hs))
	msgBody.AddKludge("UUID", fmt.Sprintf("%s", u1))
	msgBody.AddKludge("TID", newTID)
	if em.Reply != "" {
		msgBody.AddKludge("REPLY", em.Reply)
	}
	//
	msgBody.SetRaw(rawMsg)
	//
	if err5 := pw.WriteMessage(msgBody); err5 != nil {
		return "", err5
	}

	return packetName, nil
}

func (self *TosserManager) WriteEchoMessage(em *EchoMessage) error {

	inbound, _ := self.SetupManager.Get("main", "Inbound")
	outbound, _ := self.SetupManager.Get("main", "Outbound")
	tempOutbound, _ := self.SetupManager.Get("main", "TempOutbound")

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

	var params struct {
		Outbound string
		From string
		FromName string
		TearLine string
		Origin string
	}

	/* Create packet name */
	params.Outbound, _ = self.SetupManager.Get("main", "Outbound")
	params.From, _ = self.SetupManager.Get("main", "Address")
	params.FromName, _ = self.SetupManager.Get("main", "RealName")

	origin, _ := self.SetupManager.Get("main", "Origin")
	origin1 := self.prepareOrigin(origin)
	params.Origin = origin1

	TearLine, _ := self.SetupManager.Get("main", "TearLine")
	params.TearLine = TearLine

	/* Create packet name */
	pktName := self.makePacketName()
	name := path.Join(params.Outbound, pktName)
	log.Printf("Write Netmail packet %s", name)

	/* Open outbound packet */
	pw, err1 := packet.NewPacketWriter(name)
	if err1 != nil {
		return err1
	}
	defer pw.Close()

	/* Write packet header */
	pktHeader := packet.NewPacketHeader()
	pktHeader.OrigAddr.SetAddr(params.From)
	pktHeader.DestAddr.SetAddr(nm.ToAddr)

	if err := pw.WritePacketHeader(pktHeader); err != nil {
		return err
	}

	/* Encode message */
	newSubject, err1 := self.CharsetManager.Encode([]rune(nm.Subject))
	if err1 != nil {
		return err1
	}
	newTo, err2 := self.CharsetManager.Encode([]rune(nm.To))
	if err2 != nil {
		return err2
	}
	newFrom, err3 := self.CharsetManager.Encode([]rune(params.FromName))
	if err3 != nil {
		return err3
	}

	/* Prepare packet message */
	msgHeader := packet.NewPacketMessageHeader()
	msgHeader.OrigAddr.SetAddr(params.From)
	msgHeader.DestAddr.SetAddr(nm.ToAddr)
	msgHeader.SetAttribute(packet.PacketAttrDirect)
	msgHeader.SetToUserName(newTo)
	msgHeader.SetFromUserName(newFrom)
	msgHeader.SetSubject(newSubject)
	var now *fidotime.FidoDate = fidotime.NewFidoDate()
	now.SetNow()
	msgHeader.SetTime(now)

	if err := pw.WriteMessageHeader(msgHeader); err != nil {
		return err
	}

	/* Message UUID */
	u1 := uuid.NewV4()
	//	u1, err4 := uuid.NewV4()
	//	if err4 != nil {
	//		return err4
	//	}

	/* Prepare new message */
	t := tmpl.NewTemplate()
	newTearLine, _ := t.Render(params.TearLine)
	newOrigin, _ := t.Render(params.Origin)
	newTID, _ := t.Render("Golden/{GOLDEN_PLATFORM} {GOLDEN_VERSION} {GOLDEN_RELEASE_DATE} ({GOLDEN_RELEASE_HASH})")

	/* Construct message content */
	msgContent := self.MessageManager.NewMessageContent()
	msgContent.SetCharset("CP866")
	msgContent.AddLine(nm.GetBody())
	msgContent.AddLine("")
	msgContent.AddLine(fmt.Sprintf("--- %s", newTearLine))
	msgContent.AddLine(fmt.Sprintf(" * Origin: %s (%s)", newOrigin, params.From))
	rawMsg := msgContent.Pack()

	/* Calculate checksumm */
	h := crc32.NewIEEE()
	h.Write(rawMsg)
	hs := h.Sum32()
	log.Printf("crc32 = %+v", hs)

	/* Write message body */
	msgBody := packet.NewMessageBody()

	/* Cross network NETMAIL */
	destinationAddress := fmt.Sprintf("%d:%d/%d", msgHeader.DestAddr.Zone, msgHeader.DestAddr.Net, msgHeader.DestAddr.Node)
	originAddress := fmt.Sprintf("%d:%d/%d", msgHeader.OrigAddr.Zone, msgHeader.OrigAddr.Net,  msgHeader.OrigAddr.Node)

	msgBody.AddKludge("INTL", fmt.Sprintf("%s %s", destinationAddress, originAddress))
	msgBody.AddKludge("FMPT", fmt.Sprintf("%d", msgHeader.OrigAddr.Point))
	msgBody.AddKludge("TOPT", fmt.Sprintf("%d", msgHeader.DestAddr.Point))
	msgBody.AddKludge("CHRS", "CP866 2")
	msgBody.AddKludge("MSGID", fmt.Sprintf("%s %08x", params.From, hs))
	msgBody.AddKludge("UUID", fmt.Sprintf("%s", u1))
	msgBody.AddKludge("TID", newTID)

	/* Set message body */
	msgBody.SetRaw(rawMsg)

	/* Write message in packet */
	if err := pw.WriteMessage(msgBody); err != nil {
		return err
	}

	return nil
}
