package installer

import (
	"database/sql"
	"log"
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
	_, err1 := conn.Exec(sqlStmt)
	log.Printf("err1 = %+v", err1)

	query3 := "ALTER TABLE \"netmail\" ADD \"nmMsgId\" CHAR(64)"
	log.Printf("query = %s", query3)
	_, err2 := conn.Exec(query3)
	log.Printf("err2 = %+v", err2)

}
