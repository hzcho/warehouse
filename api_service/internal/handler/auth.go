package handler

import (
	"api_service/internal/domain/net/request"
	"api_service/internal/domain/service"
	"io"

	"net/http"

	"github.com/gin-gonic/gin"
)

type Auth struct {
	authService service.Auth
}

func NewAuth(authService service.Auth) *Auth {
	return &Auth{
		authService: authService,
	}
}

func (c *Auth) SignUp(ctx *gin.Context) {
	req := request.SignUp{}

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "poor registration request structure")
		return
	}

	resp, err := c.authService.SignUp(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Has something happened")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Failed to read response body")
		return
	}

	ctx.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (c *Auth) SignIn(ctx *gin.Context) {
	req := request.SignIn{}

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "poor registration request structure")
		return
	}

	resp, err := c.authService.SignIn(ctx, req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Has something happened")
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Failed to read response body")
		return
	}

	ctx.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (c *Auth) RefreshToken(ctx *gin.Context) {
	tokens := request.RefreshToken{}

	if err := ctx.ShouldBind(&tokens); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, "poor registration request structure")
		return
	}

	resp, err := c.authService.RefreshToken(ctx, tokens)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Has something happened")
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, "Failed to read response body")
		return
	}

	ctx.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
