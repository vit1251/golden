package mapper

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/registry"
	"log"
)

type FileAreaMapper struct {
	Mapper
}

func NewFileAreaMapper(r *registry.Container) *FileAreaMapper {
	newFileAreaMapper := new(FileAreaMapper)
	newFileAreaMapper.SetRegistry(r)
	return newFileAreaMapper
}

func (self *FileAreaMapper) GetAreas() ([]FileArea, error) {

	storageManager := self.restoreStorageManager()

	var areas []FileArea

	/* Restore parameters */
	sqlStmt := "SELECT `areaName`, `areaSummary` FROM `filearea` ORDER BY `areaId` ASC"

	var params []interface{}

	err1 := storageManager.Query(sqlStmt, params, func(rows *sql.Rows) error {

		var areaName string
		var areaSummary string

		err3 := rows.Scan(&areaName, &areaSummary)
		if err3 != nil {
			return err3
		}

		area := NewFileArea()
		area.SetName(areaName)
		area.SetSummary(areaSummary)

		areas = append(areas, *area)

		return nil
	})

	return areas, err1
}

func (self *FileAreaMapper) UpdateFileAreasWithFileCount(fileAreas []FileArea) ([]FileArea, error) {

	storageManager := self.restoreStorageManager()

	sqlStmt := "SELECT `fileArea`, count(`fileName`) AS `fileCount` FROM `file` GROUP BY `fileArea` ORDER BY `fileArea` ASC"

	var params []interface{}

	metrics := make(map[string]int)

	err1 := storageManager.Query(sqlStmt, params, func(rows *sql.Rows) error {

		var fileArea string
		var fileCount int

		err2 := rows.Scan(&fileArea, &fileCount)
		if err2 != nil {
			return err2
		}

		metrics[fileArea] = fileCount

		return nil
	})

	/* Populate metrics */
	var result []FileArea
	for _, a := range fileAreas {
		var areaName string = a.GetName()
		var areaCount int
		if count, ok := metrics[areaName]; ok {
			areaCount = count
		}
		a.Count = areaCount
		result = append(result, a)
	}

	return result, err1
}

func (self *FileAreaMapper) CreateFileArea(a *FileArea) error {

	storageManager := self.restoreStorageManager()

	query := "INSERT INTO `filearea` (`areaName`, `areaSummary`, `areaCharset`) VALUES (?, ?, ?)"

	var params []interface{}
	params = append(params, a.GetName())
	params = append(params, a.GetSummary())
	params = append(params, a.GetCharset())

	/* Create area */
	err1 := storageManager.Exec(query, params, func (e sql.Result, err error) error {

		if err != nil {
			log.Printf("Fail on CreateFileArea with error: err = %+v", err)
			return err
		}

		return nil

	})

	return err1
}

func (self *FileAreaMapper) GetAreaByName(areaName string) (*FileArea, error) {

	storageManager := self.restoreStorageManager()

	var result *FileArea

	query1 := "SELECT `areaName`, `areaSummary`, `areaCharset` FROM `filearea` WHERE `areaName` = ?"

	var params []interface{}
	params = append(params, areaName)

	err1 := storageManager.Query(query1, params, func(rows *sql.Rows) error {

		var areaName1 string
		var areaSummary1 string
		var areaCharset1 string

		err3 := rows.Scan(&areaName1, &areaSummary1, &areaCharset1)
		if err3 != nil {
			return err3
		}

		area := NewFileArea()
		area.SetName(areaName1)
		area.SetSummary(areaSummary1)
		area.SetCharset(areaCharset1)

		result = area

		return nil
	})

	return result, err1
}

func (self FileAreaMapper) GetMessageNewCount() (int, error) {
	return 0, nil
}
