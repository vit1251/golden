package file

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/storage"
	"log"
	"time"
)

type FileManager struct {
	conn *sql.DB
}

type FileArea struct {
	Name    string
	Path    string
	Summary string
	Count   int
}

func NewFileArea() *FileArea {
	fa := new(FileArea)
	return fa
}

func NewFileManager(sm *storage.StorageManager) *FileManager {
	fm := new(FileManager)
	fm.conn = sm.GetConnection()
	return fm
}

func (self *FileManager) GetAreas() ([]*FileArea, error) {

	var areas []*FileArea

	/* Restore parameters */
	sqlStmt := "SELECT `areaName`, `areaPath`, `areaSummary` FROM `filearea`"
	rows, err2 := self.conn.Query(sqlStmt)
	if err2 != nil {
		return nil, err2
	}
	defer rows.Close()
	for rows.Next() {

		var areaName string
		var areaPath string
		var areaSummary string

		err3 := rows.Scan(&areaName, &areaPath, &areaSummary)
		if err3 != nil {
			return nil, err3
		}

		/**/
		if areaSummary == "" {
			areaSummary = "Нет описания"
		}

		area := NewFileArea()
		area.Name = areaName
		area.Path = areaPath
		area.Summary = areaSummary

		areas = append(areas, area)

	}

	return areas, nil
}

func (self *FileManager) GetAreas2() ([]*FileArea, error) {

	var result []*FileArea

	/* Step 2. Start SQL transaction */
	ConnTransaction, err := self.conn.Begin()
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
		if err2 != nil{
			return nil, err2
		}

		area := NewFileArea()
		area.Name = name
		area.Count = count

		result = append(result, area)
	}

	ConnTransaction.Commit()

	return result, nil
}

func (self *FileManager) CreateArea(a *FileArea) error {

	log.Printf("Create file area: %+v", a)

	/* Prepare SQL request */
	query := "INSERT INTO `filearea` (`areaName`, `areaType`, `areaPath`, `areaSummary`, `areaOrder`) VALUES ( ?, '', ?, ?, 0)"

	/* Create area */
	if _, err := self.conn.Exec(query, a.Name, a.Path, a.Summary); err != nil {
		return err
	}

	return nil
}

func (self *FileManager) GetFileHeaders(echoTag string) ([]*TicFile, error) {

	var result []*TicFile

	/* Step 2. Start SQL transaction */
	ConnTransaction, err := self.conn.Begin()
	if err != nil {
		return nil, err
	}

	sqlStmt := "SELECT `fileName`, `fileDesc`, `fileTime` FROM `file` WHERE `fileArea` = $1"
	log.Printf("sql = %q echoTag = %q", sqlStmt, echoTag)
	rows, err1 := ConnTransaction.Query(sqlStmt, echoTag)
	if err1 != nil {
		log.Printf("error on query: err = %+v", err1)
		return nil, err1
	}
	defer rows.Close()
	for rows.Next() {

		var fileName string
		var fileDesc string
		var fileTime *int64

		err2 := rows.Scan(&fileName, &fileDesc, &fileTime)
		if err2 != nil {
			log.Printf("error on scan: err = %+v", err2)
			return nil, err2
		}

		tic := NewTicFile()
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

func (self *FileManager) RegisterFile(tic *TicFile) (error) {

	var unixTime int64 = time.Now().Unix()

	/* Insert new one area */
	sqlStmt1 := "INSERT INTO `file` ( `fileName`, `fileArea`, `fileDesc`, `fileTime` ) VALUES ( ?, ?, ?, ? )"
	stmt1, err2 := self.conn.Prepare(sqlStmt1)
	if err2 != nil {
		return err2
	}
	_, err3 := stmt1.Exec(tic.File, tic.Area, tic.Desc, unixTime )
	log.Printf("err3 = %+v", err3)
	if err3 != nil {
		return err3
	}

	return nil
}

func (self *FileManager) GetAreaByName(areaName string) (*FileArea, error) {

	var result *FileArea

	query1 := "SELECT `areaName`, `areaPath`, `areaSummary` FROM `filearea` WHERE `areaName` = $1"
	log.Printf("query = %s params = %s", query1, areaName)
	rows, err2 := self.conn.Query(query1, areaName)
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
		area.Name = areaName1
		area.Path = areaPath1
		area.Summary = areaSummary1

		result = area

	}

	return result, nil
}

