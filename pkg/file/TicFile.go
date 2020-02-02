package file

type TicFile struct {
	From string
	To   string
	File string
	Area string
	Desc string
}

func NewTicFile() *TicFile {
	tic := new(TicFile)
	return tic
}