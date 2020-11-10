package mapper

type FileArea struct {
	name    string
	Path    string
	Summary string
	charset string
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

func (self *FileArea) SetSummary(summary string) {
	self.Summary = summary
}

func (self FileArea) GetSummary() string {
	return self.Summary
}

func (self FileArea) GetCharset() string {
	if self.charset == "" {
		self.charset = "UTF-8"
	}
	return self.charset
}

func (self *FileArea) SetCharset(charset string) {
	self.charset = charset
}
