package installer

import "database/sql"

func migration_20201023031400_Up(conn *sql.DB) error {
	query1 := "ALTER TABLE `message` ADD `msgPacket` BLOB"
	_, err := conn.Exec(query1)
	return err
}

func init() {
	migrations.Set("2020-10-23 03:14:00", nil, migration_20201023031400_Up)
}
