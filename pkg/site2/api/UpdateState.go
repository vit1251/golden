package api

import (
	"encoding/json"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
)

type UpdateStateAction struct {
        Action
}

type messageStatus struct {
        CommonResponse
	NetMessageCount  int
	EchoMessageCount int
	FileCount        int
}

func NewUpdateStateAction(r *registry.Container) *UpdateStateAction {
        o := new(UpdateStateAction)
        o.Action.Type = "SUMMARY"
        o.SetRegistry(r)
        o.Handle = o.processRequest
        return o
}

func (self *UpdateStateAction) processRequest() []byte {

    	mapperManager := mapper.RestoreMapperManager(self.GetRegistry())

    	netMapper := mapperManager.GetNetmailMapper()
    	echoMapper := mapperManager.GetEchoMapper()
    	fileMapper := mapperManager.GetFileMapper()

    	newNetCount, _ := netMapper.GetMessageNewCount()
    	newEchoCount, _ := echoMapper.GetMessageNewCount()
    	newFileCount, _ := fileMapper.GetFileNewCount()

    	resp := new(messageStatus)
    	resp.CommonResponse.Type = self.Type

    	resp.NetMessageCount = newNetCount
    	resp.EchoMessageCount = newEchoCount
    	resp.FileCount = newFileCount

        /* Done */
        out, _ := json.Marshal(resp)
        return out
}
