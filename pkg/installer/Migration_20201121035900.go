package installer

import "database/sql"

func migration_20201121035900_Up(conn *sql.DB) error {

	query1 := "ALTER TABLE `filearea` RENAME COLUMN `areaType` TO `areaMode`"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}

	query2 := "UPDATE `filearea` SET `areaMode` = \"\""
	if _, err := conn.Exec(query2); err != nil {
		return err
	}

	return nil

}

func init() {
	migrations.Set("2020-11-21 03:59:00", nil, migration_20201121035900_Up)
}

