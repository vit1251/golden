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
	migrationDate := time.Date(2020, time.November, 15, 2, 9, 0, 0, migrationLocation)
	migrations.Register(migrationDate,
		nil,
		func(conn *sql.DB) error {

			query1 := "INSERT INTO `settings` (`section`,`name`,`value`) VALUES (?, ?, ?)"
			if _, err := conn.Exec(query1, "netmail", "Charset", "CP866"); err != nil {
				return err
			}

			return nil

		},
	)
}
