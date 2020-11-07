package mailer

import (
	"fmt"
	cmn "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/eventbus"
	"github.com/vit1251/golden/pkg/mailer/cache"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/stat"
	"log"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

type MailerManager struct {
	registry      *registry.Container
	event         chan bool
	procIteration int
	running       bool
	mailerActive  bool
}

func NewMailerManager(r *registry.Container) *MailerManager {
	newMailerManager := new(MailerManager)
	newMailerManager.registry = r
	newMailerManager.event = make(chan bool)
	newMailerManager.mailerActive = false
	eventBus := newMailerManager.restoreEventBus()
	eventBus.Register(newMailerManager)
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

	configManager := self.restoreConfigManager()

	mailerIntParam, _ := configManager.Get("mailer", "Interval")

	mailerInt, _ := strconv.ParseInt(mailerIntParam, 10, 32)

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

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func (self *MailerManager) processMailer() error {

	log.Printf("MailerManager: processMailer")

	configManager := self.restoreConfigManager()
	statManager := self.restoreStatManager()
	eventBus := self.restoreEventBus()

	/* Directory */
	inb := cmn.GetInboundDirectory()
	outb := cmn.GetOutboundDirectory()
	TempOutbound := cmn.GetTempOutboundDirectory()
	TempInbound := cmn.GetTempInboundDirectory()
	Temp := cmn.GetTempDirectory()

	/* Construct node address */
	netAddr, _ := configManager.Get("main", "NetAddr")
	password, _ := configManager.Get("main", "Password")
	address, _ := configManager.Get("main", "Address")
	Country, _ := configManager.Get("main", "Country")
	City, _ := configManager.Get("main", "City")
	realName, _ := configManager.Get("main", "RealName")
	stationName, _ := configManager.Get("main", "StationName")

	/* */
	newAddress := fmt.Sprintf("%s@fidonet", address)

	/* Get parameters */
	m := NewMailer(configManager)
	m.SetTempOutbound(TempOutbound)
	m.SetTempInbound(TempInbound)
	m.SetTemp(Temp)
	m.SetServerAddr(netAddr)
	m.SetInboundDirectory(inb)
	m.SetOutboundDirectory(outb)
	m.SetAddr(newAddress)
	m.SetSecret(password)
	m.SetUserName(realName)
	m.SetStationName(stationName)
	if City != "" && Country != "" {
		m.SetLocation(fmt.Sprintf("%s, %s", City, Country))
	}

	/* Populate outbound queue */
	mo := cache.NewMailerOutbound(self.registry)
	items, err2 := mo.GetItems()
	if err2 != nil {
		return nil
	}
	for _, item := range items {
		m.AddOutbound(item)
	}

	/* Start mailer */
	log.Printf("--- Mailer start ---")
	m.Start()

	/* Wait mailer complete */
	m.Wait()
	log.Printf("--- Mailer complete ---")

	/* Complete start tosser */
	if err := statManager.RegisterOutSession(); err != nil {
		log.Printf("Fail on mailer routine: err = %+v", err)
	}

	/* Start tossing */
	eventBus.Event("Tosser")
	eventBus.Event("Tracker")

	return nil
}

func (self *MailerManager) restoreConfigManager() *setup.ConfigManager {

	managerPtr := self.registry.Get("ConfigManager")
	if manager, ok := managerPtr.(*setup.ConfigManager); ok {
		return manager
	} else {
		panic("no config manager")
	}
}

func (self *MailerManager) restoreStatManager() *stat.StatManager {

	managerPtr := self.registry.Get("StatManager")
	if manager, ok := managerPtr.(*stat.StatManager); ok {
		return manager
	} else {
		panic("no stat manager")
	}

}

func (self *MailerManager) restoreEventBus() *eventbus.EventBus {

	managerPtr := self.registry.Get("EventBus")
	if manager, ok := managerPtr.(*eventbus.EventBus); ok {
		return manager
	} else {
		panic("no eventbus manager")
	}

}

