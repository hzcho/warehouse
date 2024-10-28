package controller

import (
	"auth/internal/domain/net/request"
	"auth/internal/domain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	userUseCase usecase.User
}

func NewUser(userUseCase usecase.User) *User {
	return &User{
		userUseCase: userUseCase,
	}
}

func (c *User) GetUsers(ctx *gin.Context) {
	req := request.GetUsersFilter{}

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "poor registration request structure")
		return
	}

	users, err := c.userUseCase.GetUsers(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Has something happened")
		return
	}

	ctx.JSON(http.StatusOK, users)
}
