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
	migrationDate := time.Date(2020, time.November, 21, 3, 59, 0, 0, migrationLocation)
	migrations.Register(migrationDate,
		func(conn *sql.DB) error {
			return fmt.Errorf("not implemented")
		},
		func(conn *sql.DB) error {

			query1 := "ALTER TABLE `filearea` RENAME COLUMN `areaType` TO `areaMode`"
			if _, err := conn.Exec(query1); err != nil {
				return err
			}

			query2 := "UPDATE `filearea` SET `areaMode` = \"\""
			if _, err := conn.Exec(query2); err != nil {
				return err
			}

			return nil

		},
	)
}
