package customer_requests_service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (s *CustomerRequestsService) PatchCustomerRequest(ctx context.Context, requestID uuid.UUID, patch domain.CustomerRequestPatch) error {
	requestDetailed, err := s.requestsRepo.GetCustomerRequest(ctx, requestID)
	if err != nil {
		return fmt.Errorf("failed to get customer request from repo: %w", err)
	}
	request := requestDetailed.CustomerRequest

	if err := patch.Validate(); err != nil {
		return fmt.Errorf("invalid patch: %w", err)
	}

	if err := request.ApplyPatch(patch); err != nil {
		return fmt.Errorf("failed to apply patch: %w", err)
	}

	if err := s.requestsRepo.PatchCustomerRequest(ctx, requestID, request); err != nil {
		return fmt.Errorf("failed to patch in repo: %w", err)
	}

	return nil
}
