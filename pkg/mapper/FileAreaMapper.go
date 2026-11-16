package mapper

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
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

func (self *FileAreaMapper) GetAreas() ([]*FileArea, error) {

	storageManager := storage.RestoreStorageManager(self.registry)

	var areas []*FileArea

	query1 := "SELECT `areaName`, `areaMode`, `areaSummary`, `areaCharset` FROM `filearea` ORDER BY `areaOrder` ASC, `areaId` ASC"
	var args []interface{}

	err1 := storageManager.Query(query1, args, func(rows *sql.Rows) error {

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

		areas = append(areas, area)

		return nil
	})
	if err1 != nil {
            return nil, err1
	}

        if err2 := self.UpdateFileAreasWithFileCount(areas); err2 != nil {
            return nil, err2
	}

        if err3 := self.UpdateNewFileAreasWithFileCount(areas); err3 != nil {
            return nil, err3
	}

	return areas, nil
}

func (self *FileAreaMapper) UpdateFileAreasWithFileCount(fileAreas []*FileArea) error {

	storageManager := storage.RestoreStorageManager(self.registry)

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
	for _, fileArea := range fileAreas {
		var areaName string = fileArea.GetName()
		var areaCount int
		if count, ok := metrics[areaName]; ok {
			areaCount = count
		}
		fileArea.SetCount(areaCount)
	}

	return err1
}

func (self *FileAreaMapper) UpdateNewFileAreasWithFileCount(fileAreas []*FileArea) error {

	storageManager := storage.RestoreStorageManager(self.registry)

	sqlStmt := "SELECT `fileArea`, count(`fileName`) AS `fileCount` FROM `file` WHERE `fileViewCount` = 0 GROUP BY `fileArea` ORDER BY `fileArea` ASC"
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

	for _, fileArea := range fileAreas {
		var areaName string = fileArea.GetName()
		var areaNewCount int
		if count, ok := metrics[areaName]; ok {
			areaNewCount = count
		}
		fileArea.SetNewCount(areaNewCount)
	}

	return err1
}

func (self *FileAreaMapper) CreateFileArea(a *FileArea) error {

	storageManager := storage.RestoreStorageManager(self.registry)

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

	storageManager := storage.RestoreStorageManager(self.registry)

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

func (self *FileAreaMapper) UpdateArea(a *FileArea) error {
    storageManager := storage.RestoreStorageManager(self.registry)
    query := "UPDATE `filearea` SET `areaSummary` = ?, `areaCharset` = ? WHERE `areaName` = ?"
    var params []interface{}
    params = append(params, a.GetSummary())
    params = append(params, a.GetCharset())
    params = append(params, a.GetName())
    return storageManager.Exec(query, params, func(e sql.Result, err error) error {
        return err
    })
}
