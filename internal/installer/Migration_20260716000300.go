package installer

import (
    "database/sql"
    "time"
)

func init() {
    migrationLocation, _ := time.LoadLocation("Europe/Moscow")
    migrationDate := time.Date(2026, time.July, 16, 0, 3, 0, 0, migrationLocation)
    migrations.Register(
        migrationDate,
        func(conn *sql.DB) error { return nil },
        func(conn *sql.DB) error {
            queries := []string{
                "ALTER TABLE `file` ADD COLUMN `fileOrigin` VARCHAR(64) NOT NULL DEFAULT ''",
                "ALTER TABLE `file` ADD COLUMN `fileFrom` VARCHAR(64) NOT NULL DEFAULT ''",
                "ALTER TABLE `file` ADD COLUMN `fileTo` VARCHAR(64) NOT NULL DEFAULT ''",
                "ALTER TABLE `file` ADD COLUMN `fileSize` INTEGER NOT NULL DEFAULT 0",
                "ALTER TABLE `file` ADD COLUMN `fileCrc` VARCHAR(16) NOT NULL DEFAULT ''",
                "ALTER TABLE `file` ADD COLUMN `fileLDesc` TEXT NOT NULL DEFAULT ''",
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
