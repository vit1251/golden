package installer

import (
	"database/sql"
	"fmt"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
	"log"
	"sort"
	"time"
)

type MigrationManager struct {
	conn            *sql.DB
	registry        *registry.Container
	applyMigrations []string
}

func NewMigrationManager(registry *registry.Container) *MigrationManager {
	mm := new(MigrationManager)
	mm.registry = registry
	mm.prepareConnection()
	mm.checkMigrationSchema()
	mm.restoreApplyMigrations()
	return mm
}

func (self *MigrationManager) prepareConnection() {
	service := self.registry.Get("StorageManager")
	if storageManager, ok := service.(*storage.StorageManager); ok {
		self.conn = storageManager.GetConnection()
	} else {
		panic("no storage service")
	}
}

var migrations *MigrationList = NewMigrationList()

func (self *MigrationManager) checkMigrationSchema() {
	query1 := "CREATE TABLE IF NOT EXISTS `migration` (" +
		"    migrationKey CHAR(10) NOT NULL PRIMARY KEY," +
		"    migrationDate INTEGER NOT NULL" +
		")"
	if _, err := self.conn.Exec(query1); err != nil {
		log.Panicf("Fail initialization migration subsystem: err = %+v", err)
	}
}

func (self *MigrationManager) restoreApplyMigrations() error {

	query := "SELECT `migrationKey` FROM `migration`"
	if rows, err1 := self.conn.Query(query); err1 == nil {
		for rows.Next() {
			var migrationKey string
			err2 := rows.Scan(&migrationKey)
			if err2 != nil {
				rows.Close()
				return err2
			}
			self.applyMigrations = append(self.applyMigrations, migrationKey)
		}
		rows.Close()
	} else {
		return err1
	}

	return nil
}

func (self *MigrationManager) registerMigration(migrationKey string) {
	unixTime := time.Now().Unix()
	query1 := "INSERT INTO `migration` (`migrationKey`,`migrationDate`) VALUES (?,?)"
	if _, err := self.conn.Exec(query1, migrationKey, unixTime); err != nil {
		log.Panicf("Fail initialization migration subsystem: err = %+v", err)
	}
}

func (self *MigrationManager) isApply(migrationKey string) bool {
	for _, applyMigrationKey := range self.applyMigrations {
		if applyMigrationKey == migrationKey {
			return true
		}
	}
	return false
}

func (self *MigrationManager) Check() {

	/* Show apply migration */
	log.Printf("MigrationManager: Apply %d migrations: migrations = %q", len(self.applyMigrations), self.applyMigrations)

	/* Get migations */
	keys := migrations.GetList()

	/* Setup migration orders */
	sort.Strings(keys)

	/* Process migrations */
	for _, migrationKey := range keys {
		m := migrations.GetByKey(migrationKey)
		if m != nil {
			err1 := self.checkMigration(migrationKey, m)
			if err1 != nil {
				panic(err1)
			}
		}
	}

}

func (self *MigrationManager) checkMigration(migrationKey string, m *Migration) error {
	var err error
	var status string
	if self.isApply(m.ID) {
		status = "SKIP"
	} else {
		err = m.Up(self.conn)
		if err != nil {
			status = fmt.Sprintf("FAIL ( %s )", err)
		} else {
			self.registerMigration(migrationKey)
			status = "PASS"
		}
	}
	log.Printf("MigrationManager: Apply migration: key = %s status = %s", m.ID, status)
	return err
}
