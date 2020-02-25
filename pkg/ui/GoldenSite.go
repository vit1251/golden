package ui

import (
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

func (self *GoldenSite) Start() (error) {

	log.Printf("Golden web service start")

	/* Register actions */
	self.WebSite.Register("/", NewWelcomeAction())
	self.WebSite.Register("/echo", NewAreaAction())
	self.WebSite.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}", NewEchoAction())
	self.WebSite.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/update", NewEchoUpdateAction())
	self.WebSite.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/compose", NewEchoComposeAction())
	self.WebSite.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/compose/complete", NewEchoComposeCompleteAction())
	self.WebSite.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/view", NewEchoViewAction())
	self.WebSite.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/reply", NewEchoReplyAction())
	self.WebSite.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/reply/complete", NewReplyCompleteAction())
	self.WebSite.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/remove", NewRemoveAction())
	self.WebSite.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/remove/complete", NewRemoveCompleteAction())
	self.WebSite.Register("/file", NewFileAreaAction())
	self.WebSite.Register("/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}", NewFileAreaViewAction())
	self.WebSite.Register("/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/tic/{file:[A-Za-z0-9\\.\\-\\_]+}/view", NewFileAreaDownloadAction())
	self.WebSite.Register("/netmail", NewNetmailAction())
	self.WebSite.Register("/netmail/compose", NewNetmailComposeAction())
	self.WebSite.Register("/netmail/compose/complete", NewNetmailComposeCompleteAction())
	self.WebSite.Register("/stat", NewStatAction())
	self.WebSite.Register("/service", NewServiceManageAction())
	self.WebSite.Register("/api/service/start", NewApiServiceStartAction())
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
