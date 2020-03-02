package main

import (
	"github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/charset"
	"github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/file"
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

	/* Create managers */
	if err := self.Container.Provide(charset.NewCharsetManager); err != nil {
		panic(err)
	}
	if err := self.Container.Provide(storage.NewStorageManager); err != nil {
		panic(err)
	}
	if err := self.Container.Provide(setup.NewSetupManager); err != nil {
		panic(err)
	}
	if err := self.Container.Provide(msg.NewMessageManager); err != nil {
		panic(err)
	}
	if err := self.Container.Provide(area.NewAreaManager); err != nil {
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
	if err := self.Container.Provide(tosser.NewTosserManager); err != nil {
		panic(err)
	}

	/* Check periodic message */
	self.Container.Invoke(func(cm*charset.CharsetManager, am *area.AreaManager, mm *msg.MessageManager, sm *stat.StatManager, setm *setup.SetupManager, fm *file.FileManager) {
		newTosser := tosser.NewTosser(cm, am, mm, sm, setm, fm)
		newTosser.Toss()
	})

	/* Initialize and start application */
	self.Container.Invoke(func(cm *charset.CharsetManager, nm *netmail.NetmailManager, am *area.AreaManager, mm *msg.MessageManager, sm *stat.StatManager, setm *setup.SetupManager, fm *file.FileManager, tm *tosser.TosserManager) {
		master := common.GetMaster()
		master.CharsetManager = cm
		master.NetmailManager = nm
		master.SetupManager = setm
		master.AreaManager = am
		master.MessageManager = mm
		master.FileManager = fm
		master.StatManager = sm
		master.TosserManager = tm
	})

	/* Start service */
	self.Container.Invoke(func() {
		newGoldenSite := ui.NewGoldenSite()
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
