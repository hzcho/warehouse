package handler

import (
	"net/http"
	"strconv"
	"warehouse/internal/domain/net/request"
	"warehouse/internal/domain/usecase"
	"warehouse/pkg/token"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	categoryUseCase usecase.Category
}

func NewCategory(categoryUseCase usecase.Category) *Category {
	return &Category{
		categoryUseCase: categoryUseCase,
	}
}

func (h *Category) GetAll(c *gin.Context) {
	tknClaims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty claims"})
		return
	}

	claims, ok := tknClaims.(token.AuthInfo)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid token claims"})
		return
	}

	var req request.GetCategories

	if name := c.Query("name"); name != "" {
		req.Name = &name
	}

	if page := c.Query("page"); page != "" {
		pageInt, err := strconv.Atoi(page)
		if err == nil {
			req.Page = pageInt
		}
	}

	if limit := c.Query("limit"); limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err == nil {
			req.Limit = limitInt
		}
	}

	categories, err := h.categoryUseCase.GetAll(c, claims, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get products"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (h *Category) Create(c *gin.Context) {
	tknClaims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty claims"})
		return
	}
	claims, ok := tknClaims.(token.AuthInfo)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token isn't string"})
		return
	}

	var req request.CreateCategory

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	productID, err := h.categoryUseCase.Create(c, claims, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product_id": productID})
}

func (h *Category) Delete(c *gin.Context) {
	tknClaims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty claims"})
		return
	}
	claims, ok := tknClaims.(token.AuthInfo)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token isn't string"})
		return
	}

	stringId := c.Param("id")
	if stringId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	id, err := primitive.ObjectIDFromHex(stringId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	deletedId, err := h.categoryUseCase.Delete(c, claims, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted_category_id": deletedId})
}
