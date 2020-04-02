package installer

import (
	"database/sql"
	"log"
)

type Migration_20200402_0001 struct {
}

func (m *Migration_20200402_0001) Up(conn *sql.DB) {

	query1 :=  "CREATE TABLE IF NOT EXISTS message (" +
		"    msgId INTEGER NOT NULL PRIMARY KEY," +
		"    msgMsgId CHAR(16) NOT NULL," +
		"    msgHash CHAR(16) NOT NULL," +
		"    msgDate INTEGER NOT NULL," +
		"    msgViewCount INTEGER DEFAULT 0," +
		"    msgArea CHAR(64) NOT NULL," +
		"    msgFrom TEXT NOT NULL," +
		"    msgTo TEXT NOT NULL," +
		"    msgSubject TEXT NOT NULL," +
		"    msgContent TEXT NOT NULL" +
		")"
	log.Printf("query = %+v", query1)
	if _, err := conn.Exec(query1); err != nil {
		log.Printf("Error create \"message\" storage: err = %+v", err)
	}

	/* Create index on msgHash */
	query2 := "CREATE INDEX IF NOT EXISTS \"idx_message_msgHash\" ON \"message\" (\"msgHash\" ASC)"
	if _, err := conn.Exec(query2); err != nil {
		log.Printf("Error create \"message\" storage: err = %+v", err)
	}

	/* Create index on msgHash */
	query3 := "CREATE INDEX IF NOT EXISTS \"idx_message_msgArea\" ON \"message\" (\"msgArea\" ASC)"
	if _, err := conn.Exec(query3); err != nil {
		log.Printf("Error create \"message\" storage: err = %+v", err)
	}

}
