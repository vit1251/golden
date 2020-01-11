package area

import (
	"github.com/vit1251/golden/pkg/msg"
	"log"
)

type AreaManager struct {
	AreaList   AreaList
}

func NewAreaManager() (*AreaManager) {
	am := new(AreaManager)
	return am
}

func (self *AreaManager) Reset() {
	self.AreaList.Reset()
}

func (self *AreaManager) Register(a *Area) {
	self.AreaList.Areas = append(self.AreaList.Areas, a)
}

func (self *AreaManager) GetAreas() ([]*Area) {
	self.Rescan()
	return self.AreaList.Areas
}

func (self *AreaManager) GetAreaByName(echoTag string) (*Area, error) {
	self.Rescan()
	return self.AreaList.SearchByName(echoTag)
}

func (self *AreaManager) Rescan() {

	/* Open message base */
	messageManager := msg.NewMessageManager()

	/* Preload echo areas */
	areas, err3 := messageManager.GetAreaList2()
	if err3 != nil {
		panic(err3)
	}

	/* Reset areas */
	self.Reset()
	for _, area := range areas {
		log.Printf("area = %q", area)
		a := NewArea()
		a.Name = area.Name
		a.MessageCount = area.Count
		self.Register(a)
	}

}
