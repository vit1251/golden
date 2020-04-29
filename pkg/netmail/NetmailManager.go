package netmail

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/storage"
	"log"
)

type NetmailManager struct {
	StorageManager  *storage.StorageManager
}

func NewNetmailManager(sm *storage.StorageManager) *NetmailManager {
	nm := new(NetmailManager)
	nm.StorageManager = sm
	return nm
}

func (self *NetmailManager) GetMessageHeaders() ([]*NetmailMessage, error) {

	var result []*NetmailMessage

	query1 := "SELECT `nmId`, `nmHash`, `nmSubject`, `nmViewCount`, `nmFrom`, `nmTo`, `nmDate` FROM `netmail` ORDER BY `nmDate` ASC, `nmId` ASC"
	var params []interface{}

	self.StorageManager.Query(query1, params, func(rows *sql.Rows) error {

		var ID string
		var msgHash *string
		var subject string
		var from string
		var to string
		var msgDate int64
		var viewCount int

		err2 := rows.Scan(&ID, &msgHash, &subject, &viewCount, &from, &to, &msgDate)
		if err2 != nil{
			return err2
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

		return nil
	})

	return result, nil
}

func (self *NetmailManager) Write(msg *NetmailMessage) error {

	query1 := "INSERT INTO `netmail` " +
		"(nmMsgId, nmHash, nmFrom, nmTo, nmSubject, nmBody, nmDate) " +
		"VALUES " +
		"(?, ?, ?, ?, ?, ?, ?)"
	var params []interface{}
	params = append(params, msg.MsgID)
	params = append(params, msg.Hash)
	params = append(params, msg.From)
	params = append(params, msg.To)
	params = append(params, msg.Subject)
	params = append(params, msg.Content)
	params = append(params, msg.UnixTime)

	self.StorageManager.Exec(query1, params, func(err error) {
		//return err
	})

	return nil
}

func (self *NetmailManager) GetMessageByHash(msgHash string) (*NetmailMessage, error) {

	var result *NetmailMessage

	query1 := "SELECT `nmId`, `nmHash`, `nmSubject`, `nmViewCount`, `nmFrom`, `nmTo`, `nmDate`, `nmBody` FROM `netmail` WHERE `nmHash` = ?"
	var params []interface{}
	params = append(params, msgHash)

	self.StorageManager.Query(query1, params, func(rows *sql.Rows) error {

		var ID string
		var msgHash *string
		var subject string
		var from string
		var to string
		var body string
		var msgDate int64
		var viewCount int

		err2 := rows.Scan(&ID, &msgHash, &subject, &viewCount, &from, &to, &msgDate, &body)
		if err2 != nil{
			return err2
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
		msg.SetContent(body)

		result = msg

		return nil
	})

	return result, nil
}

func (self *NetmailManager) ViewMessageByHash(msgHash string) error {

	query1 := "UPDATE `netmail` SET `nmViewCount` = `nmViewCount` + 1 WHERE `nmHash` = $1"
	var params []interface{}
	params = append(params, msgHash)

	self.StorageManager.Exec(query1, params, func(err error) {
		log.Printf("Insert complete with: err = %+v", err)
	})

	return nil
}

func (self *NetmailManager) GetMessageNewCount() (int, error) {

	var newMessageCount int

	query1 := "SELECT count(`nmId`) FROM `netmail` WHERE `nmViewCount` = 0"
	var params []interface{}

	self.StorageManager.Query(query1, params, func(rows *sql.Rows) error {

		err1 := rows.Scan(&newMessageCount)
		if err1 != nil {
			return err1
		}
		return nil
	})

	return newMessageCount, nil
}
