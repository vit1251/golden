package installer

import (
	"database/sql"
)

func migration_000701_Up(conn *sql.DB) error {
	query1 := "CREATE UNIQUE INDEX `uniq_stat_statDate` ON `stat` (`statDate` ASC)"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}

func init() {
	migrations.Set("2020-10-23 00:07:01", nil, migration_000701_Up)
}
