package routing

import (
	"warehouse/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, hnds *handler.Handlers) {
	api := router.Group("/api/v1")
	{
		products := api.Group("/products")
		{
			products.GET("/:id", hnds.Product.GetById)
			products.POST("", hnds.Product.Create)
			products.PUT("/:id", hnds.Product.Update)
			products.DELETE("/:id", hnds.Product.Delete)
		}
	}
}
