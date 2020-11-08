package installer

import (
	"database/sql"
)

func migration_20201108092500_Up(conn *sql.DB) error {

	query1 := "CREATE TABLE `twit` (" +
		"    twitId INTEGER NOT NULL PRIMARY KEY," +
		"    twitName VARCHAR(64) NOT NULL" +
		")"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}

	return nil
}

func init() {
	migrations.Set("2020-11-08 09:25:00", nil, migration_20201108092500_Up)
}
