package usecase

import (
	"context"
	"fmt"
	"notification/internal/domain/model"
	"notification/internal/domain/notifier"
	"notification/internal/domain/service"

	"github.com/sirupsen/logrus"
)

type Warehouse struct {
	emailNotifier notifier.Email
	authService   service.Auth
	log           *logrus.Logger
}

func NewWarehouse(emailNotifier notifier.Email, authService service.Auth, log *logrus.Logger) *Warehouse {
	return &Warehouse{
		emailNotifier: emailNotifier,
		authService:   authService,
		log:           log,
	}
}

func (u *Warehouse) MinValue(ctx context.Context, event model.MinValue) error {
	log := u.log.WithFields(logrus.Fields{
		"op": "/internal/usecase/warehouse/MinValue",
	})
	log.Info(event)

	filter := model.GetUsersFilter{
		Role: "employee",
	}

	users, err := u.authService.GetUsers(ctx, filter)
	if err != nil {
		log.Error(err)
		return err
	}

	for _, user := range users.Users {
		log.Infof("email %s", user.Email)

		message := model.EmailMessage{
			ToEmail: []string{user.Email},
			Subject: "Уведомление о минимальном лимите",
			Body: fmt.Sprintf(
				"Внимание, %s, достигнут минимальный лимит товара %s: %d\nТекущее количество товара: %d",
				user.Login,
				event.ProductName,
				event.MinStockLevel,
				event.StockLevel,
			),
		}

		err := u.emailNotifier.SendMessage(ctx, message)
		if err != nil {
			log.Error(err)
			return err
		}
	}

	return nil
}
