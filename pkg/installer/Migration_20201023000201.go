package installer

import (
	"database/sql"
)

func migration_000201_Up(conn *sql.DB) error {
	query1 := "CREATE UNIQUE INDEX `uniq_area_areaName` ON `area` (`areaName` ASC)"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}

func init() {
	migrations.Set("2020-10-23 00:02:01", nil, migration_000201_Up)
}