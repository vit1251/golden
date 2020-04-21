package storage

import (
	"database/sql"
	"log"
	"os/user"
	"path/filepath"
)

type StorageManager struct {
	conn *sql.DB
}

func (self *StorageManager) GetConnection() *sql.DB {
	return self.conn
}

func (self *StorageManager) Close() {
	self.conn.Close()
}

func NewStorageManager() *StorageManager {

	sm := new(StorageManager)

	/* Initialize storage */
	usr, err1 := user.Current()
	if err1 != nil {
		panic(err1)
	}
	userHomeDir := usr.HomeDir
	log.Printf("userHomeDir = %+v", userHomeDir)
	userStoragePath := filepath.Join(userHomeDir, "golden.sqlite3")
	log.Printf("userStoragePath = %+v", userStoragePath)
	db, err2 := sql.Open("sqlite3", userStoragePath)
	if err2 != nil {
		panic(err2)
	}
	sm.conn = db
	log.Printf("db = %+v", db)

	return sm
}

func (self *StorageManager) Query(sql string, params []interface{}, cb func()) {

}

func (self *StorageManager) Exec(query string, params []interface{}, f func(err error)) {
	log.Printf("sql = %+v params = %+v", query, params)
	_, err1 := self.conn.Exec(query, params...)
	f(err1)
}
