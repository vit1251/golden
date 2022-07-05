package mapper

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
	"github.com/vit1251/golden/pkg/utils"
	"log"
)

type EchoAreaMapper struct {
	Mapper
}

func NewEchoAreaMapper(r *registry.Container) *EchoAreaMapper {
	newEchoAreaMapper := new(EchoAreaMapper)
	newEchoAreaMapper.SetRegistry(r)
	return newEchoAreaMapper
}

func (self *EchoAreaMapper) Register(a *Area) error {

	storageManager := storage.RestoreStorageManager(self.registry)

	var areaName string = a.GetName()
	var areaType string = ""
	var areaPath string = ""
	var areaSummary string = a.GetSummary()
	var areaOrder int64 = a.GetOrder()
	var areaIndex string = a.GetAreaIndex()

	/* Make unique area index */
	if areaIndex == "" {
		areaIndex = utils.IndexHelper_makeUUID()
		a.SetAreaIndex(areaIndex)
	}

	//                                   1          2           3            4             5            6
	query1 := "INSERT INTO `area` ( `areaName`, `areaType`, `areaPath`, `areaSummary`, `areaOrder`, `areaIndex` ) VALUES ( ?, ?, ?, ?, ?, ? )"
	var params []interface{}
	params = append(params, areaName)    // 1
	params = append(params, areaType)    // 2
	params = append(params, areaPath)    // 3
	params = append(params, areaSummary) // 4
	params = append(params, areaOrder)   // 5
	params = append(params, areaIndex)   // 6

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		log.Printf("Insert complete with: err = %+v", err)
		return err
	})

	return err1
}

func (self *EchoAreaMapper) GetAreas() ([]Area, error) {

	storageManager := storage.RestoreStorageManager(self.registry)
	mapperManager := RestoreMapperManager(self.registry)
	echoMapper := mapperManager.GetEchoMapper()

	var result []Area

	query1 := "SELECT `areaIndex`, `areaName`, `areaSummary`, `areaCharset`, `areaOrder` FROM `area` ORDER BY `areaOrder` ASC, `areaName` ASC"
	var params []interface{}

	storageManager.Query(query1, params, func(rows *sql.Rows) error {

		var areaIndex string
		var areaName string
		var areaSummary string
		var areaCharset string
		var areaOrder int64

		err3 := rows.Scan(&areaIndex, &areaName, &areaSummary, &areaCharset, &areaOrder)
		if err3 != nil {
			return err3
		}

		area := NewArea()
		area.SetAreaIndex(areaIndex)
		area.SetName(areaName)
		area.SetSummary(areaSummary)
		area.SetCharset(areaCharset)
		area.SetOrder(areaOrder)

		result = append(result, *area)

		return nil
	})

	newAreas, err1 := echoMapper.UpdateAreaMessageCounters(result)
	if err1 != nil {
		return nil, err1
	}

	return newAreas, nil
}

func (self *EchoAreaMapper) GetAreaByName(areaName string) (*Area, error) {

	storageManager := storage.RestoreStorageManager(self.registry)

	var result *Area

	//                                                                                        1
	query1 := "SELECT `areaIndex`, `areaName`, `areaSummary`, `areaCharset`, `areaOrder` FROM `area` WHERE `areaName` = ?"
	var params []interface{}
	params = append(params, areaName) // 1

	storageManager.Query(query1, params, func(rows *sql.Rows) error {

		var areaIndex string
		var origAreaName string
		var areaSummary string
		var areaCharset string
		var areaOrder int64

		err3 := rows.Scan(&areaIndex, &origAreaName, &areaSummary, &areaCharset, &areaOrder)
		if err3 != nil {
			return err3
		}

		area := NewArea()
		area.SetAreaIndex(areaIndex)
		area.SetName(origAreaName)
		area.SetSummary(areaSummary)
		area.SetCharset(areaCharset)
		area.SetOrder(areaOrder)

		result = area

		return nil
	})

	return result, nil

}

func (self *EchoAreaMapper) GetAreaByAreaIndex(areaIndex string) (*Area, error) {

	storageManager := storage.RestoreStorageManager(self.registry)

	var result *Area

	//                                                                                        1
	query1 := "SELECT `areaIndex`, `areaName`, `areaSummary`, `areaCharset`, `areaOrder` FROM `area` WHERE `areaIndex` = ?"
	var params []interface{}
	params = append(params, areaIndex) // 1

	storageManager.Query(query1, params, func(rows *sql.Rows) error {

		var newAreaIndex string
		var areaName string
		var areaSummary string
		var areaCharset string
		var areaOrder int64

		err3 := rows.Scan(&newAreaIndex, &areaName, &areaSummary, &areaCharset, &areaOrder)
		if err3 != nil {
			return err3
		}

		area := NewArea()
		area.SetAreaIndex(newAreaIndex)
		area.SetName(areaName)
		area.SetSummary(areaSummary)
		area.SetCharset(areaCharset)
		area.SetOrder(areaOrder)

		result = area

		return nil
	})

	return result, nil

}

func (self *EchoAreaMapper) Update(area *Area) error {
	if area.areaIndex == "" {
		return self.updateByAreaName(area)
	} else {
		// TODO - implement update by `areaIndex` here ...
		return self.updateByAreaName(area)
	}
}

func (self *EchoAreaMapper) updateByAreaName(area *Area) error {

	log.Printf("EchoAreaMapper: Update: area = %+v", area)

	storageManager := storage.RestoreStorageManager(self.registry)

	var areaIndex string = area.GetAreaIndex()
	var areaSummary string = area.GetSummary()
	var areaCharset string = area.GetCharset()
	var areaOrder int64 = area.GetOrder()
	var areaName string = area.GetName()

	//                           1                2                  3                  4                     5
	query1 := "UPDATE `area` SET `areaIndex` = ?, `areaSummary` = ?, `areaCharset` = ?, `areaOrder` = ? WHERE `areaName` = ?"
	var params []interface{}
	params = append(params, areaIndex)   // 1
	params = append(params, areaSummary) // 2
	params = append(params, areaCharset) // 3
	params = append(params, areaOrder)   // 4
	params = append(params, areaName)    // 5

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		log.Printf("Insert complete with: err = %+v", err)
		return nil
	})

	return err1
}

func (self *EchoAreaMapper) RemoveAreaByName(echoName string) error {

	storageManager := storage.RestoreStorageManager(self.registry)

	//                                               1
	query1 := "DELETE FROM `area` WHERE `areaName` = ?"
	var params []interface{}
	params = append(params, echoName) // 1

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		log.Printf("Insert complete with: err = %+v", err)
		return nil
	})

	return err1
}
