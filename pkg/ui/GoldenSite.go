package ui

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/ui/action"
	"github.com/vit1251/golden/pkg/ui/action/api"
	"go.uber.org/dig"
	"log"
	"net/http"
	"time"
)

type IAction interface {
	SetContainer(c *dig.Container)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Route struct {
	pattern string  /* Regilar expression         */
	action  IAction /* Processing callback        */
}

type WebSite struct {
	routes []Route
	rtr    *mux.Router
}

type GoldenSite struct {
	Container *dig.Container
	port      int
	WebSite   *WebSite /* Web site common type   */
	rtr       *mux.Router
	addr      string
}

func NewService(c *dig.Container) *GoldenSite {

	site := new(GoldenSite)
	site.addr = "127.0.0.1"
	site.port = 8080
	site.Container = c
	site.rtr = mux.NewRouter()

	return site
}

func (self *GoldenSite) Register(pattern string, a IAction) {

	/* Register owner */
	a.SetContainer(self.Container)

	/* Register */
	actionFunc := a.ServeHTTP
	self.rtr.HandleFunc(pattern, actionFunc)

	/* Create router */
	//	r := Route{}
	//	r.pattern = pattern
	//	r.action = a
	//	self.routes = append(self.routes, r)

}

func (self *GoldenSite) registerFrontend() {
	self.Register("/", action.NewWelcomeAction())
	self.Register("/echo", action.NewEchoIndexAction())
	self.Register("/echo/create", action.NewEchoCreateAction())
	self.Register("/echo/create/complete", action.NewEchoCreateCompleteAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}", action.NewEchoMsgIndexAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/remove", action.NewEchoRemoveAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/remove/complete", action.NewEchoRemoveCompleteAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/purge", action.NewEchoPurgeAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/purge/complete", action.NewEchoPurgeCompleteAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/update", action.NewEchoUpdateAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/update/complete", action.NewEchoUpdateCompleteAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/compose", action.NewEchoComposeAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/compose/complete", action.NewEchoComposeCompleteAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/view", action.NewEchoViewAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/reply", action.NewEchoReplyAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/reply/complete", action.NewReplyCompleteAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/remove", action.NewEchoMsgRemoveAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/remove/complete", action.NewEchoMsgRemoveCompleteAction())
	self.Register("/file", action.NewFileAreaIndexAction())
	self.Register("/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}", action.NewFileAreaViewAction())
	self.Register("/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/tic/{file:[A-Za-z0-9\\.\\-\\_]+}/view", action.NewFileAreaDownloadAction())
	self.Register("/netmail", action.NewNetmailAction())
	self.Register("/netmail/{msgid:[A-Za-z0-9+]+}/view", action.NewNetmailViewAction())
	self.Register("/netmail/{msgid:[A-Za-z0-9+]+}/reply", action.NewNetmailReplyAction())
	self.Register("/netmail/{msgid:[A-Za-z0-9+]+}/remove", action.NewNetmailRemoveAction())
	self.Register("/netmail/compose", action.NewNetmailComposeAction())
	self.Register("/netmail/compose/complete", action.NewNetmailComposeCompleteAction())
	self.Register("/stat", action.NewStatAction())
	self.Register("/setup", action.NewSetupAction())
	self.Register("/setup/complete", action.NewSetupCompleteAction())
	self.Register("/help", action.NewHelpAction())
	self.Register("/static/{name:[A-Za-z0-9\\.\\_\\-]+}", action.NewStaticAction())
}

func (self *GoldenSite) registerBackend() {
	self.Register("/api/stat", api.NewStatAction())
	self.Register("/api/echo/create", api.NewEchoCreateAction())
	self.Register("/api/netmail/remove", api.NewNetmailRemoveAction())
}

func (self *GoldenSite) Start() error {
	self.registerFrontend()
	self.registerBackend()
	serviceAddr := fmt.Sprintf("%s:%d", self.addr, self.port)
	srv := &http.Server{
		Handler: self.rtr,
		Addr:    serviceAddr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 10 * time.Minute,
		ReadTimeout:  10 * time.Minute,
	}
	err := srv.ListenAndServe()
	return err
}

func (self *GoldenSite) Stop() error {

	log.Printf("Golden web service stop")

	//
	//	webSite.Stop()

	return nil
}

func (self *GoldenSite) SetPort(port int) {
	self.port = port
}
