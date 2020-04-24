package installer

import (
	"database/sql"
)

type Migration_000101 struct {
	IMigration
}

func (m *Migration_000101) Up(conn *sql.DB) error {
	query1 := "CREATE UNIQUE INDEX `uniq_settings` ON `settings` (`section`, `name`)"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}
