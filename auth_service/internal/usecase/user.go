package usecase

import (
	"auth/internal/domain/net/request"
	"auth/internal/domain/net/response"
	"auth/internal/domain/repository"
	"context"

	"github.com/sirupsen/logrus"
)

type User struct {
	userRepo repository.User
	log      *logrus.Logger
}

func NewUser(userRepo repository.User, log *logrus.Logger) *User {
	return &User{
		userRepo: userRepo,
		log:      log,
	}
}

func (u *User) GetUsers(ctx context.Context, filter request.GetUsersFilter) (response.Users, error) {
	log := u.log.WithFields(logrus.Fields{
		"op":    "internal/usecase/auth/GetUsers",
		"login": filter.Role,
	})
	users, err := u.userRepo.GetUsers(ctx, filter)
	if err != nil {
		log.Error(err)
		return response.Users{}, err
	}

	var baseUserInfos []response.BaseUserInfo
	for _, v := range users {
		u := response.BaseUserInfo{
			ID:          v.ID,
			Login:       v.Login,
			Role:        v.Role,
			PhoneNumber: v.PhoneNumber,
			Email:       v.Email,
		}

		baseUserInfos = append(baseUserInfos, u)
	}

	return response.Users{
		Users: baseUserInfos,
	}, nil
}
