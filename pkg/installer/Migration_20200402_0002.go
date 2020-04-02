package installer

import (
	"database/sql"
	"log"
)

type Migration_20200402_0002 struct {
}

func (m *Migration_20200402_0002) Up(conn *sql.DB) {
	query1 := "CREATE TABLE IF NOT EXISTS area (" +
		"    areaId INTEGER NOT NULL PRIMARY KEY," +
		"    areaName CHAR(64) NOT NULL," +
		"    areaType CHAR(64) NOT NULL," +
		"    areaPath CHAR(64) NOT NULL," +
		"    areaSummary CHAR(64) NOT NULL," +
		"    areaOrder INTEGER NOT NULL" +
		")"
	log.Printf("sqlStmt = %s", query1)
	if _, err := conn.Exec(query1); err != nil {
		log.Printf("Error create \"area\" storage: err = %+v", err)
	}

	/* Create index on msgHash */
	query2 := "CREATE UNIQUE INDEX IF NOT EXISTS \"uniq_area_areaName\" ON \"area\" (\"areaName\" ASC)"
	if _, err := conn.Exec(query2); err != nil {
		log.Printf("Error create \"area\" storage: err = %+v", err)
	}

}
