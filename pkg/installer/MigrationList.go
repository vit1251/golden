package installer

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/site/utils"
	"time"
)

type MigrationList struct {
	Migrations []Migration
}

func NewMigrationList() *MigrationList {
	return new(MigrationList)
}

func (self *MigrationList) Register(migrationTime time.Time, down func(*sql.DB) error, up func(*sql.DB) error) {
	migrationName := utils.DateHelper_renderDateWithSecond(migrationTime)
	self.Set(migrationName, down, up)
}

// Deprecated: Set is deprecated.
func (self *MigrationList) Set(migrationTime string, down func(*sql.DB) error, up func(*sql.DB) error) {
	m := NewMigration()
	m.ID = migrationTime
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
