package mailer

import "github.com/vit1251/golden/pkg/registry"

const MAILER_MANAGER_ID = "MailerManager"

func RestoreMailerManager(r *registry.Container) *MailerManager {
	managerPtr := r.Get(MAILER_MANAGER_ID)
	if manager, ok := managerPtr.(*MailerManager); ok {
		return manager
	} else {
		panic("no mailer manager")
	}
}
