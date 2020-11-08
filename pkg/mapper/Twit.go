package mapper

type Twit struct {
	id   string
	name string
}

func NewTwit() *Twit {
	return new(Twit)
}

func (self *Twit) SetId(id string) {
	self.id = id
}

func (self *Twit) SetName(name string) {
	self.name = name
}

func (self Twit) GetName() string {
	return self.name
}

func (self Twit) GetId() string {
	return self.id
}
