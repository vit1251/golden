package ui

import(
	"github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/msg"
)

type GoldenSite struct {
	WebSite *WebSite
}

func NewGoldenSite() (*GoldenSite) {
	site := new(GoldenSite)
	site.WebSite = NewWebSite()
	return site
}

func (self *GoldenSite) SetMessageManager(mm *msg.MessageManager) {
	self.WebSite.MessageManager = mm
}

func (self *GoldenSite) SetAreaManager(am *area.AreaManager) {
	self.WebSite.AreaManager = am
}

func (self *GoldenSite) SetSetupManager(sm *setup.SetupManager) {
	self.WebSite.SetupManager = sm
}

func (self *GoldenSite) Start() (error) {
	//
	self.WebSite.Register("/", new(WelcomeAction))
	self.WebSite.Register("/echo/{echoname:[A-Z0-9\\.\\-]+}", new(EchoAction))
	self.WebSite.Register("/echo/{echoname:[A-Z0-9\\.\\-]+}/message/compose", new(ComposeAction))
	self.WebSite.Register("/echo/{echoname:[A-Z0-9\\.\\-]+}/message/compose/complete", new(ComposeCompleteAction))
	self.WebSite.Register("/echo/{echoname:[A-Z0-9\\.\\-]+}/message/{msgid:[A-Za-z0-9+]+}/view", new(ViewAction))
	self.WebSite.Register("/echo/{echoname:[A-Z0-9\\.\\-]+}/message/{msgid:[A-Za-z0-9+]+}/reply", new(ReplyAction))
	self.WebSite.Register("/echo/{echoname:[A-Z0-9\\.\\-]+}/message/{msgid:[A-Za-z0-9+]+}/reply/complete", new(ReplyCompleteAction))
	self.WebSite.Register("/echo/{echoname:[A-Z0-9\\.\\-]+}/message/{msgid:[A-Za-z0-9+]+}/remove", new(RemoveAction))
	self.WebSite.Register("/echo/{echoname:[A-Z0-9\\.\\-]+}/message/{msgid:[A-Za-z0-9+]+}/remove/complete", new(RemoveCompleteAction))
	self.WebSite.Register("/setup", new(SetupAction))
	self.WebSite.Register("/setup/complete", new(SetupCompleteAction))
	//
	err := self.WebSite.Start()
	//
	return err
}

func (self *GoldenSite) Stop() (error) {
	//
//	webSite.Stop()
	//
	return nil
}
