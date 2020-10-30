package util

import "github.com/vit1251/golden/pkg/mailer/cache"

type Directory struct {
	items []cache.FileEntry
}

func (self *Directory) Push(entry cache.FileEntry) {
	self.items = append(self.items, entry)
}

func (self *Directory) RemoveByName(name string) {
	var newItems []cache.FileEntry
	for _, item := range self.items {
		if item.Name != name {
			newItems = append(newItems, item)
		}
	}
	self.items = newItems
}

func (self Directory) IsEmpty() bool {
	return len(self.items) == 0
}
