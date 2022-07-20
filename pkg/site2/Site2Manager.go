package site2

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/site2/api"
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

type Site2Manager struct {
	registry *registry.Container
	port     int
	WebSite  *WebSite /* Web site common type   */
	addr     string
}

func NewSite2Manager(registry *registry.Container) *Site2Manager {

	site := new(Site2Manager)
	site.addr = "127.0.0.1"
	site.port = 8081
	site.registry = registry

	return site
}

func (self *Site2Manager) ContainerMiddleware(a IAction) IAction {
	a.SetContainer(self.registry)
	return a
}

func Register(router *mux.Router, pattern string, a IAction) {
	router.HandleFunc(pattern, a.ServeHTTP)
}

func (self *Site2Manager) createRouter() *mux.Router {

	router := mux.NewRouter()

	Register(router, "/api/v1", self.ContainerMiddleware(api.NewCommandStream()))

	return router

}

func (self *Site2Manager) Start() {

	log.Printf("Site2Manager: Start HTTP service: addr = %s port = %d", self.addr, self.port)

	/* Start service */
	go self.run()

}

func (self *Site2Manager) run() {

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

func (self *Site2Manager) Stop() error {
	/* Step 1. Stop HTTP service */
	// TODO - stop HTTP service ...
	return nil
}

func (self *Site2Manager) SetPort(port int) {
	self.port = port
}

func (self *Site2Manager) GetLocation() string {
        return fmt.Sprintf("http://%s:%d\n", self.addr, self.port)
}
