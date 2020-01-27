package area

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/setup"
	"log"
)

type AreaManager struct {
	Path       string       /* Area path        */
}

func NewAreaManager() (*AreaManager) {
	am := new(AreaManager)
	basePath := setup.GetBasePath()
	am.Path = basePath
	am.Rescan()
	return am
}

func (self *AreaManager) checkSchema(conn *sql.DB) (error) {

	sqlStmt := "CREATE TABLE IF NOT EXISTS area (" +
		    "    areaId INTEGER NOT NULL PRIMARY KEY," +
		    "    areaName CHAR(64) NOT NULL," +
		    "    areaType CHAR(64) NOT NULL," +
		    "    areaPath CHAR(64) NOT NULL," +
		    "    areaSummary CHAR(64) NOT NULL," +
		    "    areaOrder INTEGER NOT NULL," +
		    "    UNIQUE(areaName)" +
		    ")"
	log.Printf("sqlStmt = %s", sqlStmt)
	_, err1 := conn.Exec(sqlStmt)
	log.Printf("createSchema: err = %v", err1)
//	if err1 != nil {
//	}

	return err1

}

func (self *AreaManager) Rescan() {

	/* Open message base */
	messageManager := msg.NewMessageManager()

	/* Preload echo areas */
	areas, err3 := messageManager.GetAreaList2()
	if err3 != nil {
		panic(err3)
	}

	/* Reset areas */
	for _, area := range areas {
		log.Printf("area = %q", area)
		a := NewArea()
		a.Name = area.Name
		a.MessageCount = area.Count
		self.Register(a)
	}

}

func (self *AreaManager) updateMsgCount(areas []*Area) {

	/* Open message base */
	messageManager := msg.NewMessageManager()

	var msgCount int
	var msgNewCount int

	for _, area := range areas {
		msgs, err1 := messageManager.GetMessageHeaders(area.Name)
		if err1 != nil {
			panic(err1)
		}

		/* Reset counter */
		msgCount = 0
		msgNewCount = 0
		
		/* Restart*/
		for _, m := range msgs {
			if m.ViewCount > 0 {
				msgCount += 1
			} else {
				msgCount += 1
				msgNewCount += 1
			}
		}
		area.MessageCount = msgCount
		area.NewMessageCount = msgNewCount
	}

}

func (self *AreaManager) Register(a *Area) error {

	/* Open SQL storage */
	db, err1 := sql.Open("sqlite3", self.Path)
	if err1 != nil {
		return err1
	}
	defer db.Close()

	/* Insert new one area */
	sqlStmt1 := "INSERT INTO `area` ( `areaName`, `areaType`, `areaPath`, `areaSummary`, `areaOrder` ) VALUES ( ?, '', '', '', 0 )"
	stmt1, err2 := db.Prepare(sqlStmt1)
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

func (self *AreaManager) GetAreas() ([]*Area, error) {

	var areas []*Area

	/* Open SQL storage */
	db, err1 := sql.Open("sqlite3", self.Path)
	if err1 != nil {
		return nil, err1
	}
	defer db.Close()

	/* Check schema */
	self.checkSchema(db)

	/* Restore parameters */
	sqlStmt := "SELECT `areaName` FROM `area`"
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

		area := NewArea()
		area.Name = areaName

		areas = append(areas, area)

	}

	/* Update metric count */
	self.updateMsgCount(areas)

	return areas, nil
}


func (self *AreaManager) GetAreaByName(echoTag string) (*Area, error) {

	var result *Area

	/* Open SQL storage */
	db, err1 := sql.Open("sqlite3", self.Path)
	if err1 != nil {
		return nil, err1
	}
	defer db.Close()

	/* Restore parameters */
	sqlStmt := "SELECT `areaName` FROM `area` WHERE `areaName` = ?"
	rows, err2 := db.Query(sqlStmt, echoTag)
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

		area := NewArea()
		area.Name = areaName

		result = area

	}

	return result, nil

}

