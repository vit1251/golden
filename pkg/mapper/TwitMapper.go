package mapper

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/registry"
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

	storageManager := self.restoreStorageManager()

	var result []Twit

	query1 := "SELECT `twitId`, `twitName` FROM `twit` ORDER BY `twitId` ASC"
	var params []interface{}

	storageManager.Query(query1, params, func(rows *sql.Rows) error {

		var twitId string
		var twitName string

		err2 := rows.Scan(&twitId, &twitName)
		if err2 != nil{
			return err2
		}

		newTwit := NewTwit()
		newTwit.SetName(twitName)

		result = append(result, *newTwit)

		return nil
	})

	return result, nil
}

func (self TwitMapper) RegisterTwitByName(twitName string) error {

	storageManager := self.restoreStorageManager()

	query := "INSERT INTO `twit` (`twitName`) VALUES ( ? )"

	var params []interface{}
	params = append(params, twitName)

	err1 := storageManager.Exec(query, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}