package main

import (
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/mailer"
	"github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/tosser"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/ui"
	"time"
	"log"
)

type Application struct {
	AreaManager       *area.AreaManager
	SetupManager      *setup.SetupManager
	MessageManager    *msg.MessageManager
	Version            string                /* Golden Point version string  */
}

func NewApplication() (*Application) {
	app := new(Application)
	app.Version = "1.2.3"
	return app
}

func (self *Application) GetAreaManager() (*area.AreaManager) {
	return self.AreaManager
}

func (self *Application) GetSetupManager() (*setup.SetupManager) {
	return self.SetupManager
}

func (self *Application) Periodic() {

	/* Prepare mailer */
//	m := mailer.NewMailer()
	m := mailer.NewMailerCompat()

	/* Prepare tosser */
	inboundDirectory, err1 := self.SetupManager.Get("main", "Inbound", ".")
	log.Printf("err1 = %+v", err1)
	workInboundDirectory, err2 := self.SetupManager.Get("main", "TempInbound", ".")
	log.Printf("err2 = %+v", err2)

	log.Printf("inboundDirectory = %s", inboundDirectory)
	log.Printf("workInboundDirectory = %s", workInboundDirectory)

	newTosser := tosser.NewTosser()
	newTosser.SetInboundDirectory(inboundDirectory)
	newTosser.SetWorkInboundDirectory(workInboundDirectory)

	log.Printf("Mailer service start")

	/* Main processing */
	for {
		log.Printf("Check new mail")

		/* Check new message */
		m.Check()

		/* Toss new message */
		newTosser.Toss()

		/* Wait 10 min. */
		time.Sleep( 10 * time.Minute )
	}

	log.Printf("Mailer service complete")

}

func (self *Application) Run() {

	/* Create managers */
	self.SetupManager = setup.NewSetupManager()
	self.AreaManager = area.NewAreaManager()
	self.MessageManager = msg.NewMessageManager()

	/* Check periodic message */
	go self.Periodic()

	/* Start user interface Web-service */
	newGoldenSite := ui.NewGoldenSite()
	newGoldenSite.SetSetupManager(self.SetupManager)
	newGoldenSite.SetAreaManager(self.AreaManager)
	newGoldenSite.SetMessageManager(self.MessageManager)
	newGoldenSite.SetVersion(self.Version)
	err := newGoldenSite.Start()
	if err != nil {
		panic(err)
	}

	/* Wait completion */
	log.Printf("Complete.")

}
