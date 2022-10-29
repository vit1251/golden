package action

import (
	"fmt"
	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"net/http"
	"strings"
)

type ServiceAction struct {
	Action
}

type Service struct {
	name string /* Service name */
	URL  string /* Service page */
}

func NewServiceAction() *ServiceAction {
	return new(ServiceAction)
}

func (self *ServiceAction) makeServices() []Service {
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

func (self ServiceAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	services := self.makeServices()

	/* Render */
	bw := widgets2.NewBaseWidget()

	vBox := widgets2.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
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

func (self ServiceAction) renderRow(s Service) widgets2.IWidget {

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
