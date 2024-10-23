package handler

import (
	"io"
	"net/http"
	"warehouse/internal/domain/net/request"
	"warehouse/internal/domain/usecase"

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

func (h *Product) GetById(c *gin.Context) {
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

	product, err := h.productUseCase.GetById(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func (h *Product) Create(c *gin.Context) {
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

	if form.File["images"] == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No images uploaded"})
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

	productID, err := h.productUseCase.Create(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product_id": productID})
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

	productID, err := h.productUseCase.Update(c, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product_id": productID})
}

func (h *Product) Delete(c *gin.Context) {
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

	deletedId, err := h.productUseCase.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deleted_product_id": deletedId})
}
