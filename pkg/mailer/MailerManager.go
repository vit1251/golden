package mailer

import (
	"fmt"
	"github.com/vit1251/golden/pkg/audio"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/stat"
	"go.uber.org/dig"
	"log"
	"time"
)

type MailerManager struct {
	Container    *dig.Container
	AudioManager *audio.AudioManager
	event         chan bool
}

func NewMailerManager(c *dig.Container) *MailerManager {
	mm := new(MailerManager)
	mm.Container = c
	mm.AudioManager = audio.NewAudioManager()
	mm.event = make(chan bool)
	go mm.run()
	return mm
}

func (self *MailerManager) Start() {
	self.event <- true
}

func (self *MailerManager) run() {
	for alive := true; alive; {
		timer := time.NewTimer(5 * time.Minute)
		select {
		case <-self.event:
		case <-timer.C:
			log.Printf("Mailer start")
			self.AudioManager.Play("sess_start.mp3")
			if err := self.processMailer(); err != nil {
				log.Printf("err = %+v", err)
			}
			self.AudioManager.Play("sess_stop.mp3")
			log.Printf("Mailer complete")
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
	m.Start()

	/* Wait complete */
	m.Wait()

	/* Complete start tosser */
	if err := statManager.RegisterOutSession(); err != nil {
		log.Printf("Fail on mailer routine: err = %+v", err)
	}

	return nil
}