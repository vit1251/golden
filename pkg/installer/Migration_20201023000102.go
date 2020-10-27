package installer

import (
	"database/sql"
)

func migration_000102_Up(conn *sql.DB) error {

	query1 := "INSERT INTO `settings` (`section`,`name`,`value`) VALUES (?, ?, ?)"

	if _, err := conn.Exec(query1, "main", "RealName", "John Smith"); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "Country", "Russia"); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "City", "Moscow"); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "Origin", "Yo Adrian, I Did It! (c) Rocky II"); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "TearLine", "Golden/{GOLDEN_PLATFORM}-{GOLDEN_ARCH} {GOLDEN_VERSION} {GOLDEN_RELEASE_DATE} ({GOLDEN_RELEASE_HASH})"); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "Address", ""); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "Link", ""); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "NetAddr", ""); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "Inbound", ""); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "Outbound", ""); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "TempInbound", ""); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "Temp", ""); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "TempOutbound", ""); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "FileBox", ""); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "Password", ""); err != nil {
	}

	return nil
}

func init() {
	migrations.Set("2020-10-23 00:01:02", nil, migration_000102_Up)
}