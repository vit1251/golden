package installer

import "database/sql"

func migration_20201104180700_Up(conn *sql.DB) error {

	query1 := "ALTER TABLE `message` ADD `msgOrigAddr` varchar(32) DEFAULT ''"
	if _, err := conn.Exec(query1); err != nil {
		return err
	}

	return nil

}

func init() {
	migrations.Set("2020-11-04 18:07:00", nil, migration_20201104180700_Up)
}