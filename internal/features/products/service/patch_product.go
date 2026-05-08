package products_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *ProductsService) PatchProduct(ctx context.Context, id uuid.UUID, patch domain.ProductPatch) (domain.BaseProduct, error) {
	product, err := s.productRepo.GetProduct(ctx, id)
	if err != nil {
		return domain.BaseProduct{}, fmt.Errorf("failed to get product from repo: %w", err)
	}

	if err := product.ApplyPatch(patch); err != nil {
		return domain.BaseProduct{}, fmt.Errorf("failed to apply product patch: %w", err)
	}

	patchedProduct, err := s.productRepo.PatchProduct(ctx, id, product)
	if err != nil {
		return domain.BaseProduct{}, fmt.Errorf("failed to patch product in repo: %w", err)
	}

	return patchedProduct, nil

}
