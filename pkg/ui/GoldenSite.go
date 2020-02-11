package ui

import (
	"github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/file"
	"github.com/vit1251/golden/pkg/msg"
	"github.com/vit1251/golden/pkg/setup"
	"log"
)

type GoldenSite struct {
	WebSite *WebSite   /* Web site common type   */
}

func NewGoldenSite() (*GoldenSite) {
	site := new(GoldenSite)
	site.WebSite = NewWebSite()
	return site
}

func (self *GoldenSite) SetMessageManager(mm *msg.MessageManager) {
	self.WebSite.MessageManager = mm
}

func (self *GoldenSite) SetFileManager(manager *file.FileManager) {
	self.WebSite.FileAreaManager = manager
}

func (self *GoldenSite) SetAreaManager(am *area.AreaManager) {
	self.WebSite.AreaManager = am
}

func (self *GoldenSite) SetSetupManager(sm *setup.SetupManager) {
	self.WebSite.SetupManager = sm
}

func (self *GoldenSite) Start() (error) {

	log.Printf("Golden web service start")

	/* Register actions */
	self.WebSite.Register("/", NewWelcomeAction())
	self.WebSite.Register("/echo", NewAreaAction())
	self.WebSite.Register("/echo/{echoname:[A-Z0-9\\.\\-]+}", NewEchoAction())
	self.WebSite.Register("/echo/{echoname:[A-Z0-9\\.\\-]+}/update", NewEchoUpdateAction())
	self.WebSite.Register("/echo/{echoname:[A-Z0-9\\.\\-]+}/message/compose", NewEchoComposeAction())
	self.WebSite.Register("/echo/{echoname:[A-Z0-9\\.\\-]+}/message/compose/complete", NewEchoComposeCompleteAction())
	self.WebSite.Register("/echo/{echoname:[A-Z0-9\\.\\-]+}/message/{msgid:[A-Za-z0-9+]+}/view", NewViewAction())
	self.WebSite.Register("/echo/{echoname:[A-Z0-9\\.\\-]+}/message/{msgid:[A-Za-z0-9+]+}/reply", NewReplyAction())
	self.WebSite.Register("/echo/{echoname:[A-Z0-9\\.\\-]+}/message/{msgid:[A-Za-z0-9+]+}/reply/complete", NewReplyCompleteAction())
	self.WebSite.Register("/echo/{echoname:[A-Z0-9\\.\\-]+}/message/{msgid:[A-Za-z0-9+]+}/remove", NewRemoveAction())
	self.WebSite.Register("/echo/{echoname:[A-Z0-9\\.\\-]+}/message/{msgid:[A-Za-z0-9+]+}/remove/complete", NewRemoveCompleteAction())
	self.WebSite.Register("/file", NewFileAreaAction())
	self.WebSite.Register("/file/{echoname:[A-Z0-9\\.\\-]+}", NewFileAreaViewAction())
	self.WebSite.Register("/netmail", NewNetmailAction())
	self.WebSite.Register("/netmail/compose", NewNetmailComposeAction())
	self.WebSite.Register("/netmail/compose/complete", NewNetmailComposeCompleteAction())
	self.WebSite.Register("/stat", NewStatAction())
	self.WebSite.Register("/service", NewServiceManageAction())
	self.WebSite.Register("/service/complete", NewServiceManageCompleteAction())
	self.WebSite.Register("/setup", NewSetupAction())
	self.WebSite.Register("/setup/complete", NewSetupCompleteAction())
	//
	err := self.WebSite.Start()
	//
	return err
}

func (self *GoldenSite) Stop() (error) {

	log.Printf("Golden web service stop")

	//
//	webSite.Stop()

	return nil
}
