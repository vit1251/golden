package setup

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vit1251/golden/pkg/storage"
	"log"
)

type ParamType int

const ParamString ParamType = 1

type ConfigValue struct {
	Summary    string         /* Parameter summary     */
	Section    string         /* Parameter section     */
	Name       string         /* Parameter name        */
	Value      string         /* Parameter value       */
	IsSet      bool           /* Parameter exists mark */
	Type       ParamType      /* Parameter value type  */
}

func (self *ConfigValue) SetValue(value string) {
	self.Value = value
}

type ConfigManager struct {
	Params []*ConfigValue
	StorageManager *storage.StorageManager
}

func NewConfigManager(sm *storage.StorageManager) *ConfigManager {
	sem := new(ConfigManager)
	sem.StorageManager = sm

	/* Set item i18n */
	sem.Register("main", "RealName", "Realname is you English version your real name (example: Dmitri Kamenski)")
	sem.Register("main", "Origin", "Origin was provide BBS station name and network address")
	sem.Register("main", "TearLine", "Tearline provide person sign in all their messages")
	sem.Register("main", "Inbound", "Directory where store incoming packets")
	sem.Register("main", "TempInbound", "Directory where should be process incoming packets")
	sem.Register("main", "TempOutbound", "Directory where process outbound packet")
	sem.Register("main", "Temp", "Temp directory where process packet")
	sem.Register("main", "Outbound", "Directory where store outbound packet")
	sem.Register("main", "Address", "FidoNet network point address (i.e. POINT address)")
	sem.Register("main", "NetAddr", "FidoNet network BOSS address (example: f24.n5023.z2.binkp.net:24554)")
	sem.Register("main", "Password", "FidoNet point password")
	sem.Register("main", "Link", "FidoNet uplink provide (i.e. BOSS address)")
	sem.Register("main", "Country", "Country where user is seat")
	sem.Register("main", "City", "City where user is seat")
	sem.Register("main", "FileBox", "Directory where store inbound file area files")
	sem.Register("main", "StationName", "Station name is your nickname")

	/* Overwrite user parameters */
	err2 := sem.Restore()
	if err2 != nil {
		panic(err2)
	}
	return sem
}

func (self *ConfigManager) GetParams() []*ConfigValue {
	return self.Params
}

func (self *ConfigManager) Set(section string, name string, value string) error {

	var updateCount int = 0

	for _, param := range self.Params {
		if param.Section == section && param.Name == name {
			param.SetValue(value)
			updateCount += 1
		}
	}
	if updateCount == 0 {
		log.Printf("config: parameter %s in section %s does not exists", name, section)
	} else {
		log.Printf("config: parameter %s in section %s update %d time(s)", name, section, updateCount)
	}

	return nil
}

func (self *ConfigManager) Get(section string, name string, defaultValue string) (string, error) {
	var result string = defaultValue
	for _, param := range self.Params {
		if param.Section == section && param.Name == name {
			result = param.Value
		}
	}
	return result, nil
}

func (self *ConfigManager) Register(section string, name string, summary string) error {

	param := new(ConfigValue)
	param.Section = section
	param.Name = name
	param.Summary = summary

	self.Params = append(self.Params, param)

	return nil
}

func (self *ConfigManager) Restore() error {

	query1 := "SELECT `section`, `name`, `value` FROM `settings`"
	var params []interface{}

	self.StorageManager.Query(query1, params, func(rows *sql.Rows) error {

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

func (self *ConfigManager) UpdateValue(value string, section string, name string) error {
	query1 := "UPDATE `settings` SET `value` = ? WHERE `section` = ? AND `name` = ?"
	var params []interface{}
	params = append(params, value)
	params = append(params, section)
	params = append(params, name)
	err1 := self.StorageManager.Exec(query1, params, func(result sql.Result, err error) error {
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

func (self *ConfigManager) InsertValue(value string, section string, name string) error {
	query1 := "INSERT INTO `settings` (`section`, `name`, `value`) VALUES (?, ?, ?)"
	var params []interface{}
	params = append(params, section)
	params = append(params, name)
	params = append(params, value)
	err1 := self.StorageManager.Exec(query1, params, func (result sql.Result, err error) error {
		log.Printf("ConfigManager: InsertValue: Exec: err = %+v", err)
		return nil
	})
	return err1
}

func (self *ConfigManager) Store() error {
	for _, param := range self.Params {
		err1 := self.UpdateValue(param.Value, param.Section, param.Name)
		log.Printf("ConfigManager: Store: UpdateValue: err1 = %+v", err1)
		if err1 != nil {
			err2 := self.InsertValue(param.Value, param.Section, param.Name)
			log.Printf("ConfigManager: Store: InsertValue: err2 = %+v", err2)
		}
	}
	return nil
}
