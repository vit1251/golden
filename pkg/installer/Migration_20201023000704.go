package installer

import (
	"database/sql"
	"time"
)

func init() {
	migrationLocation, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		panic(err)
	}
	migrationDate := time.Date(2020, time.October, 23, 0, 7, 4, 0, migrationLocation)
	migrations.Register(migrationDate,
		nil,
		func(conn *sql.DB) error {
			query1 := "ALTER TABLE \"stat\" ADD \"statPacketIn\" INTEGER DEFAULT 0"
			if _, err := conn.Exec(query1); err != nil {
				return err
			}
			return nil
		},
	)
}
