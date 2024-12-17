package installer

import (
	"database/sql"
)

func migration_000101_Up(conn *sql.DB) error {
	query1 := "CREATE UNIQUE INDEX `uniq_settings` ON `settings` (`section`, `name`)"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}

func init() {
	migrations.Set("2020-10-23 00:01:01", nil, migration_000101_Up)
}
