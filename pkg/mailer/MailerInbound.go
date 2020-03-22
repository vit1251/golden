package mailer

import (
	"github.com/vit1251/golden/pkg/setup"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

type MailerInbound struct {
	SetupManager	*setup.SetupManager
}

type MailerInboundRecType int

const TypeUnknown MailerInboundRecType = 0
const TypeNetmail MailerInboundRecType = 1
const TypeARCmail MailerInboundRecType = 2
const TypeTICmail MailerInboundRecType = 3

type MailerInboundRec struct {
	Type         MailerInboundRecType    /**/
	AbsolutePath string                  /**/
	Name         string                  /**/
}

func (self *MailerInboundRec) SetAbsolutePath(absolutePath string) {
	self.AbsolutePath = absolutePath
}

func (self *MailerInboundRec) SetType(nodeType MailerInboundRecType) {
	self.Type = nodeType
}

func (self *MailerInboundRec) SetName(name string) {
	self.Name = name
}

func NewMailerInboundRec() *MailerInboundRec {
	return new(MailerInboundRec)
}

func NewMailerInbound(sm *setup.SetupManager) *MailerInbound {
	mi := new(MailerInbound)
	mi.SetupManager = sm
	return mi
}

func (self *MailerInbound) nodeTypePrediction(name string) (MailerInboundRecType) {

	var result MailerInboundRecType = TypeUnknown

	/* Check on packet message (direct Netmail) */
	var nodeExtension string = filepath.Ext(name)

	nodeExtension = strings.ToUpper(nodeExtension)

	if nodeExtension == ".PKT" {
		result = TypeNetmail
	}

	/* Check on ARCmail message */
	var arcExtensionPrefixList []string = []string{".MO", ".TU", ".WE", ".TH", ".FR", ".SA", ".SU"}
	for _, arcExtPrefix := range arcExtensionPrefixList {
		if strings.HasPrefix(nodeExtension, arcExtPrefix) {
			result = TypeARCmail
		}
	}

	/* Check on TIC message */
	if nodeExtension == ".TIC" {
		result = TypeTICmail
	}

	return result
}

func (self *MailerInbound) Scan() ([]*MailerInboundRec, error) {

	var result []*MailerInboundRec

	inb, err1 := self.SetupManager.Get("main", "Inbound", ".")

	items, err1 := ioutil.ReadDir(inb)
	if err1 != nil {
		return nil, err1
	}

	for _, item := range items {
		absPath := path.Join(inb, item.Name())
		itemMode := item.Mode()
		if itemMode.IsRegular() {
			rec := NewMailerInboundRec()
			rec.SetAbsolutePath(absPath)
			rec.SetName(item.Name())
			rec.SetType(self.nodeTypePrediction(absPath))
			result = append(result, rec)
		}
	}

	return result, nil

}
