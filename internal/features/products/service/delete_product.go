package products_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func (s *ProductsService) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	if err := s.productRepo.DeleteProduct(ctx, id); err != nil {
		return fmt.Errorf("failed to delete product: %w", err)
	}
	return nil
}

func (s *ProductsService) DeleteProducts(ctx context.Context, IDs []uuid.UUID) error {
	if len(IDs) == 0 {
		return fmt.Errorf("IDs list cannot be empty: %w", core_errors.ErrInvalidArgument)
	}

	if err := s.productRepo.DeleteProducts(ctx, IDs); err != nil {
		return fmt.Errorf("failed to delete products: %w", err)
	}

	return nil
}
