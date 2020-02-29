package file

import (
	"database/sql"
	"log"
	"time"
)

type FileManager struct {
	conn *sql.DB
}

type FileArea struct {
	Name string
	Path string
	Summary string
}

func NewFileArea() *FileArea {
	fa := new(FileArea)
	return fa
}

func NewFileManager(conn *sql.DB) *FileManager {
	fm := new(FileManager)
	fm.conn = conn
	fm.checkSchema()
	return fm
}

func (self *FileManager) checkSchema() error {

	/* Create file area */
	sqlStmt := "CREATE TABLE IF NOT EXISTS filearea (" +
		"    areaId INTEGER NOT NULL PRIMARY KEY," +
		"    areaName CHAR(64) NOT NULL," +
		"    areaType CHAR(64) NOT NULL," +
		"    areaPath CHAR(64) NOT NULL," +
		"    areaSummary CHAR(64) NOT NULL," +
		"    areaOrder INTEGER NOT NULL" +
		")"
	log.Printf("sqlStmt = %s", sqlStmt)
	self.conn.Exec(sqlStmt)

	/* Create file */
	sqlStmt1 := "CREATE TABLE IF NOT EXISTS file (" +
		"    fileId INTEGER NOT NULL PRIMARY KEY," +
		"    fileName CHAR(64) NOT NULL," +
		"    fileArea CHAR(64) NOT NULL," +
		"    fileTime INTEGER NOT NULL," +
		"    fileDesc TEXT" +
		")"
	log.Printf("sqlStmt = %s", sqlStmt1)
	self.conn.Exec(sqlStmt1)

	/* Create index on msgHash */
	query3 := "CREATE INDEX IF NOT EXISTS \"idx_file_fileArea\" ON \"file\" (\"fileArea\" ASC)"
	if _, err := self.conn.Exec(query3); err != nil {
		log.Printf("Error create \"file\" storage: err = %+v", err)
	}

	return nil
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

