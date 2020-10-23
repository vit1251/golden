package installer

import (
	"database/sql"
)

func migration_000301_Up(conn *sql.DB) error {
	query2 := "CREATE INDEX `idx_message_msgHash` ON `message` (`msgHash` ASC)"
	if _, err := conn.Exec(query2); err != nil {
		return err
	}
	return nil
}

func init() {
	migrations.Set("2020-10-23 00:03:01", nil, migration_000301_Up)
}
