package product_categories_service

import (
	"context"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

type ProductCategoriesService struct {
	categoryRepository ProductCategoriesRepository
}

type ProductCategoriesRepository interface {
	CreateProductCategory(ctx context.Context, category domain.ProductCategory) (domain.ProductCategory, error)
}

func NewProductCategoriesService(categoryRepository ProductCategoriesRepository) *ProductCategoriesService {
	return &ProductCategoriesService{
		categoryRepository: categoryRepository,
	}
}
