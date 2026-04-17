package product_categories_service

import (
	"context"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *ProductCategoriesService) CreateProductCategory(ctx context.Context, category domain.ProductCategory) (domain.ProductCategory, error) {
	if err := category.Validate(); err != nil {
		return domain.ProductCategory{}, fmt.Errorf("validate product category domain: %w", err)
	}

	category, err := s.categoryRepository.CreateProductCategory(ctx, category)
	if err != nil {
		return domain.ProductCategory{}, fmt.Errorf("create product category: %w", err)
	}

	return category, nil // with init id

}
