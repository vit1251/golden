package main

import (
	"flag"
	"github.com/vit1251/golden/pkg/charset"
	cmn "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/eventbus"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/installer"
	"github.com/vit1251/golden/pkg/mailer"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/netmail"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/site"
	"github.com/vit1251/golden/pkg/stat"
	"github.com/vit1251/golden/pkg/storage"
	"github.com/vit1251/golden/pkg/tosser"
	"log"
	"os"
	"os/signal"
	path2 "path"
	"syscall"
)

type Application struct {
	Container *registry.Container
}

func NewApplication() *Application {
	app := new(Application)
	app.Container = registry.NewContainer()
	return app
}

func (self *Application) processMigration() {

	migrationManagerPtr := self.Container.Get("MigrationManager")
	if migrationManager, ok := migrationManagerPtr.(*installer.MigrationManager); ok {
		migrationManager.Check()
	} else {
		panic("no migration manager")
	}

}

func (self *Application) startTosserService() {

	tosserManagerPtr := self.Container.Get("TosserManager")
	if tosserManager, ok := tosserManagerPtr.(*tosser.TosserManager); ok {
		tosserManager.Start()
	} else {
		panic("no tosser manager")
	}

}

func (self *Application) Run() {

	logBaseDirectory := cmn.GetLogDirectory()
	logPath := path2.Join(logBaseDirectory, "golden.log")
	stream, err1 := os.OpenFile(logPath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err1 != nil {
		log.Printf("Error while open debug.log: err = %+v", err1)
	}
	defer stream.Close()
	log.SetOutput(stream)
	log.SetFlags(log.Ltime | log.Ldate)

	/* Parse parameters */
	var servicePort int
	flag.IntVar(&servicePort, "P", 8080, "Set HTTP service port")
	flag.Parse()

	/* Start storage service */
	self.Container.Register("EventBus", eventbus.NewEventBus(self.Container))
	self.Container.Register("StorageManager", storage.NewStorageManager(self.Container))
	self.Container.Register("MigrationManager", installer.NewMigrationManager(self.Container))
	self.Container.Register("ConfigManager", setup.NewConfigManager(self.Container))

	self.Container.Register("CharsetManager", charset.NewCharsetManager(self.Container))

	self.Container.Register("MessageManager", msg.NewMessageManager(self.Container))
	self.Container.Register("AreaManager", msg.NewAreaManager(self.Container))
	self.Container.Register("FileManager", file.NewFileManager(self.Container))
	self.Container.Register("NetmailManager", netmail.NewNetmailManager(self.Container))
	self.Container.Register("StatManager", stat.NewStatManager(self.Container))

	self.Container.Register("TosserManager",	tosser.NewTosserManager(self.Container))
	self.Container.Register("MailerManager", mailer.NewMailerManager(self.Container))

	self.Container.Register("SiteManager", site.NewSiteManager(self.Container))

	/* Initialize migrations */
	self.processMigration()

	/* Start tosser service */
	self.startTosserService()

	/* Start site service */
	siteManager := self.restoreSiteManager()
	siteManager.SetPort(servicePort)
	siteManager.Start()

	/* Start mailer service */
	mailerManager := self.restoreMailerManager()
	mailerManager.Start()

	/* Wait system interrupt marker */
	self.waitInterrupt()

	/* Stop mailer service */
	//mailerManager.Stop()

	/* Stop tosser service */
	//self.stopTosserService()

	/* Stop storage service */
	self.stopStorageServcie()

	/* Wait */
	log.Printf("Complete.")

}

func (self *Application) waitInterrupt() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}

func (self *Application) stopStorageServcie() {

	log.Printf("Application: Sync storage.")

	storageManager := self.restoreStorageManager()
	storageManager.Close()

}

func (self *Application) restoreSiteManager() *site.SiteManager {

	siteManagerPtr := self.Container.Get("SiteManager")
	if siteManager, ok := siteManagerPtr.(*site.SiteManager); ok {
		return siteManager
	} else {
		panic("no site manager")
	}

}

func (self *Application) restoreStorageManager() *storage.StorageManager {

	managerPtr := self.Container.Get("StorageManager")
	if manager, ok := managerPtr.(*storage.StorageManager); ok {
		return manager
	} else {
		panic("no storage manager")
	}

}

func (self *Application) restoreMailerManager() *mailer.MailerManager {

	managerPtr := self.Container.Get("MailerManager")
	if manager, ok := managerPtr.(*mailer.MailerManager); ok {
		return manager
	} else {
		panic("no mailer manager")
	}

}
