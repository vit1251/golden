package main

import (
	"flag"
	"github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/charset"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/installer"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/netmail"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/stat"
	"github.com/vit1251/golden/pkg/storage"
	"github.com/vit1251/golden/pkg/tosser"
	"github.com/vit1251/golden/pkg/ui"
	"go.uber.org/dig"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Application struct {
	Container *dig.Container
}

func NewApplication() *Application {
	app := new(Application)
	app.Container = dig.New()
	return app
}

func (self *Application) Run() {

	stream, err1 := os.OpenFile("debug.log", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		log.Printf("Error while open debug.log: err = %+v", err1)
	}
	defer stream.Close()
	log.SetOutput(stream)
	log.SetFlags(log.Ltime|log.Ldate)

	/* Parse parameters */
	var servicePort int
	flag.IntVar(&servicePort, "P", 8080, "Set HTTP service port")
	flag.Parse()
	log.Printf("servicePort - %+v", servicePort)

	/* Create managers */
	if err := self.Container.Provide(installer.NewMigrationManager); err != nil {
		panic(err)
	}
	if err := self.Container.Provide(charset.NewCharsetManager); err != nil {
		panic(err)
	}
	if err := self.Container.Provide(storage.NewStorageManager); err != nil {
		panic(err)
	}
	if err := self.Container.Provide(setup.NewConfigManager); err != nil {
		panic(err)
	}
	if err := self.Container.Provide(msg.NewMessageManager); err != nil {
		panic(err)
	}
	if err := self.Container.Provide(func() *area.AreaManager {
		return area.NewAreaManager(self.Container)
	}); err != nil {
		panic(err)
	}
	if err := self.Container.Provide(file.NewFileManager); err != nil {
		panic(err)
	}
	if err := self.Container.Provide(stat.NewStatManager); err != nil {
		panic(err)
	}
	if err := self.Container.Provide(netmail.NewNetmailManager); err != nil {
		panic(err)
	}
	if err := self.Container.Provide(func() *tosser.TosserManager {
		result := tosser.NewTosserManager(self.Container)
		return result
	}); err != nil {
		panic(err)
	}

	/* Migrations */
	self.Container.Invoke(func(mm *installer.MigrationManager) {
		mm.Check()
	})

	/* Check periodic message */
	self.Container.Invoke(func() {
		newTosser := tosser.NewTosser(self.Container)
		newTosser.Toss()
	})

	/* Start service */
	self.Container.Invoke(func() {
		newGoldenSite := ui.NewGoldenSite(self.Container)
		newGoldenSite.SetPort(servicePort)
		go newGoldenSite.Start()
	})

	/* Wait sigs */
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	/* Block until a signal is received. */
	<-sigs

	/* Sync storage */
	self.Container.Invoke(func(sm *storage.StorageManager) {
		log.Printf("Sync storage.")
		sm.Close()
	})

	/* Wait */
	log.Printf("Complete.")

}
