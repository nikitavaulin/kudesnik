package products_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *ProductsService) GetProduct(ctx context.Context, id uuid.UUID, category domain.ProductCategoryCode) (domain.Product, error) {
	product, err := s.getProduct(ctx, id, category)
	if err != nil {
		return nil, fmt.Errorf("failed to get product from repository: %w", err)
	}
	return product, nil
}

func (s *ProductsService) getProduct(ctx context.Context, id uuid.UUID, category domain.ProductCategoryCode) (domain.Product, error) {
	switch category {
	case domain.WindowsCategory:
		window, err := s.productRepo.GetWindow(ctx, id)
		return &window, err

	case domain.EntranceDoorsCategory:
		door, err := s.productRepo.GetEntranceDoor(ctx, id)
		return &door, err

	case domain.InteriorDoorsCategory:
		door, err := s.productRepo.GetInteriorDoor(ctx, id)
		return &door, err

	default:
		product, err := s.productRepo.GetProduct(ctx, id)
		return &product, err
	}
}
