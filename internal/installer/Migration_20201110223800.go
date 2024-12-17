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
	migrationDate := time.Date(2020, time.November, 10, 22, 38, 0, 0, migrationLocation)
	migrations.Register(migrationDate,
		nil,
		func(conn *sql.DB) error {
			query1 := "ALTER TABLE `filearea` RENAME COLUMN `areaPath` TO `areaCharset`"
			if _, err := conn.Exec(query1); err != nil {
				return err
			}
			query2 := "UPDATE `filearea` SET `areaCharset` = \"CP866\""
			if _, err := conn.Exec(query2); err != nil {
				return err
			}
			return nil
		},
	)
}
