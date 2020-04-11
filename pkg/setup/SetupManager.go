package setup

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/vit1251/golden/pkg/storage"
	"log"
	"os"
	"runtime"
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
	conn *sql.DB
}

func NewConfigManager(sm *storage.StorageManager) *ConfigManager {
	sem := new(ConfigManager)
	sem.conn = sm.GetConnection()
	/* Set parameter */
	sem.Register("main", "RealName", "Realname is you English version your real name (example: Dmitri Kamenski)")
	sem.Register("main", "Origin", "Origin was provide BBS station name and network address")
	sem.Register("main", "TearLine", "Tearline provide person sign in all their messages")
	sem.Register("main", "Inbound", "Directory where store incoming packets")
	sem.Register("main", "TempInbound", "Directory where should be process incoming packets")
	sem.Register("main", "TempOutbound", "Directory where process outbound packet")
	sem.Register("main", "Outbound", "Directory where store outbound packet")
	sem.Register("main", "Address", "FidoNet network point address (i.e. POINT address)")
	sem.Register("main", "NetAddr", "FidoNet network BOSS address (example: f24.n5023.z2.binkp.net:24554)")
	sem.Register("main", "Password", "FidoNet point password")
	sem.Register("main", "Link", "FidoNet uplink provide (i.e. BOSS address)")
	sem.Register("main", "Country", "Country where user is seat")
	sem.Register("main", "City", "City where user is seat")
	sem.Register("main", "FileBox", "Directory where store inbound file area files")

	/* Recover default parameters */
	sem.restoreDefault()

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

func (self *ConfigManager) Set(section string, name string, value string) (error) {

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

func (self *ConfigManager) Register(section string, name string, summary string) (error) {

	param := new(ConfigValue)
	param.Section = section
	param.Name = name
	param.Summary = summary

	self.Params = append(self.Params, param)

	return nil
}

func (self *ConfigManager) Audit(msg string) (error) {

	/* Store audit message in parameter storage */

	return nil

}

func (self *ConfigManager) restoreDefault() error {

	self.Set("main", "RealName", "Alice Cooper")
	self.Set("main", "Country", "Russia")
	self.Set("main", "City", "Moscow")
	self.Set("main", "Origin", "Yo Adrian, I Did It! (c) Rocky II")
	self.Set("main", "Address", "2:5030/1592.15")
	self.Set("main", "Link", "2:5030/1592.0")

	ver := "1.2.11"
	stamp := "2020-04-11 14:30 MSK"
	branch := "master"

	if runtime.GOOS == "windows" {

		/* Tearline */
		arch := "WIN"

		newTearline := fmt.Sprintf("Golden/%s %s %s (%s)", arch, ver, stamp, branch)
		self.Set("main", "TearLine", newTearline)

		/* Directory */
		inboundDir := ".\\Inbound"
		os.MkdirAll(inboundDir, os.ModePerm)
		self.Set("main", "Inbound", inboundDir)

		outboundDir := ".\\Outbound"
		os.MkdirAll(outboundDir, os.ModePerm)
		self.Set("main", "Outbound", outboundDir)

		boxDir := ".\\Files"
		os.MkdirAll(outboundDir, os.ModePerm)
		self.Set("main", "FileBox", boxDir)

	} else if runtime.GOOS == "linux" {

		arch := "LNX"
		newTearline := fmt.Sprintf("Golden/%s %s %s (%s)", arch, ver, stamp, branch)
		self.Set("main", "TearLine", newTearline)

		/* Directory */
		inboundDir := "/var/spool/ftn/inb"
		os.MkdirAll(inboundDir, os.ModePerm)
		self.Set("main", "Inbound", inboundDir)

		outboundDir := "/var/spool/ftn/outb"
		os.MkdirAll(outboundDir, os.ModePerm)
		self.Set("main", "Outbound", outboundDir)

		boxDir := "/var/spool/ftn/files"
		os.MkdirAll(outboundDir, os.ModePerm)
		self.Set("main", "FileBox", boxDir)

	} else {

		arch := "UNKNOWN"
		newTearline := fmt.Sprintf("Golden/%s %s %s (%s)", arch, ver, stamp, branch)
		self.Set("main", "TearLine", newTearline)

		/* Directory */
		inboundDir := "/var/spool/ftn/inb"
		os.MkdirAll(inboundDir, os.ModePerm)
		self.Set("main", "Inbound", inboundDir)

		outboundDir := "/var/spool/ftn/outb"
		os.MkdirAll(outboundDir, os.ModePerm)
		self.Set("main", "Outbound", outboundDir)

		boxDir := "/var/spool/ftn/files"
		os.MkdirAll(outboundDir, os.ModePerm)
		self.Set("main", "FileBox", boxDir)

	}

	return nil
}

func (self *ConfigManager) Restore() error {

	/* Restore parameters */
	sqlStmt := "SELECT `section`, `name`, `value` FROM `settings`"
	rows, err2 := self.conn.Query(sqlStmt)
	if err2 != nil {
		return err2
	}
	defer rows.Close()
	for rows.Next() {

		var section string
		var name string
		var value string

		err3 := rows.Scan(&section, &name, &value)
		if err3 != nil {
			return err3
		}
		self.Set(section, name, value)
	}

	return nil
}

func (self *ConfigManager) Store() (error) {

	/* Prepare update query */
	stmt1, err1 := self.conn.Prepare("UPDATE `settings` SET `value` = ? WHERE `section` = ? AND `name` = ?")
	if err1 != nil {
		return err1
	}

	/* Prepare insert query */
	stmt2, err2 := self.conn.Prepare("INSERT INTO `settings` (`section`, `name`, `value`) VALUES (?, ?, ?)")
	if err2 != nil {
		return err2
	}

	/* Store parameters */
	for _, param := range self.Params {
		result, err3 := stmt1.Exec(param.Value, param.Section, param.Name)
		updateCount, err4 := result.RowsAffected()
		if err4 != nil {
			return err4
		}
		log.Printf("updateCount = %+v err3 = %+v", updateCount, err3)
		if updateCount == 0 {
			_, err5 := stmt2.Exec(param.Section, param.Name, param.Value)
			if err5 != nil {
				return err5
			}
		}
	}

	return nil

}
