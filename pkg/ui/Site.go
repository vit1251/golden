package ui

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/area"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/msg"
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
	routes        []Route
	rtr            *mux.Router
	AreaManager    *area.AreaManager
	SetupManager   *setup.SetupManager
	MessageManager *msg.MessageManager
	version         string
}

func (self *WebSite) SetVersion(Version string) {
	self.version = Version
}

func (self *WebSite) GetVersion() (string) {
	return self.version
}

func (self *WebSite) GetMessageManager() (*msg.MessageManager) {
	return self.MessageManager
}

func (self *WebSite) GetAreaManager() (*area.AreaManager) {
	return self.AreaManager
}

func (self *WebSite) GetSetupManager() (*setup.SetupManager) {
	return self.SetupManager
}

func NewWebSite() (*WebSite) {

	/* Create new onw web application */
	webSite := new(WebSite)

	/* Create router */
	rtr := mux.NewRouter()
	staticDir := "./static"
	rtr.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))
	webSite.rtr = rtr

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
	    WriteTimeout: 10 * time.Minute,
	    ReadTimeout:  10 * time.Minute,
	}
	err := srv.ListenAndServe()
	return err
}

func (self *WebSite) Stop() (error) {
	return nil
}
