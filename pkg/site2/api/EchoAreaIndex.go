package api

import (
	"encoding/json"
//	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
	"log"
)

type EchoAreaIndexAction struct {
	Action
}

func NewEchoAreaIndexAction(r *registry.Container) *EchoAreaIndexAction {
        o := new(EchoAreaIndexAction)
        o.Action.Type = "ECHO_AREA_INDEX"
        o.SetRegistry(r)
        o.Handle = o.processRequest
        return o
}

type echoItem struct {
	Name             string       `json:"name"`
}

type echoAreaIndexResponse struct {
	CommonResponse
	Items []echoItem `json:"items"`
}

func (self *EchoAreaIndexAction) processRequest() []byte {

//    	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())

    	resp := new(echoAreaIndexResponse)
    	resp.CommonResponse.Type = self.Type

    	    na := echoItem{}
    	    na.Name = ""
    	    resp.Items = append(resp.Items, na)

    	log.Printf("resp = %+v", resp)

        /* Done */
        out, _ := json.Marshal(resp)
        return out

}
