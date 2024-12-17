package installer

import "database/sql"

func migration_000601_Up(conn *sql.DB) error {
	query1 := "ALTER TABLE `netmail` ADD `nmMsgId` CHAR(64)"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}

func init() {
	migrations.Set("2020-10-23 00:06:01", nil, migration_000601_Up)
}
