package ui

import (
	"log"
	"github.com/vit1251/golden/pkg/msgapi/squish"
//	"github.com/vit1251/golden/pkg/msgapi/sqlite"
	"github.com/vit1251/golden/pkg/config"
//	"errors"
)

type Application struct {
	config   *config.Config          /* Application confiuration */
}

func (self *Application) readConfig() {
	//
	self.config = new(config.Config)
	areaList, err1 := config.ReadAreas("/etc/hpt/config")
	if err1 != nil {
		panic(err1)
	}
	self.config.AreaList = areaList
}

func (self *Application) scanHeaders() {

	/* Prepare base plugins*/
//	plugins := make(map[AreaType]MessageBase)
//	plugins[config.AreaTypeMsg] = 
//	plugins[config.AreaTypeSquish] = 
//	plugins[config.AreaTypeSqlite] = 

	/* Re-scan area */
	for _, area := range self.config.AreaList.Areas {
		log.Printf("Scan %s: path = %s", area.Name, area.Path)
//		if area.Type == config.AreaTypeMsg {
//			var msgBase = msgapi.FidoMessageBase{}
//			msgBase.ReadBase(area.Path)
//			messageCount := msgBase.GetMessageCount()
//			area.MessageCount = messageCount
//		} else
		if area.Type == config.AreaTypeSquish {
			msgBase := new(squish.SquishMessageBase)
			msgBase.ReadBase(area.Path)
			messageCount := msgBase.GetMessageCount()
			area.MessageCount = messageCount
		} else {
			log.Printf("Fail on scan message base %s", area.Name)
		}
	}
}

func (self *Application) Run() {
	/* Read service parameters */
	self.readConfig()
	/* Rescan message bases */
	self.scanHeaders()
	/* Start user interface Web-service */
	StartSite(self)
}
