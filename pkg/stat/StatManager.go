package stat

import "github.com/vit1251/golden/pkg/setup"

type StatManager struct {
	Path string
}

func NewStatManager() (*StatManager) {
	sm := new(StatManager)
	sm.Path = setup.GetBasePath()
	return sm
}

func (self *StatManager) RegisterArchmailPacket(filename string) (error) {
	return nil
}

func (self *StatManager) RegisterPacket(filename string) (error) {
	return nil
}

func (self *StatManager) RegisterMessage(msgid string) (error) {
	return nil
}

func (self *StatManager) RegisterFile(filename string) (error) {
	return nil
}

func (self *StatManager) GetTotalMessageProcessCount() (int, error) {
	return 0, nil
}

func (self *StatManager) GetTodayMessageProcessCount() (int, error) {
	return 0, nil
}
