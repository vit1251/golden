package installer

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/storage"
)

type MigrationManager struct {
	conn *sql.DB
}

func NewMigrationManager(sm *storage.StorageManager) *MigrationManager {
	mm := new(MigrationManager)
	mm.conn = sm.GetConnection()
	return mm
}

func (mm *MigrationManager) Check() {
	new(Migration_20200402_0000).Up(mm.conn)
	new(Migration_20200402_0001).Up(mm.conn)
	new(Migration_20200402_0002).Up(mm.conn)
	new(Migration_20200402_0003).Up(mm.conn)
	new(Migration_20200402_0004).Up(mm.conn)
	new(Migration_20200402_0005).Up(mm.conn)
	new(Migration_20200417_0006).Up(mm.conn)
}
