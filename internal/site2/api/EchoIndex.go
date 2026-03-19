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

type echoIndex struct {
	Type  string          `json:"type"`  // Тип сообщения
	Areas []echoIndexArea `json:"areas"` // Список конференций
}

func AreasFindAll(registry *registry.Container) ([]mapper.Area, error) {
	mapperManager := mapper.RestoreMapperManager(registry)
	echoAreaMapper := mapperManager.GetEchoAreaMapper()
	return echoAreaMapper.GetAreas()
}

func (a *EchoIndexAction) processRequest(req []byte) []byte {

	// Шаг 0. Подготовка менеджера зависимостей
	registry := a.GetRegistry()

	// Шаг 1. Получаем список конференций
	areas, err1 := AreasFindAll(registry)
	if err1 != nil {
		log.Printf("Fail on GetAreas")
		return nil
	}

	var areas2 []echoIndexArea = make([]echoIndexArea, 0)
	for _, a := range areas {
		var na echoIndexArea
		na.Name = a.GetName()
		na.Summary = a.GetSummary()
		na.MessageCount = a.GetMessageCount()
		na.NewMessageCount = a.GetNewMessageCount()
		na.Order = a.GetOrder()
		na.AreaIndex = a.GetAreaIndex()
		areas2 = append(areas2, na)
	}

	// Шаг 3. Отправка ответа
	out, err2 := json.Marshal(&echoIndex{
		Type:  "ECHO_INDEX",
		Areas: areas2,
	})
	if err2 != nil {
		log.Printf("JSON encode issue")
		return nil
	}
	return out
}
