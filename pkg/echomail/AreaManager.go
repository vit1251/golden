package echomail

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
	"log"
)

type AreaManager struct {
	registry       *registry.Container
}

func NewAreaManager(r *registry.Container) *AreaManager {
	am := new(AreaManager)
	am.registry = r
	return am
}

func (self *AreaManager) Register(a *Area) error {

	storageManager := self.restoreStorageManager()

	var areaName string = a.GetName()

	query1 := "INSERT INTO `area` ( `areaName`, `areaType`, `areaPath`, `areaSummary`, `areaOrder` ) VALUES ( ?, '', '', '', 0 )"
	var params []interface{}
	params = append(params, areaName)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		log.Printf("Insert complete with: err = %+v", err)
		return nil
	})

	return err1
}

func (self *AreaManager) GetAreas() ([]Area, error) {

	storageManager := self.restoreStorageManager()
	messageManager := self.restoreMessageManager()

	var result []Area

	query1 := "SELECT `areaName`, `areaSummary`, `areaCharset` FROM `area` ORDER BY `areaName` ASC"
	var params []interface{}

	storageManager.Query(query1, params, func(rows *sql.Rows) error {

		var areaName string
		var areaSummary string
		var areaCharset string

		err3 := rows.Scan(&areaName, &areaSummary, &areaCharset)
		if err3 != nil {
			return err3
		}

		area := NewArea()
		area.SetName(areaName)
		area.Summary = areaSummary
		area.Charset = areaCharset

		result = append(result, *area)

		return nil
	})

	newAreas, err1 := messageManager.UpdateAreaMessageCounters(result)
	if err1 != nil {
		return nil, err1
	}

	return newAreas, nil
}


func (self *AreaManager) GetAreaByName(echoTag string) (*Area, error) {

	storageManager := self.restoreStorageManager()

	var result *Area

	/* Restore parameters */
	query1 := "SELECT `areaName`, `areaSummary`, `areaCharset` FROM `area` WHERE `areaName` = ?"
	var params []interface{}
	params = append(params, echoTag)

	storageManager.Query(query1, params, func(rows *sql.Rows) error {

		var areaName string
		var areaSummary string
		var areaCharset string

		err3 := rows.Scan(&areaName, &areaSummary, &areaCharset)
		if err3 != nil {
			return err3
		}

		area := NewArea()
		area.SetName(areaName)
		area.Summary = areaSummary
		area.Charset = areaCharset

		result = area

		return nil
	})

	return result, nil

}

func (self *AreaManager) Update(area *Area) error {

	log.Printf("AreaManager: Update: area = %+v", area)

	storageManager := self.restoreStorageManager()

	query1 := "UPDATE `area` SET `areaSummary` = ?, `areaCharset` = ? WHERE `areaName` = ?"
	var params []interface{}
	params = append(params, area.Summary) // 1
	params = append(params, area.Charset) // 2
	params = append(params, area.GetName()) // 3

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		log.Printf("Insert complete with: err = %+v", err)
		return nil
	})

	return err1
}

func (self *AreaManager) RemoveAreaByName(echoName string) error {

	storageManager := self.restoreStorageManager()

	query1 := "DELETE FROM `area` WHERE `areaName` = ?"
	var params []interface{}
	params = append(params, echoName)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		log.Printf("Insert complete with: err = %+v", err)
		return nil
	})

	return err1
}

func (self *AreaManager) restoreMessageManager() *MessageManager {

	managerPtr := self.registry.Get("MessageManager")
	if manager, ok := managerPtr.(*MessageManager); ok {
		return manager
	} else {
		panic("no message manager")
	}

}

func (self *AreaManager) restoreStorageManager() *storage.StorageManager {

	managerPtr := self.registry.Get("StorageManager")
	if manager, ok := managerPtr.(*storage.StorageManager); ok {
		return manager
	} else {
		panic("no message manager")
	}

}
