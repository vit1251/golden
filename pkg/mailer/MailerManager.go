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
	netAddr, err1 := setupManager.Get("main", "NetAddr", "")
	if err1 != nil {
		return err1
	}
	password, err2 := setupManager.Get("main", "Password", "")
	if err2 != nil {
		return err2
	}
	address, err3 := setupManager.Get("main", "Address", "")
	if err3 != nil {
		return err3
	}
	inb, err4 := setupManager.Get("main", "Inbound", "")
	if err4 != nil {
		return err4
	}
	outb, err5 := setupManager.Get("main", "Outbound", "")
	if err5 != nil {
		return err5
	}
	TempOutbound, err6 := setupManager.Get("main", "TempOutbound", "")
	if err6 != nil {
		return err6
	}
	TempInbound, err6 := setupManager.Get("main", "TempInbound", "")
	if err6 != nil {
		return err6
	}
	Temp, err6 := setupManager.Get("main", "Temp", "")
	if err6 != nil {
		return err6
	}
	Country, err6 := setupManager.Get("main", "Country", "")
	if err6 != nil {
		return err6
	}
	City, err6 := setupManager.Get("main", "City", "")
	if err6 != nil {
		return err6
	}
	realName, err1 := setupManager.Get("main", "RealName", "")
	if err1 != nil {
		return err1
	}
	stationName, err1 := setupManager.Get("main", "StationName", "")
	if err1 != nil {
		return err1
	}

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
	m.Start()

	/* Wait complete */
	m.Wait()

	/* Complete start tosser */
	if err := statManager.RegisterOutSession(); err != nil {
		log.Printf("Fail on mailer routine: err = %+v", err)
	}

	return nil
}