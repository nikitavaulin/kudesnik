package product_categories_repository

import (
	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

type ProductCategoriesModel struct {
	ID                uuid.UUID
	CategoryName      string
	InstallationPrice float64
}

func domainsFromModels(categoriesModels []ProductCategoriesModel) []domain.ProductCategory {
	categories := make([]domain.ProductCategory, len(categoriesModels))
	for i, model := range categoriesModels {
		categories[i] = *domain.NewProductCategory(
			model.ID,
			model.CategoryName,
			model.InstallationPrice,
		)
	}
	return categories
}
