package mapper

import (
	"database/sql"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
	"log"
	"time"
)

type StatMailerMapper struct {
	Mapper
}

type StatMailer struct {
	SessionID    int64  /* Timestamp in nanoseconds      */
	SessionStart int64  /* Session start in milliseconds */
	SessionStop  int64  /* Session stop in milliseconds  */
	Status       string /* Summary report                */
}

func (self StatMailer) GetDuration() uint64 {
	return uint64(self.SessionStop - self.SessionStart)
}

func (self *StatMailerMapper) GetMailerSummary() ([]StatMailer, error) {

	var result []StatMailer

	storageManager := storage.RestoreStorageManager(self.registry)

	query1 := "SELECT `statMailerSessionStart`, `statMailerSessionStop`, `statMailerSummary` FROM `stat_mailer` ORDER BY `statMailerId` DESC LIMIT 10"
	var params []interface{}

	err1 := storageManager.Query(query1, params, func(rows *sql.Rows) error {

		var statMailerSessionStart int64
		var statMailerSessionStop int64
		var statMailerSummary string

		err2 := rows.Scan(&statMailerSessionStart, &statMailerSessionStop, &statMailerSummary)
		if err2 != nil {
			return err2
		}

		summary := new(StatMailer)
		summary.SessionStart = statMailerSessionStart
		summary.SessionStop = statMailerSessionStop
		summary.Status = statMailerSummary
		result = append(result, *summary)

		return nil
	})

	return result, err1
}

func (self *StatMailerMapper) UpdateSession(mailerReport *StatMailer) error {
	var sessionID int64 = mailerReport.SessionID
	if sessionID == 0 {
		sessionID = time.Now().UnixNano()
		err1 := self.insertMailerReport(sessionID, mailerReport)
		if err1 != nil {
			log.Printf("error insert mailer stat: err = %#v", err1)
		} else {
			mailerReport.SessionID = sessionID
		}
	} else {
		err1 := self.updateMailerReport(mailerReport)
		if err1 != nil {
			log.Printf("error insert mailer stat: err = %#v", err1)
		}
	}
	return nil
}

func (self *StatMailerMapper) insertMailerReport(sessionID int64, report *StatMailer) error {
	storageManager := storage.RestoreStorageManager(self.registry)

	var query string
	query += "INSERT INTO `stat_mailer` "
	query += "   ( "
	query += "       `statMailerSessionStart`,  "
	query += "       `statMailerSessionStop`,   "
	query += "       `statMailerSummary`,  "
	query += "       `statMailerId`  "
	query += "   ) VALUES ("
	query += "       $1, "
	query += "       $2, "
	query += "       $3, "
	query += "       $4  "
	query += "   )"
	var params []interface{}
	params = append(params, report.SessionStart) // 1
	params = append(params, report.SessionStop)  // 2
	params = append(params, report.Status)       // 3
	params = append(params, sessionID)           // 4
	err1 := storageManager.Exec(query, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}

func (self *StatMailerMapper) updateMailerReport(report *StatMailer) error {

	storageManager := storage.RestoreStorageManager(self.registry)

	var query string
	query += "UPDATE `stat_mailer` "
	query += "   SET "
	query += "       `statMailerSessionStart` = $1, "
	query += "       `statMailerSessionStop`  = $2, "
	query += "       `statMailerSummary`      = $3  "
	query += " WHERE "
	query += "       `statMailerId` = $4"
	var params []interface{}
	params = append(params, report.SessionStart) // 1
	params = append(params, report.SessionStop)  // 2
	params = append(params, report.Status)       // 3
	params = append(params, report.SessionID)    // 4
	err1 := storageManager.Exec(query, params, func(result sql.Result, err error) error {
		return err
	})

	return err1
}

func NewStatMailerMapper(r *registry.Container) *StatMailerMapper {
	newStatMapper := new(StatMailerMapper)
	newStatMapper.SetRegistry(r)
	return newStatMapper
}
