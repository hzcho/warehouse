package controller

import (
	"auth/internal/domain/net/request"
	"auth/internal/domain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	authUseCase usecase.Auth
}

func NewAuth(authUseCase usecase.Auth) *Auth {
	return &Auth{
		authUseCase: authUseCase,
	}
}

func (c *Auth) SignUp(ctx *gin.Context) {
	req := request.SignUp{}

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "poor registration request structure")
		return
	}

	if err := c.authUseCase.SignUp(ctx, req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Has something happened")
		return
	}

	ctx.JSON(http.StatusOK, "the user is registered")
}

func (c *Auth) SignIn(ctx *gin.Context) {
	req := request.SignIn{}

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "poor registration request structure")
		return
	}

	tokens, err := c.authUseCase.SignIn(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Has something happened")
		return
	}

	ctx.JSON(http.StatusOK, tokens)
}

func (c *Auth) RefreshToken(ctx *gin.Context) {
	tokens := request.RefreshToken{}

	if err := ctx.ShouldBind(&tokens); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "poor registration request structure")
		return
	}

	newTokens, err := c.authUseCase.RefreshToken(ctx, tokens)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Has something happened")
		return
	}

	ctx.JSON(http.StatusOK, newTokens)
}