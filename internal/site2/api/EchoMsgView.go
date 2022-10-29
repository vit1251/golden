package api

import (
	"encoding/json"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
	"log"
)

type EchoMsgViewAction struct {
	Action
}

func NewEchoMsgViewAction(r *registry.Container) *EchoMsgViewAction {
	o := new(EchoMsgViewAction)
	o.Action.Type = "ECHO_MSG_VIEW"
	o.SetRegistry(r)
	o.Handle = o.processRequest
	return o
}

type echoMsg struct {
	Body string `json:"body"`
}

type echoMsgViewRequest struct {
	commonRequest
	EchoTag string `json:"echoTag"`
	MsgId   string `json:"msgId"`
}

type echoMsgViewResponse struct {
	CommonResponse
	EchoMsg echoMsg `json:"message"`
}

func (self *EchoMsgViewAction) processRequest(body []byte) []byte {

	/**/
	req := echoMsgViewRequest{}
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

	/* Step 2. Get message headers by current area name */
	currentAreaName := currentArea.GetName()
	msgHash := req.MsgId
	message, err2 := echoMapper.GetMessageByHash(currentAreaName, msgHash)
	if message == nil || err2 != nil {
		log.Printf("No message: msgId = %+v", msgHash)
		return nil
	}
	log.Printf("message = %+v", message)

	/* Update message view counter */
	err3 := echoMapper.ViewMessageByHash(currentAreaName, msgHash)
	if err3 != nil {
		return nil
	}

	/* Step 3. Populate API response */
	resp := new(echoMsgViewResponse)
	resp.CommonResponse.Type = self.Type

	resp.EchoMsg.Body = message.Content

	/* Done */
	out, _ := json.Marshal(resp)
	return out
}
