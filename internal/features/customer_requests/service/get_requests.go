package customer_requests_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *CustomerRequestsService) GetCustomerRequests(
	ctx context.Context,
	status *domain.CustomerRequestStatus,
	phone *string,
	adminID *uuid.UUID,
	limit, offset *int,
) ([]domain.CustomerRequestForList, error) {
	requests, err := s.requestsRepo.GetCustomerRequests(ctx, status, phone, adminID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get requests from repo: %w", err)
	}
	return requests, nil
}
