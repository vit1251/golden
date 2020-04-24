package installer

import (
	"database/sql"
)

type Migration_000701 struct {
	IMigration
}

func (m *Migration_000701) Up(conn *sql.DB) error {
	query1 := "CREATE UNIQUE INDEX `uniq_stat_statDate` ON `stat` (`statDate` ASC)"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}
