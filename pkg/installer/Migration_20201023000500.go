package installer

import (
	"database/sql"
)

func migration_000500_Up(conn *sql.DB) error {
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

func init() {
	migrations.Set("2020-10-23 00:05:00", nil, migration_000500_Up)
}