package mailer

import (
	"fmt"
	cmn "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/config"
	"github.com/vit1251/golden/pkg/eventbus"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/queue"
	"github.com/vit1251/golden/pkg/registry"
	"log"
	"strconv"
	"time"
)

type MailerManager struct {
	registry.Service
	event         chan bool
	procIteration int
	running       bool
	mailerActive  bool
}

func NewMailerManager(r *registry.Container) *MailerManager {
	/* Step 1. Create new EventBus instance */
	newMailerManager := new(MailerManager)
	newMailerManager.SetRegistry(r)
	newMailerManager.event = make(chan bool)
	newMailerManager.mailerActive = false
	/* Step 2. Register mailer manager in EventBus */
	eventBus := eventbus.RestoreEventBus(r)
	eventBus.Register(newMailerManager)
	//
	return newMailerManager
}

func (self *MailerManager) HandleEvent(event string) {
	if event == "Mailer" {
		if self.event != nil {
			self.event <- true
		}
	}
}

func (self *MailerManager) Start() {
	log.Printf("MailerManager: Start")
	go self.run()
	go self.processTimer()
}

func (self *MailerManager) GetMailerInterval() int {

	configManager := config.RestoreConfigManager(self.GetRegistry())

	newConfig := configManager.GetConfig()

	mailerInt, _ := strconv.ParseInt(newConfig.Mailer.Interval, 10, 32)

	/* Minimum 5 minute */
	if mailerInt > 0 {
		if mailerInt < 5 {
			mailerInt = 5
		}
	}

	return int(mailerInt)

}

func (self *MailerManager) IsAutoMailer() bool {
	var result bool = true
	if self.GetMailerInterval() == 0 {
		result = false
	}
	if self.mailerActive {
		log.Printf("Mailer already in progress. Skip.")
		result = false
	}
	return result
}

func (self *MailerManager) processTimer() {
	log.Printf(" * Mailer timer routine start")
	self.running = true
	for self.running {
		autoMailer := self.IsAutoMailer()
		if autoMailer {
			log.Printf(" * Mailer auto event")
			if self.event != nil {
				self.event <- true
			}
		}
		/* Wait next iteration */
		mailerInt := self.GetMailerInterval()
		if mailerInt == 0 {
			time.Sleep(1 * time.Minute) // This loop only for support user change 0 -> positive value in setup without reloading process.
		} else {
			log.Printf("Wait %d minute before next call", mailerInt)
			time.Sleep(time.Duration(mailerInt) * time.Minute)
		}
	}
	log.Printf(" * Mailer timer routine stop")
}

func (self *MailerManager) run() {
	log.Printf(" * Mailer routine start")
	for range self.event {
		self.procIteration += 1
		log.Printf(" * Mailer start (%d)", self.procIteration)
		self.mailerActive = true
		if err := self.processMailer(); err != nil {
			log.Printf("err = %+v", err)
		}
		self.mailerActive = false
		log.Printf(" * Mailer complete (%d)", self.procIteration)
	}
	log.Printf(" * Mailer routine stop")
}

func (self *MailerManager) Stop() {
	self.running = false
	close(self.event)
}

func (self *MailerManager) processMailer() error {

	queueManager := queue.RestoreQueueManager(self.GetRegistry())

	var stat mapper.StatMailer

	log.Printf("MailerManager: processMailer")

	configManager := config.RestoreConfigManager(self.GetRegistry())
	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
	statMailerMapper := mapperManager.GetStatMailerMapper()

	eventBus := eventbus.RestoreEventBus(self.GetRegistry())

	/* Directory */
	inb := cmn.GetInboundDirectory()
	outb := cmn.GetOutboundDirectory()
	TempOutbound := cmn.GetTempOutboundDirectory()
	TempInbound := cmn.GetTempInboundDirectory()
	Temp := cmn.GetTempDirectory()

	/* Construct node address */
	newConfig := configManager.GetConfig()

	/* Get parameters */
	m := NewMailer(self.GetRegistry())
	m.SetTempOutbound(TempOutbound)
	m.SetTempInbound(TempInbound)
	m.SetTemp(Temp)
	m.SetServerAddr(newConfig.Main.NetAddr)
	m.SetInboundDirectory(inb)
	m.SetOutboundDirectory(outb)
	m.SetAddr(fmt.Sprintf("%s@fidonet", newConfig.Main.Address))
	m.SetSecret(newConfig.Main.Password)
	m.SetUserName(newConfig.Main.RealName)
	m.SetStationName(newConfig.Main.StationName)
	if newConfig.Main.City != "" && newConfig.Main.Country != "" {
		m.SetLocation(fmt.Sprintf("%s, %s", newConfig.Main.City, newConfig.Main.Country))
	}

	/* Populate outbound queue */
	mo := queueManager.GetMailerOutbound()
	items, err2 := mo.GetItems()
	if err2 != nil {
		return nil
	}
	for _, item := range items {
		m.AddOutbound(item)
	}

	/* Start mailer */
	var mailerReport *MailerReport
	log.Printf("--- Mailer start ---")
	err3, mailerReport := m.Start()
	if err3 != nil {
		log.Printf("--- Mailer error: msg = %q ---", err3)
		return err3
	}

	/* Update mailer report */
	convertMailerReportToStatMailer(&stat, mailerReport)
	if err := statMailerMapper.UpdateSession(&stat); err != nil {
		log.Printf("Fail on mailer routine: err = %+v", err)
	}

	/* Wait mailer complete */
	mailerReport = m.Wait()
	log.Printf("Mailer: report = %#v", mailerReport)
	log.Printf("--- Mailer complete ---")

	/* Update mailer report */
	convertMailerReportToStatMailer(&stat, mailerReport)
	if err := statMailerMapper.UpdateSession(&stat); err != nil {
		log.Printf("Fail on mailer routine: err = %+v", err)
	}

	/* Toss new message */
	newTosserEvent := eventBus.CreateEvent("Tosser")
	eventBus.FireEvent(newTosserEvent)

	/* Tracker new items */
	newTrackerEvent := eventBus.CreateEvent("Tracker")
	eventBus.FireEvent(newTrackerEvent)

	return nil
}

func convertMailerReportToStatMailer(result *mapper.StatMailer, report *MailerReport) error {
	result.Status = report.GetStatus()
	result.SessionStart = report.GetSessionStart().UnixMilli()
	result.SessionStop = report.GetSessionStop().UnixMilli()
	return nil
}
