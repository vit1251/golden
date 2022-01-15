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
