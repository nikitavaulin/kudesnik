package product_categories_service

import (
	"context"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *ProductCategoriesService) PatchProductCategory(
	ctx context.Context,
	categoryCode domain.ProductCategoryCode,
	patch domain.ProductCategoryPatch,
) (domain.ProductCategory, error) {
	category, err := s.GetProductCategory(ctx, categoryCode)
	if err != nil {
		return domain.ProductCategory{}, fmt.Errorf("failed to get category: %w", err)
	}

	if err := category.ApplyPatch(patch); err != nil {
		return domain.ProductCategory{}, fmt.Errorf("failed to apply category patch: %w", err)
	}

	patchedCategory, err := s.categoriesRepository.PatchProductCategory(ctx, categoryCode, category)
	if err != nil {
		return domain.ProductCategory{}, fmt.Errorf("failed to patch category: %w", err)
	}

	return patchedCategory, nil
}
