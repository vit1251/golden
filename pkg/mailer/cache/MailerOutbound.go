package cache

import (
	cmn "github.com/vit1251/golden/pkg/common"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/setup"
	"io/ioutil"
	"log"
	"path"
)

type MailerOutbound struct {
	registry     *registry.Container
}

func (self *MailerOutbound) TransmitFile(filename string) {
	log.Printf("Schedule to transmit %s", filename)
}

func NewMailerOutbound(r *registry.Container) *MailerOutbound {
	mo := new(MailerOutbound)
	mo.registry = r
	return mo
}

func (self *MailerOutbound) GetItems() ([]FileEntry, error) {

	var items []FileEntry

	outboundDirectory := cmn.GetOutboundDirectory()
	files, err2 := ioutil.ReadDir(outboundDirectory)
	if err2 != nil {
		return nil, err2
	}

	for _, f := range files {
		entry := NewFileEntry()
		entry.AbsolutePath = path.Join(outboundDirectory, f.Name())
		entry.Name = f.Name()
		items = append(items, *entry)
	}

	return items, nil
}

func (self *MailerOutbound) restoreConfigManager() *setup.ConfigManager {
	managerPtr := self.registry.Get("ConfigManager")
	if manager, ok := managerPtr.(*setup.ConfigManager); ok {
		return manager
	} else {
		panic("no config manager")
	}
}
