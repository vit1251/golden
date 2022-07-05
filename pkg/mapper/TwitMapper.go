package mapper

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
)

type TwitMapper struct {
	Mapper
}

func NewTwitMapper(r *registry.Container) *TwitMapper {
	newTwitMapper := new(TwitMapper)
	newTwitMapper.SetRegistry(r)
	return newTwitMapper
}

func (self TwitMapper) GetTwitNames() ([]Twit, error) {

	storageManager := storage.RestoreStorageManager(self.registry)

	var result []Twit

	query1 := "SELECT `twitId`, `twitName` FROM `twit` ORDER BY `twitId` ASC"
	var params []interface{}

	storageManager.Query(query1, params, func(rows *sql.Rows) error {

		var twitId string
		var twitName string

		err2 := rows.Scan(&twitId, &twitName)
		if err2 != nil {
			return err2
		}

		newTwit := NewTwit()
		newTwit.SetId(twitId)
		newTwit.SetName(twitName)

		result = append(result, *newTwit)

		return nil
	})

	return result, nil
}

func (self TwitMapper) RegisterTwitByName(twitName string) error {

	storageManager := storage.RestoreStorageManager(self.registry)

	query := "INSERT INTO `twit` (`twitName`) VALUES ( ? )"

	var params []interface{}
	params = append(params, twitName)

	err1 := storageManager.Exec(query, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}

func (self TwitMapper) RemoveById(twitId string) error {

	storageManager := storage.RestoreStorageManager(self.registry)

	query1 := "DELETE FROM `twit` WHERE `twitId` = ?"
	var params []interface{}
	params = append(params, twitId)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		return nil
	})

	return err1

}
