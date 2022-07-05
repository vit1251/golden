package queue

import (
	"github.com/vit1251/golden/pkg/registry"
)

type QueueManager struct {
	registry.Service
	mi *MailerInbound
	mo *MailerOutbound
}

func (self *QueueManager) GetMailerInbound() *MailerInbound {
	return self.mi
}

func (self *QueueManager) GetMailerOutbound() *MailerOutbound {
	return self.mo
}

func NewQueueManager(r *registry.Container) *QueueManager {
	qs := new(QueueManager)
	qs.SetRegistry(r)
	qs.mi = newMailerInbound(r)
	qs.mo = newMailerOutbound(r)
	return qs
}
