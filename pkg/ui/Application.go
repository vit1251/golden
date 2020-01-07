package ui

import (
	"log"
	"github.com/vit1251/golden/pkg/msgapi/sqlite"
)

type Application struct {
	AreaList           AreaList
	MessageBase       *sqlite.MessageBase
	MessageBaseReader *sqlite.MessageBaseReader
}

func (self *Application) AreaListReset() {
}

func (self *Application) AreaListAreaRegister(area *sqlite.Area) {
	a := NewArea()
	a.Name = area.Name
	a.MessageCount = area.Count
	self.AreaList.Areas = append(self.AreaList.Areas, a)
}

func (self *Application) scanBase() {

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

	/* Preload echo areas */
	areas, err3 := messageBaseReader.GetAreaList2()
	if err3 != nil {
		panic(err3)
	}
	/* Reset areas */
	self.AreaListReset()
	for _, area := range areas {
		log.Printf("area = %q", area)
		self.AreaListAreaRegister(area)
	}

}

func (self *Application) Run() {

	/* Update settings */
	self.scanBase()

	/* Start user interface Web-service */
	self.StartSite()

}
