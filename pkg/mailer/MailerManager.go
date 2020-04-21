package mailer

import (
	"fmt"
	"github.com/vit1251/golden/pkg/audio"
	"github.com/vit1251/golden/pkg/fidotime"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/stat"
	"go.uber.org/dig"
	"log"
	"time"
)

type MailerManager struct {
	Container *dig.Container
}

func NewMailerManager(c *dig.Container) *MailerManager {
	mm := new(MailerManager)
	mm.Container = c
	return mm
}

func (self *MailerManager) Start() {
	go self.run()
}

func (self *MailerManager) run() {

	/* Update duration parse */
	dp := fidotime.NewDurationParser()
	d, err1 := dp.Parse("10m")
	if err1 != nil {
		log.Printf("Error parse")
		d = time.Minute * 15
	}

	log.Printf(" * Mailer Start")
	for {
		log.Printf("Start mailer processing")
		err := self.processMailer()
		log.Printf("err = %+v", err)
		log.Printf("Stop mailer processing")

		log.Printf("Start wait mailer: d = %+v", d)
		time.Sleep(d)
		log.Printf("Stop wait mailer")
	}
	log.Printf(" * Mailer complete")

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

	/**/
	newAddress := fmt.Sprintf("%s@fidonet", address)

	/* Get parameters */
	m := NewMailer(setupManager)
	m.SetTempOutbound(TempOutbound)
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

	/* Audio play */
	if m.InFileCount > 0 {
		am := audio.NewAudioManager()
		am.Play("you-have-new-message.mp3")
	}

	return nil
}