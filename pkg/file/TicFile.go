package file

type TicFile struct {
	From     string
	To       string
	File     string
	Area     string
	Desc     string
	SeenBy []string
}

func NewTicFile() *TicFile {
	tic := new(TicFile)
	return tic
}

func (self *TicFile) AddSeenBy(sb string) {
	self.SeenBy = append(self.SeenBy, sb)
}
