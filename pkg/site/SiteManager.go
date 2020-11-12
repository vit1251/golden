package site

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/site/action"
	"log"
	"net/http"
	"strings"
	"time"
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
	rtr    *mux.Router
}

type SiteManager struct {
	registry *registry.Container
	port      int
	WebSite   *WebSite /* Web site common type   */
	rtr       *mux.Router
	addr      string
}

func NewSiteManager(registry *registry.Container) *SiteManager {

	site := new(SiteManager)
	site.addr = "127.0.0.1"
	site.port = 8080
	site.registry = registry
	site.rtr = mux.NewRouter()

	return site
}

func (self *SiteManager) Register(pattern string, a IAction) {

	/* Register owner */
	a.SetContainer(self.registry)

	/* Register */
	actionFunc := a.ServeHTTP
	self.rtr.HandleFunc(pattern, actionFunc)

	/* Create router */
	//	r := Route{}
	//	r.pattern = pattern
	//	r.action = a
	//	self.routes = append(self.routes, r)

}

func (self *SiteManager) registerFrontend() {
	self.Register("/", action.NewWelcomeAction())
	self.Register("/echo", action.NewEchoAreaIndexAction())
	self.Register("/echo/create", action.NewEchoAreaCreateAction())
	self.Register("/echo/create/complete", action.NewEchoAreaCreateCompleteAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}", action.NewEchoMsgIndexAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/tree", action.NewEchoMsgTreeAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/remove", action.NewEchoAreaRemoveAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/remove/complete", action.NewEchoRemoveCompleteAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/purge", action.NewEchoAreaPurgeAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/purge/complete", action.NewEchoAreaPurgeCompleteAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/update", action.NewEchoAreaUpdateAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/update/complete", action.NewEchoAreaUpdateCompleteAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/compose", action.NewEchoMsgComposeAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/view", action.NewEchoMsgViewAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/dump", action.NewEchoMsgDumpAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/twit", action.NewEchoMsgTwitAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/reply", action.NewEchoMsgReplyAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/remove", action.NewEchoMsgRemoveAction())
	self.Register("/echo/{echoname:[A-Za-z0-9\\.\\-\\_]+}/message/{msgid:[A-Za-z0-9+]+}/remove/complete", action.NewEchoMsgRemoveCompleteAction())
	self.Register("/file", action.NewFileEchoIndexAction())
	self.Register("/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}", action.NewFileEchoAreaIndexAction())
	self.Register("/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/update", action.NewFileEchoUpdateAction())
	self.Register("/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/remove", action.NewFileEchoRemoveAction())
	self.Register("/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/remove/complete", action.NewFileEchoRemoveCompleteAction())
	self.Register("/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/tic/{file:[A-Za-z0-9\\.\\-\\_]+}/view", action.NewFileEchoAreaDownloadAction())
	self.Register("/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/tic/{file:[A-Za-z0-9\\.\\-\\_]+}/remove", action.NewFileEchoAreaRemoveAction())
	self.Register("/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/upload", action.NewFileEchoAreaUploadAction())
	self.Register("/file/{echoname:[A-Za-z0-9\\.\\-\\_]+}/upload/complete", action.NewFileEchoAreaUploadCompleteAction())
	self.Register("/netmail", action.NewNetmailIndexAction())
	self.Register("/netmail/{msgid:[A-Za-z0-9+]+}/view", action.NewNetmailViewAction())
	self.Register("/netmail/{msgid:[A-Za-z0-9+]+}/reply", action.NewNetmailReplyAction())
	self.Register("/netmail/{msgid:[A-Za-z0-9+]+}/remove", action.NewNetmailRemoveAction())
	self.Register("/netmail/compose", action.NewNetmailComposeAction())
	self.Register("/stat", action.NewStatAction())
	self.Register("/setup", action.NewSetupAction())
	self.Register("/setup/complete", action.NewSetupCompleteAction())
	self.Register("/assets/css/main.css", action.NewStyleAction())
	self.Register("/service", action.NewServiceAction())
	self.Register("/service/{name:[A-Za-z0-9\\.\\_\\-]+}/event", action.NewServiceEventAction())
	self.Register("/static/{name:[A-Za-z0-9\\.\\_\\-]+}", action.NewStaticAction())
	self.Register("/twit", action.NewTwitIndexAction())
	self.Register("/twit/{twitid:[A-Za-z0-9+]+}/remove", action.NewTwitRemoveCompleteAction())
	self.Register("/draft", action.NewDraftIndexAction())
	self.Register("/draft/{draftid:[A-Za-z0-9\\-+]+}/edit", action.NewDraftEditAction())
	self.Register("/draft/{draftid:[A-Za-z0-9\\-+]+}/edit/complete", action.NewDraftEditCompleteAction())
}

func (self *SiteManager) registerBackend() {
	self.Register("/api/stat", action.NewStatApiAction())
	self.Register("/api/netmail/remove", action.NewNetmailRemoveApiAction())
}

func (self *SiteManager) Start() {

	log.Printf("SiteManager: Start HTTP service: addr = %s port = %d", self.addr, self.port)

	/* Prepare routes  */
	self.registerFrontend()
	self.registerBackend()

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

	serviceAddr := fmt.Sprintf("%s:%d", self.addr, self.port)
	srv := &http.Server{
		Handler: self.rtr,
		Addr:    serviceAddr,
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 10 * time.Minute,
		ReadTimeout:  10 * time.Minute,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Printf("SiteManager: service error: err = %+v", err)
	}

}

func (self *SiteManager) Stop() error {

	log.Printf("SiteManager: Service stop")

	//
	//	webSite.Stop()

	return nil
}

func (self *SiteManager) SetPort(port int) {
	self.port = port
}
