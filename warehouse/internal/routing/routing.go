package routing

import (
	"warehouse/internal/handler"

	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine, hnds *handler.Handlers) {
	api := router.Group("/api/v1")
	{
		products := api.Group("/products", hnds.AccountIdentity)
		{
			products.GET("", hnds.Product.GetAll)
			products.GET("/:id", hnds.Product.GetById)
			products.POST("", hnds.Product.Create)
			products.PUT("/:id", hnds.Product.Update)
			products.PUT("/count/:id", hnds.Product.UpdateStockLevel)
			products.DELETE("/:id", hnds.Product.Delete)
		}
		category := api.Group("/categories", hnds.AccountIdentity)
		{
			category.GET("", hnds.Category.GetAll)
			category.POST("", hnds.Category.Create)
			category.DELETE("/:id", hnds.Category.Delete)
		}
	}
}
