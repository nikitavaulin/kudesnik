package admin_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (r *AdminRepositoryPostgres) GetAdmins(ctx context.Context, adminType *domain.Role) ([]domain.Admin, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
        SELECT 
            admin_id,
            email,
            full_name,
            password_hash,
            admin_type
        FROM kudesnik.admins
    `

	args := []interface{}{}
	argPos := 1

	if adminType != nil {
		query += " WHERE admin_type = $" + fmt.Sprint(argPos)
		args = append(args, *adminType)
		argPos++
	}

	query += `;`

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get admins from db: %w", err)
	}
	defer rows.Close()

	var admins []domain.Admin

	for rows.Next() {
		var admin domain.Admin
		err := rows.Scan(
			&admin.ID,
			&admin.Email,
			&admin.FullName,
			&admin.PasswordHash,
			&admin.AdminType,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan admin row: %w", err)
		}
		admins = append(admins, admin)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating admin rows: %w", err)
	}

	return admins, nil
}
