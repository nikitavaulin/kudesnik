package products_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func (s *ProductsService) UpdateProductVisability(ctx context.Context, id uuid.UUID, isVisible bool) error {
	if err := s.productRepo.UpdateProductVisability(ctx, id, isVisible); err != nil {
		return fmt.Errorf("failed to update product visibility: %w", err)
	}
	return nil
}

func (s *ProductsService) UpdateProductsVisability(ctx context.Context, IDs []uuid.UUID, isVisible bool) error {
	if len(IDs) == 0 {
		return fmt.Errorf("IDs list cannot be empty: %w", core_errors.ErrInvalidArgument)
	}

	if err := s.productRepo.UpdateProductsVisability(ctx, IDs, isVisible); err != nil {
		return fmt.Errorf("failed to update products visibility: %w", err)
	}

	return nil
}
