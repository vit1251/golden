package common

import (
	"github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/charset"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/netmail"
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
	NetmailManager *netmail.NetmailManager
	CharsetManager *charset.CharsetManager
}

var master *GoldenMaster

func GetMaster() *GoldenMaster {
	if master == nil {
		master = new(GoldenMaster)
	}
	return master
}
