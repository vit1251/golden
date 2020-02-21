package netmail

import (
	"database/sql"
)

type NetmailManager struct {
	conn *sql.DB
}

func NewNetmailManager(conn *sql.DB) *NetmailManager {
	nm := new(NetmailManager)
	nm.conn = conn
	nm.checkSchema()
	return nm
}

func (self *NetmailManager) checkSchema() error {

	sqlStmt := "CREATE TABLE IF NOT EXISTS netmail (" +
		"    nmId INTEGER NOT NULL PRIMARY KEY," +
		"    nmFrom CHAR(64) NOT NULL," +
		"    nmTo CHAR(64) NOT NULL," +
		"    nmBody TEXT NOT NULL," +
		"    nmTime INTEGER NOT NULL" +
		")"
	self.conn.Exec(sqlStmt)

	return nil

}
