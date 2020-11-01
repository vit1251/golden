package echomail

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
	"log"
)

type MessageManager struct {
	registry       *registry.Container
}

func NewMessageManager(r *registry.Container) *MessageManager {
	mm := new(MessageManager)
	mm.registry = r
	return mm
}

func (self *MessageManager) GetAreaList() ([]string, error) {

	storageManager := self.restoreStorageManager()

	var result []string

	query1 := "SELECT DISTINCT(`msgArea`) AS `name` FROM `message` ORDER BY `name` ASC"
	var params []interface{}

	storageManager.Query(query1, params, func(rows *sql.Rows) error {
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

	storageManager := self.restoreStorageManager()

	var result []*Area

	query1 := "SELECT `msgArea`, count(`msgId`) AS `msgCount` FROM `message` GROUP BY `msgArea` ORDER BY `msgArea` ASC"
	var params []interface{}

	storageManager.Query(query1, params, func(rows *sql.Rows) error {
		var name string
		var count int
		err2 := rows.Scan(&name, &count)
		if err2 != nil {
			return err2
		}
		a := NewArea()
		a.SetName(name)
		a.MessageCount = count
		result = append(result, a)
		return nil
	})

	return result, nil
}

func (self *MessageManager) GetAreaList3() ([]*Area, error) {

	storageManager := self.restoreStorageManager()

	var result []*Area

	query1 := "SELECT `msgArea`, count(`msgId`) AS `msgCount` FROM `message` WHERE `msgViewCount` = 0 GROUP BY `msgArea` ORDER BY `msgArea` ASC"
	var params []interface{}

	storageManager.Query(query1, params, func(rows *sql.Rows) error {
		var name string
		var count int

		err2 := rows.Scan(&name, &count)
		if err2 != nil{
			return err2
		}
		a := NewArea()
		a.SetName(name)
		a.NewMessageCount = count
		result = append(result, a)
		return nil
	})

	return result, nil
}

func (self *MessageManager) GetMessageHeaders(echoTag string) ([]*msg.Message, error) {

	storageManager := self.restoreStorageManager()

	var result []*msg.Message

	query1 := "SELECT `msgId`, `msgArea`, `msgHash`, `msgSubject`, `msgViewCount`, `msgFrom`, `msgTo`, `msgDate` FROM `message` WHERE `msgArea` = $1 ORDER BY `msgDate` ASC, `msgId` ASC"
	var params []interface{}
	params = append(params, echoTag)

	storageManager.Query(query1, params, func(rows *sql.Rows) error {

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

		msg := msg.NewMessage()
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

func (self *MessageManager) GetMessageByHash(echoTag string, msgHash string) (*msg.Message, error) {

	storageManager := self.restoreStorageManager()

	var result *msg.Message

	query1 := "SELECT `msgId`, `msgArea`, `msgMsgId`, `msgHash`, `msgSubject`, `msgFrom`, `msgTo`, `msgContent`, `msgDate`, `msgPacket` FROM `message` WHERE `msgArea` = $1 AND `msgHash` = $2"
	var params []interface{}
	params = append(params, echoTag)
	params = append(params, msgHash)

	storageManager.Query(query1, params, func(rows *sql.Rows) error {

		var ID string
		var msgMsgId string
		var msgArea string
		var msgHash *string
		var subject string
		var from string
		var to string
		var content string
		var packet []byte
		var written int64

		err1 := rows.Scan(&ID, &msgArea, &msgMsgId, &msgHash, &subject, &from, &to, &content, &written, &packet)
		if err1 != nil{
			return err1
		}
		log.Printf("subject = %q", subject)

		/**/
		msg := msg.NewMessage()
		msg.SetArea(msgArea)
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
		msg.SetPacket(packet)

		/* Save result */
		result = msg

		return nil
	})

	return result, nil
}

func (self *MessageManager) ViewMessageByHash(echoTag string, msgHash string) error {

	storageManager := self.restoreStorageManager()

	query1 := "UPDATE `message` SET `msgViewCount` = `msgViewCount` + 1 WHERE `msgArea` = $1 AND `msgHash` = $2"
	var params []interface{}
	params = append(params, echoTag)
	params = append(params, msgHash)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		log.Printf("Insert complete with: err = %+v", err)
		return nil
	})

	return err1

}

func (self *MessageManager) RemoveMessageByHash(echoTag string, msgHash string) error {

	storageManager := self.restoreStorageManager()

	query1 := "DELETE FROM `message` WHERE `msgArea` = $1 AND `msgHash` = $2"
	var params []interface{}
	params = append(params, echoTag)
	params = append(params, msgHash)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		log.Printf("Insert complete with: err = %+v", err)
		return nil
	})

	return err1
}

func (self *MessageManager) IsMessageExistsByHash(echoTag string, msgHash string) (bool, error) {

	storageManager := self.restoreStorageManager()

	var result bool = false

	query1 := "SELECT `msgId` FROM `message` WHERE `msgArea` = $1 AND `msgHash` = $2"
	var params []interface{}
	params = append(params, echoTag)
	params = append(params, msgHash)

	storageManager.Query(query1, params, func(rows *sql.Rows) error {
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

func (self *MessageManager) Write(msg *msg.Message) (error) {

	storageManager := self.restoreStorageManager()

	/* Step 3. Make prepare SQL insert query */
	query1 := "INSERT INTO message " +
	           "(msgMsgId, msgHash, msgArea, msgFrom, msgTo, msgSubject, msgContent, msgDate, msgPacket) " +
	           "VALUES " +
	           "(?, ?, ?, ?, ?, ?, ?, ?, ?)"

	var params []interface{}
	params = append(params, msg.MsgID)
	params = append(params, msg.Hash)
	params = append(params, msg.Area)
	params = append(params, msg.From)
	params = append(params, msg.To)
	params = append(params, msg.Subject)
	params = append(params, msg.Content)
	params = append(params, msg.UnixTime)
	params = append(params, msg.Packet)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		log.Printf("Insert complete with: err = %+v", err)
		return nil
	})

	return err1

}

func (self *MessageManager) GetMessageNewCount() (int, error) {

	storageManager := self.restoreStorageManager()

	var newMessageCount int

	query1 := "SELECT count(`msgId`) FROM `message` WHERE `msgViewCount` = 0"
	var params []interface{}

	storageManager.Query(query1, params, func(rows *sql.Rows) error {

		err1 := rows.Scan(&newMessageCount)
		if err1 != nil {
			return err1
		}
		return nil
	})

	return newMessageCount, nil
}

func (self *MessageManager) RemoveMessagesByAreaName(echoTag string) error {

	storageManager := self.restoreStorageManager()

	query1 := "DELETE FROM `message` WHERE `msgArea` = $1"
	var params []interface{}
	params = append(params, echoTag)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		log.Printf("Insert complete with: err = %+v", err)
		return nil
	})

	return err1

}

func (self *MessageManager) restoreStorageManager() *storage.StorageManager {

	managerPtr := self.registry.Get("StorageManager")
	if manager, ok := managerPtr.(*storage.StorageManager); ok {
		return manager
	} else {
		panic("no storage manager")
	}

}
