package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/netmail"
	"github.com/vit1251/golden/pkg/ui/widgets"
	"log"
	"net/http"
)

type NetmailAction struct {
	Action
}

func NewNetmailAction() (*NetmailAction) {
	nm := new(NetmailAction)
	return nm
}

func (self *NetmailAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var netmailManager *netmail.NetmailManager
	self.Container.Invoke(func(nm *netmail.NetmailManager) {
		netmailManager = nm
	})

	/* Message headers */
	msgHeaders, err1 := netmailManager.GetMessageHeaders()
	if err1 != nil {
		response := fmt.Sprintf("Fail on GetMessageHeaders on NetmailManager: err = %+v", err1)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}
	log.Printf("msgHeader = %+v", msgHeaders)

	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets.NewMainMenuWidget()
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets.NewVBoxWidget()

	amw := widgets.NewActionMenuWidget().
		Add(widgets.NewMenuAction().
			SetLink("/netmail/compose").
			SetIcon("icofont-edit").
			SetLabel("Compose"))

	containerVBox.Add(amw)
	container.SetWidget(containerVBox)
	vBox.Add(container)

	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
	}

}
