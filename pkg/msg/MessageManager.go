package msg

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vit1251/golden/pkg/charset"
	"github.com/vit1251/golden/pkg/storage"
	"log"
)

type MessageManager struct {
	conn           *sql.DB
	CharsetManager *charset.CharsetManager
	StorageManager *storage.StorageManager
}

func NewMessageManager(sm *storage.StorageManager, cm *charset.CharsetManager) *MessageManager {
	mm := new(MessageManager)
	mm.conn = sm.GetConnection()
	mm.CharsetManager = cm
	mm.StorageManager = sm
	return mm
}

func (self *MessageManager) GetAreaList() ([]string, error) {

	var result []string

	/* Step 2. Start SQL transaction */
	ConnTransaction, err := self.conn.Begin()
	if err != nil {
		return nil, err
	}

	/* Step 3. Make SQL query */
	sqlStmt := "SELECT DISTINCT(`msgArea`) AS `name` FROM `message` ORDER BY `name` ASC"
	rows, err1 := ConnTransaction.Query(sqlStmt)
	if err1 != nil {
		return nil, err1
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		err2 := rows.Scan(&name)
		if err2 != nil{
			return nil, err2
		}
		result = append(result, name)
	}

	/* Step 4. Release SQL transaction */
	ConnTransaction.Commit()

	return result, nil
}

func (self *MessageManager) GetAreaList2() ([]*Area, error) {

	var result []*Area

	/* Step 2. Start SQL transaction */
	ConnTransaction, err := self.conn.Begin()
	if err != nil {
		return nil, err
	}

	sqlStmt := "SELECT `msgArea`, count(`msgId`) AS `msgCount` FROM `message` GROUP BY `msgArea` ORDER BY `msgArea` ASC"
	rows, err1 := ConnTransaction.Query(sqlStmt)
	if err1 != nil {
		return nil, err1
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var count int
		err2 := rows.Scan(&name, &count)
		if err2 != nil{
			return nil, err2
		}
		a := NewArea()
		a.SetName(name)
		a.Count = count
		result = append(result, a)
	}

	ConnTransaction.Commit()

	return result, nil
}

func (self *MessageManager) GetAreaList3() ([]*Area, error) {

	var result []*Area

	/* Step 2. Start SQL transaction */
	ConnTransaction, err := self.conn.Begin()
	if err != nil {
		return nil, err
	}

	sqlStmt := "SELECT `msgArea`, count(`msgId`) AS `msgCount` FROM `message` WHERE `msgViewCount` = 0 GROUP BY `msgArea` ORDER BY `msgArea` ASC"
	rows, err1 := ConnTransaction.Query(sqlStmt)
	if err1 != nil {
		return nil, err1
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var count int
		err2 := rows.Scan(&name, &count)
		if err2 != nil{
			return nil, err2
		}
		a := NewArea()
		a.SetName(name)
		a.MsgNewCount = count
		result = append(result, a)
	}

	ConnTransaction.Commit()

	return result, nil
}

