package product_categories_service

import (
	"context"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_validation "github.com/nikitavaulin/kudesnik/internal/core/tools/validation"
)

func (s *ProductCategoriesService) GetProductCategories(ctx context.Context, limit, offset *int) ([]domain.ProductCategory, error) {
	if err := core_validation.ValidateLimitOffset(limit, offset); err != nil {
		return nil, fmt.Errorf(
			"wrong limit/offset: %v, %w",
			err,
			core_errors.ErrInvalidArgument,
		)
	}

	categories, err := s.categoriesRepository.GetProductCategories(ctx, limit, offset)
	if err != nil {
		return []domain.ProductCategory{}, fmt.Errorf("failed to get categories from repo: %w", err)
	}

	return categories, nil
}
