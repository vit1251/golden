package installer

import (
	"database/sql"
)

type Migration_000700 struct {
	IMigration
}

func (m *Migration_000700) Up(conn *sql.DB) error {
	query1 := "CREATE TABLE `stat` (" +
		"    statId INTEGER NOT NULL PRIMARY KEY," +
		"    statDate CHAR(10) NOT NULL," +
		"    statMessageRXcount INTEGER DEFAULT 0," +
		"    statMessageTXcount INTEGER DEFAULT 0," +
		"    statFileRXcount INTEGER DEFAULT 0," +
		"    statFileTXcount INTEGER DEFAULT 0" +
		")"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}
