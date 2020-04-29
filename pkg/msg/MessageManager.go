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
	mm.CharsetManager = cm
	mm.StorageManager = sm
	return mm
}

func (self *MessageManager) GetAreaList() ([]string, error) {

	var result []string

	query1 := "SELECT DISTINCT(`msgArea`) AS `name` FROM `message` ORDER BY `name` ASC"
	var params []interface{}

	self.StorageManager.Query(query1, params, func(rows *sql.Rows) error {
		var name string
		err2 := rows.Scan(&name)
		if err2 != nil {
			return err2
		}
		result = append(result, name)
		return nil
	})

	return result, nil
}

func (self *MessageManager) GetAreaList2() ([]*Area, error) {

	var result []*Area

	query1 := "SELECT `msgArea`, count(`msgId`) AS `msgCount` FROM `message` GROUP BY `msgArea` ORDER BY `msgArea` ASC"
	var params []interface{}

	self.StorageManager.Query(query1, params, func(rows *sql.Rows) error {
		var name string
		var count int
		err2 := rows.Scan(&name, &count)
		if err2 != nil {
			return err2
		}
		a := NewArea()
		a.SetName(name)
		a.Count = count
		result = append(result, a)
		return nil
	})

	return result, nil
}

func (self *MessageManager) GetAreaList3() ([]*Area, error) {

	var result []*Area

	query1 := "SELECT `msgArea`, count(`msgId`) AS `msgCount` FROM `message` WHERE `msgViewCount` = 0 GROUP BY `msgArea` ORDER BY `msgArea` ASC"
	var params []interface{}

	self.StorageManager.Query(query1, params, func(rows *sql.Rows) error {
		var name string
		var count int

		err2 := rows.Scan(&name, &count)
		if err2 != nil{
			return err2
		}
		a := NewArea()
		a.SetName(name)
		a.MsgNewCount = count
		result = append(result, a)
		return nil
	})

	return result, nil
}

func (self *MessageManager) GetMessageHeaders(echoTag string) ([]*Message, error) {

	var result []*Message

	query1 := "SELECT `msgId`, `msgArea`, `msgHash`, `msgSubject`, `msgViewCount`, `msgFrom`, `msgTo`, `msgDate` FROM `message` WHERE `msgArea` = $1 ORDER BY `msgDate` ASC, `msgId` ASC"
	var params []interface{}
	params = append(params, echoTag)

	self.StorageManager.Query(query1, params, func(rows *sql.Rows) error {

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
			return err2
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

		return nil
	})

	return result, nil
}

func (self *MessageManager) GetMessageByHash(echoTag string, msgHash string) (*Message, error) {

	var result *Message

	query1 := "SELECT `msgId`, `msgMsgId`, `msgHash`, `msgSubject`, `msgFrom`, `msgTo`, `msgContent`, `msgDate` FROM `message` WHERE `msgArea` = $1 AND `msgHash` = $2"
	var params []interface{}
	params = append(params, echoTag)
	params = append(params, msgHash)

	self.StorageManager.Query(query1, params, func(rows *sql.Rows) error {

		var ID string
		var msgMsgId string
		var msgHash *string
		var subject string
		var from string
		var to string
		var content string
		var written int64

		err1 := rows.Scan(&ID, &msgMsgId, &msgHash, &subject, &from, &to, &content, &written)
		if err1 != nil{
			return err1
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

		/* Save result */
		result = msg

		return nil
	})

	return result, nil
}

func (self *MessageManager) ViewMessageByHash(echoTag string, msgHash string) error {

	query1 := "UPDATE `message` SET `msgViewCount` = `msgViewCount` + 1 WHERE `msgArea` = $1 AND `msgHash` = $2"
	var params []interface{}
	params = append(params, echoTag)
	params = append(params, msgHash)

	self.StorageManager.Exec(query1, params, func(err error) {
		log.Printf("Insert complete with: err = %+v", err)
	})

	return nil

}

func (self *MessageManager) RemoveMessageByHash(echoTag string, msgHash string) error {

	query1 := "DELETE FROM `message` WHERE `msgArea` = $1 AND `msgHash` = $2"
	var params []interface{}
	params = append(params, echoTag)
	params = append(params, msgHash)

	self.StorageManager.Exec(query1, params, func(err error) {
		log.Printf("Insert complete with: err = %+v", err)
	})

	return nil
}

func (self *MessageManager) IsMessageExistsByHash(echoTag string, msgHash string) (bool, error) {

	var result bool = false

	query1 := "SELECT `msgId` FROM `message` WHERE `msgArea` = $1 AND `msgHash` = $2"
	var params []interface{}
	params = append(params, echoTag)
	params = append(params, msgHash)

	self.StorageManager.Query(query1, params, func(rows *sql.Rows) error {
		var ID string
		err1 := rows.Scan(&ID)
		if err1 != nil {
			return err1
		}
		if ID != "" {
			result = true
		}
		return nil
	})

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

func (self *MessageManager) GetMessageNewCount() (int, error) {

	var newMessageCount int

	query1 := "SELECT count(`msgId`) FROM `message` WHERE `msgViewCount` = 0"
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

func (self *MessageManager) RemoveMessagesByAreaName(echoTag string) error {

	query1 := "DELETE FROM `message` WHERE `msgArea` = $1"
	var params []interface{}
	params = append(params, echoTag)

	self.StorageManager.Exec(query1, params, func(err error) {
		log.Printf("Insert complete with: err = %+v", err)
	})

	return nil

}
