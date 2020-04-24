package installer

import (
	"database/sql"
)

type Migration_000702 struct {
	IMigration
}

func (m *Migration_000702) Up(conn *sql.DB) error {
	query1 := "ALTER TABLE \"stat\" ADD \"statSessionIn\" INTEGER DEFAULT 0"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}
