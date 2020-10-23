package installer

import (
	"database/sql"
)

func migration_000300_Up(conn *sql.DB) error {
	query1 :=  "CREATE TABLE `message` (" +
		"    `msgId` INTEGER NOT NULL PRIMARY KEY," +
		"    `msgMsgId` CHAR(16) NOT NULL," +
		"    `msgHash` CHAR(16) NOT NULL," +
		"    `msgDate` INTEGER NOT NULL," +
		"    `msgViewCount` INTEGER DEFAULT 0," +
		"    `msgArea` CHAR(64) NOT NULL," +
		"    `msgFrom` TEXT NOT NULL," +
		"    `msgTo` TEXT NOT NULL," +
		"    `msgSubject` TEXT NOT NULL," +
		"    `msgContent` TEXT NOT NULL" +
		")"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}

func init() {
	migrations.Set("2020-10-23 00:03:00", nil, migration_000300_Up)
}