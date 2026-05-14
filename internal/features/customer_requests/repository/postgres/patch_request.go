package customer_requests_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_postgres_pool "github.com/nikitavaulin/kudesnik/internal/core/repository/postgres/pool"
)

func (r *CustomerRequestsRepository) PatchCustomerRequest(ctx context.Context, requestID uuid.UUID, request domain.CustomerRequest) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		UPDATE kudesnik.customer_requests 
		SET 
			desired_date = $1,
			desired_time = $2,
			extra_comment = $3,
			version = version + 1
		WHERE request_id = $4 AND version = $5
		RETURNING version
	`

	var newVersion int
	err := r.pool.QueryRow(ctx, query,
		request.DesiredDate,
		request.DesiredTime,
		request.ExtraComment,
		requestID,
		request.Version,
	).Scan(&newVersion)

	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return fmt.Errorf("customer request with id %s not found: %w", requestID, core_errors.ErrNotFound)
		}
		return fmt.Errorf("customer request version mismatch: expected %d: %w", request.Version, core_errors.ErrConflict)
	}

	return nil
}
