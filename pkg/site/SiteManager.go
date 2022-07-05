package site

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/site/action"
	"github.com/vit1251/golden/pkg/site/action/api"
	"log"
	"net/http"
	"strings"
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
	Register(router, "/", self.ContainerMiddleware(action.NewWelcomeAction()))

	/* Echo section */
	Register(router, "/echo", self.ContainerMiddleware(action.NewEchoAreaIndexAction()))
	Register(router, "/echo/create", self.ContainerMiddleware(action.NewEchoAreaCreateAction()))
	Register(router, "/echo/create/complete", self.ContainerMiddleware(action.NewEchoAreaCreateCompleteAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}", self.ContainerMiddleware(action.NewEchoMsgIndexAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/tree", self.ContainerMiddleware(action.NewEchoMsgTreeAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/remove", self.ContainerMiddleware(action.NewEchoAreaRemoveAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/remove/complete", self.ContainerMiddleware(action.NewEchoRemoveCompleteAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/purge", self.ContainerMiddleware(action.NewEchoAreaPurgeAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/purge/complete", self.ContainerMiddleware(action.NewEchoAreaPurgeCompleteAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/mark", self.ContainerMiddleware(action.NewEchoAreaMarkAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/mark/complete", self.ContainerMiddleware(action.NewEchoAreaMarkCompleteAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/update", self.ContainerMiddleware(action.NewEchoAreaUpdateAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/update/complete", self.ContainerMiddleware(action.NewEchoAreaUpdateCompleteAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/compose", self.ContainerMiddleware(action.NewEchoMsgComposeAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/view", self.ContainerMiddleware(action.NewEchoMsgViewAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/dump", self.ContainerMiddleware(action.NewEchoMsgDumpAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/twit", self.ContainerMiddleware(action.NewEchoMsgTwitAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/reply", self.ContainerMiddleware(action.NewEchoMsgReplyAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/remove", self.ContainerMiddleware(action.NewEchoMsgRemoveAction()))
	Register(router, "/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/remove/complete", self.ContainerMiddleware(action.NewEchoMsgRemoveCompleteAction()))

	/* File section */
	Register(router, "/file", self.ContainerMiddleware(action.NewFileEchoIndexAction()))
	Register(router, "/file/create", self.ContainerMiddleware(action.NewFileEchoCreateAction()))
	Register(router, "/file/create/complete", self.ContainerMiddleware(action.NewFileEchoCreateCompleteAction()))
	Register(router, "/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}", self.ContainerMiddleware(action.NewFileEchoAreaIndexAction()))
	Register(router, "/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/update", self.ContainerMiddleware(action.NewFileEchoUpdateAction()))
	Register(router, "/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/remove", self.ContainerMiddleware(action.NewFileEchoRemoveAction()))
	Register(router, "/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/remove/complete", self.ContainerMiddleware(action.NewFileEchoRemoveCompleteAction()))
	Register(router, "/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/tic/{file:[A-Za-z0-9\\.\\-\\_]+}/view", self.ContainerMiddleware(action.NewFileEchoAreaViewAction()))
	Register(router, "/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/tic/{file:[A-Za-z0-9\\.\\-\\_]+}/download", self.ContainerMiddleware(action.NewFileEchoAreaDownloadAction()))
	Register(router, "/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/tic/{file:[A-Za-z0-9\\.\\-\\_]+}/remove", self.ContainerMiddleware(action.NewFileEchoAreaRemoveAction()))
	Register(router, "/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/upload", self.ContainerMiddleware(action.NewFileEchoAreaUploadAction()))
	Register(router, "/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/upload/complete", self.ContainerMiddleware(action.NewFileEchoAreaUploadCompleteAction()))

	/* Netmail section */
	Register(router, "/netmail", self.ContainerMiddleware(action.NewNetmailIndexAction()))
	Register(router, "/netmail/{msgid:[A-Za-z0-9+]+}/view", self.ContainerMiddleware(action.NewNetmailViewAction()))
	Register(router, "/netmail/{msgid:[A-Za-z0-9+]+}/dump", self.ContainerMiddleware(action.NewNetmailDumpAction()))
	Register(router, "/netmail/{msgid:[A-Za-z0-9+]+}/reply", self.ContainerMiddleware(action.NewNetmailReplyAction()))
	Register(router, "/netmail/{msgid:[A-Za-z0-9+]+}/remove", self.ContainerMiddleware(action.NewNetmailRemoveAction()))
	Register(router, "/netmail/{msgid:[A-Za-z0-9+]+}/attach/{attidx:[0-9]+}/view", self.ContainerMiddleware(action.NewNetmailAttachViewAction()))
	Register(router, "/netmail/compose", self.ContainerMiddleware(action.NewNetmailComposeAction()))

	/* Setup section */
	Register(router, "/setup", self.ContainerMiddleware(action.NewSetupAction()))
	Register(router, "/setup/complete", self.ContainerMiddleware(action.NewSetupCompleteAction()))

	/* Service section */
	Register(router, "/service", self.ContainerMiddleware(action.NewServiceAction()))
	Register(router, "/service/mailer/stat", self.ContainerMiddleware(action.NewServiceMailerAction()))
	Register(router, "/service/mailer/event", self.ContainerMiddleware(action.NewServiceMailerEventAction()))
	Register(router, "/service/toss/stat", self.ContainerMiddleware(action.NewServiceTossAction()))
	Register(router, "/service/toss/event", self.ContainerMiddleware(action.NewServiceTossEventAction()))
	Register(router, "/service/tracker/stat", self.ContainerMiddleware(action.NewServiceTrackerAction()))
	Register(router, "/service/tracker/event", self.ContainerMiddleware(action.NewServiceTrackerEventAction()))

	/* Twit -> AddressBook */
	Register(router, "/twit", self.ContainerMiddleware(action.NewTwitIndexAction()))
	Register(router, "/twit/{twitid:[A-Za-z0-9+]+}/remove", self.ContainerMiddleware(action.NewTwitRemoveCompleteAction()))

	/* Draft section */
	Register(router, "/draft", self.ContainerMiddleware(action.NewDraftIndexAction()))
	Register(router, "/draft/{draftid:[A-Za-z0-9\\-+]+}/edit", self.ContainerMiddleware(action.NewDraftEditAction()))
	Register(router, "/draft/{draftid:[A-Za-z0-9\\-+]+}/edit/complete", self.ContainerMiddleware(action.NewDraftEditCompleteAction()))

	/* Static section */
	Register(router, "/assets/css/main.css", self.ContainerMiddleware(action.NewStyleAction()))
	Register(router, "/static/{name:[A-Za-z0-9\\.\\_\\-]+}", self.ContainerMiddleware(action.NewStaticAction()))

	/* Classic HTTP API */
	Register(router, "/api/netmail/remove", self.ContainerMiddleware(action.NewNetmailRemoveApiAction()))

	/* Modern Web-Socket command stream */
	Register(router, "/api/v1", self.ContainerMiddleware(api.NewCommandStream()))

	return router

}

func (self *SiteManager) Start() {

	log.Printf("SiteManager: Start HTTP service: addr = %s port = %d", self.addr, self.port)

	/* Start service */
	go self.run()

	/* Report */
	var report strings.Builder

	report.WriteString("Golden Point is running at:\n")
	report.WriteString("\n")
	report.WriteString(fmt.Sprintf("    http://%s:%d\n", self.addr, self.port))
	report.WriteString("\n")
	report.WriteString("Note: You MUST setup your instalattion on first run.\n")
	report.WriteString("      Please open `Setup` section initially.\n")

	fmt.Printf("%s", report.String())

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
