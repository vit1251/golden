package api

import (
	"encoding/json"
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

type echoMsgHeader struct {
	//        ID string
	//        MsgID string
	Hash string `json:"hash"`
	//        Area string
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	//        UnixTime string
	DateWritten string `json:"date"`
	ViewCount   int    `json:"view_count"`
	//        Reply string
	//        FromAddr string
}

type echoMsgIndexResponse struct {
	CommonResponse
	Headers []echoMsgHeader `json:"headers"`
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
	resp := new(echoMsgIndexResponse)
	resp.CommonResponse.Type = self.Type

	for _, messageHeader := range messageHeaders {
		log.Printf("messageHeader = %+v", messageHeader)
		msgHeader := echoMsgHeader{}
		msgHeader.Hash = messageHeader.Hash
		msgHeader.From = messageHeader.From
		msgHeader.Subject = messageHeader.Subject
		msgHeader.DateWritten = messageHeader.DateWritten.String()
		resp.Headers = append(resp.Headers, msgHeader)
	}

	/* Done */
	out, _ := json.Marshal(resp)
	return out
}
