package sqlite

import (
	"log"
)

type MessageBaseReader struct {
	MessageBase *MessageBase
}

func NewMessageBaseReader(mBase *MessageBase) (*MessageBaseReader, error) {
	mBaseReader := new(MessageBaseReader)
	mBaseReader.MessageBase = mBase
	return mBaseReader, nil
}

func (self *MessageBaseReader) GetAreaList() ([]string, error) {

	var result []string

	/* Step 1. Create message base session (i.e. SQL service connection) */
	mBaseSession, err1 := self.MessageBase.Open()
	if err1 != nil {
		return nil, err1
	}
	defer mBaseSession.Close()

	/* Step 2. Start SQL transaction */
	ConnTransaction, err := mBaseSession.Conn.Begin()
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

func (self *MessageBaseReader) GetAreaList2() ([]*Area, error) {

	var result []*Area

	/* Step 1. Create message base session (i.e. SQL service connection) */
	mBaseSession, err1 := self.MessageBase.Open()
	if err1 != nil {
		return nil, err1
	}
	defer mBaseSession.Close()

	ConnTransaction, err := mBaseSession.Conn.Begin()
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

func (self *MessageBaseReader) GetMessageHeaders(echoTag string) ([]*Message, error) {

	var result []*Message

	/* Step 1. Create message base session (i.e. SQL service connection) */
	mBaseSession, err1 := self.MessageBase.Open()
	if err1 != nil {
		return nil, err1
	}
	defer mBaseSession.Close()

	/* Step 2. */
	ConnTransaction, err := mBaseSession.Conn.Begin()
	if err != nil {
		return nil, err
	}

	sqlStmt := "SELECT `msgId`, `msgHash`, `msgSubject`, `msgFrom`, `msgTo`, `msgDate` FROM `message` WHERE `msgArea` = $1 ORDER BY `msgDate` ASC, `msgId` ASC"
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
		err2 := rows.Scan(&ID, &msgHash, &subject, &from, &to, &msgDate)
		if err2 != nil{
			return nil, err2
		}
		log.Printf("subject = %q", subject)
		msg := NewMessage()
		if msgHash != nil {
			msg.SetMsgID(*msgHash)
		}
		msg.SetSubject(subject)
		msg.SetID(ID)
		msg.SetFrom(from)
		msg.SetTo(to)
		msg.SetUnixTime(msgDate)
		result = append(result, msg)

	}

	ConnTransaction.Commit()

	return result, nil
}

func (self *MessageBaseReader) GetMessage(echoTag string, msgId string) (*Message, error) {

	var result *Message

	/* Step 1. Create message base session (i.e. SQL service connection) */
	mBaseSession, err1 := self.MessageBase.Open()
	if err1 != nil {
		return nil, err1
	}
	defer mBaseSession.Close()

	/* Step 2. */
	ConnTransaction, err := mBaseSession.Conn.Begin()
	if err != nil {
		return nil, err
	}

	sqlStmt := "SELECT `msgId`, `msgSubject`, `msgFrom`, `msgTo`, `msgContent` FROM `message` WHERE `msgArea` = $1 AND `msgId` = $2"
	log.Printf("sql = %q echoTag = %q", sqlStmt, echoTag)
	rows, err1 := ConnTransaction.Query(sqlStmt, echoTag, msgId)
	if err1 != nil {
		return nil, err1
	}
	defer rows.Close()
	for rows.Next() {

		var ID string
		var subject string
		var from string
		var to string
		var content string
		err2 := rows.Scan(&ID, &subject, &from, &to, &content)
		if err2 != nil{
			return nil, err2
		}
		log.Printf("subject = %q", subject)
		msg := NewMessage()
		msg.Subject = subject
		msg.ID = ID
//		msg.Hash = // TODO - save message UUID in database on tossing
		msg.From = from
		msg.To = to
		msg.Content = content
		result = msg

	}

	ConnTransaction.Commit()

	return result, nil
}
