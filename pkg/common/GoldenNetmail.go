package common

type GoldenNetmail struct {
	FromAddr   string
	From       string
	ToAddr     string
	To         string
	Subject    string
	MsgID      string
	ReplyID    string
	Attributes int
	Body       string              /* Message body         */
}

func NewGoldenNetmail() (*GoldenNetmail) {
	gnm := new(GoldenNetmail)
	return gnm
}
