package installer

import (
	"database/sql"
	"fmt"
	"time"
)

func init() {
	migrationLocation, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		panic(err)
	}
	migrationDate := time.Date(2020, time.October, 23, 0, 7, 5, 0, migrationLocation)
	migrations.Register(migrationDate,
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
