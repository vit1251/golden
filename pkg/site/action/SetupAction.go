package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/mapper"
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
	name string
	value string
	summary string
}

type setupSection struct {
	Name     string
	Params   []setupParam
}

func (self *setupSection) Register(c *mapper.Config, section string, name string, summary string) {

	paramValue, _ := c.Get(section, name)

	newSetupParam := new(setupParam)
	newSetupParam.section = section
	newSetupParam.name = name
	newSetupParam.summary = summary
	newSetupParam.value = paramValue

	self.Params = append(self.Params, *newSetupParam)

}

func (self SetupAction) makeNetworkingSection(c *mapper.Config) setupSection {

	/* Section header */
	netSetupSession := new(setupSection)
	netSetupSession.Name = "Networking"

	/* Section options */
	netSetupSession.Register(c,"main", "Address", "FidoNet network point address (i.e. POINT address)")
	netSetupSession.Register(c,"main", "NetAddr", "FidoNet network BOSS address (example: f24.n5023.z2.binkp.net:24554)")
	netSetupSession.Register(c,"main", "Password", "FidoNet point password")
	netSetupSession.Register(c,"main", "Link", "FidoNet uplink provide (i.e. BOSS address)")

	return *netSetupSession
}

func (self SetupAction) makeUserSection(c *mapper.Config) setupSection {

	/* Section header */
	userSetupSession := new(setupSection)
	userSetupSession.Name = "User options"

	/* Section options */
	userSetupSession.Register(c,"main", "RealName", "Realname is you English version your real name (example: Dmitri Kamenski)")
	userSetupSession.Register(c,"main", "Country", "Country where user is seat")
	userSetupSession.Register(c,"main", "City", "City where user is seat")
	userSetupSession.Register(c,"main", "Origin", "Origin was provide BBS station name and network address")
	userSetupSession.Register(c,"main", "TearLine", "Tearline provide person sign in all their messages")

	return *userSetupSession
}

func (self SetupAction) makeOtherSection(c *mapper.Config) setupSection {

	/* Section header */
	otherSetupSession := new(setupSection)
	otherSetupSession.Name = "Other"

	/* Section options */
	otherSetupSession.Register(c,"main", "StationName", "Station name is your nickname")
	otherSetupSession.Register(c,"mailer", "Interval", "Polling interval in minutes")

	return *otherSetupSession
}

func (self SetupAction) makeDirectMessageSection(c *mapper.Config) setupSection {

	/* Section header */
	directMessageSetupSection := new(setupSection)
	directMessageSetupSection.Name = "Netmail options"

	/* Section options */
	directMessageSetupSection.Register(c,"netmail", "Charset", "Netmail charset")

	return *directMessageSetupSection
}

func (self SetupAction) makeSections(c *mapper.Config) []setupSection {

	var setupSections []setupSection

	/* Networking section */
	networkingSection := self.makeNetworkingSection(c)
	setupSections = append(setupSections, networkingSection)

	/* User section */
	userSection := self.makeUserSection(c)
	setupSections = append(setupSections, userSection)

	/* Direct message section */
	directMessageSection := self.makeDirectMessageSection(c)
	setupSections = append(setupSections, directMessageSection)

	/* Other section */
	otherSection := self.makeOtherSection(c)
	setupSections = append(setupSections, otherSection)

	return setupSections

}

func (self SetupAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	configMapper := mapperManager.GetConfigMapper()

	/* Get params */
	config, _ := configMapper.GetConfig()

	/* Make sections */
	setupSections := self.makeSections(config)

	/* Render */
	bw := widgets.NewBaseWidget()

	vBox := widgets.NewVBoxWidget()
	bw.SetWidget(vBox)

	mmw := self.makeMenu()
	vBox.Add(mmw)

	container := widgets.NewDivWidget()
	container.SetClass("container")

	containerVBox := widgets.NewVBoxWidget()

	container.SetWidget(containerVBox)

	vBox.Add(container)

	setupForm := widgets.NewFormWidget().
		SetMethod("POST").
		SetAction("/setup/complete")

	sections := widgets.NewVBoxWidget()

	/* Make sections with settings */
	sectionsWithSettings := self.makeSectionsWithSettings(setupSections)
	sections.Add(sectionsWithSettings)

	/* Add save action */
	sections.Add(widgets.NewFormButtonWidget().SetTitle("Save").SetType("submit"))

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
	mainDiv.SetWidget(mainDivBox)

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
