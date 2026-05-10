package products_service

import (
	"context"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *ProductsService) CreateProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	if err := product.Validate(); err != nil {
		return nil, fmt.Errorf("create product validation: %w", err)
	}

	product, err := s.createProduct(ctx, product)
	if err != nil {
		return nil, fmt.Errorf("failed to create product in repository : %w", err)
	}

	return product, nil
}

func (s *ProductsService) createProduct(ctx context.Context, product domain.Product) (domain.Product, error) {
	switch p := product.(type) {
	case *domain.Window:
		window, err := s.productRepo.CreateWindow(ctx, *p)
		return &window, err

	case *domain.ProductBase:
		productBase, err := s.productRepo.CreateProduct(ctx, *p)
		return &productBase, err

	case *domain.EntranceDoor:
		door, err := s.productRepo.CreateEntranceDoor(ctx, *p)
		return &door, err

	case *domain.InteriorDoor:
		door, err := s.productRepo.CreateInteriorDoor(ctx, *p)
		return &door, err

	default:
		return nil, fmt.Errorf("unknown product category")
	}
}
