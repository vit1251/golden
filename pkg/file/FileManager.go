package file

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
	"log"
	"time"
)

type FileManager struct {
	registry      *registry.Container
}

type FileArea struct {
	name    string
	Path    string
	Summary string
	Count   int
}

func (self *FileArea) SetName(name string) {
	self.name = name
}

func (self *FileArea) Name() string {
	return self.name
}

func NewFileArea() *FileArea {
	fa := new(FileArea)
	return fa
}

func NewFileManager(r *registry.Container) *FileManager {
	fm := new(FileManager)
	fm.registry = r
	return fm
}

func (self *FileManager) GetAreas() ([]*FileArea, error) {

	storageManager := self.restoreStorageManager()

	var areas []*FileArea

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

		areas = append(areas, area)

		return nil
	})

	return areas, err1
}

func (self *FileManager) GetAreas2() ([]*FileArea, error) {

	storageManager := self.restoreStorageManager()
	conn := storageManager.GetConnection() // TODO

	var result []*FileArea

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

		result = append(result, area)
	}

	ConnTransaction.Commit()

	return result, nil
}

func (self *FileManager) CreateFileArea(a *FileArea) error {

	storageManager := self.restoreStorageManager()
	conn := storageManager.GetConnection() // TODO

	var name string = a.Name()
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

func (self *FileManager) GetFileHeaders(echoTag string) ([]*TicFile, error) {

	storageManager := self.restoreStorageManager()
	conn := storageManager.GetConnection() // TODO

	var result []*TicFile

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

		tic := NewTicFile()
		tic.SetArea(fileArea)
		tic.Desc = fileDesc
		tic.File = fileName
		if fileTime != nil {
			tic.SetUnixTime(*fileTime)
		}

		result = append(result, tic)
	}

	ConnTransaction.Commit()

	return result, nil
}

func (self *FileManager) CheckFileExists(tic *TicFile) (bool, error) {
	return true, nil
}

func (self *FileManager) RegisterFile(tic *TicFile) error {

	storageManager := self.restoreStorageManager()
	conn := storageManager.GetConnection() // TODO

	var unixTime int64 = time.Now().Unix()

	/* Insert new one area */
	sqlStmt1 := "INSERT INTO `file` ( `fileName`, `fileArea`, `fileDesc`, `fileTime` ) VALUES ( ?, ?, ?, ? )"
	stmt1, err2 := conn.Prepare(sqlStmt1)
	if err2 != nil {
		return err2
	}
	areaName := tic.GetArea()
	_, err3 := stmt1.Exec(tic.File, areaName, tic.Desc, unixTime)
	log.Printf("err3 = %+v", err3)
	if err3 != nil {
		return err3
	}

	return nil
}

func (self *FileManager) GetAreaByName(areaName string) (*FileArea, error) {

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

func (self *FileManager) GetMessageNewCount() (int, error) {
	return 0, nil
}

func (self *FileManager) restoreStorageManager() *storage.StorageManager {
	managerPtr := self.registry.Get("StorageManager")
	if manager, ok := managerPtr.(*storage.StorageManager); ok {
		return manager
	} else {
		panic("no storage manager")
	}
}
