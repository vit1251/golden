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
	sqlStmt := "SELECT `areaName`, `areaMode`, `areaSummary`, `areaCharset` FROM `filearea` ORDER BY `areaOrder` ASC, `areaId` ASC"

	var params []interface{}

	err1 := storageManager.Query(sqlStmt, params, func(rows *sql.Rows) error {

		var areaName string
		var areaMode string
		var areaSummary string
		var areaCharset string

		err3 := rows.Scan(&areaName, &areaMode, &areaSummary, &areaCharset)
		if err3 != nil {
			return err3
		}

		area := NewFileArea()
		area.SetName(areaName)
		area.SetMode(areaMode)
		area.SetSummary(areaSummary)
		area.SetCharset(areaCharset)

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

	var areaName string = a.GetName()
	var areaMode string = a.GetMode()
	var areaSummary string = a.GetSummary()
	var areaCharset string = a.GetCharset()
	var areaOrder int = a.GetOrder()

	query := "INSERT INTO `filearea` (`areaName`, `areaMode`, `areaSummary`, `areaCharset`, `areaOrder`) VALUES (?, ?, ?, ?, ?)"

	var params []interface{}
	params = append(params, areaName)
	params = append(params, areaMode)
	params = append(params, areaSummary)
	params = append(params, areaCharset)
	params = append(params, areaOrder)

	/* Create area */
	err1 := storageManager.Exec(query, params, func(e sql.Result, err error) error {

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

	query1 := "SELECT `areaName`, `areaMode`, `areaSummary`, `areaCharset` FROM `filearea` WHERE `areaName` = ?"

	var params []interface{}
	params = append(params, areaName)

	err1 := storageManager.Query(query1, params, func(rows *sql.Rows) error {

		var areaName string
		var areaMode string
		var areaSummary string
		var areaCharset string

		err3 := rows.Scan(&areaName, &areaMode, &areaSummary, &areaCharset)
		if err3 != nil {
			return err3
		}

		area := NewFileArea()
		area.SetName(areaName)
		area.SetMode(areaMode)
		area.SetSummary(areaSummary)
		area.SetCharset(areaCharset)

		result = area

		return nil
	})

	return result, err1
}

func (self FileAreaMapper) GetMessageNewCount() (int, error) {
	return 0, nil
}
