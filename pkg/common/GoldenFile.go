package common

type GoldenFile struct {
	From     string     /* From */
	To       string
	File     string
	Area     string
	Desc     string
	Origin   string
	Size     string
	CRC      string
	Path   []string
}

func NewGoldenFile() (*GoldenFile) {
	gf := new(GoldenFile)
	return gf
}
