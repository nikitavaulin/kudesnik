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
	CreateProduct(ctx context.Context, product domain.ProductBase) (domain.ProductBase, error)
	CreateWindow(ctx context.Context, product domain.Window) (domain.Window, error)

	GetProducts(ctx context.Context, categoryID *uuid.UUID, limit, offset *int) ([]domain.ProductBase, error)
	GetProduct(ctx context.Context, id uuid.UUID) (domain.ProductBase, error)
	GetWindow(ctx context.Context, id uuid.UUID) (domain.Window, error)

	DeleteProduct(ctx context.Context, id uuid.UUID) error
	DeleteProducts(ctx context.Context, IDs []uuid.UUID) error

	UpdateProductVisability(ctx context.Context, id uuid.UUID, isVisible bool) error
	UpdateProductsVisability(ctx context.Context, IDs []uuid.UUID, isVisible bool) error

	PatchProduct(ctx context.Context, id uuid.UUID, product domain.ProductBase) (domain.ProductBase, error)
	PatchWindow(ctx context.Context, id uuid.UUID, product domain.Window) (domain.Window, error)
}

func NewProductsService(productRepo ProductsRepository) *ProductsService {
	return &ProductsService{
		productRepo: productRepo,
	}
}
