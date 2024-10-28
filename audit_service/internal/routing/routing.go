package routing

import (
	"audit/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, handlers *handler.Handlers) {
	api := router.Group("/api/v1")
	{
		operations := api.Group("/operations")
		{
			operations.GET("/:id", handlers.GetById)
			operations.GET("", handlers.GetAll)
		}
	}
}
