package ui

type Area struct {
	Name         string
	Path         string
	MessageCount int
}

func NewArea() *Area {
	a := new(Area)
	return a
}
