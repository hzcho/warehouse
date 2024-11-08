package handler

import (
	"api_service/internal/domain/net/request"
	"api_service/internal/domain/service"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	productService service.Product
}

func NewProduct(productUseCase service.Product) *Product {
	return &Product{
		productService: productUseCase,
	}
}

func (h *Product) GetAll(c *gin.Context) {
	token, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty token"})
		return
	}
	strToken, ok := token.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token isn't a string"})
		return
	}

	var filter request.GetAllFilter

	if productName := c.Query("product_name"); productName != "" {
		filter.ProductName = &productName
	}

	if page := c.Query("page"); page != "" {
		pageInt, err := strconv.Atoi(page)
		if err == nil {
			filter.Page = &pageInt
		}
	}

	if limit := c.Query("limit"); limit != "" {
		limitInt, err := strconv.Atoi(limit)
		if err == nil {
			filter.Limit = &limitInt
		}
	}

	resp, err := h.productService.GetAll(c.Request.Context(), strToken, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get products"})
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

func (h *Product) GetById(c *gin.Context) {
	stringId := c.Param("id")
	if stringId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

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

	resp, err := h.productService.GetById(c, strToken, stringId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product"})
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Failed to read response body")
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (h *Product) Create(c *gin.Context) {
	var req request.CreateProduct

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request form data"})
		return
	}

	token, exists := c.Get("token")
	if !exists {
		fmt.Println("not exists token")
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty token"})
		return
	}
	strToken, ok := token.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token isn't string"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	req.Images = form.File["images"]

	resp, err := h.productService.Create(c, strToken, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (h *Product) Update(c *gin.Context) {
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

	var req request.UpdateUser
	req.ID = &id
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

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

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	temp := form.File["images"]
	req.Images = &temp

	resp, err := h.productService.Update(c, strToken, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}

func (h *Product) UpdateStockLevel(c *gin.Context) {
	stringId := c.Param("id")
	if stringId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	token, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty token"})
		return
	}
	strToken, ok := token.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token isn't a string"})
		return
	}

	var req request.UpdateStockLevel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	req.Id = stringId

	resp, err := h.productService.UpdateCount(c.Request.Context(), strToken, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stock level"})
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

func (h *Product) Delete(c *gin.Context) {
	stringId := c.Param("id")
	if stringId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

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

	resp, err := h.productService.Delete(c, strToken, stringId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, "Failed to read response body")
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
