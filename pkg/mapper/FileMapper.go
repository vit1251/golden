package mapper

import (
    "database/sql"
    cmn "github.com/vit1251/golden/internal/common"
    "log"
    "os"
    "path/filepath"
    "github.com/huandu/go-sqlbuilder"
    "github.com/vit1251/golden/pkg/registry"
    "github.com/vit1251/golden/pkg/storage"
)

type FileMapper struct {
	Mapper
}

func NewFileMapper(r *registry.Container) *FileMapper {
	newFileMapper := new(FileMapper)
	newFileMapper.SetRegistry(r)
	return newFileMapper
}

func (m *FileMapper) GetFileHeaders(echoTag string) ([]File, error) {

    storageManager := storage.RestoreStorageManager(m.registry)
    var result []File

    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("fileArea", "indexName", "fileName", "fileDesc", "fileTime", "fileViewCount", "fileOrigin", "fileFrom", "fileTo", "fileSize", "fileCrc")
    sb.From("file")
    sb.Where(sb.Equal("fileArea", echoTag), sb.Equal("fileArchived", 0))
    query1, args := sb.Build()

    err := storageManager.Query(query1, args, func(rows *sql.Rows) error {

		var fileArea string
		var indexName string
		var fileName string
		var fileDesc string
		var fileTime *int64
		var fileViewCount int
		var fileOrigin string
		var fileFrom string
		var fileTo string
		var fileSize int64
		var fileCrc string

		err2 := rows.Scan(&fileArea, &indexName, &fileName, &fileDesc, &fileTime, &fileViewCount,
		    &fileOrigin,
		    &fileFrom,
		    &fileTo,
		    &fileSize,
		    &fileCrc,
		)
		if err2 != nil {
			return err2
		}

		newFile := NewFile()
		newFile.SetArea(fileArea)
		newFile.SetDesc(fileDesc)
		newFile.SetFile(indexName)
		newFile.SetOrigName(fileName)
		newFile.SetViewCount(fileViewCount)
		if fileTime != nil {
			newFile.SetUnixTime(*fileTime)
		}

		result = append(result, *newFile)
		return nil
	})

	return result, err
}

func (self *FileMapper) CheckFileExists(tic File) (bool, error) {
	return true, nil
}

