package action

import (
	"github.com/gorilla/mux"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/um"
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

	urlManager := um.RestoreUrlManager(self.GetRegistry())
	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
	echoAreaMapper := mapperManager.GetEchoAreaMapper()

	/* Parse POST parameters */
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	/* ... */
	vars := mux.Vars(r)
	areaIndex := vars["echoname"]
	log.Printf("areaIndex = %v", areaIndex)

	/* ... */
	area, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
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
	echoIndexAddr := urlManager.CreateUrl("/echo/{area_index}").
		SetParam("area_index", area.GetAreaIndex()).
		Build()

	http.Redirect(w, r, echoIndexAddr, 303)

}
