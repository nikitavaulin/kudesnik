package product_categories_service

import (
	"context"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func (s *ProductCategoriesService) GetProductCategories(ctx context.Context, limit, offset *int) ([]domain.ProductCategory, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf(
			"limit must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf(
			"offset must be non-negative: %w",
			core_errors.ErrInvalidArgument,
		)
	}

	categories, err := s.categoriesRepository.GetProductCategories(ctx, limit, offset)
	if err != nil {
		return []domain.ProductCategory{}, fmt.Errorf("failed to get categories from repo: %w", err)
	}
	return categories, nil
}
