package installer

import (
	"database/sql"
)

type Migration_20200402_0004 struct {
}

func (m *Migration_20200402_0004) Up(conn *sql.DB) {
	sqlStmt := "CREATE TABLE IF NOT EXISTS netmail (" +
		"    nmId INTEGER NOT NULL PRIMARY KEY," +
		"    nmHash CHAR(64) NOT NULL," +
		"    nmFrom CHAR(64) NOT NULL," +
		"    nmTo CHAR(64) NOT NULL," +
		"    nmSubject CHAR(512) NOT NULL," +
		"    nmBody TEXT NOT NULL," +
		"    nmDate INTEGER NOT NULL," +
		"    nmViewCount INTEGER DEFAULT 0" +
		")"
	conn.Exec(sqlStmt)
}
