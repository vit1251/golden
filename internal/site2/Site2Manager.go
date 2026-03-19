package site2

import (
	"fmt"
	"log"
	"net/http"

	api2 "github.com/vit1251/golden/internal/site2/api"
	"github.com/vit1251/golden/pkg/registry"
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

func (s *Site2Manager) createRouter() *http.ServeMux {

	mux := http.NewServeMux()

	/* Section 1. Describe API patterns */
	mux.Handle("GET /{$}", NewIndexAction(s.registry))
	mux.Handle("GET /api/v1", api2.NewCommandStream(s.registry))

	/* Section 2. Static directories */
	mux.Handle("GET /static/{name}", http.StripPrefix("/static/", http.FileServer(staticFileSystem())))
	mux.Handle("GET /public/{name}", http.StripPrefix("/public/", http.FileServer(publicFileSystem())))

	return mux

}

func (s *Site2Manager) Start() {

	log.Printf("Site2Manager: Start HTTP service: addr = %s port = %d", s.addr, s.port)

	/* Start service */
	go s.run()

}

func (s *Site2Manager) run() {

	/* Step 1. Create router */
	router := s.createRouter()

	/* Step 2. Start HTTP service */
	serviceAddr := fmt.Sprintf("%s:%d", s.addr, s.port)
	srv := &http.Server{
		Handler: router,
		Addr:    serviceAddr,
	}
	err1 := srv.ListenAndServe()
	if err1 != nil {
		panic(err1)
	}

}

func (s *Site2Manager) Stop() error {
	/* Step 1. Stop HTTP service */
	// TODO - stop HTTP service ...
	return nil
}

func (s *Site2Manager) SetPort(port int) {
	s.port = port
}

func (s *Site2Manager) GetLocation() string {
	return fmt.Sprintf("http://%s:%d\n", s.addr, s.port)
}