func (self *MessageManager) GetMessageHeaders(echoTag string) ([]*Message, error) {

	var result []*Message

	/* Step 2. Start SQL transaction */
	ConnTransaction, err := self.conn.Begin()
	if err != nil {
		return nil, err
	}

	sqlStmt := "SELECT `msgId`, `msgArea`, `msgHash`, `msgSubject`, `msgViewCount`, `msgFrom`, `msgTo`, `msgDate` FROM `message` WHERE `msgArea` = $1 ORDER BY `msgDate` ASC, `msgId` ASC"
	log.Printf("sql = %q echoTag = %q", sqlStmt, echoTag)
	rows, err1 := ConnTransaction.Query(sqlStmt, echoTag)
	if err1 != nil {
		return nil, err1
	}
	defer rows.Close()
	for rows.Next() {

		var ID string
		var msgHash *string
		var subject string
		var area string
		var from string
		var to string
		var msgDate int64
		var viewCount int

		err2 := rows.Scan(&ID, &area, &msgHash, &subject, &viewCount, &from, &to, &msgDate)
		if err2 != nil{
			return nil, err2
		}

		msg := NewMessage()
		if msgHash != nil {
			msg.SetMsgHash(*msgHash)
		}
		msg.SetArea(area)
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

func (self *MessageManager) GetMessageByHash(echoTag string, msgHash string) (*Message, error) {

	var result *Message

	/* Step 2. Start SQL transaction */
	ConnTransaction, err := self.conn.Begin()
	if err != nil {
		return nil, err
	}

	sqlStmt := "SELECT `msgId`, `msgMsgId`, `msgHash`, `msgSubject`, `msgFrom`, `msgTo`, `msgContent`, `msgDate` FROM `message` WHERE `msgArea` = $1 AND `msgHash` = $2"
	log.Printf("sql = %+v params = ( %+v, %+v )", sqlStmt, echoTag, msgHash)
	rows, err1 := ConnTransaction.Query(sqlStmt, echoTag, msgHash)
	if err1 != nil {
		return nil, err1
	}
	defer rows.Close()
	for rows.Next() {

		var ID string
		var msgMsgId string
		var msgHash *string
		var subject string
		var from string
		var to string
		var content string
		var written int64

		err2 := rows.Scan(&ID, &msgMsgId, &msgHash, &subject, &from, &to, &content, &written)
		if err2 != nil{
			return nil, err2
		}
		log.Printf("subject = %q", subject)

		/**/
		msg := NewMessage()
		msg.SetMsgID(msgMsgId)
		msg.SetSubject(subject)
		msg.SetID(ID)
		msg.SetUnixTime(written)
		if msgHash != nil {
			msg.SetMsgHash(*msgHash)
		}
		msg.SetFrom(from)
		msg.SetTo(to)
		msg.SetContent(content)

		//
		result = msg
	}

	ConnTransaction.Commit()

	return result, nil
}

func (self *MessageManager) ViewMessageByHash(echoTag string, msgHash string) (error) {

	/* Step 2. Start SQL transaction */
	ConnTransaction, err := self.conn.Begin()
	if err != nil {
		return err
	}

	sqlStmt := "UPDATE `message` SET `msgViewCount` = `msgViewCount` + 1 WHERE `msgArea` = $1 AND `msgHash` = $2"
	log.Printf("sql = %+v params = ( %+v, %+v )", sqlStmt, echoTag, msgHash)
	result, err3 := ConnTransaction.Exec(sqlStmt, echoTag, msgHash)
	if err3 != nil {
		return err3
	}
	log.Printf("result = %+v", result)

	ConnTransaction.Commit()

	return nil

}

func (self *MessageManager) RemoveMessageByHash(echoTag string, msgHash string) (error) {

	/* Step 2. Start SQL transaction */
	ConnTransaction, err := self.conn.Begin()
	if err != nil {
		return err
	}

	sqlStmt := "DELETE FROM `message` WHERE `msgArea` = $1 AND `msgHash` = $2"
	log.Printf("sql = %+v params = ( %+v, %+v )", sqlStmt, echoTag, msgHash)
	result, err3 := ConnTransaction.Exec(sqlStmt, echoTag, msgHash)
	if err3 != nil {
		return err3
	}
	log.Printf("result = %+v", result)

	ConnTransaction.Commit()

	return nil
}

func (self *MessageManager) IsMessageExistsByHash(echoTag string, msgHash string) (bool, error) {

	var result bool = false

	/* Step 2. Start SQL transaction */
	ConnTransaction, err := self.conn.Begin()
	if err != nil {
		return result, err
	}

	sqlStmt := "SELECT `msgId` FROM `message` WHERE `msgArea` = $1 AND `msgHash` = $2"
	log.Printf("sql = %+v params = ( %+v, %+v )", sqlStmt, echoTag, msgHash)
	rows, err1 := ConnTransaction.Query(sqlStmt, echoTag, msgHash)
	if err1 != nil {
		return result, err1
	}
	defer rows.Close()
	for rows.Next() {

		var ID string
		err2 := rows.Scan(&ID)
		if err2 != nil{
			return result, err2
		}
		if ID != "" {
			result = true
		}

	}

	ConnTransaction.Commit()

	return result, nil
}

func (self *MessageManager) Write(msg *Message) (error) {

	/* Step 3. Make prepare SQL insert query */
	query1 := "INSERT INTO message " +
	           "(msgMsgId, msgHash, msgArea, msgFrom, msgTo, msgSubject, msgContent, msgDate) " +
	           "VALUES " +
	           "(?, ?, ?, ?, ?, ?, ?, ?)"

	var params []interface{}
	params = append(params, msg.MsgID)
	params = append(params, msg.Hash)
	params = append(params, msg.Area)
	params = append(params, msg.From)
	params = append(params, msg.To)
	params = append(params, msg.Subject)
	params = append(params, msg.Content)
	params = append(params, msg.UnixTime)

	self.StorageManager.Exec(query1, params, func(err error) {
		log.Printf("Insert complete with: err = %+v", err)
	})

	return nil

}
