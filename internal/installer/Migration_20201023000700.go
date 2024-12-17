package installer

import (
	"database/sql"
)

func migration_000700_Up(conn *sql.DB) error {
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

func init() {
	migrations.Set("2020-10-23 00:07:00", nil, migration_000700_Up)
}
