package installer

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/storage"
	"log"
	"sort"
)

type MigrationManager struct {
	conn *sql.DB
}

func NewMigrationManager(sm *storage.StorageManager) *MigrationManager {
	mm := new(MigrationManager)
	mm.conn = sm.GetConnection()
	return mm
}

var migrations *MigrationList = NewMigrationList()

func (mm *MigrationManager) Check() {

	/* Get migations */
	keys := migrations.GetList()

	/* Setup migration orders */
	sort.Strings(keys)

	/* Process migrations */
	for _, migrationKey := range keys {
		m := migrations.GetByKey(migrationKey)
		if m != nil {
			log.Printf("Process migration: id = %q - Up", m.ID)
			if err := m.Up(mm.conn); err != nil {
				log.Printf("Error in migration: ID = %s err = %+v", m.ID, err)
			}
		}
	}

}
