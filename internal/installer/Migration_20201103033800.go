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
	migrationDate := time.Date(2020, time.November, 3, 3, 38, 0, 0, migrationLocation)
	migrations.Register(migrationDate,
		nil,
		func(conn *sql.DB) error {

			/* Add "nmOrigAddr" */
			query1 := "ALTER TABLE `netmail` ADD `nmOrigAddr` varchar(64) DEFAULT ''"
			if _, err := conn.Exec(query1); err != nil {
				return err
			}

			/* Add "nmOrigAddr" */
			query2 := "ALTER TABLE `netmail` ADD `nmDestAddr` varchar(64) DEFAULT ''"
			if _, err := conn.Exec(query2); err != nil {
				return err
			}

			return nil
		},
	)
}
