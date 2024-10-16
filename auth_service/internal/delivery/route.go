package delivery

import (
	"auth/internal/controller"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, ctrs *controller.Controllers) {
	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/signup", ctrs.Auth.SignUp)
			auth.POST("/signin", ctrs.Auth.SignIn)
			auth.POST("/refresh", ctrs.Auth.RefreshToken)
		}
	}
}
