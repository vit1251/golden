package cache


type FileEntryType int

const TypeUnknown FileEntryType = 0
const TypeNetmail FileEntryType = 1
const TypeARCmail FileEntryType = 2
const TypeTICmail FileEntryType = 3

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
