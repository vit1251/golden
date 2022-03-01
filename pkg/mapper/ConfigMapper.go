package mapper

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
	"log"
)

type ConfigMapper struct {
	Mapper
}

func NewConfigMapper(r *registry.Container) *ConfigMapper {
	newConfigMapper := new(ConfigMapper)
	newConfigMapper.SetRegistry(r)
	return newConfigMapper
}

func (self ConfigMapper) restoreStorageManager() *storage.StorageManager {
	storageManagerPtr := self.registry.Get("StorageManager")
	if storageManager, ok := storageManagerPtr.(*storage.StorageManager); ok {
		return storageManager
	} else {
		panic("no storage manager")
	}
}

func (self ConfigMapper) GetConfigFromDatabase() (*Config, error) {
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
		//
		log.Printf("param: section = %s name = %s value = %s", section, name, value)
		config.params = append(config.params, ConfigValue{
			section: section,
			name:    name,
			value:   value,
		})
		return nil
	})
	return config, nil
}

func (self *ConfigMapper) SetConfigToDatabase(config *Config) {
	log.Printf("Update parameters: config = %#v", config)
	for _, param := range config.params {
		log.Printf("Update parameter: section = %s name = %s value = %s", param.section, param.name, param.value)
		err1 := self.updateParameter(param.section, param.name, param.value)
		log.Printf("Update error: %#v", err1)
		if err1 != nil {
		} else {
			err2 := self.insertParameter(param.section, param.name, param.value)
			log.Printf("Insert error: %#v", err2)
		}
	}
}

func (self ConfigMapper) insertParameter(section string, name string, value string) error {
	return nil
}

func (self ConfigMapper) updateParameter(section string, name string, value string) error {

	storageManager := self.restoreStorageManager()

	query1 := "UPDATE `settings` SET `value` = ? WHERE `section` = ? AND `name` = ?"
	var params []interface{}
	params = append(params, value)   // 1 - value
	params = append(params, section) // 2 - section
	params = append(params, name)    // 3 name

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}
