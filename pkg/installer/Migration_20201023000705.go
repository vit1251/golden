package installer

import (
	"database/sql"
	"fmt"
)

func init() {
	migrations.Set("2020-10-23 00:07:05",
		func(conn *sql.DB) error {
			return fmt.Errorf("not implemented")
		},
		func(conn *sql.DB) error {
			query1 := "ALTER TABLE \"stat\" ADD \"statPacketOut\" INTEGER DEFAULT 0"
			if _, err := conn.Exec(query1); err != nil {
				return err
			}
			return nil
		},
	)
}
