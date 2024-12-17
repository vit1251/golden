package installer

import (
	"database/sql"
)

func migration_000302_Up(conn *sql.DB) error {
	query3 := "CREATE INDEX `idx_message_msgArea` ON `message` (`msgArea` ASC)"
	if _, err := conn.Exec(query3); err != nil {
		return err
	}
	return nil
}

func init() {
	migrations.Set("2020-10-23 00:03:02", nil, migration_000302_Up)
}
