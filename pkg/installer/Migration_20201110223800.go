package installer

import (
	"database/sql"
)

func migration_20201110223800_Up(conn *sql.DB) error {

	query1 := "ALTER TABLE `filearea` RENAME COLUMN `areaPath` TO `areaCharset`"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}

	query2 := "UPDATE `filearea` SET `areaCharset` = \"CP866\""
	if _, err := conn.Exec(query2); err != nil {
		return err
	}

	return nil
}

func init() {
	migrations.Set("2020-11-10 22:38:00", nil, migration_20201110223800_Up)
}
