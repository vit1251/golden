package installer

import (
	"database/sql"
	"time"
)

func init() {
	migrationLocation, err := time.LoadLocation("Europe/Moscow")
	if err != nil {
		panic(err)
	}
	migrationDate := time.Date(2020, time.November, 4, 18, 3, 0, 0, migrationLocation)
	migrations.Register(migrationDate,
		nil,
		func(conn *sql.DB) error {

			query1 := "ALTER TABLE `message` ADD `msgReply` VARCHAR(16) DEFAULT ''"
			if _, err := conn.Exec(query1); err != nil {
				return err
			}

			query2 := "CREATE INDEX `idx_message_msgReply` ON `message` (`msgReply` ASC)"
			if _, err := conn.Exec(query2); err != nil {
				return err
			}

			return nil
		},
	)
}
