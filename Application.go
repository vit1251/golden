package main

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/stat"
	"github.com/vit1251/golden/pkg/tosser"
	"github.com/vit1251/golden/pkg/ui"
	"go.uber.org/dig"
	"log"
	"os/user"
	"path/filepath"
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

	/* Create user storage */
	if err := self.Container.Provide(func() (*sql.DB, error) {
		usr, err1 := user.Current()
		if err1 != nil {
			return nil, err1
		}
		userHomeDir := usr.HomeDir
		log.Printf("userHomeDir = %+v", userHomeDir)
		userStoragePath := filepath.Join(userHomeDir, "golden.sqlite3")
		log.Printf("userStoragePath = %+v", userStoragePath)
		db, err2 := sql.Open("sqlite3", userStoragePath)
		return db, err2
	}); err != nil {
		panic(err)
	}

	/* Create managers */
	if err := self.Container.Provide(setup.NewSetupManager); err != nil {
		panic(err)
	}
	if err := self.Container.Provide(msg.NewMessageManager); err != nil {
		panic(err)
	}
	if err := self.Container.Provide(func(conn *sql.DB, mm *msg.MessageManager) *area.AreaManager{
		return area.NewAreaManager(conn, mm)
	}); err != nil {
		panic(err)
	}
	if err := self.Container.Provide(file.NewFileManager); err != nil {
		panic(err)
	}
	if err := self.Container.Provide(stat.NewStatManager); err != nil {
		panic(err)
	}
	if err := self.Container.Provide(func(sm *setup.SetupManager) *tosser.TosserManager {
		return tosser.NewTosserManager(sm)
	}); err != nil {
		panic(err)
	}

	/* Check periodic message */
	self.Container.Invoke(func(mm *msg.MessageManager, sm *stat.StatManager, setm *setup.SetupManager, fm *file.FileManager) {
		newTosser := tosser.NewTosser(mm, sm, setm, fm)
		newTosser.Toss()
	})

	/* Initialize master container */
	self.Container.Invoke(func(am *area.AreaManager, mm *msg.MessageManager, sm *stat.StatManager, setm *setup.SetupManager, fm *file.FileManager, tm *tosser.TosserManager) {
		master := common.GetMaster()
		master.SetupManager = setm
		master.AreaManager = am
		master.MessageManager = mm
		master.FileManager = fm
		master.StatManager = sm
		master.TosserManager = tm
	})

	/* Start Web-service */
	newGoldenSite := ui.NewGoldenSite()
	err := newGoldenSite.Start()
	if err != nil {
		panic(err)
	}

	/* Wait */
	log.Printf("Complete.")

}
