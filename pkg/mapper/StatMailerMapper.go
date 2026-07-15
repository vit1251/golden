package mapper

import (
    "log"
    "time"
    "database/sql"
    "github.com/huandu/go-sqlbuilder"
    "github.com/vit1251/golden/pkg/registry"
    "github.com/vit1251/golden/pkg/storage"
)

type StatMailerMapper struct {
	Mapper
}

type StatMailer struct {
    SessionID    int64    /* Timestamp in nanoseconds      */
    SessionStart int64    /* Session start in milliseconds */
    SessionStop  int64    /* Session stop in milliseconds  */
    Status       string   /* Summary report                */
    FileRXcount  int      /**/
    FileTXcount  int      /**/
}

func (self StatMailer) GetDuration() uint64 {
    return uint64(self.SessionStop - self.SessionStart)
}

func (self *StatMailerMapper) GetMailerSummary() ([]StatMailer, error) {

    now := time.Now()
    startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.UTC)

    var result []StatMailer

    storageManager := storage.RestoreStorageManager(self.registry)

    sb := sqlbuilder.NewSelectBuilder()
    sb.Select("statMailerSessionStart", "statMailerSessionStop", "statMailerSummary", "statMailerFileRXcount", "statMailerFileTXcount")
    sb.From("stat_mailer")
    sb.Where(sb.GE("statMailerSessionStart", startOfMonth.UnixMilli()))
    sb.OrderBy("statMailerId DESC")
    query1, args := sb.Build()

    err1 := storageManager.Query(query1, args, func(rows *sql.Rows) error {

	var statMailerSessionStart int64
	var statMailerSessionStop int64
	var statMailerSummary string
	var statMailerFileRXcount int
	var statMailerFileTXcount int

	err2 := rows.Scan(&statMailerSessionStart, &statMailerSessionStop, &statMailerSummary, &statMailerFileRXcount, &statMailerFileTXcount)
	if err2 != nil {
	    return err2
	}

	summary := new(StatMailer)
	summary.SessionStart = statMailerSessionStart
	summary.SessionStop = statMailerSessionStop
	summary.Status = statMailerSummary
	summary.FileRXcount = statMailerFileRXcount
	summary.FileTXcount = statMailerFileTXcount
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

    ib := sqlbuilder.NewInsertBuilder()
    ib.InsertInto("stat_mailer")
    ib.Cols("statMailerSessionStart", "statMailerSessionStop", "statMailerSummary", "statMailerFileRXcount", "statMailerFileTXcount", "statMailerId")
    ib.Values(report.SessionStart, report.SessionStop, report.Status, report.FileRXcount, report.FileTXcount, sessionID)
    query, args := ib.Build()

    err1 := storageManager.Exec(query, args, func(result sql.Result, err error) error {
	return err
    })

    return err1
}

func (self *StatMailerMapper) updateMailerReport(report *StatMailer) error {

    storageManager := storage.RestoreStorageManager(self.registry)

    ub := sqlbuilder.NewUpdateBuilder()
    ub.Update("stat_mailer")
    ub.Set(
	ub.Assign("statMailerSessionStart", report.SessionStart),
        ub.Assign("statMailerSessionStop", report.SessionStop),
	ub.Assign("statMailerSummary", report.Status),
        ub.Assign("statMailerFileRXcount", report.FileRXcount),
	ub.Assign("statMailerFileTXcount", report.FileTXcount),
    )
    ub.Where(ub.Equal("statMailerId", report.SessionID))
    query, args := ub.Build()

    err1 := storageManager.Exec(query, args, func(result sql.Result, err error) error {
	return err
    })

    return err1
}

func NewStatMailerMapper(r *registry.Container) *StatMailerMapper {
    newStatMapper := new(StatMailerMapper)
    newStatMapper.SetRegistry(r)
    return newStatMapper
}
