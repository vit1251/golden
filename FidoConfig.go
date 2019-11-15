package main

import (
	"os"
	"log"
	"bufio"
	"strings"
)

func processLine(row string, areaList *AreaList) (error) {
	//
	scanner := bufio.NewScanner(strings.NewReader(row))
	scanner.Split(bufio.ScanWords)
	//
	scanner.Scan() // TODO - check result ...
	KeyWord := scanner.Text()
	KeyWord = strings.ToUpper(KeyWord)
	//
	if KeyWord == "ECHOAREA" || KeyWord == "NETMAILAREA" || KeyWord == "LOCALAREA" || KeyWord == "DUPEAREA" || KeyWord == "BADAREA" {
		//
		scanner.Scan() // TODO - check result ...
		AreaEchoID := scanner.Text()
		//
		scanner.Scan()
		AreaPath := scanner.Text()
		//
		for scanner.Scan() {
			// TODO - process options ...
		}
		log.Printf("AreaEchoID = %s AreaPath = %s", AreaEchoID, AreaPath)
		//
		a := new(Area)
		a.Name = AreaEchoID
		a.Path = AreaPath
		a.Type = AreaTypeSquish // TODO - replace on ...
		//
		areaList.Areas = append(areaList.Areas, a)
	} else if KeyWord == "INCLUDE" {
		//
		scanner.Scan() // TODO - check result ...
		includeFile := scanner.Text()
		//
		err := UpdateAreas(includeFile, areaList)
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

func debugUpdateAreas(areaList *AreaList) (error) {
	//
	//
	a1 := new(Area)
	a1.Name = "DIRECT"
	a1.Path = "/var/spool/ftn/netmail"
	a1.Type = AreaTypeNetmail
	areaList.Areas = append(areaList.Areas, a1)
	//
	a2 := new(Area)
	a2.Name = "RU.UNIX.BSD"
	a2.Path = "/var/spool/ftn/msgbase/ru.unix.bsd"
	a2.Type = AreaTypeSquish
	areaList.Areas = append(areaList.Areas, a2)
	//
	a3 := new(Area)
	a3.Name = "NETMAIL"
	a3.Path = "/var/spool/ftn/msgbase/netmail"
	a3.Type = AreaTypeSquish
	areaList.Areas = append(areaList.Areas, a3)
	//
	a4 := new(Area)
	a4.Name = "HOBBIT.TEST"
	a4.Path = "/var/spool/ftn/msgbase/hobbit.test"
	a4.Type = AreaTypeSquish
	areaList.Areas = append(areaList.Areas, a4)
	//

	//
	return nil
}

func UpdateAreas(path string, areaList *AreaList) (error) {
	//
	stream, err := os.Open(path)
	if err != nil {
		panic(err)
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
			err1 := processLine(row, areaList)
			if err1 != nil {
				panic(err1)
			}
		}
	}
	//
	return nil
}

func ReadAreas(path string) (*AreaList, error) {
	//
	areaList := new(AreaList)
	//
	err1 := UpdateAreas(path, areaList)
	if err1 != nil {
		panic(err1)
	}
	//
	return areaList, nil
}

