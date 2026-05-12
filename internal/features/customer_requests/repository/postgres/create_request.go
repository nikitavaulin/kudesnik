package customer_requests_repository_postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (r *CustomerRequestsRepository) CreateCustomerRequest(ctx context.Context, request domain.CustomerRequest, customer domain.Customer) (uuid.UUID, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	createCustomerQuery := `
        INSERT INTO kudesnik.customers (customer_phone_number, customer_name)
        VALUES ($1, $2)
        ON CONFLICT (customer_phone_number) DO NOTHING
    `

	_, err = tx.Exec(ctx, createCustomerQuery,
		customer.PhoneNumber,
		customer.FullName,
	)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to create customer: %w", err)
	}

	createRequestQuery := `
        INSERT INTO kudesnik.customer_requests (
            request_id, desired_date, desired_time,
            extra_comment, customer_phone_number, created_at,
            status, product_id
        ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
    `

	_, err = tx.Exec(ctx, createRequestQuery,
		request.ID,
		request.DesiredDate,
		request.DesiredTime,
		request.ExtraComment,
		request.CustomerPhoneNumber,
		request.CreatedAt,
		request.Status,
		request.ProductID,
	)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to create customer request: %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return uuid.UUID{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return request.ID, nil
}
