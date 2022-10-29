package site

import (
	"fmt"
	"github.com/gorilla/mux"
	action2 "github.com/vit1251/golden/internal/site/action"
	"github.com/vit1251/golden/pkg/registry"
	"log"
	"net/http"
)

type IAction interface {
	SetContainer(r *registry.Container)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Route struct {
	pattern string  /* Regilar expression         */
	action  IAction /* Processing callback        */
}

type WebSite struct {
	routes []Route
}

type SiteManager struct {
	registry *registry.Container
	port     int
	WebSite  *WebSite /* Web site common type   */
	addr     string
}

func NewSiteManager(registry *registry.Container) *SiteManager {

	site := new(SiteManager)
	site.addr = "127.0.0.1"
	site.port = 8080
	site.registry = registry

	return site
}

func (self *SiteManager) ContainerMiddleware(a IAction) IAction {
	a.SetContainer(self.registry)
	return a
}

func Register(router *mux.Router, pattern string, a IAction) {
	router.HandleFunc(pattern, a.ServeHTTP)
}

func (self *SiteManager) createRouter() *mux.Router {

	router := mux.NewRouter()

	/* Welcome */
	Register(router, "/", self.ContainerMiddleware(action2.NewWelcomeAction()))

	/* Echo section */
	Register(router, "/echo", self.ContainerMiddleware(action2.NewEchoAreaIndexAction()))
	Register(router, "/echo/create", self.ContainerMiddleware(action2.NewEchoAreaCreateAction()))
	Register(router, "/echo/create/complete", self.ContainerMiddleware(action2.NewEchoAreaCreateCompleteAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}", self.ContainerMiddleware(action2.NewEchoMsgIndexAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/tree", self.ContainerMiddleware(action2.NewEchoMsgTreeAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/remove", self.ContainerMiddleware(action2.NewEchoAreaRemoveAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/remove/complete", self.ContainerMiddleware(action2.NewEchoRemoveCompleteAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/purge", self.ContainerMiddleware(action2.NewEchoAreaPurgeAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/purge/complete", self.ContainerMiddleware(action2.NewEchoAreaPurgeCompleteAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/mark", self.ContainerMiddleware(action2.NewEchoAreaMarkAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/mark/complete", self.ContainerMiddleware(action2.NewEchoAreaMarkCompleteAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/update", self.ContainerMiddleware(action2.NewEchoAreaUpdateAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/update/complete", self.ContainerMiddleware(action2.NewEchoAreaUpdateCompleteAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/compose", self.ContainerMiddleware(action2.NewEchoMsgComposeAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/view", self.ContainerMiddleware(action2.NewEchoMsgViewAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/dump", self.ContainerMiddleware(action2.NewEchoMsgDumpAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/twit", self.ContainerMiddleware(action2.NewEchoMsgTwitAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/reply", self.ContainerMiddleware(action2.NewEchoMsgReplyAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/remove", self.ContainerMiddleware(action2.NewEchoMsgRemoveAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/remove/complete", self.ContainerMiddleware(action2.NewEchoMsgRemoveCompleteAction()))

	/* File section */
	Register(router, "/file", self.ContainerMiddleware(action2.NewFileEchoIndexAction()))
	Register(router, "/file/create", self.ContainerMiddleware(action2.NewFileEchoCreateAction()))
	Register(router, "/file/create/complete", self.ContainerMiddleware(action2.NewFileEchoCreateCompleteAction()))
	Register(router, "/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}", self.ContainerMiddleware(action2.NewFileEchoAreaIndexAction()))
	Register(router, "/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/update", self.ContainerMiddleware(action2.NewFileEchoUpdateAction()))
	Register(router, "/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/remove", self.ContainerMiddleware(action2.NewFileEchoRemoveAction()))
	Register(router, "/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/remove/complete", self.ContainerMiddleware(action2.NewFileEchoRemoveCompleteAction()))
	Register(router, "/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/tic/{file:[A-Za-z0-9\\.\\-\\_]+}/view", self.ContainerMiddleware(action2.NewFileEchoAreaViewAction()))
	Register(router, "/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/tic/{file:[A-Za-z0-9\\.\\-\\_]+}/download", self.ContainerMiddleware(action2.NewFileEchoAreaDownloadAction()))
	Register(router, "/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/tic/{file:[A-Za-z0-9\\.\\-\\_]+}/remove", self.ContainerMiddleware(action2.NewFileEchoAreaRemoveAction()))
	Register(router, "/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/upload", self.ContainerMiddleware(action2.NewFileEchoAreaUploadAction()))
	Register(router, "/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/upload/complete", self.ContainerMiddleware(action2.NewFileEchoAreaUploadCompleteAction()))

	/* Netmail section */
	Register(router, "/netmail", self.ContainerMiddleware(action2.NewNetmailIndexAction()))
	Register(router, "/netmail/{msgid:[A-Za-z0-9+]+}/view", self.ContainerMiddleware(action2.NewNetmailViewAction()))
	Register(router, "/netmail/{msgid:[A-Za-z0-9+]+}/dump", self.ContainerMiddleware(action2.NewNetmailDumpAction()))
	Register(router, "/netmail/{msgid:[A-Za-z0-9+]+}/reply", self.ContainerMiddleware(action2.NewNetmailReplyAction()))
	Register(router, "/netmail/{msgid:[A-Za-z0-9+]+}/remove", self.ContainerMiddleware(action2.NewNetmailRemoveAction()))
	Register(router, "/netmail/{msgid:[A-Za-z0-9+]+}/attach/{attidx:[0-9]+}/view", self.ContainerMiddleware(action2.NewNetmailAttachViewAction()))
	Register(router, "/netmail/compose", self.ContainerMiddleware(action2.NewNetmailComposeAction()))

	/* Setup section */
	Register(router, "/setup", self.ContainerMiddleware(action2.NewSetupAction()))
	Register(router, "/setup/complete", self.ContainerMiddleware(action2.NewSetupCompleteAction()))

	/* Service section */
	Register(router, "/service", self.ContainerMiddleware(action2.NewServiceAction()))
	Register(router, "/service/mailer/stat", self.ContainerMiddleware(action2.NewServiceMailerAction()))
	Register(router, "/service/mailer/event", self.ContainerMiddleware(action2.NewServiceMailerEventAction()))
	Register(router, "/service/toss/stat", self.ContainerMiddleware(action2.NewServiceTossAction()))
	Register(router, "/service/toss/event", self.ContainerMiddleware(action2.NewServiceTossEventAction()))
	Register(router, "/service/tracker/stat", self.ContainerMiddleware(action2.NewServiceTrackerAction()))
	Register(router, "/service/tracker/event", self.ContainerMiddleware(action2.NewServiceTrackerEventAction()))

	/* Twit -> AddressBook */
	Register(router, "/twit", self.ContainerMiddleware(action2.NewTwitIndexAction()))
	Register(router, "/twit/{twitid:[A-Za-z0-9+]+}/remove", self.ContainerMiddleware(action2.NewTwitRemoveCompleteAction()))

	/* Draft section */
	Register(router, "/draft", self.ContainerMiddleware(action2.NewDraftIndexAction()))
	Register(router, "/draft/{draftid:[A-Za-z0-9\\-+]+}/edit", self.ContainerMiddleware(action2.NewDraftEditAction()))
	Register(router, "/draft/{draftid:[A-Za-z0-9\\-+]+}/edit/complete", self.ContainerMiddleware(action2.NewDraftEditCompleteAction()))

	/* Static section */
	Register(router, "/assets/css/main.css", self.ContainerMiddleware(action2.NewStyleAction()))
	Register(router, "/static/{name:[A-Za-z0-9\\.\\_\\-]+}", self.ContainerMiddleware(NewStaticAction()))

	/* Classic HTTP API */
	Register(router, "/api/netmail/remove", self.ContainerMiddleware(action2.NewNetmailRemoveApiAction()))

	return router

}

func (self *SiteManager) Start() {

	log.Printf("SiteManager: Start HTTP service: addr = %s port = %d", self.addr, self.port)

	/* Start service */
	go self.run()

}

func (self *SiteManager) run() {

	/* Step 1. Create router */
	router := self.createRouter()

	/* Step 2. Start HTTP service */
	serviceAddr := fmt.Sprintf("%s:%d", self.addr, self.port)
	srv := &http.Server{
		Handler: router,
		Addr:    serviceAddr,
	}
	err1 := srv.ListenAndServe()
	if err1 != nil {
		panic(err1)
	}

}

func (self *SiteManager) Stop() error {
	/* Step 1. Stop HTTP service */
	// TODO - stop HTTP service ...
	return nil
}

func (self *SiteManager) SetPort(port int) {
	self.port = port
}

func (self *SiteManager) GetLocation() string {
	return fmt.Sprintf("http://%s:%d\n", self.addr, self.port)
}
