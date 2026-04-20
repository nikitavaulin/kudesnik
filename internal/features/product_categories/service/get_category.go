package product_categories_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *ProductCategoriesService) GetProductCategory(ctx context.Context, categoryID uuid.UUID) (domain.ProductCategory, error) {
	category, err := s.categoriesRepository.GetProductCategory(ctx, categoryID)
	if err != nil {
		return domain.ProductCategory{}, fmt.Errorf("failed to get category from repository: %w", err)
	}
	return category, nil
}
