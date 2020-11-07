package mapper

import (
	"database/sql"
	cmn "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
	"log"
	"path/filepath"
)

type FileMapper struct {
	Mapper
}

func NewFileMapper(r *registry.Container) *FileMapper {
	newFileMapper := new(FileMapper)
	newFileMapper.SetRegistry(r)
	return newFileMapper
}

func (self *FileMapper) GetAreas() ([]FileArea, error) {

	storageManager := self.restoreStorageManager()

	var areas []FileArea

	/* Restore parameters */
	sqlStmt := "SELECT `areaName`, `areaPath`, `areaSummary` FROM `filearea`"

	var params []interface{}
	//params = append(params, echoName)

	err1 := storageManager.Query(sqlStmt, params, func(rows *sql.Rows) error {

		var areaName string
		var areaPath string
		var areaSummary string

		err3 := rows.Scan(&areaName, &areaPath, &areaSummary)
		if err3 != nil {
			return err3
		}

		/**/
		if areaSummary == "" {
			areaSummary = "Нет описания"
		}

		area := NewFileArea()
		area.SetName(areaName)
		area.Path = areaPath
		area.Summary = areaSummary

		areas = append(areas, *area)

		return nil
	})

	return areas, err1
}

func (self *FileMapper) GetAreasWithFileCount() ([]FileArea, error) {

	storageManager := self.restoreStorageManager()

	conn := storageManager.GetConnection() // TODO

	var result []FileArea

	/* Step 2. Start SQL transaction */
	ConnTransaction, err := conn.Begin()
	if err != nil {
		return nil, err
	}

	sqlStmt := "SELECT `fileArea`, count(`fileName`) AS `msgCount` FROM `file` GROUP BY `fileArea` ORDER BY `fileArea` ASC"
	rows, err1 := ConnTransaction.Query(sqlStmt)
	if err1 != nil {
		return nil, err1
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var count int
		err2 := rows.Scan(&name, &count)
		if err2 != nil {
			return nil, err2
		}

		area := NewFileArea()
		area.SetName(name)
		area.Count = count

		result = append(result, *area)
	}

	ConnTransaction.Commit()

	return result, nil
}

func (self *FileMapper) CreateFileArea(a *FileArea) error {

	storageManager := self.restoreStorageManager()
	conn := storageManager.GetConnection() // TODO

	var name string = a.GetName()
	var path string = a.Path
	var summary string = a.Summary

	log.Printf("Create file area: %+v", a)

	/* Prepare SQL request */
	query := "INSERT INTO `filearea` (`areaName`, `areaType`, `areaPath`, `areaSummary`, `areaOrder`) VALUES (?, ?, ?, ?, ?)"

	/* Create area */
	_, err1 := conn.Exec(query, name, "sqlite3", path, summary, 0)
	if err1 != nil {
		log.Printf("Fail on CreateFileArea with error: err = %+v", err1)
	}

	return err1
}

func (self *FileMapper) GetFileHeaders(echoTag string) ([]File, error) {

	storageManager := self.restoreStorageManager()
	conn := storageManager.GetConnection() // TODO

	var result []File

	/* Step 2. Start SQL transaction */
	ConnTransaction, err := conn.Begin()
	if err != nil {
		return nil, err
	}

	sqlStmt := "SELECT `fileArea`, `fileName`, `fileDesc`, `fileTime` FROM `file` WHERE `fileArea` = $1"
	log.Printf("sql = %q echoTag = %q", sqlStmt, echoTag)
	rows, err1 := ConnTransaction.Query(sqlStmt, echoTag)
	if err1 != nil {
		log.Printf("error on query: err = %+v", err1)
		return nil, err1
	}
	defer rows.Close()
	for rows.Next() {

		var fileArea string
		var fileName string
		var fileDesc string
		var fileTime *int64

		err2 := rows.Scan(&fileArea, &fileName, &fileDesc, &fileTime)
		if err2 != nil {
			log.Printf("error on scan: err = %+v", err2)
			return nil, err2
		}

		newFile := NewFile()
		newFile.SetArea(fileArea)
		newFile.SetDesc(fileDesc)
		newFile.SetFile(fileName)
		if fileTime != nil {
			newFile.SetUnixTime(*fileTime)
		}

		result = append(result, *newFile)
	}

	ConnTransaction.Commit()

	return result, nil
}

func (self *FileMapper) CheckFileExists(tic File) (bool, error) {
	return true, nil
}

func (self *FileMapper) RegisterFile(tic File) error {

	storageManager := self.restoreStorageManager()

	query1 := "INSERT INTO `file` ( `fileName`, `fileArea`, `fileDesc`, `fileTime` ) VALUES ( ?, ?, ?, ? )"

	var params []interface{}
	params = append(params, tic.GetFile())
	params = append(params, tic.GetArea())
	params = append(params, tic.GetDesc())
	params = append(params, tic.GetUnixTime())

	storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		return err
	})

	return nil
}

func (self *FileMapper) GetAreaByName(areaName string) (*FileArea, error) {

	storageManager := self.restoreStorageManager()
	conn := storageManager.GetConnection() // TODO

	var result *FileArea

	query1 := "SELECT `areaName`, `areaPath`, `areaSummary` FROM `filearea` WHERE `areaName` = $1"
	log.Printf("query = %s params = %s", query1, areaName)
	rows, err2 := conn.Query(query1, areaName)
	if err2 != nil {
		return nil, err2
	}
	defer rows.Close()

	for rows.Next() {

		var areaName1 string
		var areaPath1 string
		var areaSummary1 string

		err3 := rows.Scan(&areaName1, &areaPath1, &areaSummary1)
		if err3 != nil {
			return nil, err3
		}

		log.Printf("row: areaName = %s areaPath = %s areaSummary = %s", areaName1, areaPath1, areaSummary1)

		area := NewFileArea()
		area.SetName(areaName1)
		area.Path = areaPath1
		area.Summary = areaSummary1

		result = area

	}

	return result, nil
}

func (self FileMapper) GetMessageNewCount() (int, error) {
	return 0, nil
}

func (self FileMapper) restoreStorageManager() *storage.StorageManager {
	managerPtr := self.registry.Get("StorageManager")
	if manager, ok := managerPtr.(*storage.StorageManager); ok {
		return manager
	} else {
		panic("no storage manager")
	}
}

func (self FileMapper) GetFileAbsolutePath(areaName string, name string) string {
	boxDirectory := cmn.GetFilesDirectory()
	path := filepath.Join(boxDirectory, areaName, name)
	return path
}

func (self *FileMapper) GetFileBoxAbsolutePath(areaName string) string {
	boxDirectory := cmn.GetFilesDirectory()
	path := filepath.Join(boxDirectory, areaName)
	return path
}

func (self *FileMapper) RemoveFilesByAreaName(areaName string) error {
	storageManager := self.restoreStorageManager()

	query1 := "DELETE FROM `file` WHERE `fileArea` = $1"
	var params []interface{}
	params = append(params, areaName)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}

func (self *FileMapper) RemoveAreaByName(areaName string) error {
	storageManager := self.restoreStorageManager()

	query1 := "DELETE FROM `filearea` WHERE `areaName` = $1"
	var params []interface{}
	params = append(params, areaName)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}
