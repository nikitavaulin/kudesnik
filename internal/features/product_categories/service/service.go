package product_categories_service

import (
	"context"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

type ProductCategoriesService struct {
	categoriesRepository ProductCategoriesRepository
}

type ProductCategoriesRepository interface {
	CreateProductCategory(ctx context.Context, category domain.ProductCategory) (domain.ProductCategory, error)

	GetProductCategories(ctx context.Context, limit, offset *int) ([]domain.ProductCategory, error)

	GetProductCategory(ctx context.Context, categoryID uuid.UUID) (domain.ProductCategory, error)

	DeleteProductCategory(ctx context.Context, categoryID uuid.UUID) error

	PatchProductCategory(ctx context.Context, categoryID uuid.UUID, category domain.ProductCategory) (domain.ProductCategory, error)
}

func NewProductCategoriesService(categoriesRepository ProductCategoriesRepository) *ProductCategoriesService {
	return &ProductCategoriesService{
		categoriesRepository: categoriesRepository,
	}
}
