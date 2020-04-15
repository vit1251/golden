package installer

import (
"database/sql"
	"log"
)

type Migration_20200402_0005 struct {
}

func (m *Migration_20200402_0005) Up(conn *sql.DB) {
	query1 := "CREATE TABLE IF NOT EXISTS `stat` (" +
		"    statId INTEGER NOT NULL PRIMARY KEY," +
		"    statDate CHAR(10) NOT NULL," +
		"    statMessageRXcount INTEGER DEFAULT 0," +
		"    statMessageTXcount INTEGER DEFAULT 0," +
		"    statFileRXcount INTEGER DEFAULT 0," +
		"    statFileTXcount INTEGER DEFAULT 0" +
		")"
	log.Printf("query = %s", query1)
	conn.Exec(query1)

	query2 := "CREATE UNIQUE INDEX IF NOT EXISTS \"uniq_stat_statDate\" ON \"stat\" (\"statDate\" ASC)"
	log.Printf("query = %s", query2)
	conn.Exec(query2)

	query3 := "ALTER TABLE \"stat\" ADD \"statSessionIn\" INTEGER DEFAULT 0"
	log.Printf("query = %s", query3)
	conn.Exec(query3)

	query4 := "ALTER TABLE \"stat\" ADD \"statSessionOut\" INTEGER DEFAULT 0"
	log.Printf("query = %s", query4)
	conn.Exec(query4)
}
