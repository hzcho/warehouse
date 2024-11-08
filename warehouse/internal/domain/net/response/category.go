package response

import "warehouse/internal/domain/model"

type GetCategories struct {
	Page       int
	Limit      int
	Categories []model.Category
}
