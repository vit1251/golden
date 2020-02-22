package stat

import (
	"database/sql"
	"fmt"
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

	Dupe int

	PacketReceived  int
	PacketSent      int

	MessageReceived int
	MessageSent     int

}

func NewStatManager(conn *sql.DB) *StatManager {
	sm := new(StatManager)
	sm.conn = conn
	sm.checkSchema()
	sm.createStat()
	return sm
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

func (self *StatManager) checkSchema() {

	query1 := "CREATE TABLE IF NOT EXISTS `stat` (" +
		"    statId INTEGER NOT NULL PRIMARY KEY," +
		"    statDate CHAR(10) NOT NULL," +
		"    statMessageRXcount INTEGER DEFAULT 0," +
		"    statMessageTXcount INTEGER DEFAULT 0," +
		"    statFileRXcount INTEGER DEFAULT 0," +
		"    statFileTXcount INTEGER DEFAULT 0" +
		")"
	log.Printf("query = %s", query1)
	self.conn.Exec(query1)

	query2 := "CREATE UNIQUE INDEX \"uniq_stat_statDate\" ON \"stat\" (\"statDate\" ASC)"
	log.Printf("query = %s", query2)
	self.conn.Exec(query2)

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
