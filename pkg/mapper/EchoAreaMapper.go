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

	query1 := "INSERT INTO `area` ( `areaName`, `areaType`, `areaPath`, `areaSummary`, `areaOrder` ) VALUES ( ?, '', '', '', 0 )"
	var params []interface{}
	params = append(params, areaName)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		log.Printf("Insert complete with: err = %+v", err)
		return nil
	})

	return err1
}

func (self *EchoAreaMapper) GetAreas() ([]Area, error) {

	storageManager := self.restoreStorageManager()
	mapperManager := self.restoreMapperManager()
	echoMapper := mapperManager.GetEchoMapper()

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

	newAreas, err1 := echoMapper.UpdateAreaMessageCounters(result)
	if err1 != nil {
		return nil, err1
	}

	return newAreas, nil
}


func (self *EchoAreaMapper) GetAreaByName(echoTag string) (*Area, error) {

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

func (self *EchoAreaMapper) Update(area *Area) error {

	log.Printf("EchoAreaMapper: Update: area = %+v", area)

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

func (self *EchoAreaMapper) RemoveAreaByName(echoName string) error {

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
