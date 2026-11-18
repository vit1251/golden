package installer

import (
    "database/sql"
    "time"
)

func init() {
    migrationLocation, _ := time.LoadLocation("Europe/Moscow")
    migrationDate := time.Date(2026, time.November, 18, 3, 48, 0, 0, migrationLocation)
    migrations.Register(
        migrationDate,
        func(conn *sql.DB) error { return nil },
        func(conn *sql.DB) error {
            queries := []string{
                "ALTER TABLE `stat_mailer` ADD COLUMN `statMailerByteRXcount` INTEGER NOT NULL DEFAULT 0",
                "ALTER TABLE `stat_mailer` ADD COLUMN `statMailerByteTXcount` INTEGER NOT NULL DEFAULT 0",
            }
            for _, q := range queries {
                if _, err := conn.Exec(q); err != nil {
                    return err
                }
            }
            return nil
        },
    )
}
