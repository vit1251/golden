package installer

import "database/sql"

func migration_20211201003500_Up(conn *sql.DB) error {

	query1 := "INSERT INTO `settings` (`section`,`name`,`value`) VALUES (?, ?, ?)"
	if _, err := conn.Exec(query1, "echomail", "Charset", "CP866"); err != nil {
		return err
	}

	return nil

}

func init() {
	migrations.Set("2021-12-01 00:35:00", nil, migration_20211201003500_Up)
}
