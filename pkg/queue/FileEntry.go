package queue

type FileEntryType int

const (
	TypeUnknown FileEntryType = 0
	TypeNetmail FileEntryType = 1
	TypeARCmail FileEntryType = 2
	TypeTICmail FileEntryType = 3
)

type FileEntry struct {
	Type         FileEntryType
	AbsolutePath string
	Name         string
}

func (self *FileEntry) SetAbsolutePath(absolutePath string) {
	self.AbsolutePath = absolutePath
}

func (self *FileEntry) SetType(nodeType FileEntryType) {
	self.Type = nodeType
}

func (self *FileEntry) SetName(name string) {
	self.Name = name
}

func NewFileEntry() *FileEntry {
	return new(FileEntry)
}
