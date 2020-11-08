package installer

import (
	"database/sql"
)

func migration_20201108120300_Up(conn *sql.DB) error {

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
}

func init() {
	migrations.Set("2020-11-08 12:03:00", nil, migration_20201108120300_Up)
}
