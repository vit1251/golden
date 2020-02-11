package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/mailer"
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

	/* Get Web-site action */
	webSite := self.Site

	/* Get setup manager */
	setupManager := webSite.GetSetupManager()
	params := setupManager.GetParams()
	log.Printf("params = %+v", params)

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

	/**/
	newAddress := fmt.Sprintf("%s@fidonet", address)

	/* Get parameters */
	m := mailer.NewMailer()
	m.SetServerAddr(netAddr)
	m.SetInboundDirectory(inb)
	m.SetOutboundDirectory(outb)
	m.SetAddr(newAddress)
	m.SetSecret(password)
	m.Start()
	m.Wait()

	return nil
}

func (self *ServiceManageCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* ... */
	self.Start()

	/* Redirect */
	newLocation := fmt.Sprintf("/service")
	http.Redirect(w, r, newLocation, 303)

}
