package config

import (
	"os"
	"log"
	"bufio"
	"strings"
)

type HuskyImport struct {
	AreaList *AreaList
}

func (self *HuskyImport) processNetMail() {

}

func (self *HuskyImport) processEchoMail() {

}

func (self *HuskyImport) processLine(row string) (error) {
	//
	var params []string
	scanner := bufio.NewScanner(strings.NewReader(row))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		param := scanner.Text()
		params = append(params, param)
	}
	log.Printf("params = %v", params)
	//
	return self.processParams(params)
}

func (self *HuskyImport) processParams(params []string) (error) {
	//
	if len(params) == 0 {
		return nil
	}
	//
	KeyWord := params[0]
	//
	if strings.EqualFold(KeyWord, "NETMAILAREA") {

		AreaEchoID := params[1]
		AreaPath := params[2]

		a := new(Area)
		a.Name = AreaEchoID
		a.Path = AreaPath
		a.Type = AreaTypeMsg
		//
		self.registerArea(a)

	} else if strings.EqualFold(KeyWord, "DUPEAREA") {
	} else if strings.EqualFold(KeyWord, "LOCALAREA") {
	} else if strings.EqualFold(KeyWord, "BADAREA") {
	} else if strings.EqualFold(KeyWord, "ECHOAREA") {
		//
		AreaEchoID := params[1]
		AreaPath := params[2]
		//
		log.Printf("AreaEchoID = %s AreaPath = %s", AreaEchoID, AreaPath)
		//
		a := new(Area)
		a.Name = AreaEchoID
		a.Path = AreaPath
		a.Type = AreaTypeSquish
		//
		self.registerArea(a)
		//
	} else if strings.EqualFold(KeyWord, "INCLUDE") {
		//
		includeFile := params[1]
		//
		err := self.UpdateAreas(includeFile)
		if err != nil {
			return err
		}
		//
	} else {
		log.Printf("Unknown keyword %s", KeyWord)
	}
	//
	return nil
}

func (self *HuskyImport) registerArea(a *Area) (error) {
	self.AreaList.Areas = append(self.AreaList.Areas, a)
	return nil
}

func (self *HuskyImport) debugUpdateAreas() (error) {
	//
	//
	a1 := new(Area)
	a1.Name = "DIRECT"
	a1.Path = "/var/spool/ftn/netmail"
	a1.Type = AreaTypeMsg
	self.registerArea(a1)
	//
	a2 := new(Area)
	a2.Name = "RU.UNIX.BSD"
	a2.Path = "/var/spool/ftn/msgbase/ru.unix.bsd"
	a2.Type = AreaTypeSquish
	self.registerArea(a2)
	//
	a3 := new(Area)
	a3.Name = "NETMAIL"
	a3.Path = "/var/spool/ftn/msgbase/netmail"
	a3.Type = AreaTypeSquish
	self.registerArea(a3)
	//
	a4 := new(Area)
	a4.Name = "HOBBIT.TEST"
	a4.Path = "/var/spool/ftn/msgbase/hobbit.test"
	a4.Type = AreaTypeSquish
	self.registerArea(a4)
	//
	return nil
}

func (self *HuskyImport) UpdateAreas(filename string) (error) {
	//
	stream, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer stream.Close()
	//
	scanner := bufio.NewScanner(stream)
	for scanner.Scan() {
		row := scanner.Text()
		log.Printf("row = %s", row)
		if len(row) > 0 {
			if row[0] == '#' {
				log.Printf("Comment. Skip.")
				continue
			}
			err1 := self.processLine(row)
			if err1 != nil {
				return err1
			}
		}
	}
	//
	return nil
}

func NewHuskyImport() (*HuskyImport) {
	hi := new(HuskyImport)
	return hi
}

func (self *HuskyImport) ReadAreas(filename string) (*AreaList, error) {
	self.AreaList = new(AreaList)
	err1 := self.UpdateAreas(filename)
	if err1 != nil {
		return nil, err1
	}
	return self.AreaList, nil
}
