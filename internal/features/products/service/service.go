package products_service

import (
	"context"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

type ProductsService struct {
	productRepo ProductsRepository
}

type ProductsRepository interface {
	CreateProduct(ctx context.Context, product domain.BaseProduct) (domain.BaseProduct, error)
	GetProduct(ctx context.Context, id uuid.UUID) (domain.BaseProduct, error)
	GetProducts(ctx context.Context, categoryID *uuid.UUID, limit, offset *int) ([]domain.BaseProduct, error)
	DeleteProduct(ctx context.Context, id uuid.UUID) error
	DeleteProducts(ctx context.Context, IDs []uuid.UUID) error
	UpdateProductVisability(ctx context.Context, id uuid.UUID, isVisible bool) error
	UpdateProductsVisability(ctx context.Context, IDs []uuid.UUID, isVisible bool) error
}

func NewProductsService(productRepo ProductsRepository) *ProductsService {
	return &ProductsService{
		productRepo: productRepo,
	}
}
