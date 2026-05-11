package admin_repository_postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func (r *AdminRepositoryPostgres) GetAdminByEmail(ctx context.Context, email string) (domain.Admin, error) {
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
		WHERE email = $1
	`

	var admin domain.Admin
	err := r.pool.QueryRow(ctx, query, email).Scan(
		&admin.ID,
		&admin.Email,
		&admin.FullName,
		&admin.PasswordHash,
		&admin.AdminType,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Admin{}, fmt.Errorf("admin with email %s: %w", email, core_errors.ErrNotFound)
		}
		return domain.Admin{}, fmt.Errorf("failed to get admin by email: %w", err)
	}

	return admin, nil
}

func (r *AdminRepositoryPostgres) GetAdminByID(ctx context.Context, id uuid.UUID) (domain.Admin, error) {
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
		WHERE admin_id = $1
	`

	var admin domain.Admin
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&admin.ID,
		&admin.Email,
		&admin.FullName,
		&admin.PasswordHash,
		&admin.AdminType,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Admin{}, fmt.Errorf("admin with ID = %v: %w", id, core_errors.ErrNotFound)
		}
		return domain.Admin{}, fmt.Errorf("failed to get admin by ID: %w", err)
	}

	return admin, nil
}
