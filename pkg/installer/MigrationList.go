package installer

import "database/sql"

type MigrationList struct {
	Migrations []Migration
}

func NewMigrationList() *MigrationList {
	return new(MigrationList)
}

func (self *MigrationList) Set(migrationName string, down func(*sql.DB) error, up func(*sql.DB) error) {
	m := NewMigration()
	m.ID = migrationName
	m.Down = down
	m.Up = up
	self.Migrations = append(self.Migrations, *m)
}

func (self *MigrationList) GetList() []string {
	var result []string
	for _, m := range self.Migrations {
		result = append(result, m.ID)
	}
	return result
}

func (self *MigrationList) GetByKey(key string) *Migration {
	for _, m := range self.Migrations {
		if m.ID == key {
			return &m
		}
	}
	return nil
}
