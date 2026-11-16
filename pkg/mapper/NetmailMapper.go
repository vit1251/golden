package mapper

import (
    "database/sql"

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

    err := storageManager.Query(query1, args, func(rows *sql.Rows) error {
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
    if err != nil {
	return nil, err
    }

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

    err := storageManager.Query(query1, args, func(rows *sql.Rows) error {
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
    if err != nil {
	return nil, err
    }

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

    err := storageManager.Query(query1, args, func(rows *sql.Rows) error {
        return rows.Scan(&count)
    })
    if err != nil {
	return 0, err
    }

    return count, nil
}

func (self *NetmailMapper) Write(msg *NetmailMsg) error {
	storageManager := storage.RestoreStorageManager(self.registry)

	ib := sqlbuilder.NewInsertBuilder()
	ib.InsertInto("netmail")
	ib.Cols("nmMsgId", "nmHash", "nmFrom", "nmTo", "nmSubject", "nmBody", "nmDate", "nmOrigAddr", "nmDestAddr", "nmPacket")
	ib.Values(msg.MsgID, msg.Hash, msg.From, msg.To, msg.Subject, msg.Content, msg.UnixTime, msg.OrigAddr, msg.DestAddr, msg.GetPacket())
	query1, args := ib.Build()

	return storageManager.Exec(query1, args, func(result sql.Result, err error) error {
		return err
	})

}

func (self *NetmailMapper) GetMessageByHash(msgHash string) (*NetmailMsg, error) {
	storageManager := storage.RestoreStorageManager(self.registry)

	var result *NetmailMsg

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("nmId", "nmMsgId", "nmSubject", "nmViewCount", "nmFrom", "nmTo", "nmDate", "nmBody", "nmOrigAddr", "nmDestAddr", "nmPacket")
	sb.From("netmail")
	sb.Where(sb.Equal("nmHash", msgHash))
	query1, args := sb.Build()

	err := storageManager.Query(query1, args, func(rows *sql.Rows) error {
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
	if err != nil {
	    return nil, err
	}

	return result, nil
}

func (self *NetmailMapper) ViewMessageByHash(msgHash string) error {
	storageManager := storage.RestoreStorageManager(self.registry)

	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("netmail")
	ub.SetMore("nmViewCount = nmViewCount + 1")
	ub.Where(ub.Equal("nmHash", msgHash))
	query1, args := ub.Build()

	return storageManager.Exec(query1, args, func(result sql.Result, err error) error {
		return err
	})
}

func (self *NetmailMapper) GetMessageNewCount() (int, error) {
	storageManager := storage.RestoreStorageManager(self.registry)

	var newMessageCount int

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(nmId)")
	sb.From("netmail")
	sb.Where(sb.Equal("nmViewCount", 0), sb.Equal("nmArchived", 0))
	query1, args := sb.Build()

	err := storageManager.Query(query1, args, func(rows *sql.Rows) error {
		err1 := rows.Scan(&newMessageCount)
		if err1 != nil {
			return err1
		}
		return nil
	})
	if err != nil {
	    return 0, err
	}

	return newMessageCount, nil
}

func (self *NetmailMapper) RemoveMessageByHash(msgHash string) error {
	storageManager := storage.RestoreStorageManager(self.registry)

	db := sqlbuilder.NewDeleteBuilder()
	db.DeleteFrom("netmail")
	db.Where(db.Equal("nmHash", msgHash))
	query1, args := db.Build()

	return storageManager.Exec(query1, args, func(result sql.Result, err error) error {
		return err
	})
}

func (m *NetmailMapper) ArchiveMessageByHash(msgHash string) error {
    storageManager := storage.RestoreStorageManager(m.registry)

    ub := sqlbuilder.NewUpdateBuilder()
    ub.Update("netmail")
    ub.Set(ub.Assign("nmArchived", 1))
    ub.Where(ub.Equal("nmHash", msgHash))
    query1, args := ub.Build()

    return storageManager.Exec(query1, args, func(result sql.Result, err error) error {
        return err
    })
}
