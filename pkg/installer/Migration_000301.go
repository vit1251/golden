package installer

import (
	"database/sql"
)

type Migration_000301 struct {
	IMigration
}

func (m *Migration_000301) Up(conn *sql.DB) error {
	query2 := "CREATE INDEX `idx_message_msgHash` ON `message` (`msgHash` ASC)"
	if _, err := conn.Exec(query2); err != nil {
		return err
	}
	return nil
}
