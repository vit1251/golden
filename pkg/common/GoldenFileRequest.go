package common

type GoldenFileRequest struct {
}

func NewGoldenFileRequest() (*GoldenFileRequest) {
	gfr := new(GoldenFileRequest)
	return gfr
}

func (self *GoldenFileRequest) Pack(filename string) (error) {
	return nil
}
