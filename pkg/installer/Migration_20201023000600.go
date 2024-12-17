package installer

import "database/sql"

func migration_000600_Up(conn *sql.DB) error {
	query1 := "CREATE TABLE `netmail` (" +
		"    nmId INTEGER NOT NULL PRIMARY KEY," +
		"    nmHash CHAR(64) NOT NULL," +
		"    nmFrom CHAR(64) NOT NULL," +
		"    nmTo CHAR(64) NOT NULL," +
		"    nmSubject CHAR(512) NOT NULL," +
		"    nmBody TEXT NOT NULL," +
		"    nmDate INTEGER NOT NULL," +
		"    nmViewCount INTEGER DEFAULT 0" +
		")"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}

func init() {
	migrations.Set("2020-10-23 00:06:00", nil, migration_000600_Up)
}
