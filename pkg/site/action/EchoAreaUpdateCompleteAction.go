package action

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type EchoAreaUpdateCompleteAction struct {
	Action
}

func NewEchoAreaUpdateCompleteAction() *EchoAreaUpdateCompleteAction {
	euc := new(EchoAreaUpdateCompleteAction)
	return euc
}

func (self *EchoAreaUpdateCompleteAction) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	mapperManager := self.restoreMapperManager()
	echoAreaMapper := mapperManager.GetEchoAreaMapper()

	/* Parse POST parameters */
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	/* ... */
	vars := mux.Vars(r)
	echoTag := vars["echoname"]
	log.Printf("echoTag = %v", echoTag)

	/* ... */
	area, err1 := echoAreaMapper.GetAreaByName(echoTag)
	if err1 != nil {
		panic(err1)
	}
	log.Printf("area = %+v", area)

	/* Update summary */
	area.Summary = r.PostForm.Get("summary")

	/* Update charset */
	newCharset := r.PostForm.Get("charset")
	area.SetCharset(newCharset)

	/* Update order */
	newOrder := r.PostForm.Get("order")
	order, err2 := strconv.ParseInt(newOrder, 10, 64)
	if err2 != nil {
		log.Printf("Error parsing: %s as int64", newOrder)
	}
	area.SetOrder(order)

	/* Update area property */
	err3 := echoAreaMapper.Update(area)
	if err3 != nil {
		panic(err3)
	}

	/* Render */
	newLocation := fmt.Sprintf("/echo/%s", echoTag)
	http.Redirect(w, r, newLocation, 303)

}
