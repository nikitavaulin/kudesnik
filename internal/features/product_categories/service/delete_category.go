package product_categories_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *ProductCategoriesService) DeleteProductCategory(ctx context.Context, categoryID uuid.UUID) error {
	if err := s.categoriesRepository.DeleteProductCategory(ctx, categoryID); err != nil {
		return fmt.Errorf("failed to delete category: %w", err)
	}
	return nil
}
