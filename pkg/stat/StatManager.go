package stat

import (
	"database/sql"
	"log"
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

var stat Stat

func NewStatManager(conn *sql.DB) *StatManager {
	sm := new(StatManager)
	sm.conn = conn
	sm.checkSchema()
	return sm
}

func (self *StatManager) RegisterNetmail(filename string) (error) {
	self.createStat()
	stat.NetmailReceived += 1
	return nil
}

func (self *StatManager) RegisterARCmail(filename string) (error) {
	self.createStat()
	stat.EchomailReceived += 1
	return nil
}

func (self *StatManager) RegisterFile(filename string) (error) {

	self.createStat()

	query1 := "UPDATE `stat` SET `statFileRXcount` = `statFileRXcount` + 1 WHERE `statDate` = ?"
	statDate := "2020-02-21"
	self.conn.Exec(query1, statDate)

	return nil
}

func (self *StatManager) GetStat() (*Stat, error) {
	return &stat, nil
}

func (self *StatManager) RegisterPacket(p string) error {

	//statDate := "2020-02-21"

	self.createStat()
	stat.PacketReceived += 1
	return nil
}

func (self *StatManager) RegisterDupe() error {
	self.createStat()
	stat.Dupe += 1
	return nil
}

func (self *StatManager) RegisterMessage() error {

	/* Create record */
	self.createStat()

	/* Update */
	query1 := "UPDATE `stat` SET `statMessageRXcount` = `statMessageRXcount` + 1 WHERE `statDate` = ?"
	statDate := "2020-02-21"
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

func (self *StatManager) createStat() {
	query1 := "INSERT INTO `stat` (`statDate`) VALUES ( ? )"
	statDate := "2020-02-21"
	self.conn.Exec(query1, statDate)
}
