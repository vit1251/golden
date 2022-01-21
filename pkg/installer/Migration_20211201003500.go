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
	migrationDate := time.Date(2021, time.December, 1, 0, 35, 0, 0, migrationLocation)
	migrations.Register(migrationDate,
		func(conn *sql.DB) error {
			return fmt.Errorf("not implemented")
		},
		func(conn *sql.DB) error {
			query1 := "INSERT INTO `settings` (`section`,`name`,`value`) VALUES (?, ?, ?)"
			if _, err := conn.Exec(query1, "echomail", "Charset", "CP866"); err != nil {
				return err
			}
			return nil
		},
	)
}
