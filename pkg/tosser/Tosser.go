package tosser

import (
	"github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/charset"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/stat"
	"go.uber.org/dig"
	"log"
	"time"
)

type Tosser struct {
	MessageManager *msg.MessageManager
	StatManager    *stat.StatManager
	SetupManager   *setup.ConfigManager
	FileManager    *file.FileManager
	AreaManager    *area.AreaManager
	CharsetManager *charset.CharsetManager
}

func NewTosser(c *dig.Container) *Tosser {
	tosser := new(Tosser)
	c.Invoke(func(cm *charset.CharsetManager, am *area.AreaManager, mm *msg.MessageManager, sm *stat.StatManager, setm *setup.ConfigManager, fm *file.FileManager) {
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

	tosserStart := time.Now()
	log.Printf("Start tosser session")

	err1 := self.ProcessInbound()
	if err1 != nil {
		log.Printf("err = %+v", err1)
	}
	err2 := self.ProcessOutbound()
	if err2 != nil {
		log.Printf("err = %+v", err2)
	}

	log.Printf("Stop tosser session")
	elapsed := time.Since(tosserStart)
	log.Printf("Tosser session: %+v", elapsed)
}