func (m *FileMapper) RegisterFile(tic File) error {
    storageManager := storage.RestoreStorageManager(m.registry)
    ib := sqlbuilder.NewInsertBuilder()
    ib.InsertInto("file")
    ib.Cols("indexName", "fileName", "fileArea", "fileDesc", "fileTime", "fileOrigin", "fileFrom", "fileTo", "fileSize", "fileCrc", "fileLDesc")
    ib.Values(tic.GetFile(), tic.GetOrigName(), tic.GetArea(), tic.GetDesc(), tic.GetUnixTime(), tic.GetOrigin(), tic.GetFrom(), tic.GetTo(), tic.GetSize(), tic.GetCrc(), tic.GetLDesc())
    query, args := ib.Build()
    return storageManager.Exec(query, args, func(result sql.Result, err error) error {
        return err
    })
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

func (self *FileMapper) RemoveFileByName(indexName string) error {
	storageManager := storage.RestoreStorageManager(self.registry)

	query1 := "DELETE FROM `file` WHERE `indexName` = ?"
	var params []interface{}
	params = append(params, indexName)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}

func (self *FileMapper) RemoveFilesByAreaName(areaName string) error {
	storageManager := storage.RestoreStorageManager(self.registry)

	query1 := "DELETE FROM `file` WHERE `fileArea` = $1"
	var params []interface{}
	params = append(params, areaName)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}

func (self *FileMapper) RemoveAreaByName(areaName string) error {
	storageManager := storage.RestoreStorageManager(self.registry)

	query1 := "DELETE FROM `filearea` WHERE `areaName` = $1"
	var params []interface{}
	params = append(params, areaName)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}

func (self *FileMapper) ViewFileByIndexName(fileArea string, indexName string) error {

	storageManager := storage.RestoreStorageManager(self.registry)

	query1 := "UPDATE `file` SET `fileViewCount` = `fileViewCount` + 1 WHERE `fileArea` = $1 AND `indexName` = $2"
	var params []interface{}
	params = append(params, fileArea)
	params = append(params, indexName)

	err1 := storageManager.Exec(query1, params, func(result sql.Result, err error) error {
		log.Printf("Insert complete with: err = %+v", err)
		return nil
	})

	return err1

}

func (m *FileMapper) GetFileNewCount() (int, error) {

    storageManager := storage.RestoreStorageManager(m.registry)

    var newMessageCount int

    query1 := "SELECT count(`fileId`) FROM `file` WHERE `fileViewCount` = 0 AND `fileArchived` = 0"
    var params []interface{}

    err := storageManager.Query(query1, params, func(rows *sql.Rows) error {
	err1 := rows.Scan(&newMessageCount)
	if err1 != nil {
	    return err1
	}
	return nil
    })

    return newMessageCount, err
}

func (m *FileMapper) GetFileByIndexName(echoTag string, indexName string) (*File, error) {
    storageManager := storage.RestoreStorageManager(m.registry)
    var result *File

    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("fileArea", "indexName", "fileName", "fileDesc", "fileTime", "fileViewCount", "fileOrigin", "fileFrom", "fileTo", "fileSize", "fileCrc")
    sb.From("file")
    sb.Where(sb.Equal("fileArea", echoTag), sb.Equal("indexName", indexName))
    query1, args := sb.Build()

    err := storageManager.Query(query1, args, func(rows *sql.Rows) error {
        var fileArea, indexName, fileName, fileDesc, fileOrigin, fileFrom, fileTo, fileCrc string
        var fileTime *int64
        var fileViewCount int
        var fileSize int64

        err2 := rows.Scan(&fileArea, &indexName, &fileName, &fileDesc, &fileTime, &fileViewCount, &fileOrigin, &fileFrom, &fileTo, &fileSize, &fileCrc)
        if err2 != nil { return err2 }

        newFile := NewFile()
        newFile.SetArea(fileArea)
        newFile.SetDesc(fileDesc)
        newFile.SetFile(indexName)
        newFile.SetOrigName(fileName)
        newFile.SetViewCount(fileViewCount)
        newFile.SetOrigin(fileOrigin)
        newFile.SetFrom(fileFrom)
        newFile.SetTo(fileTo)
        newFile.SetSize(fileSize)
        newFile.SetCrc(fileCrc)
        if fileTime != nil {
            newFile.SetUnixTime(*fileTime)
        }
        newFile.SetAbsolutePath(m.GetFileAbsolutePath(fileArea, indexName))
        result = newFile
        return nil
    })

    return result, err
}

func (m *FileMapper) ArchiveFileByName(indexName string) error {
    storageManager := storage.RestoreStorageManager(m.registry)
    query1 := "UPDATE `file` SET `fileArchived` = 1 WHERE `indexName` = ?"
    var params []interface{}
    params = append(params, indexName)
    return storageManager.Exec(query1, params, func(result sql.Result, err error) error {
        return err
    })
}

func (m *FileMapper) PurgeArchivedFiles(areaName string) error {
    storageManager := storage.RestoreStorageManager(m.registry)

    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("indexName")
    sb.From("file")
    sb.Where(sb.Equal("fileArea", areaName), sb.Equal("fileArchived", 1))
    query1, args := sb.Build()

    var files []string
    storageManager.Query(query1, args, func(rows *sql.Rows) error {
        var name string
        if err := rows.Scan(&name); err != nil {
            return err
        }
        files = append(files, name)
        return nil
    })

    delQuery := "DELETE FROM `file` WHERE `fileArea` = ? AND `fileArchived` = 1"
    var delParams []interface{}
    delParams = append(delParams, areaName)
    if err := storageManager.Exec(delQuery, delParams, func(result sql.Result, err error) error {
        return err
    }); err != nil {
        return err
    }

    for _, name := range files {
        p := m.GetFileAbsolutePath(areaName, name)
        if err := os.Remove(p); err != nil {
            log.Printf("Purge: remove %s: %v", p, err)
        }
    }

    return nil
}

func (m *FileMapper) GetArchivedFileCount(areaName string) (int, error) {
    storageManager := storage.RestoreStorageManager(m.registry)

    var count int

    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("COUNT(*)")
    sb.From("file")
    sb.Where(sb.Equal("fileArea", areaName), sb.Equal("fileArchived", 1))
    query1, args := sb.Build()

    err := storageManager.Query(query1, args, func(rows *sql.Rows) error {
        return rows.Scan(&count)
    })

    return count, err
}
