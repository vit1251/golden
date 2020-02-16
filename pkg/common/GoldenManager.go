package common

import (
	"github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/stat"
	"github.com/vit1251/golden/pkg/tosser"
)

type GoldenMaster struct {
	SetupManager   *setup.SetupManager
	AreaManager    *area.AreaManager
	MessageManager *msg.MessageManager
	FileManager    *file.FileManager
	StatManager    *stat.StatManager
	TosserManager  *tosser.TosserManager
}

var master *GoldenMaster

func GetMaster() *GoldenMaster {
	if master == nil {
		master = new(GoldenMaster)
		master.SetupManager = setup.NewSetupManager()
		master.MessageManager = msg.NewMessageManager()
		master.AreaManager = area.NewAreaManager(master.MessageManager)
		master.FileManager = file.NewFileManager()
		master.StatManager = stat.NewStatManager()
		master.TosserManager = tosser.NewTosserManager(master.SetupManager)
	}
	return master
}
