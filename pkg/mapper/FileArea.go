package mapper

type FileArea struct {
	name     string
	Path     string
	summary  string
	mode     string
	charset  string
	count    int
	order    int
	newCount int
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
	self.summary = summary
}

func (self FileArea) GetSummary() string {
	return self.summary
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

func (self *FileArea) SetMode(mode string) {
	self.mode = mode
}

func (self FileArea) GetMode() string {
	return self.mode
}

func (self *FileArea) GetOrder() int {
	return self.order
}

func (self *FileArea) SetCount(count int) {
	self.count = count
}

func (self *FileArea) GetCount() int {
	return self.count
}

func (self *FileArea) SetNewCount(count int) {
	self.newCount = count
}

func (self *FileArea) GetNewCount() int {
	return self.newCount
}
