package main

import (
	"flag"
	"fmt"
	"github.com/vit1251/golden/pkg/charset"
	cmn "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/config"
	"github.com/vit1251/golden/pkg/eventbus"
	"github.com/vit1251/golden/pkg/installer"
	"github.com/vit1251/golden/pkg/mailer"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/queue"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/site"
	"github.com/vit1251/golden/pkg/storage"
	"github.com/vit1251/golden/pkg/tosser"
	"github.com/vit1251/golden/pkg/tracker"
	"github.com/vit1251/golden/pkg/um"
	"io"
	"log"
	"os"
	"os/signal"
	path2 "path"
	"syscall"
	"time"
)

type Application struct {
	registry *registry.Container /* Component registry    */
	stream   io.WriteCloser      /* Logging writer        */
}

func NewApplication() *Application {

	app := new(Application)
	app.registry = registry.NewContainer()
	return app

}

func (self *Application) makeLogName() string {
	cur := time.Now()
	return fmt.Sprintf("debug_%d%02d%02d_%02d%02d.log", cur.Year(), cur.Month(), cur.Day(), cur.Hour(), cur.Minute())
}

func (self *Application) makeLogPath() string {
	logBaseDirectory := cmn.GetLogDirectory()
	debugName := self.makeLogName()
	return path2.Join(logBaseDirectory, debugName)

}

func (self *Application) startLogging(debug bool) {

	log.SetFlags(log.Ltime | log.Ldate | log.Lmicroseconds)
	log.SetOutput(io.Discard)

	if debug {
		logPath := self.makeLogPath()
		stream, err1 := os.OpenFile(logPath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
		if err1 != nil {
			log.Printf("Error while open debug.log: err = %+v", err1)
		}
		self.stream = stream
		log.SetOutput(self.stream)
	}

}

func (self *Application) stopLogging() {
	if self.stream != nil {
		self.stream.Close()
		self.stream = nil
	}
}

func (self *Application) Run() {

	/* Parse parameters */
	var servicePort int = 8080
	var debugMode bool = false
	flag.IntVar(&servicePort, "P", 8080, "Set HTTP service port")
	flag.BoolVar(&debugMode, "debug", false, "Enable debugging mode")
	flag.Parse()

	/* Start debugging */
	self.startLogging(debugMode)
	defer self.stopLogging()

	/* Database chane path: 1.2.17 -> 1.2.18 */
	self.checkDatabaseLocation()

	/* Start storage service */
	self.registry.Register("UrlManager", um.NewUrlManager(self.registry))
	self.registry.Register(queue.QUEUE_MANAGER_ID, queue.NewQueueManager(self.registry))
	self.registry.Register("EventBus", eventbus.NewEventBus(self.registry))
	self.registry.Register("StorageManager", storage.NewStorageManager(self.registry))
	self.registry.Register("MapperManager", mapper.NewMapperManager(self.registry))
	self.registry.Register(config.CONFIG_MANAGER_ID, config.NewConfigManager(self.registry))

	self.registry.Register(installer.MIGRATION_MANAGER_ID, installer.NewMigrationManager(self.registry))

	self.registry.Register(charset.CHARSET_MANAGER_ID, charset.NewCharsetManager(self.registry))

	self.registry.Register("TrackerManager", tracker.NewTrackerManager(self.registry))
	self.registry.Register("TosserManager", tosser.NewTosserManager(self.registry))
	self.registry.Register("MailerManager", mailer.NewMailerManager(self.registry))

	self.registry.Register("SiteManager", site.NewSiteManager(self.registry))

	/* Debug message */
	cur := time.Now()
	zone, offset := cur.Zone()
	log.Printf("Time zone: %+v (%+v)", zone, offset)

	/* Initialize database (apply new migration) */
	migrationManager := installer.RestoreMigrationManager(self.registry)
	migrationManager.Check()

	/* Restore configuration */
	// TODO - get config ...

	/* Start mail processor */
	tosserManager := tosser.RestoreTosserManager(self.registry)
	tosserManager.Start()

	/* Start file processor */
	trackerManager := tracker.RestoreTrackerManager(self.registry)
	trackerManager.Start()

	/* Start UI site */
	siteManager := site.RestoreSiteManager(self.registry)
	siteManager.SetPort(servicePort)
	siteManager.Start()

	/* Start mailer */
	mailerManager := mailer.RestoreMailerManager(self.registry)
	mailerManager.Start()

	/* Wait system Ctrl+C keyboard interruption or OS terminate request */
	self.waitInterrupt()

	/* Stop mailer service */
	//mailerManager.Stop()

	/* Stop tosser service */
	//tosserService.Stop()

	/* Stop storage service */
	self.stopStorageService()

	/* Wait */
	log.Printf("Complete.")

}

func (self *Application) waitInterrupt() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}

func (self *Application) stopStorageService() {

	log.Printf("Application: Sync storage.")

	storageManager := storage.RestoreStorageManager(self.registry)
	storageManager.Close()

}

func (self *Application) checkDatabaseLocation() {
	prevStorageFile := cmn.GetPrevStorageFile()
	// TODO - check exists ...
	modernStorageFile := cmn.GetModernStorageFile()
	err2 := os.Rename(prevStorageFile, modernStorageFile)
	if err2 != nil {
		log.Printf("Move storage is error: err  = %#v", err2)
	}
}
