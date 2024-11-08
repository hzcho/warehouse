package usecase_test

import (
	"context"
	"testing"
	"warehouse/internal/domain/model"
	"warehouse/internal/domain/net/request"
	publishMocks "warehouse/internal/domain/producer/mock"
	repoMocks "warehouse/internal/domain/repository/mock"
	"warehouse/internal/usecase"
	"warehouse/pkg/token"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	minValueTopic = "min_value"
	saveOperation = "save_operation"
)

func TestProduct_GetAll(t *testing.T) {
	mockProductRepo := new(repoMocks.Product)
	mockFileStorage := new(repoMocks.FileStorage)
	mockCategory := new(repoMocks.Category)
	logger := logrus.New()
	baseURL := "http://example.com"

	product := usecase.NewProduct(nil, mockProductRepo, mockCategory, mockFileStorage, logger, baseURL)

	ctx := context.TODO()
	claims := token.AuthInfo{UserID: "123"}
	filter := request.GetAllFilter{
		ProductName: stringPtr("test product"),
		Page:        intPtr(1),
		Limit:       intPtr(10),
	}

	id := primitive.NewObjectID()
	name := "test product"
	expectedProducts := []model.Product{
		{ID: &id, Name: &name},
	}
	mockProductRepo.On("GetAll", ctx, filter).Return(expectedProducts, nil)

	result, err := product.GetAll(ctx, claims, filter)

	assert.NoError(t, err)
	assert.Equal(t, 1, result.Page)
	assert.Equal(t, 10, result.Limit)
	assert.Equal(t, expectedProducts, result.Products)

	mockProductRepo.AssertExpectations(t)
	mockFileStorage.AssertExpectations(t)
}

func TestProduct_GetById(t *testing.T) {
	mockProductRepo := new(repoMocks.Product)
	mockFileStorage := new(repoMocks.FileStorage)
	mockCategory := new(repoMocks.Category)
	logger := logrus.New()
	baseURL := "http://example.com"

	product := usecase.NewProduct(nil, mockProductRepo, mockCategory, mockFileStorage, logger, baseURL)

	ctx := context.TODO()
	claims := token.AuthInfo{UserID: "123"}
	productID := primitive.NewObjectID()
	expectedProduct := model.Product{ID: &productID, Name: stringPtr("test product")}

	mockProductRepo.On("GetById", ctx, productID).Return(expectedProduct, nil)

	result, err := product.GetById(ctx, claims, productID)

	assert.NoError(t, err)
	assert.Equal(t, expectedProduct, result)

	mockProductRepo.AssertExpectations(t)
}

func TestProduct_Update(t *testing.T) {
	mockProductRepo := new(repoMocks.Product)
	mockFileStorage := new(repoMocks.FileStorage)
	mockPublisher := new(publishMocks.Publisher)
	mockCategory := new(repoMocks.Category)
	logger := logrus.New()
	baseURL := "http://example.com"

	product := usecase.NewProduct(mockPublisher, mockProductRepo, mockCategory, mockFileStorage, logger, baseURL)

	ctx := context.TODO()
	claims := token.AuthInfo{
		UserID: "123",
		Login:  "testUser",
		Role:   "admin",
	}

	req := request.UpdateProduct{
		Name:          stringPtr("Updated Product"),
		StockLevel:    intPtr(5),
		MinStockLevel: intPtr(10),
	}

	updatedProductID := primitive.NewObjectID()
	updatedProduct := model.Product{
		ID:            &updatedProductID,
		Name:          req.Name,
		StockLevel:    req.StockLevel,
		MinStockLevel: req.MinStockLevel,
	}

	mockProductRepo.On("Update", ctx, mock.AnythingOfType("model.Product")).Return(updatedProduct, nil)

	mockPublisher.On("Produce", minValueTopic, mock.AnythingOfType("model.MinValue")).Return(nil)

	mockPublisher.On("Produce", saveOperation, mock.AnythingOfType("model.OperationLog")).Return(nil)

	result, err := product.Update(ctx, claims, req)

	assert.NoError(t, err)
	assert.Equal(t, updatedProduct, result)

	mockProductRepo.AssertExpectations(t)
	mockPublisher.AssertExpectations(t)
}

func TestProduct_UpdateCount(t *testing.T) {
	mockProductRepo := new(repoMocks.Product)
	mockPublisher := new(publishMocks.Publisher)
	mockCategory := new(repoMocks.Category)
	logger := logrus.New()
	baseURL := "http://example.com"

	product := usecase.NewProduct(mockPublisher, mockProductRepo, mockCategory, nil, logger, baseURL)

	ctx := context.TODO()
	claims := token.AuthInfo{UserID: "123"}

	productID := primitive.NewObjectID()
	currentStockLevel := 10
	req := request.UpdateStockLevel{
		Id:         productID,
		StockLevel: -5,
	}

	productInDB := model.Product{
		ID:         &productID,
		StockLevel: &currentStockLevel,
	}

	updatedProduct := productInDB
	*updatedProduct.StockLevel += req.StockLevel

	mockProductRepo.On("GetById", ctx, req.Id).Return(productInDB, nil)
	mockProductRepo.On("Update", ctx, mock.AnythingOfType("model.Product")).Return(updatedProduct, nil)

	mockPublisher.On("Produce", saveOperation, mock.AnythingOfType("model.OperationLog")).Return(nil)

	result, err := product.UpdateCount(ctx, claims, req)

	assert.NoError(t, err)
	assert.Equal(t, updatedProduct, result)
	mockProductRepo.AssertExpectations(t)
	mockPublisher.AssertExpectations(t)
}

func TestProduct_Delete(t *testing.T) {
	mockProductRepo := new(repoMocks.Product)
	mockPublisher := new(publishMocks.Publisher)
	mockCategory := new(repoMocks.Category)
	logger := logrus.New()
	baseURL := "http://example.com"

	product := usecase.NewProduct(mockPublisher, mockProductRepo, mockCategory, nil, logger, baseURL)

	ctx := context.TODO()
	claims := token.AuthInfo{
		UserID: "123",
		Role:   "admin",
	}

	productID := primitive.NewObjectID()

	mockProductRepo.On("Delete", ctx, productID).Return(productID, nil)
	mockPublisher.On("Produce", saveOperation, mock.AnythingOfType("model.OperationLog")).Return(nil)

	result, err := product.Delete(ctx, claims, productID)

	assert.NoError(t, err)
	assert.Equal(t, productID, result)
	mockProductRepo.AssertExpectations(t)
	mockPublisher.AssertExpectations(t)
}

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}
