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

func (self *EchoMapper) getAreaListCount() ([]*Area, error) {

	storageManager := storage.RestoreStorageManager(self.registry)

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

func (self *EchoMapper) getAreaListNewCount() ([]*Area, error) {

	storageManager := storage.RestoreStorageManager(self.registry)

	var result []*Area

	query1 := "SELECT `msgArea`, count(`msgId`) AS `msgCount` FROM `message` WHERE `msgViewCount` = 0 GROUP BY `msgArea` ORDER BY `msgArea` ASC"
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
		a.SetNewMessageCount(count)
		result = append(result, a)
		return nil
	})

	return result, nil
}

func (self *EchoMapper) GetMessageHeaders(echoTag string) ([]msg.Message, error) {

	storageManager := storage.RestoreStorageManager(self.registry)

	var result []msg.Message

	//	query1 := "SELECT `msgId`, `msgMsgId`, `msgReply`, `msgArea`, `msgHash`, `msgSubject`, `msgViewCount`, `msgFrom`, `msgTo`, `msgDate` FROM `message` WHERE `msgArea` = $1 ORDER BY `msgDate` ASC, `msgId` ASC"
	//	var params []interface{}
	//	params = append(params, echoTag)

	sb := sqlbuilder.NewSelectBuilder()
	sb.Select("msgId", "msgMsgId", "msgReply", "msgArea", "msgHash", "msgSubject", "msgViewCount", "msgFrom", "msgTo", "msgDate")
	sb.From("message")
	sb.Where(sb.Equal("msgArea", echoTag))
	sb.OrderBy("msgDate ASC", "msgId ASC")

	query1, args := sb.Build()

	log.Printf("EchoMapper: query = %+v args = %+v", query1, args)

	storageManager.Query(query1, args, func(rows *sql.Rows) error {

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

	return result, nil
}

func (self *EchoMapper) GetMessageByHash(echoTag string, msgHash string) (*msg.Message, error) {

	storageManager := storage.RestoreStorageManager(self.registry)

	var result *msg.Message

	query1 := "SELECT `msgId`, `msgReply`, `msgArea`, `msgMsgId`, `msgHash`, `msgSubject`, `msgFrom`, `msgOrigAddr`, `msgTo`, `msgContent`, `msgDate`, `msgPacket` FROM `message` WHERE `msgArea` = $1 AND `msgHash` = $2"
	var params []interface{}
	params = append(params, echoTag)
	params = append(params, msgHash)

	storageManager.Query(query1, params, func(rows *sql.Rows) error {

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

	return result, nil
}

func (self *EchoMapper) ViewMessageByHash(echoTag string, msgHash string) error {

	storageManager := storage.RestoreStorageManager(self.registry)

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

func (self *EchoMapper) RemoveMessageByHash(echoTag string, msgHash string) error {

	storageManager := storage.RestoreStorageManager(self.registry)

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

func (self *EchoMapper) IsMessageExistsByHash(echoTag string, msgHash string) (bool, error) {

	storageManager := storage.RestoreStorageManager(self.registry)

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

func (self *EchoMapper) RemoveMessagesByAreaName(echoTag string) error {

	storageManager := storage.RestoreStorageManager(self.registry)

	query1 := "DELETE FROM `message` WHERE `msgArea` = $1"
	var params []interface{}
	params = append(params, echoTag)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		log.Printf("Insert complete with: err = %+v", err)
		return nil
	})

	return err1

}

func (self *EchoMapper) UpdateAreaMessageCounters(areas []Area) ([]Area, error) {

	var newAreas []Area

	//log.Printf("areas = %+v", areas)

	/* Get message count */
	areas2, err1 := self.getAreaListCount()
	if err1 != nil {
		return nil, err1
	}
	//log.Printf("areas = %+v", areas2)

	/* Get message new count */
	areas3, err2 := self.getAreaListNewCount()
	if err2 != nil {
		//log.Printf("err2 = %+v", err2)
		return nil, err2
	}
	//log.Printf("areas = %+v", areas3)

	/* Update original areas values */
	for _, area := range areas {

		/* Search area count */
		for _, area2 := range areas2 {
			var areaName string = area.GetName()
			var area2Name string = area2.GetName()
			if strings.EqualFold(areaName, area2Name) {
				//log.Printf("area = '%+v' area2 = '%+v'", areaName, area2Name)
				area.MessageCount = area2.MessageCount
			}
		}

		/* Search area new count */
		for _, area3 := range areas3 {
			var areaName string = area.GetName()
			var area3Name string = area3.GetName()
			if strings.EqualFold(areaName, area3Name) {
				//log.Printf("area = '%+v' area3 = '%+v'", areaName, area3Name)
				area.SetNewMessageCount(area3.GetNewMessageCount())
			}
		}

		newAreas = append(newAreas, area)

	}

	return newAreas, nil
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
