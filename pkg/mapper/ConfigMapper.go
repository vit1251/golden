package mapper

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
	"log"
)

type ConfigMapper struct {
	Mapper
	cacheConfig *Config
}

func NewConfigMapper(r *registry.Container) *ConfigMapper {
	newConfigMapper := new(ConfigMapper)
	newConfigMapper.SetRegistry(r)
	return newConfigMapper
}

func (self *ConfigMapper) GetConfig() (*Config, error) {
	newConfig, _ := self.restore()
	return newConfig, nil
}

func (self *ConfigMapper) Set(section string, name string, value string) error {

	newConfig, _ := self.restore()

	newConfig.Set(section, name, value)

	return nil

}

func (self ConfigMapper) Get(section string, name string) (string, bool) {

	newConfig, _ := self.restore()

	value, ok := newConfig.Get(section, name)

	return value, ok
}

func (self ConfigMapper) restoreStorageManager() *storage.StorageManager {
	storageManagerPtr := self.registry.Get("StorageManager")
	if storageManager, ok := storageManagerPtr.(*storage.StorageManager); ok {
		return storageManager
	} else {
		panic("no storage manager")
	}
}

func (self ConfigMapper) restore() (*Config, error) {
	if self.cacheConfig == nil {
		self.cacheConfig, _ = self.restoreFromDatabase()
	}
	return self.cacheConfig, nil
}

func (self ConfigMapper) restoreFromDatabase() (*Config, error) {

	storageManager := self.restoreStorageManager()

	config := NewConfig()

	query1 := "SELECT `section`, `name`, `value` FROM `settings`"
	var params []interface{}

	storageManager.Query(query1, params, func(rows *sql.Rows) error {

		var section string
		var name string
		var value string

		err3 := rows.Scan(&section, &name, &value)
		if err3 != nil {
			return err3
		}
		config.Insert(section, name, value)

		return nil
	})

	return config, nil
}

func (self ConfigMapper) UpdateValue(value string, section string, name string) error {

	storageManager := self.restoreStorageManager()

	query1 := "UPDATE `settings` SET `value` = ? WHERE `section` = ? AND `name` = ?"
	var params []interface{}
	params = append(params, value)
	params = append(params, section)
	params = append(params, name)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {

		if err != nil {
			return err
		}

		updateRowCount, err2 := result.RowsAffected()
		if err2 != nil {
			return err2
		}
		log.Printf("ConfigMapper: Update: section = %+v name = %+v value = %+v affected = %d", section, name, value, updateRowCount)

		return nil
	})

	return err1
}

func (self ConfigMapper) InsertValue(value string, section string, name string) error {

	storageManager := self.restoreStorageManager()

	query1 := "INSERT INTO `settings` (`section`, `name`, `value`) VALUES (?, ?, ?)"
	var params []interface{}
	params = append(params, section)
	params = append(params, name)
	params = append(params, value)
	err1 := storageManager.Exec(query1, params, func (result sql.Result, err error) error {
		log.Printf("ConfigMapper: InsertValue: Exec: err = %+v", err)
		return nil
	})
	return err1
}

func (self *ConfigMapper) Store(newConfig *Config) error {

	/* Update parameters in storage */
	params := newConfig.GetParams()
	for _, param := range params {
		if param.IsUpdate() {
			log.Printf("ConfigMapper: update: section = %s name = %s value = %s", param.Section, param.Name, param.GetValue())
			err1 := self.UpdateValue(param.GetValue(), param.Section, param.Name)
			if err1 != nil {
				return err1
			}
		}
	}

	/* Reset cache */
	self.cacheConfig = nil

	return nil
}
