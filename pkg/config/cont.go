
package config

import (
	"errors"
)

type AreaType int

const (
	AreaTupeUnkown AreaType = 0
	AreaTypeMsg    AreaType = 1
	AreaTypeSquish AreaType = 2
)

type Area struct {
	Name         string
	Path         string
	Type         AreaType
	MessageCount int
}

type AreaList struct {
	Areas   []*Area
}

func (self *AreaList) SearchByName(echoTag string) (*Area, error) {
	for _, area := range self.Areas {
		if area.Name == echoTag {
			return area, nil
		}
	}
	return nil, errors.New("No area exists.")
}

type Config struct {
	AreaList *AreaList
}
