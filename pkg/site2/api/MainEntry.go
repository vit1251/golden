package api

import (
	"encoding/json"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
	"log"
	"net/http"
)

type commandStream struct {
	registry *registry.Container
}

func NewCommandStream() *commandStream {
	cs := new(commandStream)
	return cs
}

func (self *commandStream) SetContainer(r *registry.Container) {
	self.registry = r
}

func (self *commandStream) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		// handle error
	}
	go func() {
		defer conn.Close()

		for {
			/* Step 1. Read client data */
			req, op, err1 := wsutil.ReadClientData(conn)
			if err1 != nil {
				log.Printf("Fail on `ReadClientData` with %s", err1)
				break
			}
			/* Step 2. Debug message */
			log.Printf("CommandStream: req = %+v op = %+v", req, op)
			/* Step 3. Processing request */
			if op.IsData() {
				/* Step 1. Process user request */
				resp := self.processRequest(req)
				/* Step 2. Send user response */
				err2 := wsutil.WriteServerMessage(conn, op, resp)
				if err2 != nil {
					log.Printf("Command stream write error: err = %#v", err2)
					break
				}
			}
		}
	}()
}

type messageStatus struct {
	NetMessageCount  int
	EchoMessageCount int
	FileCount        int
}

func (self *commandStream) processRequest(req []byte) []byte {

	mapperManager := mapper.RestoreMapperManager(self.registry)
	netMapper := mapperManager.GetNetmailMapper()
	echoMapper := mapperManager.GetEchoMapper()
	fileMapper := mapperManager.GetFileMapper()

	newNetCount, _ := netMapper.GetMessageNewCount()
	newEchoCount, _ := echoMapper.GetMessageNewCount()
	newFileCount, _ := fileMapper.GetFileNewCount()

	result := messageStatus{
		NetMessageCount:  newNetCount,
		EchoMessageCount: newEchoCount,
		FileCount:        newFileCount,
	}

	/* Done */
	out, _ := json.Marshal(result)
	return out
}
