package product_categories_repository

import (
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

type ProductCategoriesModel struct {
	Code              string
	CategoryName      string
	InstallationPrice float64
}

func domainsFromModels(categoriesModels []ProductCategoriesModel) []domain.ProductCategory {
	categories := make([]domain.ProductCategory, len(categoriesModels))
	for i, model := range categoriesModels {
		categories[i] = *domain.NewProductCategory(
			model.Code,
			model.CategoryName,
			model.InstallationPrice,
		)
	}
	return categories
}

func domainFromModel(categoryModel ProductCategoriesModel) domain.ProductCategory {
	return *domain.NewProductCategory(
		categoryModel.Code,
		categoryModel.CategoryName,
		categoryModel.InstallationPrice,
	)
}
