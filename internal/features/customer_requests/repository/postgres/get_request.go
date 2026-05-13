package customer_requests_repository_postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (r *CustomerRequestsRepository) GetCustomerRequest(ctx context.Context, id uuid.UUID) (domain.CustomerRequestDetailed, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
        SELECT 
            cr.request_id,
			cr.version,
            cr.desired_date,
            cr.desired_time,
            cr.extra_comment,
            cr.status,
            cr.customer_phone_number,
            cr.product_id,
            cr.handled_at,
            cr.created_at,
            cr.handler_admin_id,
            c.customer_name,
            a.full_name as handler_admin_name
        FROM kudesnik.customer_requests cr
        LEFT JOIN kudesnik.customers c ON cr.customer_phone_number = c.customer_phone_number
        LEFT JOIN kudesnik.admins a ON cr.handler_admin_id = a.admin_id
        WHERE cr.request_id = $1
    `

	var detailed domain.CustomerRequestDetailed

	err := r.pool.QueryRow(ctx, query, id).Scan(
		&detailed.ID,
		&detailed.Version,
		&detailed.DesiredDate,
		&detailed.DesiredTime,
		&detailed.ExtraComment,
		&detailed.Status,
		&detailed.CustomerPhoneNumber,
		&detailed.ProductID,
		&detailed.HandledAt,
		&detailed.CreatedAt,
		&detailed.HandlerAdminID,
		&detailed.CustomerFullname,
		&detailed.HandlerAdminName,
	)

	if err != nil {
		return domain.CustomerRequestDetailed{}, fmt.Errorf("failed to get customer request detailed: %w", err)
	}

	return detailed, nil
}
