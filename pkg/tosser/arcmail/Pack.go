package archmail

type Pack struct {
	Name    string
	Weekday string
	Index   int
}

func NewPack() *Pack {
	return new(Pack)
}
