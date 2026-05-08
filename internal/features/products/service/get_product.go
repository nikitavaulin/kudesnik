package products_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *ProductsService) GetProduct(ctx context.Context, id uuid.UUID) (domain.BaseProduct, error) {
	product, err := s.productRepo.GetProduct(ctx, id)
	if err != nil {
		return domain.BaseProduct{}, fmt.Errorf("failed to get product from repository: %w", err)
	}
	return product, nil
}
