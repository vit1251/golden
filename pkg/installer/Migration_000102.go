package installer

import (
	"database/sql"
	"path"
)

type Migration_000102 struct {
	IMigration
}

func (self *Migration_000102) Up(conn *sql.DB) error {

	query1 := "INSERT INTO `settings` (`section`,`name`,`value`) VALUES (?,?,?)"

	if _, err := conn.Exec(query1, "main", "RealName", "Alice Cooper"); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "Country", "Russia"); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "City", "Moscow"); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "Origin", "Yo Adrian, I Did It! (c) Rocky II"); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "TearLine", "Golden/{GOLDEN_PLATFORM}-{GOLDEN_ARCH} {GOLDEN_VERSION} {GOLDEN_RELEASE_DATE} ({GOLDEN_RELEASE_HASH})"); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "Address", "2:5030/1592.15"); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "Link", "2:5030/1592.0"); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "NetAddr", "f24.n5023.z2.binkp.net:24554"); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "Inbound", path.Join(".", "Inbound")); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "Outbound", path.Join(".", "Outbound")); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "TempInbound", path.Join(".", "TempInbound")); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "TempOutbound", path.Join(".", "TempOutbound")); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "FileBox", path.Join(".", "Files")); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "Password", ""); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "Nodelist", path.Join(".", "Nodelist")); err != nil {
	}

	if _, err := conn.Exec(query1, "main", "Pointlist", path.Join(".", "Pointlist")); err != nil {
	}

	return nil
}
