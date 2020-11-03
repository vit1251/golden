package installer

import "database/sql"

func migration_20201103033800_Up(conn *sql.DB) error {

	/* Add "nmOrigAddr" */
	query1 := "ALTER TABLE `netmail` ADD `nmOrigAddr` varchar(64) DEFAULT ''"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}

	/* Add "nmOrigAddr" */
	query2 := "ALTER TABLE `netmail` ADD `nmDestAddr` varchar(64) DEFAULT ''"
	if _, err := conn.Exec(query2); err != nil {
		return err
	}

	return nil
}

func init() {
	migrations.Set("2020-11-03 03:38:00", nil, migration_20201103033800_Up)
}