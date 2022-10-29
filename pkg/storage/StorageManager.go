package storage

import (
	"context"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	cmn "github.com/vit1251/golden/internal/common"
	"github.com/vit1251/golden/pkg/registry"
	"log"
	"time"
)

type StorageManager struct {
	conn *sql.DB
}

func (self *StorageManager) GetConnection() *sql.DB {
	return self.conn
}

// / Initialize storage
func NewStorageManager(r *registry.Container) *StorageManager {
	sm := new(StorageManager)
	sm.Open()
	return sm
}

func (self *StorageManager) Open() error {

	storageFile := cmn.GetModernStorageFile()

	log.Printf("StorageManager: Open storage: path = %s", storageFile)

	db, err2 := sql.Open("sqlite3", storageFile)
	if err2 != nil {
		panic(err2)
	}
	self.conn = db

	return nil
}

func (self *StorageManager) Close() error {
	return self.conn.Close()
}

func (self *StorageManager) Query(query string, params []interface{}, f func(rows *sql.Rows) error) error {
	var errRes error = nil

	parentContext := context.Background()
	ctx, cancel := context.WithTimeout(parentContext, 5*time.Second)
	defer cancel()

	log.Printf("StorageManager: Query: sql = %+v params = %+v", query, params)

	start := time.Now()
	if rows, err1 := self.conn.QueryContext(ctx, query, params...); err1 == nil {
		for rows.Next() {
			err2 := f(rows)
			if err2 != nil {
				errRes = err2
				break
			}
		}
		err3 := rows.Close()
		if err3 != nil {
			errRes = err3
		}
	} else {
		errRes = err1
	}
	duration := time.Since(start)
	log.Printf("StorageManager: Query: duration = %+v err = %+v", duration, errRes)

	return errRes
}

func (self *StorageManager) Exec(query string, params []interface{}, f func(result sql.Result, err error) error) error {
	log.Printf("StorageManager: Exec: sql = %+v params = %+v", query, params)
	start := time.Now()
	result, err1 := self.conn.Exec(query, params...)
	duration := time.Since(start)
	log.Printf("StorageManager: Exec: duration = %+v err = %+v", duration, err1)
	return f(result, err1)
}
