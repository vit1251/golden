package sqlite

type Area struct {
	Name    string
	Count   int
}

func NewArea() *Area {
	a := new(Area)
	return a
}
