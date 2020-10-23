package installer

import (
	"database/sql"
)

func migration_000100_Up(conn *sql.DB) error {
	query1 := "CREATE TABLE `settings` (" +
		"    `section` CHAR(64) NOT NULL," +
		"    `name` CHAR(64) NOT NULL," +
		"    `value` CHAR(64) NOT NULL" +
		")"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}

func init() {
	migrations.Set("2020-10-23 00:01:00", nil, migration_000100_Up)
}
