package sqlite

import (
	"log"
)

type MessageBaseWriter struct {
	MessageBase *MessageBase
}


func NewMessageBaseWriter(mBase *MessageBase) (*MessageBaseWriter, error) {
	writer := new(MessageBaseWriter)
	writer.MessageBase = mBase
	return writer, nil
}

func (self *MessageBaseWriter) IsMessageExistsByHash(echoTag string, msgHash string) (bool, error) {

	var result bool = false

	/* Step 1. Create message base session (i.e. SQL service connection) */
	mBaseSession, err1 := self.MessageBase.Open()
	if err1 != nil {
		return result, err1
	}
	defer mBaseSession.Close()

	/* Step 2. */
	ConnTransaction, err := mBaseSession.Conn.Begin()
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

func (self *MessageBaseWriter) Write(msg *Message) (error) {

	/* Step 1. Create message base session (i.e. SQL service connection) */
	mBaseSession, err1 := self.MessageBase.Open()
	if err1 != nil {
		return err1
	}
	defer mBaseSession.Close()

	/* Step 2. Start SQL transaction */
	ConnTransaction, err2 := mBaseSession.Conn.Begin()
	if err2 != nil {
		return err2
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
