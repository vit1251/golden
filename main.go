package main

import (
	"log"
	"errors"
)

type AreaType int

const (
    AreaTypeNetmail AreaType = 1
    AreaTypeSquish  AreaType = 2
)

type Area struct {
	Name         string
	Path         string
	Type         AreaType
	MessageCount int
}

type AreaList struct {
	Areas   []*Area
}

func (self *AreaList) SearchByName(echoTag string) (*Area, error) {
	for _, area := range self.Areas {
		if area.Name == echoTag {
			return area, nil
		}
	}
	return nil, errors.New("No area exists.")
}

type Config struct {
	AreaList *AreaList
}

type Application struct {
	config   *Config          /* Application confiuration */
}

func (self *Application) readConfig() {
	//
	self.config = new(Config)
	areaList, err1 := ReadAreas("/etc/hpt/config")
	if err1 != nil {
		panic(err1)
	}
	self.config.AreaList = areaList
}

func (self *Application) scanHeaders() {
	for _, area := range self.config.AreaList.Areas {
		log.Printf("Scan %s: path = %s", area.Name, area.Path)
		if area.Type == AreaTypeNetmail {
			var msgBase = FidoMessageBase{}
			msgBase.ReadBase(area.Path)
			messageCount := msgBase.GetMessageCount()
			area.MessageCount = messageCount
		} else if area.Type == AreaTypeSquish {
			var msgBase = SquishMessageBase{}
			msgBase.ReadBase(area.Path)
			messageCount := msgBase.GetMessageCount()
			area.MessageCount = messageCount
		} else {
			log.Fatal("Fail on scan message base %s", area.Name)
		}
	}
}


func (self *Application) Run() {
	//
	self.readConfig()
	self.scanHeaders()
	self.startSite()
	//
}

func main() {
	app := new(Application)
	app.Run()
}
