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
		return err
	}

	conn.Exec(query1, "main", "Country", "Russia")
	conn.Exec(query1, "main", "City", "Moscow")
	conn.Exec(query1, "main", "Origin", "Yo Adrian, I Did It! (c) Rocky II")
	conn.Exec(query1, "main", "TearLine", "Golden/{GOLDEN_PLATFORM}-{GOLDEN_ARCH} {GOLDEN_VERSION} {GOLDEN_RELEASE_DATE} ({GOLDEN_RELEASE_HASH})")
	conn.Exec(query1, "main", "Address", "2:5030/1592.15")
	conn.Exec(query1, "main", "Link", "2:5030/1592.0")
	conn.Exec(query1, "main", "NetAddr", "f24.n5023.z2.binkp.net:24554")
	conn.Exec(query1, "main", "Inbound", path.Join(".", "Inbound"))
	conn.Exec(query1, "main", "Outbound", path.Join(".", "Outbound"))
	conn.Exec(query1, "main", "TempInbound", path.Join(".", "TempInbound"))
	conn.Exec(query1, "main", "TempOutbound", path.Join(".", "TempOutbound"))
	conn.Exec(query1, "main", "FileBox", path.Join(".", "Files"))
	conn.Exec(query1, "main", "Password", "")

	return nil
}
