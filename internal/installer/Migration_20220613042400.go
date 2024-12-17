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
	migrationDate := time.Date(2022, time.June, 13, 4, 24, 0, 0, migrationLocation)
	migrations.Register(
		migrationDate,
		func(conn *sql.DB) error {
			return fmt.Errorf("not implemented")
		},
		func(conn *sql.DB) error {
			/* Step 1. Create new `indexName` column */
			query1 := "ALTER TABLE `file` ADD `indexName` CHAR(128) NOT NULL DEFAULT ''"
			if _, err := conn.Exec(query1); err != nil {
				return err
			}
			/* Step 2. Populate `indexName` with `fileName` value */
			query2 := "UPDATE `file` SET `indexName` = `fileName` WHERE `indexName` = ''"
			if _, err := conn.Exec(query2); err != nil {
				return err
			}
			return nil
		},
	)
}
