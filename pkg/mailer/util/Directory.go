package util

import (
	"github.com/vit1251/golden/pkg/queue"
)

type Directory struct {
	items []queue.FileEntry
}

func (self *Directory) Push(entry queue.FileEntry) {
	self.items = append(self.items, entry)
}

func (self *Directory) RemoveByName(name string) {
	var newItems []queue.FileEntry
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

func (self *Directory) Contains(name string) bool {
	for _, item := range self.items {
		if item.Name == name {
			return true
		}
	}
	return false
}
