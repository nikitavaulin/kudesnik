package customer_requests_service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func (s *CustomerRequestsService) UpdateCustomerRequestStatus(ctx context.Context, reqID, adminID uuid.UUID, status domain.CustomerRequestStatus) error {
	if !domain.IsCustomerRequestStatus(status) {
		return fmt.Errorf("unknown status '%s': %w", status, core_errors.ErrInvalidArgument)
	}

	handledAt := time.Now()

	return s.requestsRepo.UpdateCustomerRequestStatus(ctx, reqID, adminID, status, handledAt)
}
