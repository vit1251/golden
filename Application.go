package main

import (
	"github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/tosser"
	"github.com/vit1251/golden/pkg/ui"
	"log"
	"time"
)

type Application struct {
}

func NewApplication() (*Application) {
	app := new(Application)
	return app
}

func (self *Application) Periodic() {

	master := common.GetMaster()

	// TODO - how determine that we stop ...
	for {
		/* Mailer */
		log.Printf("Toss process start")

		/* Tosser */
		newTosser := tosser.NewTosser(
			master.MessageManager,
			master.StatManager,
			master.SetupManager,
			master.FileManager)
		newTosser.Toss()

		log.Printf("Toss process stop")

		time.Sleep( 5 * time.Minute )
	}

	log.Printf("Mailer service complete")

}

func (self *Application) Run() {

	/* Create managers */
	common.GetMaster()

	/* Check periodic message */
	go self.Periodic()

	/* Start user interface Web-service */
	newGoldenSite := ui.NewGoldenSite()
	err := newGoldenSite.Start()
	if err != nil {
		panic(err)
	}

	/* Wait completion */
	log.Printf("Complete.")

}
