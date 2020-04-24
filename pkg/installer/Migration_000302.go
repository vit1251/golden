package installer

import (
	"database/sql"
)

type Migration_000302 struct {
	IMigration
}

func (m *Migration_000302) Up(conn *sql.DB) error {
	query3 := "CREATE INDEX `idx_message_msgArea` ON `message` (`msgArea` ASC)"
	if _, err := conn.Exec(query3); err != nil {
		return err
	}
	return nil
}
