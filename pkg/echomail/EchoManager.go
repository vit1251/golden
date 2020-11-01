package echomail

import "github.com/vit1251/golden/pkg/registry"

type EchoManager struct {
	registy *registry.Container
}

func NewEchoManager(r *registry.Container) *EchoManager {
	manager := new(EchoManager)
	manager.registy = r
	return manager
}
