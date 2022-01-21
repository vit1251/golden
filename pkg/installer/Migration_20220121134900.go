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
	migrationDate := time.Date(2022, time.January, 21, 13, 49, 0, 0, migrationLocation)
	migrations.Register(migrationDate,
		nil,
		func(db *sql.DB) error {
			query1 := "CREATE TABLE `stat_mailer` (" +
				"    statMailerId INTEGER NOT NULL PRIMARY KEY," +
				"    statMailerSessionStart INTEGER NOT NULL," +
				"    statMailerSessionStop INTEGER NOT NULL," +
				"    statMailerPacketRXcount INTEGER DEFAULT 0," +
				"    statMailerPacketTXcount INTEGER DEFAULT 0," +
				"    statMailerFileRXcount INTEGER DEFAULT 0," +
				"    statMailerFileTXcount INTEGER DEFAULT 0," +
				"    statMailerSummary VARCHAR(512) NOT NULL DEFAULT ''" +
				")"
			if _, err := db.Exec(query1); err != nil {
				return err
			}
			return nil
		},
	)
}
