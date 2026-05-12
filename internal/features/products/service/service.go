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
	CreateEntranceDoor(ctx context.Context, door domain.EntranceDoor) (domain.EntranceDoor, error)
	CreateInteriorDoor(ctx context.Context, door domain.InteriorDoor) (domain.InteriorDoor, error)
	CreateBalcony(ctx context.Context, balcony domain.Balcony) (domain.Balcony, error)

	GetProducts(ctx context.Context, categoryID *uuid.UUID, limit, offset *int) ([]domain.ProductBase, error)
	GetProduct(ctx context.Context, id uuid.UUID) (domain.ProductBase, error)
	GetProductDetailed(ctx context.Context, id uuid.UUID) (domain.ProductBaseDetailed, error)
	GetWindow(ctx context.Context, id uuid.UUID) (domain.Window, error)
	GetEntranceDoor(ctx context.Context, id uuid.UUID) (domain.EntranceDoor, error)
	GetInteriorDoor(ctx context.Context, id uuid.UUID) (domain.InteriorDoor, error)
	GetBalcony(ctx context.Context, id uuid.UUID) (domain.Balcony, error)

	DeleteProduct(ctx context.Context, id uuid.UUID) error
	DeleteProducts(ctx context.Context, IDs []uuid.UUID) error

	UpdateProductVisability(ctx context.Context, id uuid.UUID, isVisible bool) error
	UpdateProductsVisability(ctx context.Context, IDs []uuid.UUID, isVisible bool) error

	PatchProduct(ctx context.Context, id uuid.UUID, product domain.ProductBase) (domain.ProductBase, error)
	PatchWindow(ctx context.Context, id uuid.UUID, product domain.Window) (domain.Window, error)
	PatchEntranceDoor(ctx context.Context, id uuid.UUID, door domain.EntranceDoor) (domain.EntranceDoor, error)
	PatchInteriorDoor(ctx context.Context, id uuid.UUID, door domain.InteriorDoor) (domain.InteriorDoor, error)
	PatchBalcony(ctx context.Context, id uuid.UUID, balcony domain.Balcony) (domain.Balcony, error)
}

func NewProductsService(productRepo ProductsRepository) *ProductsService {
	return &ProductsService{
		productRepo: productRepo,
	}
}
