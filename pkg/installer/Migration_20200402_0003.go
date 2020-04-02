package installer

import (
	"database/sql"
	"log"
)

type Migration_20200402_0003 struct {
}

func (m *Migration_20200402_0003) Up(conn *sql.DB) {
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
	conn.Exec(sqlStmt)

	/* Create file */
	sqlStmt1 := "CREATE TABLE IF NOT EXISTS file (" +
		"    fileId INTEGER NOT NULL PRIMARY KEY," +
		"    fileName CHAR(64) NOT NULL," +
		"    fileArea CHAR(64) NOT NULL," +
		"    fileTime INTEGER NOT NULL," +
		"    fileDesc TEXT" +
		")"
	log.Printf("sqlStmt = %s", sqlStmt1)
	conn.Exec(sqlStmt1)

	/* Create index on msgHash */
	query3 := "CREATE INDEX IF NOT EXISTS \"idx_file_fileArea\" ON \"file\" (\"fileArea\" ASC)"
	if _, err := conn.Exec(query3); err != nil {
		log.Printf("Error create \"file\" storage: err = %+v", err)
	}
}

