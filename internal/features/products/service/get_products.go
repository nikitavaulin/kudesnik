package products_service

import (
	"context"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_validation "github.com/nikitavaulin/kudesnik/internal/core/tools/validation"
)

func (s *ProductsService) GetProducts(ctx context.Context, category *domain.ProductCategoryCode, order *string, limit, offset *int) ([]domain.ProductBaseDetailed, error) {
	if err := core_validation.ValidateLimitOffset(limit, offset); err != nil {
		return nil, fmt.Errorf(
			"wrong limit/offset: %v, %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	products, err := s.productRepo.GetProducts(ctx, category, order, limit, offset)
	if err != nil {
		return []domain.ProductBaseDetailed{}, fmt.Errorf("failed to get products from repo: %w", err)
	}

	return products, nil
}
