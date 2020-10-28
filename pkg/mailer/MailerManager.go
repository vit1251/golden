package mailer

import (
	"fmt"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/stat"
	"go.uber.org/dig"
	"log"
	"time"
)

type MailerManager struct {
	Container    *dig.Container
	event         chan bool
}

func NewMailerManager(c *dig.Container) *MailerManager {
	mm := new(MailerManager)
	mm.Container = c
	mm.event = make(chan bool)
	go mm.run()
	return mm
}

func (self *MailerManager) Start() {
	self.event <- true
}

func (self *MailerManager) run() {
	var procIteration int
	tick := time.NewTicker(5 * time.Minute)
	for alive := true; alive; {
		select {
		case <-self.event:
		case <-tick.C:
			procIteration += 1
			log.Printf(" * Mailer start (%d)", procIteration)
			if err := self.processMailer(); err != nil {
				log.Printf("err = %+v", err)
			}
			log.Printf(" * Mailer complete (%d)", procIteration)
		}
	}
}

func (self *MailerManager) processMailer() error {

	var setupManager *setup.ConfigManager
	var statManager *stat.StatManager
	self.Container.Invoke(func(sm *setup.ConfigManager, sm2 *stat.StatManager) {
		setupManager = sm
		statManager = sm2
	})

	/* Construct node address */
	netAddr, _ := setupManager.Get("main", "NetAddr")
	password, _ := setupManager.Get("main", "Password")
	address, _ := setupManager.Get("main", "Address")
	inb, _ := setupManager.Get("main", "Inbound")
	outb, _ := setupManager.Get("main", "Outbound")
	TempOutbound, _ := setupManager.Get("main", "TempOutbound")
	TempInbound, _ := setupManager.Get("main", "TempInbound")
	Temp, _ := setupManager.Get("main", "Temp")
	Country, _ := setupManager.Get("main", "Country")
	City, _ := setupManager.Get("main", "City")
	realName, _ := setupManager.Get("main", "RealName")
	stationName, _ := setupManager.Get("main", "StationName")

	/* */
	newAddress := fmt.Sprintf("%s@fidonet", address)

	/* Get parameters */
	m := NewMailer(setupManager)
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

	return nil
}