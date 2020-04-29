package storage

import (
	"context"
	"database/sql"
	"log"
	"os/user"
	"path/filepath"
	"time"
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

func (self *StorageManager) Query(query string, params []interface{}, f func(rows *sql.Rows) error) error {
	var errRes error = nil

	parentContext := context.Background()
	ctx, cancel := context.WithTimeout(parentContext, 5*time.Second)
	defer cancel()

	log.Printf("sql = %+v params = %+v", query, params)
	start := time.Now()
	rows, err1 := self.conn.QueryContext(ctx, query, params...)
	if err1 != nil {
		panic(err1)
		return err1
	}
	defer rows.Close()
	for rows.Next() {
		err2 := f(rows)
		if err2 != nil {
			errRes = err2
			break
		}
	}
	age := time.Since(start)
	log.Printf("Storage: query: duration = %+v", age)
	return errRes
}

func (self *StorageManager) Exec(query string, params []interface{}, f func(err error)) error {
	log.Printf("sql = %+v params = %+v", query, params)
	start := time.Now()
	_, err1 := self.conn.Exec(query, params...)
	age := time.Since(start)
	log.Printf("Storage: exec: duration = %+v", age)
	f(err1)
	return nil
}
