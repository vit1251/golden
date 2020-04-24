package installer

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/storage"
	"log"
)

type IMigration interface {
	Up(conn *sql.DB) error
}

type MigrationManager struct {
	conn *sql.DB
}

func NewMigrationManager(sm *storage.StorageManager) *MigrationManager {
	mm := new(MigrationManager)
	mm.conn = sm.GetConnection()
	return mm
}

func (mm *MigrationManager) Check() {
	var migrations []IMigration

	migrations = append(migrations, new(Migration_000100))
	migrations = append(migrations, new(Migration_000101))
	migrations = append(migrations, new(Migration_000102))
	migrations = append(migrations, new(Migration_000200))
	migrations = append(migrations, new(Migration_000201))
	migrations = append(migrations, new(Migration_000300))
	migrations = append(migrations, new(Migration_000301))
	migrations = append(migrations, new(Migration_000302))
	migrations = append(migrations, new(Migration_000400))
	migrations = append(migrations, new(Migration_000500))
	migrations = append(migrations, new(Migration_000501))
	migrations = append(migrations, new(Migration_000600))
	migrations = append(migrations, new(Migration_000601))
	migrations = append(migrations, new(Migration_000700))
	migrations = append(migrations, new(Migration_000701))
	migrations = append(migrations, new(Migration_000702))
	migrations = append(migrations, new(Migration_000703))
	migrations = append(migrations, new(Migration_000704))
	migrations = append(migrations, new(Migration_000705))

	for _, m := range migrations {
		/* Check migration exists */
		log.Printf("Make migration: %T", m)
		if err := m.Up(mm.conn); err != nil {
			log.Printf("Fail make migration %T with error: msg = %+v", m, err)
		} else {
			/* Set migration success */
		}
	}

}
