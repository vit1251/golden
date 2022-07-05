package action

import (
	"fmt"
	"github.com/vit1251/golden/pkg/config"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type SetupCompleteAction struct {
	Action
}

func NewSetupCompleteAction() *SetupCompleteAction {
	sca := new(SetupCompleteAction)
	return sca
}

func (self *SetupCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	log.Printf("Store new settings")

	configManager := config.RestoreConfigManager(self.GetRegistry())

	newConfig := configManager.GetConfig()

	/* Update parameters */
	err1 := r.ParseForm()
	if err1 != nil {
		log.Printf("Parse form error: err = %#v", err1)
	}

	/* Main */
	params := config.GetParams()
	for _, p := range params {
		self.updateParam(newConfig, r.PostForm, p.Section, p.Name)
	}

	/* Setup default parameters */
	// TODO - if netAddr == "" {
	// TODO - uplinkAddr := "2:5030/1024"
	// TODO - newNetAddr := self.createNetAddr(uplinkAddr)
	// TODO - }

	/* Dump config */
	log.Printf("Dump config %#v", newConfig)

	/* Store update */
	err2 := configManager.Store(newConfig)
	if err2 != nil {
		log.Printf("Update config error: err = %#v", err2)
	}

	/* Redirect */
	newLocation := fmt.Sprintf("/setup")
	http.Redirect(w, r, newLocation, 303)

}

func (self *SetupCompleteAction) updateParam(c *config.Config, input url.Values, section string, name string) {
	newName := fmt.Sprintf("%s.%s", section, name)
	if values, ok := input[newName]; ok {
		newValue := strings.Join(values, ",")
		log.Printf("Update: section = %s name = %s value = %#v", section, name, newValue)
		config.SetByPath(c, section, name, newValue)
	}
}
