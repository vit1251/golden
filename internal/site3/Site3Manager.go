package site3

import (
	"fmt"
	"log"
	"net/http"

	api3 "github.com/vit1251/golden/internal/site3/api"
	"github.com/vit1251/golden/pkg/registry"
)

type Route struct {
	pattern string       /* Regilar expression         */
	action  api3.IAction /* Processing callback        */
}

type WebSite struct {
	routes []Route
}

type Site3Manager struct {
	registry *registry.Container
	port     int
	WebSite  *WebSite /* Web site common type   */
	addr     string
}

func NewSite3Manager(registry *registry.Container) *Site3Manager {

	site := new(Site3Manager)
	site.addr = "127.0.0.1"
	site.port = 8082
	site.registry = registry

	return site
}

func (s *Site3Manager) createRouter() *http.ServeMux {

	mux := http.NewServeMux()

	/* Section 1. Describe API patterns */
	mux.Handle("GET /{$}", NewIndexAction(s.registry))
	mux.Handle("GET /api/v1", api3.NewCommandStream(s.registry))

	/* Section 2. Static directories */
	mux.Handle("GET /static/{name}", http.StripPrefix("/static/", http.FileServer(staticFileSystem())))
	mux.Handle("GET /public/{name}", http.StripPrefix("/public/", http.FileServer(publicFileSystem())))

	return mux

}

func (s *Site3Manager) Start() {

	log.Printf("Site3Manager: Start HTTP service: addr = %s port = %d", s.addr, s.port)

	/* Start service */
	go s.run()

}

func (s *Site3Manager) run() {

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

func (s *Site3Manager) Stop() error {
	/* Step 1. Stop HTTP service */
	// TODO - stop HTTP service ...
	return nil
}

func (s *Site3Manager) SetPort(port int) {
	s.port = port
}

func (s *Site3Manager) GetLocation() string {
	return fmt.Sprintf("http://%s:%d\n", s.addr, s.port)
}
