package installer

import (
	"database/sql"
)

func init() {
	migrations.Set(
		"2020-10-23 00:02:00",
		nil,
		func(conn *sql.DB) error {
			query1 := "CREATE TABLE `area` (" +
				"    `areaId` INTEGER NOT NULL PRIMARY KEY," +
				"    `areaName` CHAR(64) NOT NULL," +
				"    `areaType` CHAR(64) NOT NULL," +
				"    `areaPath` CHAR(64) NOT NULL," +
				"    `areaSummary` CHAR(64) NOT NULL," +
				"    `areaOrder` INTEGER NOT NULL" +
				")"
			if _, err := conn.Exec(query1); err != nil {
				return err
			}
			return nil
		})
}
