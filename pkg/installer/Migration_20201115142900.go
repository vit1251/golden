package installer

import "database/sql"

func migration_20201115142900_Up(conn *sql.DB) error {
	query1 := "ALTER TABLE `netmail` ADD `nmPacket` BLOB"
	_, err := conn.Exec(query1)
	return err
}

func init() {
	migrations.Set("2020-11-15 14:29:00", nil, migration_20201115142900_Up)
}
