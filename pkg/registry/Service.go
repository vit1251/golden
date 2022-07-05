package registry

type Service struct {
	registry *Container
}

func (self *Service) SetRegistry(r *Container) {
	self.registry = r
}

func (self *Service) GetRegistry() *Container {
	return self.registry
}
