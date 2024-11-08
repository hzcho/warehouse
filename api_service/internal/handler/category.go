package handler

import (
	"api_service/internal/domain/net/request"
	"api_service/internal/domain/service"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	categoryService service.Category
}

func NewCategory(categoryService service.Category) *Category {
	return &Category{
		categoryService: categoryService,
	}
}

func (h *Category) GetAll(c *gin.Context) {
	token, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty token"})
		return
	}
	strToken, ok := token.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token isn't string"})
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

	resp, err := h.categoryService.GetAll(c, strToken, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get products"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (h *Category) Create(c *gin.Context) {
	token, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty token"})
		return
	}
	strToken, ok := token.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token isn't string"})
		return
	}

	var req request.CreateCategory

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	resp, err := h.categoryService.Create(c, strToken, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (h *Category) Delete(c *gin.Context) {
	token, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty token"})
		return
	}
	strToken, ok := token.(string)
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

	resp, err := h.categoryService.Delete(c, strToken, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
