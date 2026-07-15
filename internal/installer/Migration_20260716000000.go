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
    migrationDate := time.Date(2026, time.July, 16, 0, 0, 0, 0, migrationLocation)
    migrations.Register(
        migrationDate,
        func(conn *sql.DB) error {
            return nil
        },
        func(conn *sql.DB) error {
            _, err := conn.Exec("DROP TABLE IF EXISTS `twit`")
            return err
        },
    )
}