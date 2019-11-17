package ui

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"log"
)

type WebSite struct {
	app *Application
}

type EchoAction struct {
	Action
}
type ViewAction struct {
	Action
}

func StartSite(app *Application) {
	//
	webSite := new(WebSite)
	webSite.app = app
	//
	router := httprouter.New()
	//
	welcomeAction := new(WelcomeAction)
	welcomeAction.Site = webSite
	router.GET("/", welcomeAction.ServeHTTP)
	//
	echoAction := new(EchoAction)
	echoAction.Site = webSite
	router.GET("/echo/:name", echoAction.ServeHTTP)
	//
	viewAction := new(ViewAction)
	viewAction.Site = webSite
	router.GET("/echo/:name/view/:msgid", viewAction.ServeHTTP)
	//
	composeAction := new(ComposeAction)
	composeAction.Site = webSite
	router.GET("/echo/:name/compose", composeAction.ServeHTTP)
	//
	composeCompleteAction := new(ComposeCompleteAction)
	composeCompleteAction.Site = webSite
	router.POST("/echo/:name/compose", composeCompleteAction.ServeHTTP)
	//
	//fs := http.FileServer(http.Dir("static"))
	//http.Handle("/static/", http.StripPrefix("/static/", fs))
	//
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", router))
}
