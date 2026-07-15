package site

import (
	"fmt"
	"log"
	"net/http"

	"github.com/vit1251/golden/internal/site/handler"
	"github.com/vit1251/golden/pkg/registry"
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

func (s *SiteManager) createRouter() *http.ServeMux {

	mux := http.NewServeMux()

	// Шаг 1. Главная страница
	mux.Handle("GET /{$}", handler.NewWelcomeHandler(s.registry))

	// Шаг 2. Раздел конференций
	mux.Handle("GET /echo", handler.NewEchoAreaIndexHandler(s.registry))
	mux.Handle("GET /echo/create", handler.NewEchoAreaCreateHandler(s.registry))
	mux.Handle("GET /echo/create/complete", handler.NewEchoAreaCreateCompleteHandler(s.registry))
	mux.Handle("GET /echo/{echoname}", handler.NewEchoMsgIndexHandler(s.registry))
	mux.Handle("GET /echo/{echoname}/tree", handler.NewEchoMsgTreeHandler(s.registry))
	mux.Handle("GET /echo/{echoname}/remove", handler.NewEchoAreaRemoveHandler(s.registry))
	mux.Handle("GET /echo/{echoname}/remove/complete", handler.NewEchoRemoveCompleteHandler(s.registry))
	mux.Handle("GET /echo/{echoname}/purge", handler.NewEchoAreaPurgeHandler(s.registry))
	mux.Handle("GET /echo/{echoname}/purge/complete", handler.NewEchoAreaPurgeCompleteHandler(s.registry))
	mux.Handle("GET /echo/{echoname}/mark", handler.NewEchoAreaMarkHandler(s.registry))
	mux.Handle("GET /echo/{echoname}/mark/complete", handler.NewEchoAreaMarkCompleteHandler(s.registry))
	mux.Handle("GET /echo/{echoname}/update", handler.NewEchoAreaUpdateHandler(s.registry))
	mux.Handle("GET /echo/{echoname}/update/complete", handler.NewEchoAreaUpdateCompleteHandler(s.registry))
	mux.Handle("GET /echo/{echoname}/message/compose", handler.NewEchoMsgComposeHandler(s.registry))
	mux.Handle("GET /echo/{echoname}/message/{msgid}/view", handler.NewEchoMsgViewHandler(s.registry))
	mux.Handle("GET /echo/{echoname}/message/{msgid}/dump", handler.NewEchoMsgDumpHandler(s.registry))
	mux.Handle("GET /echo/{echoname}/message/{msgid}/twit", handler.NewEchoMsgTwitHandler(s.registry))
	mux.Handle("GET /echo/{echoname}/message/{msgid}/reply", handler.NewEchoMsgReplyHandler(s.registry))
	mux.Handle("GET /echo/{echoname}/message/{msgid}/remove", handler.NewEchoMsgRemoveHandler(s.registry))
	mux.Handle("GET /echo/{echoname}/message/{msgid}/remove/complete", handler.NewEchoMsgRemoveCompleteHandler(s.registry))

	/* File section */
	mux.Handle("GET /file", handler.NewFileEchoIndexHandler(s.registry))
	mux.Handle("GET /file/create", handler.NewFileEchoCreateHandler(s.registry))
	mux.Handle("GET /file/create/complete", handler.NewFileEchoCreateCompleteHandler(s.registry))
	mux.Handle("GET /file/{echoname}", handler.NewFileEchoAreaIndexHandler(s.registry))
	mux.Handle("GET /file/{echoname}/update", handler.NewFileEchoUpdateHandler(s.registry))
	mux.Handle("GET /file/{echoname}/remove", handler.NewFileEchoRemoveHandler(s.registry))
	mux.Handle("GET /file/{echoname}/remove/complete", handler.NewFileEchoRemoveCompleteHandler(s.registry))
	mux.Handle("GET /file/{echoname}/tic/{file}/view", handler.NewFileEchoAreaViewHandler(s.registry))
	mux.Handle("GET /file/{echoname}/tic/{file}/download", handler.NewFileEchoAreaDownloadHandler(s.registry))
	mux.Handle("GET /file/{echoname}/tic/{file}/remove", handler.NewFileEchoAreaRemoveHandler(s.registry))
	mux.Handle("GET /file/{echoname}/upload", handler.NewFileEchoAreaUploadHandler(s.registry))
	mux.Handle("GET /file/{echoname}/upload/complete", handler.NewFileEchoAreaUploadCompleteHandler(s.registry))

	/* Netmail section */
	mux.Handle("GET /netmail", handler.NewNetmailIndexHandler(s.registry))
	mux.Handle("GET /netmail/{msgid}/view", handler.NewNetmailViewHandler(s.registry))
	mux.Handle("GET /netmail/{msgid}/dump", handler.NewNetmailDumpHandler(s.registry))
	mux.Handle("GET /netmail/{msgid}/reply", handler.NewNetmailReplyHandler(s.registry))
	mux.Handle("GET /netmail/{msgid}/remove", handler.NewNetmailRemoveHandler(s.registry))
	mux.Handle("GET /netmail/{msgid}/attach/{attidx}/view", handler.NewNetmailAttachViewHandler(s.registry))
	mux.Handle("GET /netmail/compose", handler.NewNetmailComposeHandler(s.registry))

	/* Setup section */
	mux.Handle("GET /settings", handler.NewSettingsHandler(s.registry))
	mux.Handle("POST /settings/update", handler.NewSettingsUpdateHandler(s.registry))

	/* Service section */
	mux.Handle("GET /service", handler.NewServiceHandler(s.registry))

	/* Twit -> AddressBook */
	mux.Handle("GET /twit", handler.NewTwitIndexHandler(s.registry))
	mux.Handle("GET /twit/{twitid}/remove", handler.NewTwitRemoveCompleteHandler(s.registry))

	/* Draft section */
	mux.Handle("GET /draft", handler.NewDraftIndexHandler(s.registry))
	mux.Handle("GET /draft/{draftid}/edit", handler.NewDraftEditHandler(s.registry))
	mux.Handle("POST /draft/{draftid}/edit/complete", handler.NewDraftEditCompleteHandler(s.registry))

	/* Static section */
	mux.Handle("GET /assets/css/main.css", handler.NewStyleHandler())
	mux.Handle("GET /static/{filename}", handler.NewStaticHandler())

	/* Classic HTTP API */
	//mux.Handle("GET /api/netmail/remove", handler.NewNetmailRemoveApiHandler(s.registry))

	return mux
}

func (self *SiteManager) Start() {

	log.Printf("SiteManager: Start HTTP service: addr = %s port = %d", self.addr, self.port)

	/* Start service */
	go self.run()

}

func (self *SiteManager) run() error {
	srv := &http.Server{
		Handler: self.createRouter(),
		Addr:    fmt.Sprintf("%s:%d", self.addr, self.port),
	}
	return srv.ListenAndServe()
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
