package admin_repository_postgres

import (
	"context"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (r *AdminRepositoryPostgres) CreateAdmin(ctx context.Context, admin domain.Admin) (domain.Admin, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
        INSERT INTO kudesnik.admins (admin_id, email, full_name, password_hash, admin_type)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING admin_id, email, full_name, admin_type
    `

	var createdAdmin domain.Admin
	err := r.pool.QueryRow(ctx, query,
		admin.ID,
		admin.Email,
		admin.FullName,
		admin.PasswordHash,
		admin.AdminType,
	).Scan(&admin.ID, &createdAdmin.Email, &createdAdmin.FullName, &createdAdmin.AdminType)

	if err != nil {
		return domain.Admin{}, fmt.Errorf("failed to create admin: %w", err)
	}

	return createdAdmin, nil
}
