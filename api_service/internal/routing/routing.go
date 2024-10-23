package routing

import (
	"api_service/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, hnds *handler.Handlers) {
	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/signin", hnds.Auth.SignIn)
			auth.POST("/signup", hnds.Auth.SignUp)
			auth.POST("/refresh", hnds.Auth.RefreshToken)
		}
	}
}
