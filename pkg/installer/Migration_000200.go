package installer

import (
	"database/sql"
)

type Migration_000200 struct {
	IMigration
}

func (m *Migration_000200) Up(conn *sql.DB) error {
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
}
