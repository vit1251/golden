package installer

import (
	"database/sql"
)

type Migration_000600 struct {
	IMigration
}

func (m *Migration_000600) Up(conn *sql.DB) error {
	query1 := "CREATE TABLE `netmail` (" +
		"    nmId INTEGER NOT NULL PRIMARY KEY," +
		"    nmHash CHAR(64) NOT NULL," +
		"    nmFrom CHAR(64) NOT NULL," +
		"    nmTo CHAR(64) NOT NULL," +
		"    nmSubject CHAR(512) NOT NULL," +
		"    nmBody TEXT NOT NULL," +
		"    nmDate INTEGER NOT NULL," +
		"    nmViewCount INTEGER DEFAULT 0" +
		")"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}
