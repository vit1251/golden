package installer

import "database/sql"

func migration_20201104180300_Up(conn *sql.DB) error {

	query1 := "ALTER TABLE `message` ADD `msgReply` VARCHAR(16) DEFAULT ''"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}

	query2 := "CREATE INDEX `idx_message_msgReply` ON `message` (`msgReply` ASC)"
	if _, err := conn.Exec(query2); err != nil {
		return err
	}

	return nil
}

func init() {
	migrations.Set("2020-11-04 18:03:00", nil, migration_20201104180300_Up)
}
