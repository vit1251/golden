package installer

import "database/sql"

func migration_20201115020900_Up(conn *sql.DB) error {

	query1 := "INSERT INTO `settings` (`section`,`name`,`value`) VALUES (?, ?, ?)"
	if _, err := conn.Exec(query1, "netmail", "Charset", "CP866"); err != nil {
		return err
	}

	return nil

}

func init() {
	migrations.Set("2020-11-15 02:09:00", nil, migration_20201115020900_Up)
}
