package area

import (
	"github.com/vit1251/golden/pkg/msgapi/sqlite"
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
	messageBase, err1 := sqlite.NewMessageBase()
	if err1 != nil {
		panic(err1)
	}

	/* Create message base reader */
	messageBaseReader, err2 := sqlite.NewMessageBaseReader(messageBase)
	if err2 != nil {
		panic(err2)
	}

	/* Preload echo areas */
	areas, err3 := messageBaseReader.GetAreaList2()
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
