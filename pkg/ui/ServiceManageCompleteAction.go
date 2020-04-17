package ui

import (
	"encoding/json"
	"fmt"
	"github.com/vit1251/golden/pkg/mailer"
	"github.com/vit1251/golden/pkg/setup"
	"github.com/vit1251/golden/pkg/stat"
	"github.com/vit1251/golden/pkg/tosser"
	"log"
	"net/http"
)

type ServiceManageCompleteAction struct {
	Action
}

func NewServiceManageCompleteAction() *ServiceManageCompleteAction {
	smac := new(ServiceManageCompleteAction)
	return smac
}

func (self *ServiceManageCompleteAction) Start() {
	go self.Run()
}

func (self *ServiceManageCompleteAction) Run() error {

	var setupManager *setup.ConfigManager
	var statManager *stat.StatManager
	self.Container.Invoke(func(sm *setup.ConfigManager, sm2 *stat.StatManager) {
		setupManager = sm
		statManager = sm2
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

	/* Wait complete */
	m.Wait()

	/* Complete start tosser */
	if err := statManager.RegisterOutSession(); err != nil {
		log.Printf("Fail on mailer routine: err = %+v", err)
	}

	return nil
}

func (self *ServiceManageCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

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

	} else {
		log.Printf("! Unkown service: name = %+v", service)
	}

}
