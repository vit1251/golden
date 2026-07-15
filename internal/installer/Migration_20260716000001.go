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
    migrationDate := time.Date(2026, time.July, 16, 0, 1, 0, 0, migrationLocation)
    migrations.Register(
        migrationDate,
        func(conn *sql.DB) error {
            return nil
        },
        func(conn *sql.DB) error {
            if _, err := conn.Exec("ALTER TABLE `netmail` ADD COLUMN `nmArchived` INTEGER NOT NULL DEFAULT 0"); err != nil {
                return err
            }
            if _, err := conn.Exec("ALTER TABLE `message` ADD COLUMN `msgArchived` INTEGER NOT NULL DEFAULT 0"); err != nil {
                return err
            }
            return nil
        },
    )
}
