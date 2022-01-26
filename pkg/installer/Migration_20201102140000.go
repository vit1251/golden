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
	migrationDate := time.Date(2020, time.November, 2, 14, 0, 0, 0, migrationLocation)
	migrations.Register(migrationDate,
		nil,
		func(conn *sql.DB) error {
			query1 := "ALTER TABLE `area` ADD `areaCharset` CHAR(16) NOT NULL DEFAULT \"CP866\""
			_, err := conn.Exec(query1)
			return err
		},
	)
}
