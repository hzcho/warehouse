package handler

import (
	"io"
	"net/http"
	"strconv"
	"warehouse/internal/domain/net/request"
	"warehouse/internal/domain/usecase"

	"warehouse/pkg/token"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	productUseCase usecase.Product
}

func NewProduct(productUseCase usecase.Product) *Product {
	return &Product{
		productUseCase: productUseCase,
	}
}

func (h *Product) GetAll(c *gin.Context) {
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

	products, err := h.productUseCase.GetAll(c, claims, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})
}

func (h *Product) GetById(c *gin.Context) {
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

	product, err := h.productUseCase.GetById(c, claims, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func (h *Product) Create(c *gin.Context) {
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

	var req request.CreateProduct

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	var images []io.Reader

	for _, fileHeader := range form.File["images"] {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open image"})
			return
		}
		defer file.Close()

		images = append(images, file)
	}

	req.Images = images

	productID, err := h.productUseCase.Create(c, claims, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product_id": productID})
}

func (h *Product) Update(c *gin.Context) {
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

	var req request.UpdateProduct
	req.ID = &id
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form"})
		return
	}

	var images []io.Reader

	for _, fileHeader := range form.File["images"] {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open image"})
			return
		}
		defer file.Close()

		images = append(images, file)
	}

	req.Images = &images

	updatedProduct, err := h.productUseCase.Update(c, claims, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": updatedProduct})
}

func (h *Product) UpdateStockLevel(c *gin.Context) {
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

	var req request.UpdateStockLevel
	req.Id = id
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
		return
	}

	updatedProduct, err := h.productUseCase.UpdateCount(c, claims, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": updatedProduct})
}

func (h *Product) Delete(c *gin.Context) {
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

	deletedId, err := h.productUseCase.Delete(c, claims, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted_product_id": deletedId})
}
