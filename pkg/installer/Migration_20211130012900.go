package installer

import "database/sql"

func migration_20211130012900_Up(conn *sql.DB) error {

	query1 := "INSERT INTO `settings` (`section`,`name`,`value`) VALUES (?, ?, ?)"
	if _, err := conn.Exec(query1, "main", "StationName", "N/A"); err != nil {
		return err
	}

	return nil

}

func init() {
	migrations.Set("2021-11-30 01:29:00", nil, migration_20211130012900_Up)
}
