package action

import (
	"fmt"
	"log"
	"net/http"
)

type SetupCompleteAction struct {
	Action
}

func NewSetupCompleteAction() *SetupCompleteAction {
	sca := new(SetupCompleteAction)
	return sca
}

func (self *SetupCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	configManager := self.restoreConfigManager()

	/* Setup manager operation */
	params := configManager.GetParams()
	log.Printf("params = %+v", params)

	/* Update parameters */
	r.ParseForm()
	for _, param := range params {
		newValue := r.PostForm.Get(param.Name)
		log.Printf("param: name = %s value = %s newValue = %s", param.Name, param.Value, newValue)
		param.SetValue(newValue)
	}

	/* Store update */
	err1 := configManager.Store()
	if err1 != nil {
		panic(err1)
	}

	/* Redirect */
	newLocation := fmt.Sprintf("/setup")
	http.Redirect(w, r, newLocation, 303)
}
