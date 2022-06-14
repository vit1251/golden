package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/config"
	"github.com/vit1251/golden/pkg/site/widgets"
	"net/http"
)

type SetupAction struct {
	Action
}

func NewSetupAction() *SetupAction {
	newSetupAction := new(SetupAction)
	return newSetupAction
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

func (self SetupAction) makeFidoSection(c *config.Config) setupSection {

	/* Section header */
	netSetupSession := new(setupSection)
	netSetupSession.Name = "Fidonet options"

	/* Section options */
	netSetupSession.Register(c, "main", "Address", "FidoNet point address (example: 2:5020/1024.11)")
	netSetupSession.Register(c, "main", "Password", "FidoNet point password (example: pa$$w0rd)")
	netSetupSession.Register(c, "main", "Link", "FidoNet uplink provide (example: 2:5020/1024)")

	return *netSetupSession
}

func (self SetupAction) makeGatewaySection(c *config.Config) setupSection {

	/* Section header */
	gwSetupSession := new(setupSection)
	gwSetupSession.Name = "Gateway options"

	/* Section options */
	gwSetupSession.Register(c, "main", "NetAddr", "Gateway IP address and port (example: f1024.n5020.z2.binkp.net:24554)")

	return *gwSetupSession
}

func (self SetupAction) makeUserSection(c *config.Config) setupSection {

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

func (self SetupAction) makeOtherSection(c *config.Config) setupSection {

	/* Section header */
	otherSetupSession := new(setupSection)
	otherSetupSession.Name = "Other"

	/* Section options */
	otherSetupSession.Register(c, "main", "StationName", "Station name is your nickname")
	otherSetupSession.Register(c, "mailer", "Interval", "Polling interval in minutes")

	return *otherSetupSession
}

func (self SetupAction) makeDirectMessageSection(c *config.Config) setupSection {

	/* Section header */
	directMessageSetupSection := new(setupSection)
	directMessageSetupSection.Name = "Netmail options"

	/* Section options */
	directMessageSetupSection.Register(c, "netmail", "Charset", "Netmail default charset")

	return *directMessageSetupSection
}

func (self SetupAction) makeEchomailMessageSection(c *config.Config) setupSection {

	/* Section header */
	echoSetupSection := new(setupSection)
	echoSetupSection.Name = "Echomail options"

	/* Section options */
	echoSetupSection.Register(c, "echomail", "Charset", "Echomail default charset")

	return *echoSetupSection
}

func (self SetupAction) makeSections(c *config.Config) []setupSection {

	var setupSections []setupSection

	/* Fidonet section */
	networkingSection := self.makeFidoSection(c)
	setupSections = append(setupSections, networkingSection)

	/* Gateway  section */
	gatewaySection := self.makeGatewaySection(c)
	setupSections = append(setupSections, gatewaySection)

	/* User section */
	userSection := self.makeUserSection(c)
	setupSections = append(setupSections, userSection)

	/* Direct message section */
	directMessageSection := self.makeDirectMessageSection(c)
	setupSections = append(setupSections, directMessageSection)

	/* Echomail message section */
	echoSection := self.makeEchomailMessageSection(c)
	setupSections = append(setupSections, echoSection)

	/* Other section */
	otherSection := self.makeOtherSection(c)
	setupSections = append(setupSections, otherSection)

	return setupSections

}

func (self SetupAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	configManager := self.restoreConfigManager()

	/* Get params */
	newConfig := configManager.GetConfig()

	/* Make sections */
	setupSections := self.makeSections(newConfig)

	/* Render */
	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets.NewVBoxWidget()

	container.AddWidget(containerVBox)

	vBox.Add(container)

	setupForm := widgets.NewFormWidget().
		SetMethod("POST").
		SetAction("/setup/complete")

	sections := widgets.NewVBoxWidget()

	sections.Add(widgets.NewFormButtonWidget().
		SetTitle("Save").
		SetType("submit"))
	sections.Add(widgets.NewFormButtonWidget().
		SetTitle("Discard").
		SetType("reset"))

	/* Make sections with settings */
	sectionsWithSettings := self.makeSectionsWithSettings(setupSections)
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

func (self *SetupAction) createInputField(box *widgets.VBoxWidget, param setupParam) {

	mainDiv := widgets.NewDivWidget().
		SetClass("form-group row")

	mainDivBox := widgets.NewVBoxWidget()
	mainDiv.AddWidget(mainDivBox)

	mainTitle := widgets.NewDivWidget().
		SetClass("").
		SetContent(param.name)

	mainDivBox.Add(mainTitle)

	newInputName := fmt.Sprintf("%s.%s", param.section, param.name)
	mainInput := widgets.NewFormInputWidget().
		SetTitle(param.summary).
		SetName(newInputName).
		SetValue(param.value)

	mainDivBox.Add(mainInput)

	box.Add(mainDiv)

}

func (self SetupAction) makeSectionWithSettings(s setupSection) widgets.IWidget {

	newSection := widgets.NewSectionWidget()

	newSection.SetTitle(s.Name)

	sectionVBox := widgets.NewVBoxWidget()

	for _, param := range s.Params {
		self.createInputField(sectionVBox, param)
	}

	newSection.SetWidget(sectionVBox)

	return newSection

}

func (self SetupAction) makeSectionsWithSettings(setupSections []setupSection) widgets.IWidget {

	newSections := widgets.NewVBoxWidget()

	/* Draw sections */
	for _, s := range setupSections {
		newSection := self.makeSectionWithSettings(s)
		newSections.Add(newSection)
	}

	return newSections

}
