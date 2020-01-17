package mailer

import (
	"path"
	"strings"
	"path/filepath"
	"io/ioutil"
//	"log"
)

type MailerInbound struct {
	inboundDirectory string   /* Inbound directory     */
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

func NewMailerInboundRec() (*MailerInboundRec) {
	return new(MailerInboundRec)
}

func NewMailerInbound() (*MailerInbound) {
	mi := new(MailerInbound)
	return mi
}

func (self *MailerInbound) SetInboundDirectory(inboundDirectory string) {
	self.inboundDirectory = inboundDirectory
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

	items, err1 := ioutil.ReadDir(self.inboundDirectory)
	if err1 != nil {
		return nil, err1
	}

	for _, item := range items {
		absPath := path.Join(self.inboundDirectory, item.Name())
		itemMode := item.Mode()
		if (itemMode.IsRegular()) {
			rec := NewMailerInboundRec()
			rec.SetAbsolutePath(absPath)
			rec.SetType(self.nodeTypePrediction(absPath))
			result = append(result, rec)
		}
	}

	return result, nil

}
