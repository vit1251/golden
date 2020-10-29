package registry

type Registrant struct {
	Name string
	Service interface{}
}

type Container struct {
	registrants []Registrant
}

func NewContainer() *Container {
	return new(Container)
}

func (self *Container) Register(name string, service interface{}) {
	registrant := Registrant{
		Name: name,
		Service: service,
	}
	self.registrants = append(self.registrants, registrant)
}

func (self *Container) Get(name string) interface{} {
	for _, r := range self.registrants {
		if r.Name == name {
			return r.Service
		}
	}
	return nil
}
