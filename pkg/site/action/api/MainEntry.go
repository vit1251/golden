package api

import (
	"encoding/json"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/vit1251/golden/pkg/charset"
	"github.com/vit1251/golden/pkg/config"
	"github.com/vit1251/golden/pkg/eventbus"
	"github.com/vit1251/golden/pkg/mapper"
	"github.com/vit1251/golden/pkg/registry"
	"github.com/vit1251/golden/pkg/storage"
	"github.com/vit1251/golden/pkg/tosser"
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

func (self commandStream) restoreTosserManager() *tosser.TosserManager {
	managerPtr := self.registry.Get("TosserManager")
	if manager, ok := managerPtr.(*tosser.TosserManager); ok {
		return manager
	} else {
		panic("no tosser manager")
	}
}

func (self commandStream) restoreEventBus() *eventbus.EventBus {
	managerPtr := self.registry.Get("EventBus")
	if manager, ok := managerPtr.(*eventbus.EventBus); ok {
		return manager
	} else {
		panic("no eventbus manager")
	}
}

func (self commandStream) restoreStorageManager() *storage.StorageManager {
	managerPtr := self.registry.Get("StorageManager")
	if manager, ok := managerPtr.(*storage.StorageManager); ok {
		return manager
	} else {
		panic("no storage manager")
	}
}

func (self commandStream) restoreMapperManager() *mapper.MapperManager {
	managerPtr := self.registry.Get("MapperManager")
	if manager, ok := managerPtr.(*mapper.MapperManager); ok {
		return manager
	} else {
		panic("no mapper manager")
	}
}

func (self commandStream) restoreCharsetManager() *charset.CharsetManager {
	managerPtr := self.registry.Get("CharsetManager")
	if manager, ok := managerPtr.(*charset.CharsetManager); ok {
		return manager
	} else {
		panic("no charset manager")
	}
}

func (self *commandStream) restoreConfigManager() *config.ConfigManager {
	managerPtr := self.registry.Get("ConfigManager")
	if manager, ok := managerPtr.(*config.ConfigManager); ok {
		return manager
	} else {
		panic("no config manager")
	}
}

func (self *commandStream) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		// handle error
	}
	go func() {
		defer conn.Close()

		for {
			req, op, err1 := wsutil.ReadClientData(conn)
			if err1 != nil {
				log.Printf("Command stream read error: err = %#v", err1)
				break
			}
			//
			log.Printf("CommandStream: req = %s", req)
			resp := self.processRequest(req)
			//
			err2 := wsutil.WriteServerMessage(conn, op, resp)
			if err2 != nil {
				log.Printf("Command stream write error: err = %#v", err2)
				break
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

	mapperManager := self.restoreMapperManager()
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
