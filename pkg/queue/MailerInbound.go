package queue

import (
	cmn "github.com/vit1251/golden/internal/common"
	"github.com/vit1251/golden/pkg/registry"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

type MailerInbound struct {
	registry *registry.Container
}

func newMailerInbound(registry *registry.Container) *MailerInbound {
	mi := new(MailerInbound)
	mi.registry = registry
	return mi
}

func (self *MailerInbound) nodeTypePrediction(name string) FileEntryType {

	var result FileEntryType = TypeUnknown

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

func (self *MailerInbound) Scan() ([]FileEntry, error) {

	var result []FileEntry

	inboundDirectory := cmn.GetInboundDirectory()

	items, err1 := ioutil.ReadDir(inboundDirectory)
	if err1 != nil {
		return nil, err1
	}

	for _, item := range items {
		absPath := path.Join(inboundDirectory, item.Name())
		itemMode := item.Mode()
		if itemMode.IsRegular() {
			rec := NewFileEntry()
			rec.SetAbsolutePath(absPath)
			rec.SetName(item.Name())
			rec.SetType(self.nodeTypePrediction(absPath))
			result = append(result, *rec)
		}
	}

	return result, nil

}
