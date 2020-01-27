package file

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/setup"
	"log"
)

type FileManager struct {
	Path string /* Path */
}

type FileArea struct {
	Name string /* */
}

func NewFileArea() (*FileArea) {
	fa := new(FileArea)
	return fa
}

func NewFileManager() (*FileManager) {
	fm := new(FileManager)
	fm.Path = setup.GetBasePath()
	return fm
}

func (self *FileManager) SetInboundDirectory(inboundDirectory string) {
}

func (self *FileManager) ProcessTic(filename string) {
}

func (self *FileManager) checkSchema(conn *sql.DB) (error) {

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
	conn.Exec(sqlStmt)

	return nil
}

func (self *FileManager) GetAreas() ([]*FileArea, error) {

	var areas []*FileArea

	/* Open SQL storage */
	db, err1 := sql.Open("sqlite3", self.Path)
	if err1 != nil {
		return nil, err1
	}
	defer db.Close()

	/* Check schema */
	self.checkSchema(db)

	/* Restore parameters */
	sqlStmt := "SELECT `areaName` FROM `filearea`"
	rows, err2 := db.Query(sqlStmt)
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