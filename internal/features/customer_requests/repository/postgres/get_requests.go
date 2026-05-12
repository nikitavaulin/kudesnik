package customer_requests_repository_postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (r *CustomerRequestsRepository) GetCustomerRequests(
	ctx context.Context,
	status *domain.CustomerRequestStatus,
	phone *string,
	adminID *uuid.UUID,
	limit, offset *int,
) ([]domain.CustomerRequestForList, error) {
	query := `
		SELECT 
			cr.request_id,
			cr.desired_date,
			cr.desired_time,
			cr.status,
			cr.customer_phone_number,
			c.customer_name as customer_fullname,
			cr.created_at,
			cr.handler_admin_id,
			a.full_name as handler_admin_name,
			cr.handled_at
		FROM kudesnik.customer_requests cr
		LEFT JOIN kudesnik.customers c ON cr.customer_phone_number = c.customer_phone_number
		LEFT JOIN kudesnik.admins a ON cr.handler_admin_id = a.admin_id
		WHERE 1=1
	`

	var args []interface{}
	argCounter := 1

	if status != nil {
		query += fmt.Sprintf(" AND cr.status = $%d", argCounter)
		args = append(args, string(*status))
		argCounter++
	}

	if phone != nil && *phone != "" {
		query += fmt.Sprintf(" AND cr.customer_phone_number = $%d", argCounter)
		args = append(args, *phone)
		argCounter++
	}

	if adminID != nil {
		query += fmt.Sprintf(" AND cr.handler_admin_id = $%d", argCounter)
		args = append(args, *adminID)
		argCounter++
	}

	query += " ORDER BY cr.created_at DESC"

	if limit != nil && *limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argCounter)
		args = append(args, *limit)
		argCounter++
	}

	if offset != nil && *offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", argCounter)
		args = append(args, *offset)
		argCounter++
	}

	query += ";"

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query customer requests: %w", err)
	}
	defer rows.Close()

	var requests []domain.CustomerRequestForList

	for rows.Next() {
		var req domain.CustomerRequestForList

		err := rows.Scan(
			&req.ID,
			&req.DesiredDate,
			&req.DesiredTime,
			&req.Status,
			&req.CustomerPhoneNumber,
			&req.CustomerFullname,
			&req.CreatedAt,
			&req.HandlerAdminID,
			&req.HandlerAdminName,
			&req.HandledAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan customer request: %w", err)
		}

		requests = append(requests, req)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	return requests, nil
}
