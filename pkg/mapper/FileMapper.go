package mapper

import (
	"database/sql"
	cmn "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
	"log"
	"path/filepath"
)

type FileMapper struct {
	Mapper
}

func NewFileMapper(r *registry.Container) *FileMapper {
	newFileMapper := new(FileMapper)
	newFileMapper.SetRegistry(r)
	return newFileMapper
}

func (self *FileMapper) GetFileHeaders(echoTag string) ([]File, error) {

	storageManager := self.restoreStorageManager()
	conn := storageManager.GetConnection() // TODO

	var result []File

	/* Step 2. Start SQL transaction */
	ConnTransaction, err := conn.Begin()
	if err != nil {
		return nil, err
	}

	sqlStmt := "SELECT `fileArea`, `fileName`, `fileDesc`, `fileTime` FROM `file` WHERE `fileArea` = $1"
	log.Printf("sql = %q echoTag = %q", sqlStmt, echoTag)
	rows, err1 := ConnTransaction.Query(sqlStmt, echoTag)
	if err1 != nil {
		log.Printf("error on query: err = %+v", err1)
		return nil, err1
	}
	defer rows.Close()
	for rows.Next() {

		var fileArea string
		var fileName string
		var fileDesc string
		var fileTime *int64

		err2 := rows.Scan(&fileArea, &fileName, &fileDesc, &fileTime)
		if err2 != nil {
			log.Printf("error on scan: err = %+v", err2)
			return nil, err2
		}

		newFile := NewFile()
		newFile.SetArea(fileArea)
		newFile.SetDesc(fileDesc)
		newFile.SetFile(fileName)
		if fileTime != nil {
			newFile.SetUnixTime(*fileTime)
		}

		result = append(result, *newFile)
	}

	ConnTransaction.Commit()

	return result, nil
}

func (self *FileMapper) CheckFileExists(tic File) (bool, error) {
	return true, nil
}

func (self *FileMapper) RegisterFile(tic File) error {

	storageManager := self.restoreStorageManager()

	query1 := "INSERT INTO `file` ( `fileName`, `fileArea`, `fileDesc`, `fileTime` ) VALUES ( ?, ?, ?, ? )"

	var params []interface{}
	params = append(params, tic.GetFile())
	params = append(params, tic.GetArea())
	params = append(params, tic.GetDesc())
	params = append(params, tic.GetUnixTime())

	storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		return err
	})

	return nil
}

func (self FileMapper) restoreStorageManager() *storage.StorageManager {
	managerPtr := self.registry.Get("StorageManager")
	if manager, ok := managerPtr.(*storage.StorageManager); ok {
		return manager
	} else {
		panic("no storage manager")
	}
}

func (self FileMapper) GetFileAbsolutePath(areaName string, name string) string {
	boxDirectory := cmn.GetFilesDirectory()
	path := filepath.Join(boxDirectory, areaName, name)
	return path
}

func (self *FileMapper) GetFileBoxAbsolutePath(areaName string) string {
	boxDirectory := cmn.GetFilesDirectory()
	path := filepath.Join(boxDirectory, areaName)
	return path
}

func (self *FileMapper) RemoveFileByName(fileName string) error {
	storageManager := self.restoreStorageManager()

	query1 := "DELETE FROM `file` WHERE `fileName` = ?"
	var params []interface{}
	params = append(params, fileName)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}

func (self *FileMapper) RemoveFilesByAreaName(areaName string) error {
	storageManager := self.restoreStorageManager()

	query1 := "DELETE FROM `file` WHERE `fileArea` = $1"
	var params []interface{}
	params = append(params, areaName)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}

func (self *FileMapper) RemoveAreaByName(areaName string) error {
	storageManager := self.restoreStorageManager()

	query1 := "DELETE FROM `filearea` WHERE `areaName` = $1"
	var params []interface{}
	params = append(params, areaName)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}
