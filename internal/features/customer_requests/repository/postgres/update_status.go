package customer_requests_repository_postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_postgres_pool "github.com/nikitavaulin/kudesnik/internal/core/repository/postgres/pool"
)

func (r *CustomerRequestsRepository) UpdateCustomerRequestStatus(
	ctx context.Context,
	reqID, adminID uuid.UUID,
	status domain.CustomerRequestStatus,
	handledAt time.Time,
) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	var currentVersion int64
	var currentStatus domain.CustomerRequestStatus

	getVersionQuery := `
		SELECT version, status
		FROM kudesnik.customer_requests
		WHERE request_id = $1
	`

	err := r.pool.QueryRow(ctx, getVersionQuery, reqID).Scan(&currentVersion, &currentStatus)
	if err != nil {
		if err == core_postgres_pool.ErrNoRows {
			return fmt.Errorf("customer request with id %s: %w", reqID, core_errors.ErrNotFound)
		}
		return fmt.Errorf("failed to get current version: %w", err)
	}

	var handlerAdminIDParam any = nil
	switch status {
	case domain.CompletedRequestStatus, domain.CancelledRequestStatus:
		handlerAdminIDParam = adminID
	case domain.InProgressRequestStatus:
		handlerAdminIDParam = adminID
	}

	updateQuery := `
		UPDATE kudesnik.customer_requests
		SET status = $1,
		    version = version + 1,
		    handled_at = $2,
		    handler_admin_id = $3
		WHERE request_id = $4 AND version = $5
	`

	result, err := r.pool.Exec(ctx, updateQuery,
		string(status),
		handledAt,
		handlerAdminIDParam,
		reqID,
		currentVersion,
	)
	if err != nil {
		return fmt.Errorf("failed to update request status: %w", err)
	}

	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		return fmt.Errorf("failed to update: request was modified concurrently, please retry: %w", core_errors.ErrConflict)
	}

	return nil
}
