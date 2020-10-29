package installer

import (
	"database/sql"
)

func migration_000102_Up(conn *sql.DB) error {

	type param struct {
		section string
		name string
		value string
	}

	var params []param

	params = append(params, param{"main", "RealName", "John Smith"})
	params = append(params, param{"main", "Country", "Russia"})
	params = append(params, param{"main", "City", "Moscow"})
	params = append(params, param{"main", "Origin", "Yo Adrian, I Did It! (c) Rocky II"})
	params = append(params, param{"main", "TearLine", "Golden/{GOLDEN_PLATFORM}-{GOLDEN_ARCH} {GOLDEN_VERSION} {GOLDEN_RELEASE_DATE} ({GOLDEN_RELEASE_HASH})"})
	params = append(params, param{ "main", "Address", ""})
	params = append(params, param{"main", "Link", ""})
	params = append(params, param{"main", "NetAddr", ""})
	params = append(params, param{"main", "Inbound", ""})
	params = append(params, param{"main", "Outbound", ""})
	params = append(params, param{"main", "TempInbound", ""})
	params = append(params, param{"main", "Temp", ""})
	params = append(params, param{"main", "TempOutbound", ""})
	params = append(params, param{"main", "FileBox", ""})
	params = append(params, param{"main", "Password", ""})

	/* Execute */
	query1 := "INSERT INTO `settings` (`section`,`name`,`value`) VALUES (?, ?, ?)"
	for _, param := range params {
		if _, err := conn.Exec(query1, param.section, param.name, param.value); err != nil {
			return err
		}
	}

	return nil
}

func init() {
	migrations.Set("2020-10-23 00:01:02", nil, migration_000102_Up)
}