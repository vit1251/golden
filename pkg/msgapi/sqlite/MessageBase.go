package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type MessageBase struct {
	Path string            /* Database path */
}

func NewMessageBase() (*MessageBase, error) {
	mBase := new(MessageBase)
	mBase.Path = "/var/spool/ftn/echo/base.sqlite3"
	return mBase, nil
}

type MessageBaseConnection struct {
	Conn	*sql.DB     /* Database connection */
}

func (self *MessageBase) Open() (*MessageBaseConnection, error) {
	conn := new(MessageBaseConnection)
	if db, err := sql.Open("sqlite3", self.Path); err != nil {
		return nil, err
	} else {
		conn.Conn = db
	}
	return conn, nil
}

func (self *MessageBaseConnection) Close() {
	self.Conn.Close()
	self.Conn = nil
}
