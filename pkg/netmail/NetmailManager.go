package netmail

import (
	"database/sql"
	"log"
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
		"    nmHash CHAR(64) NOT NULL," +
		"    nmFrom CHAR(64) NOT NULL," +
		"    nmTo CHAR(64) NOT NULL," +
		"    nmSubject CHAR(512) NOT NULL," +
		"    nmBody TEXT NOT NULL," +
		"    nmDate INTEGER NOT NULL," +
		"    nmViewCount INTEGER DEFAULT 0" +
		")"
	self.conn.Exec(sqlStmt)

	return nil

}

func (self *NetmailManager) GetMessageHeaders() ([]*NetmailMessage, error) {

	var result []*NetmailMessage

	/* Step 2. Start SQL transaction */
	ConnTransaction, err := self.conn.Begin()
	if err != nil {
		return nil, err
	}

	sqlStmt := "SELECT `nmId`, `nmHash`, `nmSubject`, `nmViewCount`, `nmFrom`, `nmTo`, `nmDate` FROM `netmail` ORDER BY `nmDate` ASC, `nmId` ASC"
	log.Printf("sql = %q", sqlStmt)
	rows, err1 := ConnTransaction.Query(sqlStmt)
	if err1 != nil {
		return nil, err1
	}
	defer rows.Close()
	for rows.Next() {

		var ID string
		var msgHash *string
		var subject string
		var from string
		var to string
		var msgDate int64
		var viewCount int

		err2 := rows.Scan(&ID, &msgHash, &subject, &viewCount, &from, &to, &msgDate)
		if err2 != nil{
			return nil, err2
		}

		msg := NewNetmailMessage()
		if msgHash != nil {
			msg.SetMsgID(*msgHash)
		}
		msg.SetSubject(subject)
		msg.SetID(ID)
		msg.SetFrom(from)
		msg.SetTo(to)
		msg.SetUnixTime(msgDate)
		msg.SetViewCount(viewCount)

		result = append(result, msg)

	}

	ConnTransaction.Commit()

	return result, nil
}
