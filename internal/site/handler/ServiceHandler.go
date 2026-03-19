package handler

import (
	"fmt"
	"net/http"
	"strings"

	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/pkg/registry"
)

type ServiceHandler struct {
	registry *registry.Container
}

type Service struct {
	name string /* Service name */
	URL  string /* Service page */
}

func NewServiceHandler(registry *registry.Container) *ServiceHandler {
	return &ServiceHandler{
		registry: registry,
	}
}

func (self *ServiceHandler) makeServices() []Service {
	var services []Service
	/* Mailer service */
	services = append(services, Service{
		name: "mailer",
		URL:  "/service/mailer/stat",
	})
	/* Toss service */
	services = append(services, Service{
		name: "tosser",
		URL:  "/service/toss/stat",
	})
	/* Tracker service */
	services = append(services, Service{
		name: "tracker",
		URL:  "/service/tracker/stat",
	})
	return services
}

func (self ServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	services := self.makeServices()

	/* Render */
	bw := widgets2.NewBaseWidget()

	vBox := widgets2.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := widgets2.NewMainMenuWidget()
	vBox.Add(mmw)

	container := widgets2.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets2.NewVBoxWidget()

	container.AddWidget(containerVBox)

	vBox.Add(container)

	/* Render service */
	for _, s := range services {
		newRow := self.renderRow(s)
		containerVBox.Add(newRow)
	}

	/* Render */
	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (self ServiceHandler) renderRow(s Service) widgets2.IWidget {

	/* Make message row container */
	rowWidget := widgets2.NewDivWidget().
		SetStyle("display: flex").
		SetStyle("direction: column").
		SetStyle("align-items: center")

	var classNames []string
	classNames = append(classNames, "service-index-item")
	rowWidget.SetClass(strings.Join(classNames, " "))

	/* Render service name */
	serviceName := strings.Title(s.name)
	subjectWidget := widgets2.NewDivWidget().
		SetStyle("min-width: 350px").
		SetHeight("38px").
		SetStyle("flex-grow: 1").
		SetStyle("white-space: nowrap").
		SetStyle("overflow: hidden").
		SetStyle("text-overflow: ellipsis").
		//SetStyle("border: 1px solid red").
		SetContent(serviceName)

	rowWidget.AddWidget(subjectWidget)

	/* Link container */
	navigateItem := widgets2.NewLinkWidget().
		SetLink(s.URL).
		AddWidget(rowWidget)

	return navigateItem

}
