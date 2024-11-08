package converter

import (
	"warehouse/internal/domain/model"
	"warehouse/internal/domain/net/request"
)

func ProductFromCreate(createProduct request.CreateProduct) model.Product {
	return model.Product{
		ID:            nil,
		Name:          &createProduct.Name,
		Description:   &createProduct.Description,
		CategoryName:  &createProduct.CategoryName,
		Price:         &createProduct.Price,
		StockLevel:    &createProduct.StockLevel,
		MinStockLevel: &createProduct.MinStockLevel,
		Manufacturer:  &createProduct.Manufacturer,
		Supplier:      &createProduct.Supplier,
		Weight:        &createProduct.Weight,
		Dimensions: &model.Dimensions{
			Length: createProduct.Dimensions.Length,
			Width:  createProduct.Dimensions.Width,
			Height: createProduct.Dimensions.Height,
		},
		CreatedAt: nil,
		UpdatedAt: nil,
		ImageURLs: nil,
		Tags:      &createProduct.Tags,
		IsActive:  nil,
	}
}

func ProductFromUpdate(updateUser request.UpdateProduct) model.Product {
	return model.Product{
		ID:            updateUser.ID,
		Name:          updateUser.Name,
		Description:   updateUser.Description,
		CategoryName:  updateUser.CategoryName,
		Price:         updateUser.Price,
		StockLevel:    updateUser.StockLevel,
		MinStockLevel: updateUser.MinStockLevel,
		Manufacturer:  updateUser.Manufacturer,
		Supplier:      updateUser.Supplier,
		Weight:        updateUser.Weight,
		Dimensions:    updateUser.Dimensions,
		ImageURLs:     updateUser.ImageURLs,
		Tags:          updateUser.Tags,
		CreatedAt:     nil,
		UpdatedAt:     nil,
		IsActive:      updateUser.IsActive,
	}
}
