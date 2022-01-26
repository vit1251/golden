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
	migrationDate := time.Date(2020, time.October, 23, 3, 14, 0, 0, migrationLocation)
	migrations.Register(migrationDate,
		nil,
		func(conn *sql.DB) error {
			query1 := "ALTER TABLE `message` ADD `msgPacket` BLOB"
			_, err := conn.Exec(query1)
			return err
		},
	)
}
