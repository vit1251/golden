package mapper

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/registry"
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

	storageManager := self.restoreStorageManager()

	var areaName string = a.GetName()
	var areaType string = ""
	var areaPath string = ""
	var areaSummary string = a.GetSummary()
	var areaOrder int64 = a.GetOrder()

	//                                   1          2           3            4             5
	query1 := "INSERT INTO `area` ( `areaName`, `areaType`, `areaPath`, `areaSummary`, `areaOrder` ) VALUES ( ?, ?, ?, ?, ? )"
	var params []interface{}
	params = append(params, areaName)    // 1
	params = append(params, areaType)    // 2
	params = append(params, areaPath)    // 3
	params = append(params, areaSummary) // 4
	params = append(params, areaOrder)   // 5

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		log.Printf("Insert complete with: err = %+v", err)
		return err
	})

	return err1
}

func (self *EchoAreaMapper) GetAreas() ([]Area, error) {

	storageManager := self.restoreStorageManager()
	mapperManager := self.restoreMapperManager()
	echoMapper := mapperManager.GetEchoMapper()

	var result []Area

	query1 := "SELECT `areaName`, `areaSummary`, `areaCharset`, `areaOrder` FROM `area` ORDER BY `areaOrder` ASC, `areaName` ASC"
	var params []interface{}

	storageManager.Query(query1, params, func(rows *sql.Rows) error {

		var areaName string
		var areaSummary string
		var areaCharset string
		var areaOrder int64

		err3 := rows.Scan(&areaName, &areaSummary, &areaCharset, &areaOrder)
		if err3 != nil {
			return err3
		}

		area := NewArea()
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

func (self *EchoAreaMapper) GetAreaByName(echoTag string) (*Area, error) {

	storageManager := self.restoreStorageManager()

	var result *Area

	//                                                                                        1
	query1 := "SELECT `areaName`, `areaSummary`, `areaCharset`, `areaOrder` FROM `area` WHERE `areaName` = ?"
	var params []interface{}
	params = append(params, echoTag) // 1

	storageManager.Query(query1, params, func(rows *sql.Rows) error {

		var areaName string
		var areaSummary string
		var areaCharset string
		var areaOrder int64

		err3 := rows.Scan(&areaName, &areaSummary, &areaCharset, &areaOrder)
		if err3 != nil {
			return err3
		}

		area := NewArea()
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

	log.Printf("EchoAreaMapper: Update: area = %+v", area)

	storageManager := self.restoreStorageManager()

	var areaSummary string = area.GetSummary()
	var areaCharset string = area.GetCharset()
	var areaOrder int64 = area.GetOrder()
	var areaName string = area.GetName()

	//                                           1                  2                3                    4
	query1 := "UPDATE `area` SET `areaSummary` = ?, `areaCharset` = ?, `areaOrder` = ? WHERE `areaName` = ?"
	var params []interface{}
	params = append(params, areaSummary) // 1
	params = append(params, areaCharset) // 2
	params = append(params, areaOrder)   // 3
	params = append(params, areaName)    // 4

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		log.Printf("Insert complete with: err = %+v", err)
		return nil
	})

	return err1
}

func (self *EchoAreaMapper) RemoveAreaByName(echoName string) error {

	storageManager := self.restoreStorageManager()

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
