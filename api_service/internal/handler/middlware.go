package handler

import (
	"api_service/internal/domain/usecase"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
)

type Middleware struct {
	authUseCase usecase.Auth
}

func NewMiddleware(authUseCase usecase.Auth) *Middleware {
	return &Middleware{
		authUseCase: authUseCase,
	}
}

func (m *Middleware) AccountIdentity(c *gin.Context) {
	header := c.Request.Header.Get(authorizationHeader)
	if header == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No authorization header"})
		c.Abort()
		return
	}

	parts := strings.SplitN(header, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header"})
		c.Abort()
		return
	}

	if len(parts[1]) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token is empty"})
		c.Abort()
		return
	}

	_, err := m.authUseCase.VerifyToken(c.Request.Context(), parts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}

	c.Set("token", parts[1])

	c.Next()
}
