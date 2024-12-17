package installer

import (
	"database/sql"
)

func migration_000703_Up(conn *sql.DB) error {
	query1 := "ALTER TABLE \"stat\" ADD \"statSessionOut\" INTEGER DEFAULT 0"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}

func init() {
	migrations.Set("2020-10-23 00:07:03", nil, migration_000703_Up)
}
