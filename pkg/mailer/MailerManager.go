package mailer

import (
	"fmt"
	"github.com/vit1251/golden/pkg/eventbus"
	"github.com/vit1251/golden/pkg/mailer/cache"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/stat"
	"log"
	"reflect"
	"runtime"
	"time"
)

type MailerManager struct {
	registry      *registry.Container
	event         chan bool
	procIteration int
	running       bool
}

func NewMailerManager(r *registry.Container) *MailerManager {
	mm := new(MailerManager)
	mm.registry = r
	mm.event = make(chan bool)
	return mm
}

func (self *MailerManager) Start() {
	log.Printf("MailerManager: Start ")
	go self.run()
	self.running = true
}


func (self *MailerManager) run() {

	for self.running {

		self.procIteration += 1

		log.Printf(" * Mailer start (%d)", self.procIteration)
		if err := self.processMailer(); err != nil {
			log.Printf("err = %+v", err)
		}
		log.Printf(" * Mailer complete (%d)", self.procIteration)

		/* Wait 5 minute */
		time.Sleep(5 * time.Minute)

	}

}

func (self *MailerManager) Stop() {
	self.running = false
}

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func (self *MailerManager) processMailer() error {

	log.Printf("MailerManager: processMailer")

	configManager := self.restoreConfigManager()
	statManager := self.restoreStatManager()
	eventBus := self.restoreEventBus()

	/* Construct node address */
	netAddr, _ := configManager.Get("main", "NetAddr")
	password, _ := configManager.Get("main", "Password")
	address, _ := configManager.Get("main", "Address")
	inb, _ := configManager.Get("main", "Inbound")
	outb, _ := configManager.Get("main", "Outbound")
	TempOutbound, _ := configManager.Get("main", "TempOutbound")
	TempInbound, _ := configManager.Get("main", "TempInbound")
	Temp, _ := configManager.Get("main", "Temp")
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
	eventBus.Event("Toss")

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