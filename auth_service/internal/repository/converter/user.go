package converter

import (
	"auth/internal/domain/model"
	"auth/internal/repository/dao"
)

func ToDomainUser(daoUser dao.User, daoRole dao.Role) model.User {
	return model.User{
		ID:           daoUser.ID,
		Login:        daoUser.Login,
		Password:     daoUser.PassHash,
		Role:         daoRole.Role,
		PhoneNumber:  daoUser.PhoneNumber,
		Email:        daoUser.Email,
		RefreshToken: daoUser.RefreshToken,
		TokenExpiry:  daoUser.TokenExpiry,
	}
}
