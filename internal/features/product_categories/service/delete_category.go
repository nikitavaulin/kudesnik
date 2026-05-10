package product_categories_service

import (
	"context"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *ProductCategoriesService) DeleteProductCategory(ctx context.Context, categoryCode domain.ProductCategoryCode) error {
	if err := s.categoriesRepository.DeleteProductCategory(ctx, categoryCode); err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}
	return nil
}
