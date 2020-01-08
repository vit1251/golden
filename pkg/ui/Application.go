package ui

import (
	"github.com/vit1251/golden/pkg/msgapi/sqlite"
	"github.com/vit1251/golden/pkg/area"
//	"log"
)

type Application struct {
	AreaManager       *area.AreaManager
	MessageBase       *sqlite.MessageBase
	MessageBaseReader *sqlite.MessageBaseReader
}

func (self *Application) GetAreaManager() (*area.AreaManager) {
	return self.AreaManager
}

func (self *Application) Open() {

	/* Open message base */
	messageBase, err1 := sqlite.NewMessageBase()
	if err1 != nil {
		panic(err1)
	}
	self.MessageBase = messageBase

	/* Create message base reader */
	messageBaseReader, err2 := sqlite.NewMessageBaseReader(messageBase)
	if err2 != nil {
		panic(err2)
	}
	self.MessageBaseReader = messageBaseReader

}

func (self *Application) Close() {
}

func (self *Application) Run() {

	/* Open base */
	self.Open()

	/* Update settings */
	self.AreaManager = area.NewAreaManager()

	/* Start user interface Web-service */
	self.StartSite()

}
