package file

import (
	"database/sql"
	"log"
)

type FileManager struct {
	conn *sql.DB
}

type FileArea struct {
	Name string /* */
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
		"    areaOrder INTEGER NOT NULL," +
		"    UNIQUE(areaName)" +
		")"
	log.Printf("sqlStmt = %s", sqlStmt)
	self.conn.Exec(sqlStmt)

	/* Create file */
	sqlStmt1 := "CREATE TABLE IF NOT EXISTS file (" +
		"    fileId INTEGER NOT NULL PRIMARY KEY," +
		"    fileName CHAR(64) NOT NULL," +
		"    fileArea CHAR(64) NOT NULL," +
		"    fileDesc TEXT" +
		")"
	log.Printf("sqlStmt = %s", sqlStmt1)
	self.conn.Exec(sqlStmt1)

	return nil
}

func (self *FileManager) GetAreas() ([]*FileArea, error) {

	var areas []*FileArea

	/* Restore parameters */
	sqlStmt := "SELECT `areaName` FROM `filearea`"
	rows, err2 := self.conn.Query(sqlStmt)
	if err2 != nil {
		return nil, err2
	}
	defer rows.Close()
	for rows.Next() {

		var areaName string

		err3 := rows.Scan(&areaName)
		if err3 != nil {
			return nil, err3
		}

		area := NewFileArea()
		area.Name = areaName

		areas = append(areas, area)

	}

	return areas, nil
}

func (self *FileManager) CreateFileArea(a *FileArea) error {

	log.Printf("Create file area: %+v", a)

	/* Insert new one area */
	sqlStmt1 := "INSERT INTO `file` ( `areaName`, `areaType`, `areaPath`, `areaSummary`, `areaOrder` ) VALUES ( ?, '', '', '', 0 )"
	stmt1, err2 := self.conn.Prepare(sqlStmt1)
	if err2 != nil {
		return err2
	}
	_, err3 := stmt1.Exec(a.Name)
	log.Printf("err3 = %+v", err3)
	if err3 != nil {
		return err3
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

	sqlStmt := "SELECT `fileName`, `fileDesc` FROM `file` WHERE `fileArea` = $1"
	log.Printf("sql = %q echoTag = %q", sqlStmt, echoTag)
	rows, err1 := ConnTransaction.Query(sqlStmt, echoTag)
	if err1 != nil {
		return nil, err1
	}
	defer rows.Close()
	for rows.Next() {

		var fileName string
		var fileDesc string

		err2 := rows.Scan(&fileName, &fileDesc)
		if err2 != nil{
			return nil, err2
		}

		tic := NewTicFile()
		tic.Desc = fileDesc
		tic.File = fileName

		result = append(result, tic)
	}

	ConnTransaction.Commit()

	return result, nil
}

func (self *FileManager) CheckFileExists(tic *TicFile) (bool, error) {
	return true, nil
}

func (self *FileManager) RegisterFile(tic *TicFile) (error) {

	/* Insert new one area */
	sqlStmt1 := "INSERT INTO `file` ( `fileName`, `fileArea`, `fileDesc` ) VALUES ( ?, ?, ? )"
	stmt1, err2 := self.conn.Prepare(sqlStmt1)
	if err2 != nil {
		return err2
	}
	_, err3 := stmt1.Exec(tic.File, tic.Area, tic.Desc)
	log.Printf("err3 = %+v", err3)
	if err3 != nil {
		return err3
	}

	return nil
}

