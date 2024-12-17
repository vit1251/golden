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
	migrationDate := time.Date(2020, time.November, 8, 9, 25, 0, 0, migrationLocation)
	migrations.Register(migrationDate,
		nil,
		func(conn *sql.DB) error {

			query1 := "CREATE TABLE `twit` (" +
				"    twitId INTEGER NOT NULL PRIMARY KEY," +
				"    twitName VARCHAR(64) NOT NULL" +
				")"
			if _, err := conn.Exec(query1); err != nil {
				return err
			}

			return nil
		},
	)
}
