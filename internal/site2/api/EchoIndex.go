package api

import (
	"encoding/json"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
	"log"
)

type EchoIndexAction struct {
	Action
}

func NewEchoIndexAction(r *registry.Container) *EchoIndexAction {
	o := new(EchoIndexAction)
	o.Action.Type = "ECHO_INDEX"
	o.SetRegistry(r)
	o.Handle = o.processRequest
	return o
}

type echoIndexArea struct {
	Name            string `json:"name"`
	Summary         string `json:"summary"`
	MessageCount    int    `json:"message_count"`
	NewMessageCount int    `json:"new_message_count"`
	Order           int64  `json:"order"`
	AreaIndex       string `json:"area_index"`
}

type echoIndexResponse struct {
	CommonResponse
	Areas []echoIndexArea `json:"areas"`
}

func (self *EchoIndexAction) processRequest(req []byte) []byte {

	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
	echoAreaMapper := mapperManager.GetEchoAreaMapper()

	/* Get message area */
	areas, err1 := echoAreaMapper.GetAreas()
	if err1 != nil {
		log.Printf("Fail on GetAreas")
		return nil
	}

	resp := new(echoIndexResponse)
	resp.CommonResponse.Type = self.Type

	for _, a := range areas {
		na := echoIndexArea{}
		na.Name = a.GetName()
		na.Summary = a.GetSummary()
		//    	    na.MessageCount = a.GetMessageCount()
		na.NewMessageCount = a.GetNewMessageCount()
		na.Order = a.GetOrder()
		na.AreaIndex = a.GetAreaIndex()
		resp.Areas = append(resp.Areas, na)
	}
	log.Printf("resp = %+v", resp)

	/* Done */
	out, _ := json.Marshal(resp)
	return out
}
