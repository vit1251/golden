package installer

import (
	"database/sql"
)

type Migration_000601 struct {
	IMigration
}

func (m *Migration_000601) Up(conn *sql.DB) error {
	query1 := "ALTER TABLE `netmail` ADD `nmMsgId` CHAR(64)"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}
	return nil
}
