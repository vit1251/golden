package msg

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vit1251/golden/pkg/storage"
	"go.uber.org/dig"
	"log"
	"strings"
)

type AreaManager struct {
	MessageManager *MessageManager
	StorageManager *storage.StorageManager
	Container      *dig.Container
}

func NewAreaManager(c *dig.Container) *AreaManager {
	am := new(AreaManager)
	am.Container = c
	c.Invoke(func(sm *storage.StorageManager, mm *MessageManager) {
		am.StorageManager = sm
		am.MessageManager = mm
	})
	am.Rescan()
	return am
}


func (self *AreaManager) Rescan() {

	/* Preload echo areas */
	areas, err3 := self.MessageManager.GetAreaList2()
	if err3 != nil {
		panic(err3)
	}

	/* Reset areas */
	for _, area := range areas {

		log.Printf("rescan: area = %+v", area)

		var areaName string = area.Name()

		a := NewArea()
		a.SetName(areaName)
		a.MessageCount = area.MessageCount
	}

}

func (self *AreaManager) updateMsgCount(areas []*Area) error {

	log.Printf("areas = %+v", areas)

	/* Get message count */
	areas2, err1 := self.MessageManager.GetAreaList2()
	if err1 != nil {
		log.Printf("err1 = %+v", err1)
		return err1
	}
	log.Printf("areas = %+v", areas2)

	/* Get message new count */
	areas3, err2 := self.MessageManager.GetAreaList3()
	if err2 != nil {
		log.Printf("err2 = %+v", err2)
		return err2
	}
	log.Printf("areas = %+v", areas3)

	/* Update original areas values */
	for _, area := range areas {

		/* Search area count */
		for _, area2 := range areas2 {
			var areaName string = area.Name()
			var area2Name string = area2.Name()
			if strings.EqualFold(areaName, area2Name) {
				log.Printf("area = '%+v' area2 = '%+v'", areaName, area2Name)
				area.MessageCount = area2.MessageCount
			}
		}

		/* Search area new count */
		for _, area3 := range areas3 {
			var areaName string = area.Name()
			var area3Name string = area3.Name()
			if strings.EqualFold(areaName, area3Name) {
				log.Printf("area = '%+v' area3 = '%+v'", areaName, area3Name)
				area.NewMessageCount = area3.NewMessageCount
			}
		}

	}

	return nil
}

func (self *AreaManager) Register(a *Area) error {

	var areaName string = a.Name()

	query1 := "INSERT INTO `area` ( `areaName`, `areaType`, `areaPath`, `areaSummary`, `areaOrder` ) VALUES ( ?, '', '', '', 0 )"
	var params []interface{}
	params = append(params, areaName)

	self.StorageManager.Exec(query1, params, func(err error) {
		log.Printf("Insert complete with: err = %+v", err)
	})

	return nil
}

func (self *AreaManager) GetAreas() ([]*Area, error) {

	var areas []*Area

	query1 := "SELECT `areaName`, `areaSummary` FROM `area` ORDER BY `areaName` ASC"
	var params []interface{}

	self.StorageManager.Query(query1, params, func(rows *sql.Rows) error {

		var areaName string
		var areaSummary string

		err3 := rows.Scan(&areaName, &areaSummary)
		if err3 != nil {
			return err3
		}

		if areaSummary == "" {
			areaSummary = "Нет описания"
		}

		area := NewArea()
		area.SetName(areaName)
		area.Summary = areaSummary

		areas = append(areas, area)

		return nil
	})

	self.updateMsgCount(areas)

	return areas, nil
}


func (self *AreaManager) GetAreaByName(echoTag string) (*Area, error) {

	var result *Area

	/* Restore parameters */
	query1 := "SELECT `areaName`, `areaSummary` FROM `area` WHERE `areaName` = ?"
	var params []interface{}
	params = append(params, echoTag)

	self.StorageManager.Query(query1, params, func(rows *sql.Rows) error {

		var areaName string
		var areaSummary string

		err3 := rows.Scan(&areaName, &areaSummary)
		if err3 != nil {
			return err3
		}

		area := NewArea()
		area.SetName(areaName)
		area.Summary = areaSummary

		result = area

		return nil
	})

	return result, nil

}

func (self *AreaManager) Update(area *Area) error {

	query1 := "UPDATE `area` SET `areaSummary` = ? WHERE `areaName` = ?"
	var params []interface{}
	params = append(params, area.Summary)
	params = append(params, area.Name())

	self.StorageManager.Exec(query1, params, func(err error) {
		log.Printf("Insert complete with: err = %+v", err)
	})

	return nil
}

func (self *AreaManager) RemoveAreaByName(echoName string) error {

	query1 := "DELETE FROM `area` WHERE `areaName` = ?"
	var params []interface{}
	params = append(params, echoName)

	self.StorageManager.Exec(query1, params, func(err error) {
		log.Printf("Insert complete with: err = %+v", err)
	})

	return nil
}
