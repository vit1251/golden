package action

import (
	"fmt"
	"log"
	"net/http"
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

	mapperManager := self.restoreMapperManager()
	configMapper := mapperManager.GetConfigMapper()

	/* Setup manager operation */
	config, _ := configMapper.GetConfig()

	/* Update parameters */
	r.ParseForm()

	params := config.GetParams()
	for _, param := range params {
		newFormName := fmt.Sprintf("%s.%s", param.Section, param.Name)
		if items, ok := r.PostForm[newFormName]; ok {
			curValue := param.GetValue()
			newValue := strings.Join(items, ",")

			if curValue != newValue {
				log.Printf("SetupCompleteAction: section = %s name = %s value = %s -> %s", param.Section, param.Name, curValue, newValue)
				param.SetValue(newValue)
			}

		} else {
			log.Printf("Problem with value: name = %s", newFormName)
		}
	}

	/* Store update */
	err1 := configMapper.Store(config)
	if err1 != nil {
		panic(err1)
	}

	/* Redirect */
	newLocation := fmt.Sprintf("/setup")
	http.Redirect(w, r, newLocation, 303)

}
