package api

import (
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
	"log"
)

type EchoMsgIndexAction struct {
	Action
}

func NewEchoMsgIndexAction(r *registry.Container) *EchoMsgIndexAction {
	o := new(EchoMsgIndexAction)
	o.Action.Type = "ECHO_MSG_INDEX"
	o.SetRegistry(r)
	o.Handle = o.processRequest
	return o
}

func (self *EchoMsgIndexAction) processRequest(req []byte) []byte {

	/* Step 0. Prepare mappers */
	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	echoMapper := mapperManager.GetEchoMapper()

	/* Step 1. Find current area by area UUID */
	areaIndex := "e3c002a2-fde8-407e-bbd8-0de177527484"
	currentArea, err1 := echoAreaMapper.GetAreaByAreaIndex(areaIndex)
	if currentArea == nil || err1 != nil {
		return nil
	}
	log.Printf("currentArea = %+v", currentArea)

	/* Step 2. Get message headers by current area name */
	currentAreaName := currentArea.GetName()
	messageHeaders, err2 := echoMapper.GetMessageHeaders(currentAreaName)
	if err2 != nil {
		return nil
	}
	log.Printf("messageHeaders = %+v", messageHeaders)

	/* Step 3. Populate API response */
	for _, messageHeader := range messageHeaders {
		log.Printf("messageHeader = %+v", messageHeader)
	}

	return nil

}
