package area

import (
	"errors"
)

type AreaList struct {
	Areas   []*Area
}

func (self *AreaList) Reset() {
	self.Areas = nil
}

func (self *AreaList) SearchByName(echoTag string) (*Area, error) {
	for _, area := range self.Areas {
		if area.Name == echoTag {
			return area, nil
		}
	}
	return nil, errors.New("No area exists.")
}
