package products_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_validation "github.com/nikitavaulin/kudesnik/internal/core/tools/validation"
)

func (s *ProductsService) GetProducts(ctx context.Context, categoryID *uuid.UUID, limit, offset *int) ([]domain.ProductBase, error) {
	if err := core_validation.ValidateLimitOffset(limit, offset); err != nil {
		return nil, fmt.Errorf(
			"wrong limit/offset: %v, %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	products, err := s.productRepo.GetProducts(ctx, categoryID, limit, offset)
	if err != nil {
		return []domain.ProductBase{}, fmt.Errorf("failed to get products from repo: %w", err)
	}

	return products, nil
}
