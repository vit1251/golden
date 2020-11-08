package mapper

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
	"log"
)

type NetmailMapper struct {
	Mapper
}

func NewNetmailMapper(r *registry.Container) *NetmailMapper {
	newNetmailMapper := new(NetmailMapper)
	newNetmailMapper.SetRegistry(r)
	return newNetmailMapper
}

func (self *NetmailMapper) GetMessageHeaders() ([]*NetmailMsg, error) {

	storageManager := self.restoreStorageManager()

	var result []*NetmailMsg

	query1 := "SELECT `nmId`, `nmHash`, `nmSubject`, `nmViewCount`, `nmFrom`, `nmTo`, `nmOrigAddr`, `nmDestAddr`, `nmDate` FROM `netmail` ORDER BY `nmDate` ASC, `nmId` ASC"
	var params []interface{}

	storageManager.Query(query1, params, func(rows *sql.Rows) error {

		var ID string
		var msgHash *string
		var subject string
		var from string
		var to string
		var origAddr string
		var destAddr string
		var msgDate int64
		var viewCount int

		err2 := rows.Scan(&ID, &msgHash, &subject, &viewCount, &from, &to, &origAddr, &destAddr, &msgDate)
		if err2 != nil{
			return err2
		}

		msg := NewNetmailMsg()
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
		msg.SetOrigAddr(origAddr)
		msg.SetDestAddr(destAddr)

		result = append(result, msg)

		return nil
	})

	return result, nil
}

func (self *NetmailMapper) Write(msg *NetmailMsg) error {

	storageManager := self.restoreStorageManager()

	query1 := "INSERT INTO `netmail` " +
		"(nmMsgId, nmHash, nmFrom, nmTo, nmSubject, nmBody, nmDate, nmOrigAddr, nmDestAddr) " +
		"VALUES " +
		"(?, ?, ?, ?, ?, ?, ?, ?, ?)"
	var params []interface{}
	params = append(params, msg.MsgID)
	params = append(params, msg.Hash)
	params = append(params, msg.From)
	params = append(params, msg.To)
	params = append(params, msg.Subject)
	params = append(params, msg.Content)
	params = append(params, msg.UnixTime)
	params = append(params, msg.OrigAddr)
	params = append(params, msg.DestAddr)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		return nil
	})

	return err1
}

func (self *NetmailMapper) GetMessageByHash(msgHash string) (*NetmailMsg, error) {

	storageManager := self.restoreStorageManager()

	var result *NetmailMsg

	query1 := "SELECT `nmId`, `nmMsgId`, `nmSubject`, `nmViewCount`, `nmFrom`, `nmTo`, `nmDate`, `nmBody`, `nmOrigAddr`, `nmDestAddr` FROM `netmail` WHERE `nmHash` = ?"
	var params []interface{}
	params = append(params, msgHash)

	storageManager.Query(query1, params, func(rows *sql.Rows) error {

		var nmId string
		var nmMsgId string
		var nmSubject string
		var nmViewCount int
		var nmFrom string
		var nmTo string
		var nmDate int64
		var nmBody string
		var nmOrigAddr string
		var nmDestAddr string

		err2 := rows.Scan(&nmId, &nmMsgId, &nmSubject, &nmViewCount, &nmFrom, &nmTo, &nmDate, &nmBody, &nmOrigAddr, &nmDestAddr)
		if err2 != nil{
			return err2
		}

		newMsg := NewNetmailMsg()
		newMsg.SetID(nmId)
		newMsg.SetMsgID(nmMsgId)
		newMsg.SetSubject(nmSubject)
		newMsg.SetHash(msgHash)
		newMsg.SetFrom(nmFrom)
		newMsg.SetTo(nmTo)
		newMsg.SetUnixTime(nmDate)
		newMsg.SetViewCount(nmViewCount)
		newMsg.SetContent(nmBody)
		newMsg.SetOrigAddr(nmOrigAddr)
		newMsg.SetDestAddr(nmDestAddr)

		result = newMsg

		return nil
	})

	return result, nil
}

func (self *NetmailMapper) ViewMessageByHash(msgHash string) error {

	storageManager := self.restoreStorageManager()

	query1 := "UPDATE `netmail` SET `nmViewCount` = `nmViewCount` + 1 WHERE `nmHash` = $1"
	var params []interface{}
	params = append(params, msgHash)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		log.Printf("Insert complete with: err = %+v", err)
		return nil
	})

	return err1
}

func (self *NetmailMapper) GetMessageNewCount() (int, error) {

	storageManager := self.restoreStorageManager()

	var newMessageCount int

	query1 := "SELECT count(`nmId`) FROM `netmail` WHERE `nmViewCount` = 0"
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

func (self *NetmailMapper) RemoveMessageByHash(msgHash string) error {

	storageManager := self.restoreStorageManager()

	query1 := "DELETE FROM `netmail` WHERE `nmHash` = ?"
	var params []interface{}
	params = append(params, msgHash)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		return nil
	})

	return err1
}

func (self *NetmailMapper) restoreStorageManager() *storage.StorageManager {
	managerPtr := self.registry.Get("StorageManager")
	if manager, ok := managerPtr.(*storage.StorageManager); ok {
		return manager
	} else {
		panic("no storage manager")
	}
}
