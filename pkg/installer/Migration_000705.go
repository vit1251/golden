package installer

import (
	"database/sql"
)

func migration_000705_Up(conn *sql.DB) error {
	query1 := "ALTER TABLE \"stat\" ADD \"statPacketOut\" INTEGER DEFAULT 0"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}

func init() {
	migrations.Set("2020-10-23 00:07:05", nil, migration_000705_Up)
}
