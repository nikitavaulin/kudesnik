package product_categories_service

import (
	"context"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *ProductCategoriesService) GetProductCategories(ctx context.Context, limit, offset *int) ([]domain.ProductCategory, error) {
	categories, err := s.categoriesRepository.GetProductCategories(ctx, limit, offset)
	if err != nil {
		return []domain.ProductCategory{}, fmt.Errorf("failed to get categories from repo: %w", err)
	}
	return categories, nil
}
