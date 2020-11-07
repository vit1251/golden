package mapper

import (
	"database/sql"
	"errors"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
	"log"
)

type ConfigMapper struct {
	Mapper
	Params         []*ConfigValue
}

func NewConfigMapper(r *registry.Container) *ConfigMapper {
	newConfigMapper := new(ConfigMapper)

	newConfigMapper.SetRegistry(r)

	/* Set item i18n */
	newConfigMapper.Register("main", "RealName", "Realname is you English version your real name (example: Dmitri Kamenski)")
	newConfigMapper.Register("main", "Origin", "Origin was provide BBS station name and network address")
	newConfigMapper.Register("main", "TearLine", "Tearline provide person sign in all their messages")
	newConfigMapper.Register("main", "Address", "FidoNet network point address (i.e. POINT address)")
	newConfigMapper.Register("main", "NetAddr", "FidoNet network BOSS address (example: f24.n5023.z2.binkp.net:24554)")
	newConfigMapper.Register("main", "Password", "FidoNet point password")
	newConfigMapper.Register("main", "Link", "FidoNet uplink provide (i.e. BOSS address)")
	newConfigMapper.Register("main", "Country", "Country where user is seat")
	newConfigMapper.Register("main", "City", "City where user is seat")
	newConfigMapper.Register("main", "StationName", "Station name is your nickname")
	newConfigMapper.Register("mailer", "Interval", "Mailer interval")

	/* Overwrite user parameters */
	err2 := newConfigMapper.Restore()
	if err2 != nil {
		panic(err2)
	}

	return newConfigMapper
}

func (self *ConfigMapper) GetParams() []*ConfigValue {
	return self.Params
}

func (self *ConfigMapper) Set(section string, name string, value string) error {

	//var updateCount int = 0

	for _, param := range self.Params {
		if param.Section == section && param.Name == name {
			param.SetValue(value)
			//updateCount += 1
		}
	}

	//if updateCount == 0 {
	//	log.Printf("config: parameter %s in section %s does not exists", name, section)
	//} else {
	//	log.Printf("config: parameter %s in section %s update %d time(s)", name, section, updateCount)
	//}

	return nil
}

func (self ConfigMapper) Get(section string, name string) (string, bool) {
	for _, param := range self.Params {
		if param.Section == section && param.Name == name {
			return param.Value, true
		}
	}
	return "", false
}

func (self *ConfigMapper) Register(section string, name string, summary string) error {

	param := new(ConfigValue)
	param.Section = section
	param.Name = name
	param.Summary = summary

	self.Params = append(self.Params, param)

	return nil
}

func (self ConfigMapper) restoreStorageManager() *storage.StorageManager {
	storageManagerPtr := self.registry.Get("StorageManager")
	if storageManager, ok := storageManagerPtr.(*storage.StorageManager); ok {
		return storageManager
	} else {
		panic("no storage manager")
	}
}

func (self ConfigMapper) Restore() error {

	storageManager := self.restoreStorageManager()

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
		self.Set(section, name, value)

		return nil
	})

	return nil
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
		if updateRowCount != 1 {
			return errors.New("no update config parameters")
		}
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

func (self ConfigMapper) Store() error {
	for _, param := range self.Params {
		err1 := self.UpdateValue(param.Value, param.Section, param.Name)
		if err1 != nil {
			return err1
		}
	}
	return nil
}
