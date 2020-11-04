package installer

import "database/sql"


func migration_20201104232600_Up(conn *sql.DB) error {

	query1 := "INSERT INTO `settings` (`section`,`name`,`value`) VALUES (?, ?, ?)"
	if _, err := conn.Exec(query1, "mailer", "Interval", "5"); err != nil {
		return err
	}

	return nil

}

func init() {
	migrations.Set("2020-11-04 23:26:00", nil, migration_20201104232600_Up)
}
