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
	migrationDate := time.Date(2020, time.November, 4, 18, 7, 0, 0, migrationLocation)
	migrations.Register(migrationDate,
		nil,
		func(conn *sql.DB) error {

			query1 := "ALTER TABLE `message` ADD `msgOrigAddr` varchar(32) DEFAULT ''"
			if _, err := conn.Exec(query1); err != nil {
				return err
			}

			return nil

		},
	)
}
