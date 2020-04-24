package installer

import (
	"database/sql"
)

type Migration_000500 struct {
	IMigration
}

func (m *Migration_000500) Up(conn *sql.DB) error {
	query1 := "CREATE TABLE `file` (" +
		"    fileId INTEGER NOT NULL PRIMARY KEY," +
		"    fileName CHAR(64) NOT NULL," +
		"    fileArea CHAR(64) NOT NULL," +
		"    fileTime INTEGER NOT NULL," +
		"    fileDesc TEXT" +
		")"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}

