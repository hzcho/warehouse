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
		product := api.Group("/products", hnds.Middleware.AccountIdentity)
		{
			product.GET("/", hnds.Product.GetAll)
			product.GET("/:id", hnds.Product.GetById)
			product.POST("/", hnds.Product.Create)
			product.PUT("/:id", hnds.Product.Update)
			product.PUT("/count/:id", hnds.Product.UpdateStockLevel)
			product.DELETE("/:id", hnds.Product.Delete)
		}
		category := api.Group("/categories", hnds.Middleware.AccountIdentity)
		{
			category.GET("", hnds.Category.GetAll)
			category.POST("", hnds.Category.Create)
			category.DELETE("/:id", hnds.Category.Delete)
		}
		audit := api.Group("/audit", hnds.Middleware.AccountIdentity)
		{
			audit.GET("/operations", hnds.Audit.GetAll)
		}
	}
}
