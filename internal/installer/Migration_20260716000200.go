package installer

import (
    "database/sql"
    "time"
)

func init() {
    migrationLocation, _ := time.LoadLocation("Europe/Moscow")
    migrationDate := time.Date(2026, time.July, 16, 0, 2, 0, 0, migrationLocation)
    migrations.Register(
        migrationDate,
        func(conn *sql.DB) error { return nil },
        func(conn *sql.DB) error {
            _, err := conn.Exec("ALTER TABLE `file` ADD COLUMN `fileArchived` INTEGER NOT NULL DEFAULT 0")
            return err
        },
    )
}
