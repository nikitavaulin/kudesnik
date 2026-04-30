package product_categories_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *ProductCategoriesService) PatchProductCategory(
	ctx context.Context,
	categoryID uuid.UUID,
	patch domain.ProductCategoryPatch,
) (domain.ProductCategory, error) {
	category, err := s.GetProductCategory(ctx, categoryID)
	if err != nil {
		return domain.ProductCategory{}, fmt.Errorf("failed to get category by ID: %w", err)
	}

	if err := category.ApplyPatch(patch); err != nil {
		return domain.ProductCategory{}, fmt.Errorf("failed to apply category patch: %w", err)
	}

	patchedCategory, err := s.categoriesRepository.PatchProductCategory(ctx, categoryID, category)
	if err != nil {
		return domain.ProductCategory{}, fmt.Errorf("failed to patch category: %w", err)
	}

	return patchedCategory, nil
}
