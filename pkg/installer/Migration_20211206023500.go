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
	migrationDate := time.Date(2021, time.December, 6, 2, 35, 0, 0, migrationLocation)
	migrations.Register(
		migrationDate,
		func(conn *sql.DB) error {
			return fmt.Errorf("not implemented")
		},
		func(conn *sql.DB) error {
			query1 := "ALTER TABLE `file` ADD `fileViewCount` INTEGER DEFAULT 0"
			if _, err := conn.Exec(query1); err != nil {
				return err
			}
			return nil
		},
	)
}
