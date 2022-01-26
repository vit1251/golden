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
	migrationDate := time.Date(2020, time.November, 8, 12, 3, 0, 0, migrationLocation)
	migrations.Register(migrationDate,
		nil,
		func(conn *sql.DB) error {
			query1 := "CREATE TABLE `draft` (" +
				"    draftId INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT," +
				"    draftUUID VARCHAR(128) NOT NULL," +
				"    draftDestAddr VARCHAR(64) NOT NULL DEFAULT ''," +
				"    draftDest VARCHAR(64) NOT NULL DEFAULT ''," +
				"    draftSubject VARCHAR(64) NOT NULL DEFAULT ''," +
				"    draftReply VARCHAR(64) NOT NULL DEFAULT ''," +
				"    draftArea VARCHAR(64) NOT NULL DEFAULT ''," +
				"    draftBody TEXT NOT NULL DEFAULT ''," +
				"    draftDone INTEGER NOT NULL DEFAULT 0" +
				")"
			if _, err := conn.Exec(query1); err != nil {
				return err
			}
			return nil
		},
	)
}
