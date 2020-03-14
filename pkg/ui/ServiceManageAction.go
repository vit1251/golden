package ui

import (
	"fmt"
	"github.com/vit1251/golden/pkg/ui/views"
	"net/http"
	"path/filepath"
)

type ServiceManageAction struct {
	Action
}

type ServiceInfo struct {
	Label string
	Class string
}

func NewServiceManageAction() *ServiceManageAction {
	sma := new(ServiceManageAction)
	return sma
}

func (self *ServiceManageAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	/* Services */
	var services []*ServiceInfo

	if si := new(ServiceInfo); si != nil {
		si.Label = "Mailer Service"
		si.Class = "mailer"
		services = append(services, si)
	}
	if si := new(ServiceInfo); si != nil {
		si.Label = "Tosser Service"
		si.Class = "tosser"
		services = append(services, si)
	}

	/* Render */
	doc := views.NewDocument()
	layoutPath := filepath.Join("views", "layout.tmpl")
	doc.SetLayout(layoutPath)
	pagePath := filepath.Join("views", "service_index.tmpl")
	doc.SetPage(pagePath)
	doc.SetParam("Services", services)
	if err := doc.Render(w); err != nil {
		response := fmt.Sprintf("Fail on Render: err = %+v", err)
		http.Error(w, response, http.StatusInternalServerError)
		return
	}

}

