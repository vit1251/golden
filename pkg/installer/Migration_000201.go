package installer

import (
	"database/sql"
)

type Migration_000201 struct {
	IMigration
}

func (m *Migration_000201) Up(conn *sql.DB) error {
	query1 := "CREATE UNIQUE INDEX `uniq_area_areaName` ON `area` (`areaName` ASC)"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}
