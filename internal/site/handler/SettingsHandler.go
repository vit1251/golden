package handler

import (
	"fmt"
	"net/http"

	widgets2 "github.com/vit1251/golden/internal/site/widgets"
	"github.com/vit1251/golden/pkg/config"
	"github.com/vit1251/golden/pkg/registry"
)

type SettingsHandler struct {
	registry *registry.Container
}

func NewSettingsHandler(registry *registry.Container) *SettingsHandler {
	return &SettingsHandler{
		registry: registry,
	}
}

/* Setup sections */
type setupParam struct {
	section string
	name    string
	value   string
	summary string
}

type setupSection struct {
	Name   string
	Params []setupParam
}

func (self *setupSection) Register(c *config.Config, section string, name string, summary string) {

	paramValue := config.GetByPath(c, section, name)

	newSetupParam := new(setupParam)
	newSetupParam.section = section
	newSetupParam.name = name
	newSetupParam.summary = summary
	newSetupParam.value = paramValue

	self.Params = append(self.Params, *newSetupParam)

}

func (h *SettingsHandler) makeFidoSection(c *config.Config) setupSection {

	/* Section header */
	netSetupSession := new(setupSection)
	netSetupSession.Name = "Fidonet options"

	/* Section options */
	netSetupSession.Register(c, "main", "Address", "FidoNet point address (example: 2:5020/1024.11)")
	netSetupSession.Register(c, "main", "Password", "FidoNet point password (example: pa$$w0rd)")
	netSetupSession.Register(c, "main", "Link", "FidoNet uplink provide (example: 2:5020/1024)")

	return *netSetupSession
}

func (h *SettingsHandler) makeGatewaySection(c *config.Config) setupSection {

	/* Section header */
	gwSetupSession := new(setupSection)
	gwSetupSession.Name = "Gateway options"

	/* Section options */
	gwSetupSession.Register(c, "main", "NetAddr", "Gateway IP address and port (example: f1024.n5020.z2.binkp.net:24554)")

	return *gwSetupSession
}

func (h *SettingsHandler) makeUserSection(c *config.Config) setupSection {

	/* Section header */
	userSetupSession := new(setupSection)
	userSetupSession.Name = "User options"

	/* Section options */
	userSetupSession.Register(c, "main", "RealName", "Realname is you English version your real name (example: Dmitri Kamenski)")
	userSetupSession.Register(c, "main", "Country", "Country where user is seat")
	userSetupSession.Register(c, "main", "City", "City where user is seat")
	userSetupSession.Register(c, "main", "Origin", "Origin was provide BBS station name and network address")
	userSetupSession.Register(c, "main", "TearLine", "Tearline provide person sign in all their messages")

	return *userSetupSession
}

func (h *SettingsHandler) makeOtherSection(c *config.Config) setupSection {

	/* Section header */
	otherSetupSession := new(setupSection)
	otherSetupSession.Name = "Other"

	/* Section options */
	otherSetupSession.Register(c, "main", "StationName", "Station name is your nickname")
	otherSetupSession.Register(c, "mailer", "Interval", "Polling interval in minutes")

	return *otherSetupSession
}

func (h *SettingsHandler) makeDirectMessageSection(c *config.Config) setupSection {

	/* Section header */
	directMessageSetupSection := new(setupSection)
	directMessageSetupSection.Name = "Netmail options"

	/* Section options */
	directMessageSetupSection.Register(c, "netmail", "Charset", "Netmail default charset")

	return *directMessageSetupSection
}

func (h *SettingsHandler) makeEchomailMessageSection(c *config.Config) setupSection {

	/* Section header */
	echoSetupSection := new(setupSection)
	echoSetupSection.Name = "Echomail options"

	/* Section options */
	echoSetupSection.Register(c, "echomail", "Charset", "Echomail default charset")

	return *echoSetupSection
}

func (h *SettingsHandler) makeSections(c *config.Config) []setupSection {

	var setupSections []setupSection

	/* Fidonet section */
	networkingSection := h.makeFidoSection(c)
	setupSections = append(setupSections, networkingSection)

	/* Gateway  section */
	gatewaySection := h.makeGatewaySection(c)
	setupSections = append(setupSections, gatewaySection)

	/* User section */
	userSection := h.makeUserSection(c)
	setupSections = append(setupSections, userSection)

	/* Direct message section */
	directMessageSection := h.makeDirectMessageSection(c)
	setupSections = append(setupSections, directMessageSection)

	/* Echomail message section */
	echoSection := h.makeEchomailMessageSection(c)
	setupSections = append(setupSections, echoSection)

	/* Other section */
	otherSection := h.makeOtherSection(c)
	setupSections = append(setupSections, otherSection)

	return setupSections

}

func (h *SettingsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	configManager := config.RestoreConfigManager(h.registry)

	/* Get params */
	newConfig := configManager.GetConfig()

	/* Make sections */
	setupSections := h.makeSections(newConfig)

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

	setupForm := widgets2.NewFormWidget().
		SetMethod("POST").
		SetAction("/settings/update")

	sections := widgets2.NewVBoxWidget()

	sections.Add(widgets2.NewFormButtonWidget().
		SetTitle("Save").
		SetType("submit"))
	sections.Add(widgets2.NewFormButtonWidget().
		SetTitle("Discard").
		SetType("reset"))

	/* Make sections with settings */
	sectionsWithSettings := h.makeSectionsWithSettings(setupSections)
	sections.Add(sectionsWithSettings)

	setupForm.SetWidget(sections)

	containerVBox.Add(setupForm)

	/* Render */
	if err := bw.Render(w); err != nil {
		status := fmt.Sprintf("%+v", err)
		http.Error(w, status, http.StatusInternalServerError)
		return
	}

}

func (h *SettingsHandler) createInputField(box *widgets2.VBoxWidget, param setupParam) {

	mainDiv := widgets2.NewDivWidget().
		SetClass("form-group row")

	mainDivBox := widgets2.NewVBoxWidget()
	mainDiv.AddWidget(mainDivBox)

	mainTitle := widgets2.NewDivWidget().
		SetClass("").
		SetContent(param.name)

	mainDivBox.Add(mainTitle)

	newInputName := fmt.Sprintf("%s.%s", param.section, param.name)
	mainInput := widgets2.NewFormInputWidget().
		SetTitle(param.summary).
		SetName(newInputName).
		SetValue(param.value)

	mainDivBox.Add(mainInput)

	box.Add(mainDiv)

}

func (h *SettingsHandler) makeSectionWithSettings(s setupSection) widgets2.IWidget {

	newSection := widgets2.NewSectionWidget()

	newSection.SetTitle(s.Name)

	sectionVBox := widgets2.NewVBoxWidget()

	for _, param := range s.Params {
		h.createInputField(sectionVBox, param)
	}

	newSection.SetWidget(sectionVBox)

	return newSection

}

func (h *SettingsHandler) makeSectionsWithSettings(setupSections []setupSection) widgets2.IWidget {

	newSections := widgets2.NewVBoxWidget()

	/* Draw sections */
	for _, s := range setupSections {
		newSection := h.makeSectionWithSettings(s)
		newSections.Add(newSection)
	}

	return newSections

}
