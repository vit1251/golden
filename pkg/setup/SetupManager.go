package setup

import (
	"os/user"
	"path/filepath"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type ParamType int

const ParamString ParamType = 1

type SetupParam struct {
	Summary    string         /* Parameter summary     */
	Section    string         /* Parameter section     */
	Name       string         /* Parameter name        */
	Value      string         /* Parameter value       */
	IsSet      bool           /* Parameter exists mark */
	Type       ParamType      /* Parameter value type  */
}

func (self *SetupParam) SetValue(value string) {
	self.Value = value
}

type SetupManager struct {
	Path      string       /* Param stroage path */
	Params []*SetupParam   /* Param array        */
}

func NewSetupManager() (*SetupManager) {
	sm := new(SetupManager)

	/* Search user home directory */
	usr, err1 := user.Current()
	if err1 != nil {
		panic( err1 )
	}
	userHomeDirectory := usr.HomeDir
	log.Printf("userHomeDirectory = %+v", userHomeDirectory)

	/* Set parameter storage */
	sm.Path = filepath.Join(userHomeDirectory, "golden.sqlite3")

	/* Set parameter */
	sm.Register("main", "Origin", "Origin was provide BBS station name and network address")
	sm.Register("main", "TearLine", "Tearline provide person sign in all their messages")
	sm.Register("main", "Inbound", "Directory where store incoming packets")
	sm.Register("main", "TempInbound", "Directory where should be process incoming packets")
	sm.Register("main", "TempOutbound", "Directory where process outbound packet")
	sm.Register("main", "Outbound", "Directory where store outbound packet")
	sm.Register("main", "MessageBaseDirectory", "Directory where store message base")
	sm.Register("main", "Address", "FidoNet network point address (i.e. POINT address)")
	sm.Register("main", "Link", "FidoNet uplink provide (i.e. BOSS address)")

	/* Recover default parameters */
	sm.restoreDefault()

	/* Overwrite user parameters */
	err2 := sm.Restore()
	if err2 != nil {
		log.Printf("err1 = %+v", err1)
	}

	return sm
}

func (self *SetupManager) GetParams() ([]*SetupParam) {
	return self.Params
}

func (self *SetupManager) Set(section string, name string, value string) (error) {

	var updateCount int = 0

	for _, param := range self.Params {
		if param.Section == section && param.Name == name {
			param.SetValue(value)
			updateCount += 1
		}
	}

	log.Printf("Update %s parameter %d times", name, updateCount)

	return nil
}

func (self *SetupManager) Get(section string, name string, defaultValue string) (string, error) {

	var result string = defaultValue

	for _, param := range self.Params {
		if param.Section == section && param.Name == name {
			result = param.Value
		}
	}

	return result, nil

}

func (self *SetupManager) Register(section string, name string, summary string) (error) {

	param := new(SetupParam)
	param.Section = section
	param.Name = name
	param.Summary = summary

	self.Params = append(self.Params, param)

	return nil
}

func (self *SetupManager) Audit(msg string) (error) {

	/* Store audit message in parameter storage */

	return nil

}

func (self *SetupManager) restoreDefault() (error) {

	/* Step 1. Initialize parameters */
	self.Set("main", "FirstName", "Alice")
	self.Set("main", "LastName", "Cooper")
	self.Set("main", "Country", "Russia")
	self.Set("main", "City", "Moscow")
	self.Set("main", "Origin", "Yo Adrian, I Did It! (c) Rocky II")
	self.Set("main", "TearLine", "Golden/LNX 1.2.1 2020-01-05 18:29:20 MSK (master)")
	self.Set("main", "Address", "2:5030/1592.15")
	self.Set("main", "Link", "2:5030/1592.0")
	self.Set("main", "Inbound", "/var/spool/ftn/inb")
	self.Set("main", "Outbound", "/var/spool/ftn/outb")

	return nil
}

func (self *SetupManager) checkSchema(conn *sql.DB) (error) {

	/* Step 1. Create "settings" store */
	sqlStmt1 := "CREATE TABLE IF NOT EXISTS settings (" +
		    "    section CHAR(64) NOT NULL," +
		    "    name CHAR(64) NOT NULL," +
		    "    value CHAR(64) NOT NULL," +
		    "    UNIQUE(section, name)" +
		    ")"

	log.Printf("sql = %s", sqlStmt1)

	conn.Exec(sqlStmt1)

	return nil
}

func (self *SetupManager) Restore() (error) {

	/* Open SQL storage */
	db, err1 := sql.Open("sqlite3", self.Path)
	if err1 != nil {
		return err1
	}
	defer db.Close()

	/* Check schema */
	self.checkSchema(db)

	/* Restore parameters */
	sqlStmt := "SELECT `section`, `name`, `value` FROM `settings`"
	rows, err2 := db.Query(sqlStmt)
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

func (self *SetupManager) Store() (error) {

	/* Open SQL storage */
	db, err1 := sql.Open("sqlite3", self.Path)
	if err1 != nil {
		return err1
	}
	defer db.Close()

	/* Check schema */
	self.checkSchema(db)

	/* Prepare update query */
	stmt1, err2_1 := db.Prepare("UPDATE `settings` SET `value` = ? WHERE `section` = ? AND `name` = ?")
	if err2_1 != nil {
		return err2_1
	}

	/* Prepare insert query */
	stmt2, err2_2 := db.Prepare("INSERT INTO `settings` (`section`, `name`, `value`) VALUES (?, ?, ?)")
	if err2_2 != nil {
		return err2_2
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
