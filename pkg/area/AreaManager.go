package area

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vit1251/golden/pkg/msg"
	"log"
)

type AreaManager struct {
	MessageManager *msg.MessageManager
	conn *sql.DB
}

func NewAreaManager(conn *sql.DB, mm *msg.MessageManager) (*AreaManager) {
	log.Printf("NewAreaManager: conn = %+v", conn)
	am := new(AreaManager)
	am.conn = conn
	am.MessageManager = mm
	am.checkSchema()
	am.Rescan()
	return am
}

func (self *AreaManager) checkSchema() {
	query1 := "CREATE TABLE IF NOT EXISTS area (" +
		    "    areaId INTEGER NOT NULL PRIMARY KEY," +
		    "    areaName CHAR(64) NOT NULL," +
		    "    areaType CHAR(64) NOT NULL," +
		    "    areaPath CHAR(64) NOT NULL," +
		    "    areaSummary CHAR(64) NOT NULL," +
		    "    areaOrder INTEGER NOT NULL" +
		    ")"
	log.Printf("sqlStmt = %s", query1)
	if _, err := self.conn.Exec(query1); err != nil {
		log.Printf("Error create \"area\" storage: err = %+v", err)
	}

	/* Create index on msgHash */
	query2 := "CREATE INDEX \"idx_area_areaName\" ON \"area\" (\"areaName\" ASC)"
	if _, err := self.conn.Exec(query2); err != nil {
		log.Printf("Error create \"area\" storage: err = %+v", err)
	}

}

func (self *AreaManager) Rescan() {

	/* Preload echo areas */
	areas, err3 := self.MessageManager.GetAreaList2()
	if err3 != nil {
		panic(err3)
	}

	/* Reset areas */
	for _, area := range areas {
		log.Printf("area = %q", area)
		a := NewArea()
		a.SetName(area.Name)
		a.MessageCount = area.Count
	}

}

func (self *AreaManager) updateMsgCount(areas []*Area) {

	var msgCount int
	var msgNewCount int

	for _, area := range areas {
		msgs, err1 := self.MessageManager.GetMessageHeaders(area.Name)
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

	/* Insert new one area */
	sqlStmt1 := "INSERT INTO `area` ( `areaName`, `areaType`, `areaPath`, `areaSummary`, `areaOrder` ) VALUES ( ?, '', '', '', 0 )"
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

func (self *AreaManager) GetAreas() ([]*Area, error) {

	var areas []*Area

	/* Restore parameters */
	sqlStmt := "SELECT `areaName` FROM `area`"
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

	/* Restore parameters */
	sqlStmt := "SELECT `areaName` FROM `area` WHERE `areaName` = ?"
	rows, err2 := self.conn.Query(sqlStmt, echoTag)
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

