package stat

import (
	"github.com/vit1251/golden/pkg/setup"
	"log"
)

type StatManager struct {
	Path string
}

type Stat struct {
	TicReceived      int
	TicSent          int
	EchomailReceived int
	EchomailSent     int
}

func NewStatManager() (*StatManager) {
	sm := new(StatManager)
	sm.Path = setup.GetBasePath()
	return sm
}

func (self *StatManager) RegisterNetmail(filename string) (error) {
	log.Printf("Stat: RegisterNetmail: %s", filename)

	return nil
}

func (self *StatManager) RegisterARCmail(filename string) (error) {
	log.Printf("Stat: RegisterARCmail: %s", filename)

	return nil
}

func (self *StatManager) RegisterFile(filename string) (error) {
	log.Printf("Stat: RegisterFile: %s", filename)
	return nil
}

func (self *StatManager) GetStat() (*Stat, error) {
	stat := new(Stat)
	return stat, nil
}
