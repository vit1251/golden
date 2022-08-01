package api

import (
	"encoding/json"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/vit1251/golden/pkg/registry"
	"log"
	"net/http"
)


type commandStream struct {
	registry *registry.Container
	actions []*Action
}

func NewCommandStream() *commandStream {
	cs := new(commandStream)
	return cs
}

func (self *commandStream) SetContainer(r *registry.Container) {
	self.registry = r
}

func (self *commandStream) RegisterAction(a *Action) {
    self.actions = append(self.actions, a)
}

func (self *commandStream) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		// handle error
	}

        /* Initialize actions */
        self.RegisterAction(&NewUpdateStateAction(self.registry).Action)
        self.RegisterAction(&NewEchoIndexAction(self.registry).Action)
        self.RegisterAction(&NewEchoAreaIndexAction(self.registry).Action)

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
				if resp != nil {
        				err2 := wsutil.WriteServerMessage(conn, op, resp)
	        			if err2 != nil {
		        			log.Printf("Command stream write error: err = %#v", err2)
			        		break
				        }
				}
			}
		}
	}()
}

type Request struct {
        Type string `json:"type"`
}

func (self *commandStream) processRequest(body []byte) []byte {

        req := Request{}
        json.Unmarshal(body, &req)

        log.Printf("req = %+v", req)

        /* Processing */
        for _, a := range self.actions {
            log.Printf("action = %s", a.Type)
            if a.Type == req.Type {
                log.Printf("Process")
                return a.Handle()
            }
        }

        return nil
}
