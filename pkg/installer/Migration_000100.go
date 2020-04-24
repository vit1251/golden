package installer

import (
	"database/sql"
)

type Migration_000100 struct {
	IMigration
}

func (m *Migration_000100) Up(conn *sql.DB) error {
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
