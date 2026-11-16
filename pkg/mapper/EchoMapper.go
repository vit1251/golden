package mapper

import (
	"database/sql"
	"github.com/huandu/go-sqlbuilder"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
	"log"
	"strings"
)

type EchoMapper struct {
	Mapper
}

func NewEchoMapper(r *registry.Container) *EchoMapper {
	manager := new(EchoMapper)
	manager.SetRegistry(r)
	return manager
}

func (self *EchoMapper) GetAreaList() ([]string, error) {

	storageManager := storage.RestoreStorageManager(self.registry)

	var result []string

	sb := sqlbuilder.NewSelectBuilder()
	sb.Distinct()
	sb.Select("msgArea AS name")
	sb.From("message")
	sb.OrderBy("name ASC")
	query1, args := sb.Build()

	err := storageManager.Query(query1, args, func(rows *sql.Rows) error {
		var name string
		err2 := rows.Scan(&name)
		if err2 != nil {
			return err2
		}
		result = append(result, name)
		return nil
	})
	if err != nil {
	    return nil, err
	}

	return result, nil
}

func (self *EchoMapper) getAreaListCount() ([]*Area, error) {

	storageManager := storage.RestoreStorageManager(self.registry)

	var result []*Area

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("msgArea", "COUNT(msgId) AS msgCount")
	sb.From("message")
	sb.Where(sb.Equal("msgArchived", 0))
	sb.GroupBy("msgArea")
	sb.OrderBy("msgArea ASC")
	query1, args := sb.Build()

	err := storageManager.Query(query1, args, func(rows *sql.Rows) error {
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
	if err != nil {
	    return nil, err
	}

	return result, nil
}

func (self *EchoMapper) getAreaListNewCount() ([]*Area, error) {

	storageManager := storage.RestoreStorageManager(self.registry)

	var result []*Area

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("msgArea", "COUNT(msgId) AS msgCount")
	sb.From("message")
	sb.Where(sb.Equal("msgViewCount", 0), sb.Equal("msgArchived", 0))
	sb.GroupBy("msgArea")
	sb.OrderBy("msgArea ASC")
	query1, args := sb.Build()

	err := storageManager.Query(query1, args, func(rows *sql.Rows) error {
		var name string
		var count int

		err2 := rows.Scan(&name, &count)
		if err2 != nil {
			return err2
		}
		a := NewArea()
		a.SetName(name)
		a.SetNewMessageCount(count)
		result = append(result, a)
		return nil
	})
	if err != nil {
	    return nil, err
	}

	return result, nil
}

func (m *EchoMapper) GetMessageHeaders(echoTag string) ([]msg.Message, error) {

    storageManager := storage.RestoreStorageManager(m.registry)

    var result []msg.Message

    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("msgId", "msgMsgId", "msgReply", "msgArea", "msgHash", "msgSubject", "msgViewCount", "msgFrom", "msgTo", "msgDate")
    sb.From("message")
    sb.Where(sb.Equal("msgArea", echoTag))
    sb.Where(sb.Equal("msgArchived", 0))
    sb.OrderBy("msgDate ASC", "msgId ASC")
    query1, args := sb.Build()

    log.Printf("EchoMapper: query = %+v args = %+v", query1, args)

    err := storageManager.Query(query1, args, func(rows *sql.Rows) error {
		var ID string
		var msgId string
		var reply string
		var msgHash *string
		var subject string
		var area string
		var from string
		var to string
		var msgDate int64
		var viewCount int

		err2 := rows.Scan(&ID, &msgId, &reply, &area, &msgHash, &subject, &viewCount, &from, &to, &msgDate)
		if err2 != nil {
			return err2
		}

		newMsg := msg.NewMessage()
		if msgHash != nil {
			newMsg.SetMsgHash(*msgHash)
		}
		newMsg.SetID(ID)
		newMsg.SetMsgID(msgId)
		newMsg.SetReply(reply)
		newMsg.SetArea(area)
		newMsg.SetSubject(subject)
		newMsg.SetFrom(from)
		newMsg.SetTo(to)
		newMsg.SetUnixTime(msgDate)
		newMsg.SetViewCount(viewCount)

		result = append(result, *newMsg)

		return nil
    })
    if err != nil {
	return nil, err
    }

    return result, nil
}

func (self *EchoMapper) GetMessageByHash(echoTag string, msgHash string) (*msg.Message, error) {

	storageManager := storage.RestoreStorageManager(self.registry)

	var result *msg.Message

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select(
	    "msgId", "msgReply", "msgArea", "msgMsgId", "msgHash",
	    "msgSubject", "msgFrom", "msgOrigAddr", "msgTo", "msgContent", "msgDate", "msgPacket",
	)
	sb.From("message")
	sb.Where(sb.Equal("msgArea", echoTag), sb.Equal("msgHash", msgHash))
	query1, args := sb.Build()

	err := storageManager.Query(query1, args, func(rows *sql.Rows) error {

		var ID string
		var reply string
		var msgMsgId string
		var msgArea string
		var msgHash *string
		var subject string
		var msgFrom string
		var msgOrigAddr string
		var to string
		var content string
		var packet []byte
		var written int64

		err1 := rows.Scan(&ID, &reply, &msgArea, &msgMsgId, &msgHash, &subject, &msgFrom, &msgOrigAddr, &to, &content, &written, &packet)
		if err1 != nil {
			return err1
		}
		log.Printf("subject = %q", subject)

		/**/
		newMsg := msg.NewMessage()
		newMsg.SetReply(reply)
		newMsg.SetArea(msgArea)
		newMsg.SetMsgID(msgMsgId)
		newMsg.SetSubject(subject)
		newMsg.SetID(ID)
		newMsg.SetUnixTime(written)
		if msgHash != nil {
			newMsg.SetMsgHash(*msgHash)
		}
		newMsg.SetFrom(msgFrom)
		newMsg.SetFromAddr(msgOrigAddr)
		newMsg.SetTo(to)
		newMsg.SetContent(content)
		newMsg.SetPacket(packet)

		/* Save result */
		result = newMsg

		return nil
	})
	if err != nil {
	    return nil, err
	}

	return result, nil
}

func (self *EchoMapper) ViewMessageByHash(echoTag string, msgHash string) error {

	storageManager := storage.RestoreStorageManager(self.registry)

	ub := sqlbuilder.NewUpdateBuilder()
	ub.Update("message")
	ub.SetMore("msgViewCount = msgViewCount + 1")
	ub.Where(ub.Equal("msgArea", echoTag), ub.Equal("msgHash", msgHash))
	query1, args := ub.Build()

	return storageManager.Exec(query1, args, func(result sql.Result, err error) error {
		return err
	})
}

func (self *EchoMapper) RemoveMessageByHash(echoTag string, msgHash string) error {

	storageManager := storage.RestoreStorageManager(self.registry)

	db := sqlbuilder.NewDeleteBuilder()
	db.DeleteFrom("message")
	db.Where(db.Equal("msgArea", echoTag), db.Equal("msgHash", msgHash))
	query1, args := db.Build()

	return storageManager.Exec(query1, args, func(result sql.Result, err error) error {
		return err
	})

}

func (self *EchoMapper) IsMessageExistsByHash(echoTag string, msgHash string) (bool, error) {

	storageManager := storage.RestoreStorageManager(self.registry)

	var result bool = false

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("msgId")
	sb.From("message")
	sb.Where(sb.Equal("msgArea", echoTag), sb.Equal("msgHash", msgHash))
	query1, args := sb.Build()

	err := storageManager.Query(query1, args, func(rows *sql.Rows) error {
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
	if err != nil {
	    return false, err
	}

	return result, nil
}

func (self *EchoMapper) Write(msg msg.Message) error {

	storageManager := storage.RestoreStorageManager(self.registry)

	/* Step 3. Make prepare SQL insert query */
	query1 := "INSERT INTO `message` " +
		"(`msgMsgId`, `msgReply`, `msgHash`, `msgArea`, `msgFrom`, `msgTo`, `msgSubject`, `msgContent`, `msgDate`, `msgPacket`, `msgOrigAddr`) " +
		"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	var params []interface{}
	params = append(params, msg.MsgID) // 1
	params = append(params, msg.Reply) // 2
	params = append(params, msg.Hash)
	params = append(params, msg.Area)
	params = append(params, msg.From)
	params = append(params, msg.To)
	params = append(params, msg.Subject)
	params = append(params, msg.Content)
	params = append(params, msg.UnixTime)
	params = append(params, msg.Packet)
	params = append(params, msg.FromAddr)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		if err != nil {
			log.Printf("Insert complete with: err = %+v", err)
		}
		return nil
	})

	return err1

}

func (self *EchoMapper) GetMessageNewCount() (int, error) {

	storageManager := storage.RestoreStorageManager(self.registry)

	var newMessageCount int

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("COUNT(msgId)")
	sb.From("message")
	sb.Where(sb.Equal("msgViewCount", 0), sb.Equal("msgArchived", 0))
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

func (self *EchoMapper) RemoveMessagesByAreaName(echoTag string) error {

	storageManager := storage.RestoreStorageManager(self.registry)

	db := sqlbuilder.NewDeleteBuilder()
	db.DeleteFrom("message")
	db.Where(db.Equal("msgArea", echoTag))
	query1, args := db.Build()

	return storageManager.Exec(query1, args, func(result sql.Result, err error) error {
		return err
	})

}

func (self *EchoMapper) UpdateAreaMessageCounters(areas []*Area) error {

	// Part 1. Area counts
	areaCounts, err1 := self.getAreaListCount()
	if err1 != nil {
		return err1
	}
        counts := make(map[string]int, len(areaCounts))
	for _, a := range areaCounts {
	    counts[strings.ToUpper(a.GetName())] = a.GetMessageCount()
	}
	for _, area := range areas {
	    key := strings.ToUpper(area.GetName())
	    if count, ok := counts[key]; ok {
		area.SetMessageCount(count)
	    }
	}

	// Part 2. Area new counts
	areaNewCounts, err2 := self.getAreaListNewCount()
	if err2 != nil {
		return err2
	}
	newCounts := make(map[string]int, len(areaNewCounts))
	for _, a := range areaNewCounts {
	    newCounts[strings.ToUpper(a.GetName())] = a.GetNewMessageCount()
	}
	for _, area := range areas {
	    key := strings.ToUpper(area.GetName())
	    if newCount, ok := newCounts[key]; ok {
	        area.SetNewMessageCount(newCount)
	    }
	}

	return nil
}

func (self *EchoMapper) MarkAllReadByAreaName(echoTag string) error {

	storageManager := storage.RestoreStorageManager(self.registry)

	query1 := "UPDATE `message` SET `msgViewCount` = `msgViewCount` + 1 WHERE `msgArea` = $1"
	var params []interface{}
	params = append(params, echoTag)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		log.Printf("Error: Update problme: %+v", err)
		return err
	})

	return err1

}

//
func (self *EchoMapper) GetMessageHeadersPage(echoTag string, limit, offset int) ([]msg.Message, error) {
    storageManager := storage.RestoreStorageManager(self.registry)
    var result []msg.Message

    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("msgId", "msgMsgId", "msgReply", "msgArea", "msgHash", "msgSubject", "msgViewCount", "msgFrom", "msgTo", "msgDate")
    sb.From("message")
    sb.Where(sb.Equal("msgArea", echoTag), sb.Equal("msgArchived", 0))
    sb.OrderBy("msgDate ASC", "msgId ASC")
    sb.Limit(limit)
    sb.Offset(offset)
    query1, args := sb.Build()

    err := storageManager.Query(query1, args, func(rows *sql.Rows) error {
	var ID string
	var msgId string
	var reply string
	var msgHash *string
	var subject string
	var area string
	var from string
	var to string
	var msgDate int64
	var viewCount int

	err2 := rows.Scan(&ID, &msgId, &reply, &area, &msgHash, &subject, &viewCount, &from, &to, &msgDate)
	if err2 != nil {
	    return err2
	}

	newMsg := msg.NewMessage()
	if msgHash != nil {
	    newMsg.SetMsgHash(*msgHash)
	}
	newMsg.SetID(ID)
	newMsg.SetMsgID(msgId)
	newMsg.SetReply(reply)
	newMsg.SetArea(area)
	newMsg.SetSubject(subject)
	newMsg.SetFrom(from)
	newMsg.SetTo(to)
	newMsg.SetUnixTime(msgDate)
	newMsg.SetViewCount(viewCount)

	result = append(result, *newMsg)
	return nil
    })
    if err != nil {
	return nil, err
    }

    return result, nil
}

// GetMessageCount — общее количество сообщений в эхе
func (self *EchoMapper) GetMessageCount(echoTag string) (int, error) {
    storageManager := storage.RestoreStorageManager(self.registry)

    var count int

    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("COUNT(*)")
    sb.From("message")
    sb.Where(sb.Equal("msgArea", echoTag))
    sb.Where(sb.Equal("msgArchived", 0))

    query1, args := sb.Build()

    err1 := storageManager.Query(query1, args, func(rows *sql.Rows) error {
	return rows.Scan(&count)
    })

    return count, err1
}

func (m *EchoMapper) ArchiveMessageByHash(echoTag string, msgHash string) error {

    storageManager := storage.RestoreStorageManager(m.registry)

    ub := sqlbuilder.NewUpdateBuilder()
    ub.Update("message")
    ub.Set(ub.Assign("msgArchived", 1))
    ub.Where(ub.Equal("msgArea", echoTag), ub.Equal("msgHash", msgHash))
    query1, args := ub.Build()

    return storageManager.Exec(query1, args, func(result sql.Result, err error) error {
        return err
    })
}

func (m *EchoMapper) PurgeArchivedMessages(areaName string) error {
    storageManager := storage.RestoreStorageManager(m.registry)

    db := sqlbuilder.NewDeleteBuilder()
    db.DeleteFrom("message")
    db.Where(db.Equal("msgArea", areaName), db.Equal("msgArchived", 1))
    query1, args := db.Build()

    return storageManager.Exec(query1, args, func(result sql.Result, err error) error {
        return err
    })
}

func (self *EchoMapper) GetArchivedMessageCount(areaName string) (int, error) {
    storageManager := storage.RestoreStorageManager(self.registry)

    var count int

    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("COUNT(*)")
    sb.From("message")
    sb.Where(sb.Equal("msgArea", areaName), sb.Equal("msgArchived", 1))
    query1, args := sb.Build()

    err1 := storageManager.Query(query1, args, func(rows *sql.Rows) error {
        return rows.Scan(&count)
    })

    return count, err1
}
