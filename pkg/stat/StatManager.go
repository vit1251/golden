package stat

import (
	"database/sql"
	"fmt"
	"github.com/vit1251/golden/pkg/storage"
	"log"
	"time"
)

type StatManager struct {
	conn *sql.DB
}

type Stat struct {
	TicReceived      int
	TicSent          int
	EchomailReceived int
	EchomailSent     int
	NetmailReceived  int
	NetmailSent      int

	PacketReceived   int
	PacketSent       int

	MessageReceived  int
	MessageSent      int

	SessionIn        int
	SessionOut       int
}

func NewStatManager(sm *storage.StorageManager) *StatManager {
	statm := new(StatManager)
	statm.conn = sm.GetConnection()
	statm.createStat()
	return statm
}

func (self *StatManager) RegisterNetmail(filename string) (error) {
	self.createStat()
	return nil
}

func (self *StatManager) RegisterARCmail(filename string) (error) {
	self.createStat()
	return nil
}

func (self *StatManager) RegisterFile(filename string) (error) {

	self.createStat()

	query1 := "UPDATE `stat` SET `statFileRXcount` = `statFileRXcount` + 1 WHERE `statDate` = ?"
	statDate := self.makeToday()
	self.conn.Exec(query1, statDate)

	return nil
}

func (self *StatManager) RegisterSession(in int, out int) error {

	self.createStat()

	query1 := "UPDATE `stat` SET `statFileRXcount` = `statFileRXcount` + 1 WHERE `statDate` = ?"
	statDate := self.makeToday()
	self.conn.Exec(query1, statDate)

	return nil
}

type SummaryRow struct {
	Date string
	Value int
}

func (self *StatManager) GetStatRow(statDate string) (*Stat, error) {

	var statMessageRXcount int64
	var statMessageTXcount int64

	/* Step 2. Start SQL transaction */
	ConnTransaction, err := self.conn.Begin()
	if err != nil {
		return nil, err
	}

	sqlStmt := "SELECT `statMessageRXcount`, `statMessageTXcount` FROM `stat` WHERE `statDate` = $1"
	log.Printf("sql = %q echoTag = %q", sqlStmt, statDate)
	rows, err1 := ConnTransaction.Query(sqlStmt, statDate)
	if err1 != nil {
		return nil, err1
	}
	defer rows.Close()
	for rows.Next() {

		err2 := rows.Scan(&statMessageRXcount, &statMessageTXcount)
		if err2 != nil{
			return nil, err2
		}

	}

	ConnTransaction.Commit()

	result := new(Stat)
	result.MessageReceived = int(statMessageRXcount)
	result.MessageSent = int(statMessageTXcount)
	//result.TicReceived = statFileRXcount
	//result.TicSent = statFileTXcount

	return result, nil
}

func (self *StatManager) GetStat() (*Stat, error) {
	statDate := self.makeToday()
	stat, err := self.GetStatRow(statDate)
	return stat, err
}

func (self *StatManager) RegisterPacket(p string) error {
	self.createStat()
	return nil
}

func (self *StatManager) RegisterDupe() error {
	self.createStat()
	return nil
}

func (self *StatManager) RegisterMessage() error {
	self.createStat()
	query1 := "UPDATE `stat` SET `statMessageRXcount` = `statMessageRXcount` + 1 WHERE `statDate` = ?"
	statDate := self.makeToday()
	self.conn.Exec(query1, statDate)
	return nil
}

func (self *StatManager) makeToday() string {
	currentTime := time.Now()
	//result := currentTime.Format("2006-01-02")
	result := fmt.Sprintf("%04d-%02d-%02d", currentTime.Year(), currentTime.Month(), currentTime.Day())
	return result
}

func (self *StatManager) createStat() {
	query1 := "INSERT INTO `stat` (`statDate`) VALUES ( ? )"
	statDate := self.makeToday()
	log.Printf("Create stat on %s", statDate)
	self.conn.Exec(query1, statDate)
}
