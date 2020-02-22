package msg

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type MessageManager struct {
	conn *sql.DB
}

func NewMessageManager(conn *sql.DB) *MessageManager {
	mm := new(MessageManager)
	mm.conn = conn
	mm.checkSchema()
	return mm
}

func (self *MessageManager) checkSchema() {

	query1 :=  "CREATE TABLE IF NOT EXISTS message (" +
		    "    msgId INTEGER NOT NULL PRIMARY KEY," +
		    "    msgHash CHAR(16) NOT NULL," +
			"    msgDate INTEGER NOT NULL," +
			"    msgViewCount INTEGER DEFAULT 0," +
		    "    msgArea CHAR(64) NOT NULL," +
		    "    msgFrom TEXT NOT NULL," +
		    "    msgTo TEXT NOT NULL," +
		    "    msgSubject TEXT NOT NULL," +
		    "    msgContent TEXT NOT NULL" +
		    ")"
	log.Printf("query = %+v", query1)
	if _, err := self.conn.Exec(query1); err != nil {
		log.Printf("Error create \"message\" storage: err = %+v", err)
	}

	/* Create index on msgHash */
	query2 := "CREATE INDEX \"idx_message_msgHash\" ON \"message\" (\"msgHash\" ASC)"
	if _, err := self.conn.Exec(query2); err != nil {
		log.Printf("Error create \"message\" storage: err = %+v", err)
	}

	/* Create index on msgHash */
	query3 := "CREATE INDEX \"idx_message_msgArea\" ON \"message\" (\"msgArea\" ASC)"
	if _, err := self.conn.Exec(query3); err != nil {
		log.Printf("Error create \"message\" storage: err = %+v", err)
	}

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
		a.Name = name
		a.Count = count
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

	sqlStmt := "SELECT `msgId`, `msgHash`, `msgSubject`, `msgViewCount`, `msgFrom`, `msgTo`, `msgDate` FROM `message` WHERE `msgArea` = $1 ORDER BY `msgDate` ASC, `msgId` ASC"
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
		var from string
		var to string
		var msgDate int64
		var viewCount int

		err2 := rows.Scan(&ID, &msgHash, &subject, &viewCount, &from, &to, &msgDate)
		if err2 != nil{
			return nil, err2
		}

		msg := NewMessage()
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

func (self *MessageManager) GetMessageByHash(echoTag string, msgHash string) (*Message, error) {

	var result *Message

	/* Step 2. Start SQL transaction */
	ConnTransaction, err := self.conn.Begin()
	if err != nil {
		return nil, err
	}

	sqlStmt := "SELECT `msgId`, `msgHash`, `msgSubject`, `msgFrom`, `msgTo`, `msgContent` FROM `message` WHERE `msgArea` = $1 AND `msgHash` = $2"
	log.Printf("sql = %+v params = ( %+v, %+v )", sqlStmt, echoTag, msgHash)
	rows, err1 := ConnTransaction.Query(sqlStmt, echoTag, msgHash)
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
		var content string
		err2 := rows.Scan(&ID, &msgHash, &subject, &from, &to, &content)
		if err2 != nil{
			return nil, err2
		}
		log.Printf("subject = %q", subject)
		msg := NewMessage()
		msg.Subject = subject
		msg.ID = ID
		if msgHash != nil {
			msg.Hash = *msgHash
		}
		msg.From = from
		msg.To = to
		msg.Content = content
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

	/* Step 2. Start SQL transaction */
	ConnTransaction, err := self.conn.Begin()
	if err != nil {
		return err
	}

	/* Step 3. Make prepare SQL insert query */
	sqlStmt := "INSERT INTO message "+
	           "    (msgHash, msgArea, msgFrom, msgTo, msgSubject, msgContent, msgDate) " +
	           "  VALUES " + 
	           "    (?, ?, ?, ?, ?, ?, ?)"
	log.Printf("sql = %q", sqlStmt)
	stmt, err3 := ConnTransaction.Prepare(sqlStmt)
	if err3 != nil {
		return err3
	}
	defer stmt.Close()

	/* Step 4. Invoke prepare SQL insert query */
	_, err4 := stmt.Exec(msg.Hash, msg.Area, msg.From, msg.To, msg.Subject, msg.Content, msg.UnixTime)
	if err4 != nil {
		return err4
	}

	/* Step 5. Commit SQL transaction */
	ConnTransaction.Commit()

	return nil

}
