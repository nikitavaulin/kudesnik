package customer_requests_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *CustomerRequestsService) CreateCustomerRequest(ctx context.Context, request domain.CustomerRequest, customer domain.Customer) (uuid.UUID, error) {
	if err := request.Validate(); err != nil {
		return uuid.UUID{}, fmt.Errorf("invalid customer request: %w", err)
	}

	requestID, err := s.requestsRepo.CreateCustomerRequest(ctx, request, customer)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to create customer request in repo: %w", err)
	}

	// TODO: EMAIL NOTIFY

	return requestID, nil
}
