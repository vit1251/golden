package installer

import "database/sql"

func migration_20201102140000_Up(conn *sql.DB) error {
	query1 := "ALTER TABLE `area` ADD `areaCharset` CHAR(16) NOT NULL DEFAULT \"CP866\""
	_, err := conn.Exec(query1)
	return err
}

func init() {
	migrations.Set("2020-11-02 14:00:00", nil, migration_20201102140000_Up)
}

