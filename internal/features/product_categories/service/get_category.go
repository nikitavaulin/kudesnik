package product_categories_service

import (
	"context"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *ProductCategoriesService) GetProductCategory(ctx context.Context, categoryCode domain.ProductCategoryCode) (domain.ProductCategory, error) {
	category, err := s.categoriesRepository.GetProductCategory(ctx, categoryCode)
	if err != nil {
		return domain.ProductCategory{}, fmt.Errorf("failed to get category from repository: %w", err)
	}
	return category, nil
}
