package usecase

import (
	"context"
	"fmt"
	"notification/internal/domain/model"
	"notification/internal/domain/notifier"
	"notification/internal/domain/service"
)

type Warehouse struct {
	emailNotifier notifier.Email
	authService   service.Auth
}

func NewWarehouse(emailNotifier notifier.Email, authService service.Auth) *Warehouse {
	return &Warehouse{
		emailNotifier: emailNotifier,
		authService:   authService,
	}
}

func (u *Warehouse) MinValue(ctx context.Context, event model.MinValue) error {
	filter := model.GetUsersFilter{
		Role: "employee",
	}

	users, err := u.authService.GetUsers(ctx, filter)
	if err != nil {
		return err
	}

	for _, user := range users.Users {
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
			return err
		}
	}

	return nil
}
