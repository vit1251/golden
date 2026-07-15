package mapper

import (
    "database/sql"
    "log"
    "github.com/huandu/go-sqlbuilder"
    "github.com/vit1251/golden/pkg/registry"
    "github.com/vit1251/golden/pkg/storage"
)

type NetmailMapper struct {
    Mapper
}

func NewNetmailMapper(r *registry.Container) *NetmailMapper {
    newNetmailMapper := new(NetmailMapper)
    newNetmailMapper.SetRegistry(r)
    return newNetmailMapper
}

func (m *NetmailMapper) GetMessageHeaders() ([]*NetmailMsg, error) {

    storageManager := storage.RestoreStorageManager(m.registry)

    var result []*NetmailMsg

    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("nmId", "nmMsgId", "nmHash", "nmSubject", "nmViewCount", "nmFrom", "nmTo", "nmOrigAddr", "nmDestAddr", "nmDate")
    sb.Where(sb.Equal("nmArchived", 0))
    sb.From("netmail")
    sb.OrderBy("nmDate ASC", "nmId ASC")
    query1, args := sb.Build()

    storageManager.Query(query1, args, func(rows *sql.Rows) error {
		var nmId string
		var nmMsgId string
		var nmHash string
		var nmSubject string
		var nmFrom string
		var nmTo string
		var nmOrigAddr string
		var nmDestAddr string
		var nmDate int64
		var nmViewCount int

		err2 := rows.Scan(&nmId, &nmMsgId, &nmHash, &nmSubject, &nmViewCount, &nmFrom, &nmTo, &nmOrigAddr, &nmDestAddr, &nmDate)
		if err2 != nil {
			return err2
		}

		msg := NewNetmailMsg()
		msg.SetHash(nmHash)
		msg.SetMsgID(nmMsgId)
		msg.SetSubject(nmSubject)
		msg.SetID(nmId)
		msg.SetFrom(nmFrom)
		msg.SetTo(nmTo)
		msg.SetUnixTime(nmDate)
		msg.SetViewCount(nmViewCount)
		msg.SetOrigAddr(nmOrigAddr)
		msg.SetDestAddr(nmDestAddr)

		result = append(result, msg)

		return nil
    })

    return result, nil
}

func (m *NetmailMapper) GetMessageHeadersPage(limit, offset int) ([]*NetmailMsg, error) {

    storageManager := storage.RestoreStorageManager(m.registry)

    var result []*NetmailMsg

    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("nmId", "nmMsgId", "nmHash", "nmSubject", "nmViewCount", "nmFrom", "nmTo", "nmOrigAddr", "nmDestAddr", "nmDate")
    sb.From("netmail")
    sb.Where(sb.Equal("nmArchived", 0))
    sb.OrderBy("nmDate ASC", "nmId ASC")
    sb.Limit(limit)
    sb.Offset(offset)
    query1, args := sb.Build()

    storageManager.Query(query1, args, func(rows *sql.Rows) error {
		var nmId string
		var nmMsgId string
		var nmHash string
		var nmSubject string
		var nmFrom string
		var nmTo string
		var nmOrigAddr string
		var nmDestAddr string
		var nmDate int64
		var nmViewCount int

		err2 := rows.Scan(&nmId, &nmMsgId, &nmHash, &nmSubject, &nmViewCount, &nmFrom, &nmTo, &nmOrigAddr, &nmDestAddr, &nmDate)
		if err2 != nil {
			return err2
		}

		msg := NewNetmailMsg()
		msg.SetHash(nmHash)
		msg.SetMsgID(nmMsgId)
		msg.SetSubject(nmSubject)
		msg.SetID(nmId)
		msg.SetFrom(nmFrom)
		msg.SetTo(nmTo)
		msg.SetUnixTime(nmDate)
		msg.SetViewCount(nmViewCount)
		msg.SetOrigAddr(nmOrigAddr)
		msg.SetDestAddr(nmDestAddr)

		result = append(result, msg)

		return nil
	})

    return result, nil
}

func (m *NetmailMapper) GetMessageCount() (int, error) {
    storageManager := storage.RestoreStorageManager(m.registry)

    var count int

    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("COUNT(*)")
    sb.From("netmail")
    sb.Where(sb.Equal("nmArchived", 0))
    query1, args := sb.Build()

    storageManager.Query(query1, args, func(rows *sql.Rows) error {
        return rows.Scan(&count)
    })

    return count, nil
}

func (self *NetmailMapper) Write(msg *NetmailMsg) error {
	storageManager := storage.RestoreStorageManager(self.registry)

	query1 := "INSERT INTO `netmail` " +
		"(nmMsgId, nmHash, nmFrom, nmTo, nmSubject, nmBody, nmDate, nmOrigAddr, nmDestAddr, nmPacket) " +
		"VALUES " +
		"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
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
	params = append(params, msg.GetPacket())

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		return nil
	})

	return err1
}

func (self *NetmailMapper) GetMessageByHash(msgHash string) (*NetmailMsg, error) {
	storageManager := storage.RestoreStorageManager(self.registry)

	var result *NetmailMsg

	query1 := "SELECT `nmId`, `nmMsgId`, `nmSubject`, `nmViewCount`, `nmFrom`, `nmTo`, `nmDate`, `nmBody`, `nmOrigAddr`, `nmDestAddr`, `nmPacket` FROM `netmail` WHERE `nmHash` = ?"
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
		var nmPacket []byte

		err2 := rows.Scan(&nmId, &nmMsgId, &nmSubject, &nmViewCount, &nmFrom, &nmTo, &nmDate, &nmBody, &nmOrigAddr, &nmDestAddr, &nmPacket)
		if err2 != nil {
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
		newMsg.SetPacket(nmPacket)

		result = newMsg

		return nil
	})

	return result, nil
}

func (self *NetmailMapper) ViewMessageByHash(msgHash string) error {
	storageManager := storage.RestoreStorageManager(self.registry)

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
	storageManager := storage.RestoreStorageManager(self.registry)

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
	storageManager := storage.RestoreStorageManager(self.registry)

	query1 := "DELETE FROM `netmail` WHERE `nmHash` = ?"
	var params []interface{}
	params = append(params, msgHash)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		return nil
	})

	return err1
}

func (m *NetmailMapper) ArchiveMessageByHash(msgHash string) error {
    storageManager := storage.RestoreStorageManager(m.registry)
    query1 := "UPDATE `netmail` SET `nmArchived` = 1 WHERE `nmHash` = ?"
    var params []interface{}
    params = append(params, msgHash)
    return storageManager.Exec(query1, params, func(result sql.Result, err error) error {
        return nil
    })
}
