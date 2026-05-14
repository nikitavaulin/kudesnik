package customer_requests_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

func (s *CustomerRequestsService) DeleteCustomerRequest(ctx context.Context, requestID uuid.UUID) error {
	if err := s.requestsRepo.DeleteCustomerRequest(ctx, requestID); err != nil {
		return fmt.Errorf("failed to delete in repo: %w", err)
	}
	return nil
}
