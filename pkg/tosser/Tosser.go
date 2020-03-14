package tosser

import (
	"github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/charset"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/stat"
	"go.uber.org/dig"
)

type Tosser struct {
	MessageManager *msg.MessageManager
	StatManager    *stat.StatManager
	SetupManager   *setup.SetupManager
	FileManager    *file.FileManager
	AreaManager    *area.AreaManager
	CharsetManager *charset.CharsetManager
}

func NewTosser(c *dig.Container) *Tosser {
	tosser := new(Tosser)
	c.Invoke(func(cm *charset.CharsetManager, am *area.AreaManager, mm *msg.MessageManager, sm *stat.StatManager, setm *setup.SetupManager, fm *file.FileManager) {
		tosser.CharsetManager = cm
		tosser.AreaManager = am
		tosser.MessageManager = mm
		tosser.StatManager = sm
		tosser.SetupManager = setm
		tosser.FileManager = fm
	})
	return tosser
}

func (self *Tosser) Toss() {
	self.ProcessInbound()
	self.ProcessOutbound()
}
