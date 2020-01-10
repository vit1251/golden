package common

type GoldenEchomail struct {
	FromAddr   string
	From       string
	ToAddr     string
	To         string
	Subject    string
	MsgID      string
	ReplyID    string
	Area       string              /* Areaname             */
	Attributes int
	Body       string              /* Message body         */
}

func NewGoldenEchomail() (*GoldenEchomail) {
	gem := new(GoldenEchomail)
	return gem
}
