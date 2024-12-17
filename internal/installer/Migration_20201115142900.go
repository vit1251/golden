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
	migrationDate := time.Date(2020, time.November, 15, 14, 29, 0, 0, migrationLocation)
	migrations.Register(migrationDate,
		func(conn *sql.DB) error {
			return fmt.Errorf("no implemented")
		},
		func(conn *sql.DB) error {
			query1 := "ALTER TABLE `netmail` ADD `nmPacket` BLOB"
			_, err := conn.Exec(query1)
			return err
		},
	)
}
