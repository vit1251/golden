package ui

import (
	"encoding/json"
	"fmt"
	"github.com/vit1251/golden/pkg/mailer"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/tosser"
	"log"
	"net/http"
)

type ApiServiceStartAction struct {
	Action
}

func NewApiServiceStartAction() *ApiServiceStartAction {
	smac := new(ApiServiceStartAction)
	return smac
}

func (self *ApiServiceStartAction) Start() {
	go self.Run()
}

func (self *ApiServiceStartAction) Run() error {

	var setupManager *setup.SetupManager
	self.Container.Invoke(func(sm *setup.SetupManager) {
		setupManager = sm
	})

	/* Construct node address */
	netAddr, err1 := setupManager.Get("main", "NetAddr", "")
	if err1 != nil {
		return err1
	}
	password, err2 := setupManager.Get("main", "Password", "")
	if err2 != nil {
		return err2
	}
	address, err3 := setupManager.Get("main", "Address", "")
	if err3 != nil {
		return err3
	}
	inb, err4 := setupManager.Get("main", "Inbound", "")
	if err4 != nil {
		return err4
	}
	outb, err5 := setupManager.Get("main", "Outbound", "")
	if err5 != nil {
		return err5
	}
	TempOutbound, err6 := setupManager.Get("main", "TempOutbound", "")
	if err6 != nil {
		return err6
	}

	/**/
	newAddress := fmt.Sprintf("%s@fidonet", address)

	/* Get parameters */
	m := mailer.NewMailer(setupManager)
	m.SetTempOutbound(TempOutbound)
	m.SetServerAddr(netAddr)
	m.SetInboundDirectory(inb)
	m.SetOutboundDirectory(outb)
	m.SetAddr(newAddress)
	m.SetSecret(password)
	m.Start()
	m.Wait()

	/* Complete start tosser */


	return nil
}

func (self *ApiServiceStartAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	service := r.PostForm.Get("service")
	log.Printf("service = %s", service)

	/* ... */
	if service == "tosser" {

		log.Printf("Start tosser")

		self.Container.Invoke(func(tm *tosser.TosserManager) {
			tm.Toss()
		})

		p := make(map[string]interface{})
		p["code"] = 0
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)

	} else if service == "mailer" {

		self.Start()

		p := make(map[string]interface{})
		p["code"] = 0
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(p)

	}
}
