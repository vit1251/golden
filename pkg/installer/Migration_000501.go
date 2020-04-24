package installer

import (
	"database/sql"
)

type Migration_000501 struct {
	IMigration
}

func (m *Migration_000501) Up(conn *sql.DB) error {
	query1 := "CREATE INDEX `idx_file_fileArea` ON `file` (`fileArea` ASC)"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}
