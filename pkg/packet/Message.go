package packet

import (
	"time"
)

type Message struct {
	Area      string
	From      string
	To        string
	Subject   string
	Content   string
	UnixTime  int64          /* Unit time stampe */
	Time     *time.Time
}

func NewMessage() (*Message) {
	msg := new(Message)
	return msg
}

func (self *Message) SetTime(msgTime *time.Time) (error) {

	/* Set time */
	self.Time = msgTime

	/* Set unix time */
	if msgTime != nil {
		var stamp time.Time = *msgTime
		var unixTime int64 = stamp.Unix()
		self.UnixTime = unixTime
	}

	return nil
}