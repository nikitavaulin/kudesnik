package admin_repository_postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_postgres_pool "github.com/nikitavaulin/kudesnik/internal/core/repository/postgres/pool"
)

func (r *AdminRepositoryPostgres) PatchAdmin(ctx context.Context, id uuid.UUID, patchedAdmin domain.Admin) (domain.Admin, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		UPDATE kudesnik.admins 
		SET
			email = $1,
			full_name = $2,
			admin_type = $3
		WHERE admin_id = $4
		RETURNING
			admin_id,
			email,
			full_name,
			password_hash,
			admin_type;
	`

	row := r.pool.QueryRow(
		ctx, query,
		patchedAdmin.Email,
		patchedAdmin.FullName,
		patchedAdmin.AdminType,
		id,
	)

	var resultAdmin domain.Admin

	err := row.Scan(
		&resultAdmin.ID,
		&resultAdmin.Email,
		&resultAdmin.FullName,
		&resultAdmin.PasswordHash,
		&resultAdmin.AdminType,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Admin{}, fmt.Errorf(
				"admin with id=%v: %w",
				id,
				core_errors.ErrNotFound,
			)
		}
		return domain.Admin{}, fmt.Errorf(
			"failed to scan admin after patch: id=%v: %w",
			id,
			err,
		)
	}

	return resultAdmin, nil
}
