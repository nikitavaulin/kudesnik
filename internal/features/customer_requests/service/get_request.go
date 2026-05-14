package customer_requests_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *CustomerRequestsService) GetCustomerRequest(ctx context.Context, id uuid.UUID) (domain.CustomerRequestDetailed, error) {
	request, err := s.requestsRepo.GetCustomerRequest(ctx, id)
	if err != nil {
		return domain.CustomerRequestDetailed{}, fmt.Errorf("failed to get customer request from repo: %w", err)
	}

	if request.ProductID != nil {
		product, err := s.productsRepo.GetProductDetailed(ctx, *request.ProductID)
		if err != nil {
			return domain.CustomerRequestDetailed{}, fmt.Errorf("failed to get chosen product from repo: %w", err)
		}
		request.ChosenProduct = &product
	}

	return request, nil
}
