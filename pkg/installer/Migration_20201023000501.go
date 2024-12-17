package installer

import (
	"database/sql"
)

func migration_000501_Up(conn *sql.DB) error {
	query1 := "CREATE INDEX `idx_file_fileArea` ON `file` (`fileArea` ASC)"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}

func init() {
	migrations.Set("2020-10-23 00:05:01", nil, migration_000501_Up)
}
