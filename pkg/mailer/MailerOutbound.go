package mailer

import (
	"github.com/vit1251/golden/pkg/setup"
	"io/ioutil"
	"log"
	"path"
)

type MailerOutbound struct {
	SetupManager *setup.ConfigManager
}

type Item struct {
	Name string
	AbsolutePath string
//	Type
}

func NewItem() *Item {
	i := new(Item)
	return i
}

func (self *MailerOutbound) TransmitFile(filename string) {
	log.Printf("Schedule to transmit %s", filename)
}

func NewMailerOutbound(sm *setup.ConfigManager) *MailerOutbound {
	mo := new(MailerOutbound)
	mo.SetupManager = sm
	return mo
}

func (self *MailerOutbound) GetItems() ([]*Item, error) {

	var items []*Item

	outb, err1 := self.SetupManager.Get("main", "Outbound", "")
	if err1 != nil {
		return nil, err1
	}

	files, err2 := ioutil.ReadDir(outb)
	if err2 != nil {
		return nil, err2
	}

	for _, f := range files {
		log.Printf("Oubound item %s", f.Name())
		i := NewItem()
		i.AbsolutePath = path.Join(outb, f.Name())
		i.Name = f.Name()
		items = append(items, i)
	}

	return items, nil
}
