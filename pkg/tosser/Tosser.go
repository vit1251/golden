package tosser

import (
	"github.com/vit1251/golden/pkg/registry"
	"log"
	"time"
)

type TosserService struct {
	registry.Service
}

func NewTosser(registry *registry.Container) *TosserService {
	tosser := new(TosserService)
	tosser.SetRegistry(registry)
	return tosser
}

func (self *TosserService) Toss() {

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
