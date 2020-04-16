package netmail

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/storage"
	"log"
)

type NetmailManager struct {
	conn *sql.DB
}

func NewNetmailManager(sm *storage.StorageManager) *NetmailManager {
	nm := new(NetmailManager)
	nm.conn = sm.GetConnection()
	return nm
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
			msg.SetHash(*msgHash)
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

func (self *NetmailManager) Write(msg *NetmailMessage) error {

		/* Step 2. Start SQL transaction */
		ConnTransaction, err := self.conn.Begin()
		if err != nil {
			return err
		}

		/* Step 3. Make prepare SQL insert query */
		sqlStmt := "INSERT INTO `netmail` " +
			"    (nmMsgId, nmHash, nmFrom, nmTo, nmSubject, nmBody, nmDate) " +
			"  VALUES " +
			"    (?, ?, ?, ?, ?, ?, ?)"
		log.Printf("sql = %q", sqlStmt)
		stmt, err3 := ConnTransaction.Prepare(sqlStmt)
		if err3 != nil {
			return err3
		}
		defer stmt.Close()

		/* Step 4. Invoke prepare SQL insert query */
		_, err4 := stmt.Exec(msg.MsgID, msg.Hash, msg.From, msg.To, msg.Subject, msg.Content, msg.UnixTime)
		if err4 != nil {
			return err4
		}

		/* Step 5. Commit SQL transaction */
		ConnTransaction.Commit()

		return nil

}

func (self *NetmailManager) GetMessageByHash(hash string) (*NetmailMessage, error) {

	return nil, nil
}

func (self *NetmailManager) ViewMessageByHash(hash string) error {
	return nil
}
