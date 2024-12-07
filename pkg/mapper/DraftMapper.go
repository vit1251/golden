package mapper

import (
	"database/sql"
	"log"

	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
)

type DraftMapper struct {
	Mapper
}

func NewDraftMapper(r *registry.Container) *DraftMapper {
	newDraftMapper := new(DraftMapper)
	newDraftMapper.SetRegistry(r)
	return newDraftMapper
}

func (self DraftMapper) GetDraftMessages(state int) ([]Draft, error) {
	var result []Draft

	storageManager := storage.RestoreStorageManager(self.registry)

	query1 := "SELECT `draftId`, `draftUUID`, `draftSubject`, `draftArea`, `draftBody`, `draftDone` FROM `draft` WHERE `draftDone` = ? ORDER BY `draftId` DESC"

	var params []interface{}
	params = append(params, state)

	storageManager.Query(query1, params, func(rows *sql.Rows) error {
		var draftId string
		var draftUUID string
		var draftSubject string
		var draftArea string
		var draftBody string
		var draftDone int

		err3 := rows.Scan(&draftId, &draftUUID, &draftSubject, &draftArea, &draftBody, &draftDone)
		if err3 != nil {
			return err3
		}

		newDraft := NewDraft()
		newDraft.SetId(draftId)
		newDraft.SetUUID(draftUUID)
		newDraft.SetSubject(draftSubject)
		newDraft.SetArea(draftArea)
		newDraft.SetBody(draftBody)
		newDraft.SetState(draftDone)

		result = append(result, *newDraft)

		return nil
	})

	return result, nil
}

func (self DraftMapper) RegisterNewDraft(draft Draft) error {
	storageManager := storage.RestoreStorageManager(self.registry)

	query1 := "INSERT INTO `draft` (`draftUUID`, `draftArea`,`draftBody`,`draftDest`,`draftDestAddr`,`draftSubject`,`draftReply`) VALUES (?,?,?,?,?,?,?)"
	var params []interface{}
	params = append(params, draft.GetUUID())
	params = append(params, draft.GetArea())
	params = append(params, draft.GetBody())
	params = append(params, draft.GetTo())
	params = append(params, draft.GetToAddr())
	params = append(params, draft.GetSubject())
	params = append(params, draft.GetReply())

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		if err != nil {
			log.Printf("DraftMapper: Fail on RegisterNewDraft: err = %+v", err)
		}
		return nil
	})

	return err1
}

func (self DraftMapper) GetDraftByUUID(uuid string) (*Draft, error) {
	var result *Draft

	storageManager := storage.RestoreStorageManager(self.registry)

	query1 := "SELECT `draftId`, `draftUUID`, `draftDest`, `draftDestAddr`, `draftSubject`, `draftArea`, `draftBody`, `draftDone`, `draftReply` FROM `draft` WHERE `draftUUID` = ?"

	var params []interface{}
	params = append(params, uuid)

	storageManager.Query(query1, params, func(rows *sql.Rows) error {
		var draftId string
		var draftUUID string
		var draftDest string
		var draftDestAddr string
		var draftSubject string
		var draftArea string
		var draftBody string
		var draftDone int
		var draftReply string

		err3 := rows.Scan(&draftId, &draftUUID, &draftDest, &draftDestAddr, &draftSubject, &draftArea, &draftBody, &draftDone, &draftReply)
		if err3 != nil {
			return err3
		}

		newDraft := NewDraft()
		newDraft.SetId(draftId)
		newDraft.SetUUID(draftUUID)
		newDraft.SetSubject(draftSubject)
		newDraft.SetArea(draftArea)
		newDraft.SetBody(draftBody)
		newDraft.SetState(draftDone)
		newDraft.SetTo(draftDest)
		newDraft.SetToAddr(draftDestAddr)
		newDraft.SetReply(draftReply)

		result = newDraft

		return nil
	})

	return result, nil
}

func (self DraftMapper) UpdateDraft(newDraft Draft) error {
	storageManager := storage.RestoreStorageManager(self.registry)

	query1 := "UPDATE `draft` SET `draftSubject` = ?, `draftBody` = ?, `draftDest` = ?, `draftDestAddr` = ? WHERE `draftUUID` = ?"
	var params []interface{}
	params = append(params, newDraft.GetSubject())
	params = append(params, newDraft.GetBody())
	params = append(params, newDraft.GetTo())
	params = append(params, newDraft.GetToAddr())
	params = append(params, newDraft.GetUUID()) // NOTE - Always keep at a last ...

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		if err != nil {
			log.Printf("DraftMapper: update: err = %+v", err)
		}
		return nil
	})

	return err1
}

func (self DraftMapper) RemoveByUUID(uuid string) error {
	storageManager := storage.RestoreStorageManager(self.registry)

	query1 := "DELETE FROM `draft` WHERE `draftUUID` = ?"
	var params []interface{}
	params = append(params, uuid)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		if err != nil {
			log.Printf("DraftMapper: remove: err = %+v", err)
		}
		return nil
	})

	return err1
}
