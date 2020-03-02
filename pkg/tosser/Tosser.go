package tosser

import (
	"github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/charset"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/stat"
)

type Tosser struct {
	MessageManager *msg.MessageManager
	StatManager    *stat.StatManager
	SetupManager   *setup.SetupManager
	FileManager    *file.FileManager
	AreaManager    *area.AreaManager
	CharsetManager *charset.CharsetManager
}

func NewTosser(cm *charset.CharsetManager, am *area.AreaManager, mm *msg.MessageManager, sm* stat.StatManager, setupm*setup.SetupManager, fm*file.FileManager) *Tosser {
	tosser := new(Tosser)
	tosser.CharsetManager = cm
	tosser.AreaManager = am
	tosser.MessageManager = mm
	tosser.StatManager = sm
	tosser.SetupManager = setupm
	tosser.FileManager = fm
	return tosser
}

func (self *Tosser) Toss() {
	self.ProcessInbound()
	self.ProcessOutbound()
}
