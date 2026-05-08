package products_service

import (
	"context"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *ProductsService) CreateProduct(ctx context.Context, product domain.BaseProduct) (domain.BaseProduct, error) {
	if err := product.Validate(); err != nil {
		return domain.BaseProduct{}, fmt.Errorf("create product validation: %w", err)
	}

	product, err := s.productRepo.CreateProduct(ctx, product)
	if err != nil {
		return domain.BaseProduct{}, fmt.Errorf("failed to create product in repository : %w", err)
	}

	return product, nil
}
