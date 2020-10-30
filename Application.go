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
	registry *registry.Container
}

func NewApplication() *Application {

	app := new(Application)
	app.registry = registry.NewContainer()
	return app

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
	self.registry.Register("EventBus", eventbus.NewEventBus(self.registry))
	self.registry.Register("StorageManager", storage.NewStorageManager(self.registry))
	self.registry.Register("MigrationManager", installer.NewMigrationManager(self.registry))
	self.registry.Register("ConfigManager", setup.NewConfigManager(self.registry))

	self.registry.Register("CharsetManager", charset.NewCharsetManager(self.registry))

	self.registry.Register("MessageManager", msg.NewMessageManager(self.registry))
	self.registry.Register("AreaManager", msg.NewAreaManager(self.registry))
	self.registry.Register("FileManager", file.NewFileManager(self.registry))
	self.registry.Register("NetmailManager", netmail.NewNetmailManager(self.registry))
	self.registry.Register("StatManager", stat.NewStatManager(self.registry))

	self.registry.Register("TosserManager",	tosser.NewTosserManager(self.registry))
	self.registry.Register("MailerManager", mailer.NewMailerManager(self.registry))

	self.registry.Register("SiteManager", site.NewSiteManager(self.registry))

	/* Initialize database (apply new migration) */
	migrationManager := self.restoreMigrationManager()
	migrationManager.Check()

	/* Start tosser */
	tosserManager := self.restoreTosserManager()
	tosserManager.Start()

	/* Start site */
	siteManager := self.restoreSiteManager()
	siteManager.SetPort(servicePort)
	siteManager.Start()

	/* Start mailer */
	mailerManager := self.restoreMailerManager()
	mailerManager.Start()

	/* Wait system Ctrl+C keyboard interruption or OS terminate request */
	self.waitInterrupt()

	/* Stop mailer service */
	//mailerManager.Stop()

	/* Stop tosser service */
	//tosserService.Stop()

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

/* Restore managers */

func (self *Application) restoreSiteManager() *site.SiteManager {

	siteManagerPtr := self.registry.Get("SiteManager")
	if siteManager, ok := siteManagerPtr.(*site.SiteManager); ok {
		return siteManager
	} else {
		panic("no site manager")
	}

}

func (self *Application) restoreStorageManager() *storage.StorageManager {

	managerPtr := self.registry.Get("StorageManager")
	if manager, ok := managerPtr.(*storage.StorageManager); ok {
		return manager
	} else {
		panic("no storage manager")
	}

}

func (self *Application) restoreMailerManager() *mailer.MailerManager {

	managerPtr := self.registry.Get("MailerManager")
	if manager, ok := managerPtr.(*mailer.MailerManager); ok {
		return manager
	} else {
		panic("no mailer manager")
	}

}

func (self *Application) restoreMigrationManager() *installer.MigrationManager {
	managerPtr := self.registry.Get("MigrationManager")
	if manager, ok := managerPtr.(*installer.MigrationManager); ok {
		return manager
	} else {
		panic("no migration manager")
	}
}

func (self *Application) restoreTosserManager() *tosser.TosserManager {

	managerPtr := self.registry.Get("TosserManager")
	if manager, ok := managerPtr.(*tosser.TosserManager); ok {
		return manager
	} else {
		panic("no tosser manager")
	}

}