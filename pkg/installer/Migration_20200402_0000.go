package installer

import (
	"database/sql"
	"log"
)

type Migration_20200402_0000 struct {
}

func (m *Migration_20200402_0000) Up(conn *sql.DB) {
	sqlStmt1 := "CREATE TABLE IF NOT EXISTS settings (" +
		"    section CHAR(64) NOT NULL," +
		"    name CHAR(64) NOT NULL," +
		"    value CHAR(64) NOT NULL," +
		"    UNIQUE(section, name)" +
		")"
	log.Printf("sql = %s", sqlStmt1)
	conn.Exec(sqlStmt1)
}
