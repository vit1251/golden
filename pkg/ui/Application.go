package ui

import (
	"log"
	"github.com/vit1251/golden-toss/pkg/msgapi/sqlite"
//	"errors"
)

type Application struct {
	AreaList           AreaList
	MessageBase       *sqlite.MessageBase
	MessageBaseReader *sqlite.MessageBaseReader
}

func (self *Application) AreaListReset() {
}

func (self *Application) AreaListAreaRegister(areaName string) {
	area := NewArea()
	area.Name = areaName
	self.AreaList.Areas = append(self.AreaList.Areas, area)
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
	areas, err3 := messageBaseReader.GetAreaList()
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

	self.scanBase()

	/* Start user interface Web-service */
	StartSite(self)
}
