package installer

import (
	"database/sql"
)

type Migration struct {
	ID    string
	Up    func(conn *sql.DB) error
	Down  func(conn *sql.DB) error
}

func NewMigration() *Migration {
	m := new(Migration)
	return m
}
