package mapper

type Twit struct {
	Name string
}

func NewTwit() *Twit {
	return new(Twit)
}

func (self *Twit) SetName(name string) {
	self.Name = name
}
