package api

import (
	"encoding/json"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
	"log"
)

type EchoMsgRemoveAction struct {
	Action
}

func NewEchoMsgRemoveAction(r *registry.Container) *EchoMsgRemoveAction {
	o := new(EchoMsgRemoveAction)
	o.Action.Type = "ECHO_MSG_REMOVE"
	o.SetRegistry(r)
	o.Handle = o.processRequest
	return o
}

type echoMsgRemoveRequest struct {
	commonRequest
	EchoTag string `json:"echoTag"`
	MsgId   string `json:"msgId"`
}

type echoMsgRemoveResponse struct {
	CommonResponse
	Code int `json:"code"`
}

func (self *EchoMsgRemoveAction) processRequest(body []byte) []byte {

	/**/
	req := echoMsgRemoveRequest{}
	err1 := json.Unmarshal(body, &req)
	if err1 != nil {
		log.Printf("err = %+v", err1)
	}

	/* Step 0. Prepare mappers */
	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	echoMapper := mapperManager.GetEchoMapper()

	/* Step 1. Find current area by area UUID */
	areaIndex := req.EchoTag
	currentArea, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
	if currentArea == nil || err1 != nil {
		return nil
	}
	log.Printf("currentArea = %+v", currentArea)

	/* Step 2. Remove message by hash */
	currentAreaName := currentArea.GetName()
	msgHash := req.MsgId
	err2 := echoMapper.RemoveMessageByHash(currentAreaName, msgHash)
	if err2 != nil {
		return nil
	}

	/* Step 3. Populate API response */
	resp := new(echoMsgRemoveResponse)
	resp.CommonResponse.Type = self.Type

	resp.Code = 0

	/* Done */
	out, _ := json.Marshal(resp)
	return out
}
