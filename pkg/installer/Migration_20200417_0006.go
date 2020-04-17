package installer

import (
	"database/sql"
	"log"
)

type Migration_20200417_0006 struct {
}

func (m *Migration_20200417_0006) Up(conn *sql.DB) {

	query3 := "ALTER TABLE \"stat\" ADD \"statPacketIn\" INTEGER DEFAULT 0"
	log.Printf("query = %s", query3)
	conn.Exec(query3)

	query4 := "ALTER TABLE \"stat\" ADD \"statPacketOut\" INTEGER DEFAULT 0"
	log.Printf("query = %s", query4)
	conn.Exec(query4)

}
