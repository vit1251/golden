package msg

import (
	"errors"
	"strings"
)

type AreaList struct {
	Areas   []*Area
}

func (self *AreaList) Reset() {
	self.Areas = nil
}

func (self *AreaList) SearchByName(echoTag string) (*Area, error) {
	for _, area := range self.Areas {
		var areaName string = area.Name()
		if strings.EqualFold(areaName, echoTag) {
			return area, nil
		}
	}
	return nil, errors.New("no area exists")
}
