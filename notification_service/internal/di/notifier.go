package di

import (
	"notification/internal/config"
	notifierInterface "notification/internal/domain/notifier"
	"notification/internal/notifier"
)

type Notifiers struct {
	notifierInterface.Email
}

func NewNotifiers(cfg *config.Config) *Notifiers {
	return &Notifiers{
		Email: notifier.NewEmail(cfg.SMPT),
	}
}
