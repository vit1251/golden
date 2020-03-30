package ui

import (
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/ui/api"
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
	pattern  string        /* Regilar expression         */
	action   IAction       /* Processing callback        */
}

type WebSite struct {
	routes          []Route
	rtr             *mux.Router
}

type GoldenSite struct {
	Container *dig.Container
	WebSite   *WebSite         /* Web site common type   */
	rtr       *mux.Router
}

func NewGoldenSite(c *dig.Container) *GoldenSite {

	site := new(GoldenSite)
	site.Container = c

	/* Create router */
	rtr := mux.NewRouter()
	staticDir := "./static"
	rtr.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	site.rtr = rtr

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

func (self *GoldenSite) Start() (error) {

	log.Printf("Golden web service start")

	/* Register actions */
	self.Register("/", NewWelcomeAction())
	self.Register("/echo", NewEchoIndexAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}", NewEchoMsgIndexAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/update", NewEchoUpdateAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/update/complete", NewEchoUpdateCompleteAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/compose", NewEchoComposeAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/compose/complete", NewEchoComposeCompleteAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/view", NewEchoViewAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/reply", NewEchoReplyAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/reply/complete", NewReplyCompleteAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/remove", NewRemoveAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/remove/complete", NewRemoveCompleteAction())
	self.Register("/file", NewFileAreaAction())
	self.Register("/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}", NewFileAreaViewAction())
	self.Register("/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/tic/{file:[A-Za-z0-9\\.\\-\\_]+}/view", NewFileAreaDownloadAction())
	self.Register("/netmail", NewNetmailAction())
	self.Register("/netmail/compose", NewNetmailComposeAction())
	self.Register("/netmail/compose/complete", NewNetmailComposeCompleteAction())
	self.Register("/stat", NewStatAction())
	self.Register("/stat/image", api.NewAPIStatAction())
	self.Register("/service", NewServiceManageAction())
	self.Register("/api/service/start", NewApiServiceStartAction())
	self.Register("/setup", NewSetupAction())
	self.Register("/setup/complete", NewSetupCompleteAction())

	//
	INTERFACE := "127.0.0.1:8080"

	srv := &http.Server{
		Handler:      self.rtr,
		Addr:         INTERFACE,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 10 * time.Minute,
		ReadTimeout:  10 * time.Minute,
	}
	err := srv.ListenAndServe()

	return err
}

func (self *GoldenSite) Stop() (error) {

	log.Printf("Golden web service stop")

	//
//	webSite.Stop()

	return nil
}
