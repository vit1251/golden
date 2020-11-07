package mapper

type FileArea struct {
	name    string
	Path    string
	Summary string
	Count   int
}

func NewFileArea() *FileArea {
	fa := new(FileArea)
	return fa
}

func (self *FileArea) SetName(name string) {
	self.name = name
}

func (self FileArea) GetName() string {
	return self.name
}
