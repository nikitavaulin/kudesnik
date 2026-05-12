package admin_repository_postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func (r *AdminRepositoryPostgres) UpdateAdminType(ctx context.Context, id uuid.UUID, adminType domain.Role) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
        UPDATE kudesnik.admins
        SET admin_type = $1
        WHERE admin_id = $2;
    `

	result, err := r.pool.Exec(ctx, query, adminType, id)
	if err != nil {
		return fmt.Errorf("failed to update admin type: %w", err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("admin with ID = %v: %w", id, core_errors.ErrNotFound)
	}

	return nil
}
