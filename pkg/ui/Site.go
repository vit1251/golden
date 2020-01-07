package ui

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
)

type IAction interface {
	SetSite(webSite *WebSite)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type Route struct {
	pattern  string        /* Regilar expression         */
	action   IAction       /* Processing callback        */
}

type WebSite struct {
	app      *Application
	routes  []Route
	rtr      *mux.Router
}

type EchoAction struct {
	Action
}
type ViewAction struct {
	Action
}

func NewWebSite(app *Application) (*WebSite) {
	webSite := new(WebSite)
	webSite.app = app
	webSite.rtr = mux.NewRouter()
	return webSite
}

func (self *WebSite) Register(pattern string, a IAction) {

	/* Register owner */
	a.SetSite(self)

	/* Register */
	actionFunc := a.ServeHTTP
	self.rtr.HandleFunc(pattern, actionFunc)

	/* Create router */
//	r := Route{}
//	r.pattern = pattern
//	r.action = a
//	self.routes = append(self.routes, r)

}

const INTERFACE = "127.0.0.1:8080"

func (self *WebSite) Start() (error) {
	srv := &http.Server{
	    Handler:      self.rtr,
	    Addr:         INTERFACE,
	    // Good practice: enforce timeouts for servers you create!
	    WriteTimeout: 15 * time.Second,
	    ReadTimeout:  15 * time.Second,
	}
	err := srv.ListenAndServe()
	return err
}

func (self *WebSite) Stop() (error) {
	return nil
}

func (self *Application) StartSite() (error) {
	//
	webSite := NewWebSite(self)
	//
	webSite.Register("/", new(WelcomeAction))
	webSite.Register("/echo/{echoname:[A-Z0-9\\.]+}", new(EchoAction))
	webSite.Register("/echo/{echoname:[A-Z0-9\\.]+}/message/compose", new(ComposeAction))
	webSite.Register("/echo/{echoname:[A-Z0-9\\.]+}/message/compose/complete", new(ComposeCompleteAction))
	webSite.Register("/echo/{echoname:[A-Z0-9\\.]+}/message/{msgid:[A-Za-z0-9+]+}/view", new(ViewAction))
	webSite.Register("/echo/{echoname:[A-Z0-9\\.]+}/message/{msgid:[A-Za-z0-9+]+}/reply", new(ReplyAction))
	webSite.Register("/echo/{echoname:[A-Z0-9\\.]+}/message/{msgid:[A-Za-z0-9+]+}/remove", new(RemoveAction))
	webSite.Register("/echo/{echoname:[A-Z0-9\\.]+}/message/{msgid:[A-Za-z0-9+]+}/remove/complete", new(RemoveCompleteAction))
	//
	err := webSite.Start()
	//
	return err
}

func (app *Application) StopSite() (error) {
	//
//	webSite.Stop()
	//
	return nil
}
