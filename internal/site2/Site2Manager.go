package site2

import (
	"github.com/gorilla/mux"
	api2 "github.com/vit1251/golden/internal/site2/api"
	"github.com/vit1251/golden/pkg/registry"
	"log"
        "fmt"
	"net/http"
)

type Route struct {
	pattern string       /* Regilar expression         */
	action  api2.IAction /* Processing callback        */
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

func (self *Site2Manager) ContainerMiddleware(a api2.IAction) api2.IAction {
	a.SetContainer(self.registry)
	return a
}

func Register(router *mux.Router, pattern string, a api2.IAction) {
	router.HandleFunc(pattern, a.ServeHTTP)
}

func (self *Site2Manager) createRouter() *mux.Router {

	router := mux.NewRouter()

        /* Section 1. Describe API patterns */
	Register(router, "/api/v1", self.ContainerMiddleware(api2.NewCommandStream()))
	Register(router, "/", self.ContainerMiddleware(NewIndexAction()))

        /* Section 2. Static directories */
	router.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(staticFileSystem())))
	router.PathPrefix("/public").Handler(http.StripPrefix("/public/", http.FileServer(publicFileSystem())))

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
