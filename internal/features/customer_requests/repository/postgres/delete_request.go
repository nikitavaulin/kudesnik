package customer_requests_repository_postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func (r *CustomerRequestsRepository) DeleteCustomerRequest(ctx context.Context, requestID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		DELETE FROM kudesnik.customer_requests
		WHERE request_id = $1;
	`

	result, err := r.pool.Exec(ctx, query, requestID)
	if err != nil {
		return fmt.Errorf("failed to delete customer request with ID=%s: %w", requestID, err)
	}
	if result.RowsAffected() == 0 {
		return fmt.Errorf("request with ID=%s: %w", requestID, core_errors.ErrNotFound)
	}

	return nil
}
