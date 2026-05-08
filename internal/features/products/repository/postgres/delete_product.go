package products_repository_postges

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func (r *ProductsRepositoryPostgres) DeleteProduct(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		DELETE FROM kudesnik.products
		WHERE product_id = $1;
	`

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("exec delete query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf(
			"product with id=%s: %w",
			id.String(),
			core_errors.ErrNotFound,
		)
	}

	return nil
}

func (r *ProductsRepositoryPostgres) DeleteProducts(ctx context.Context, IDs []uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		DELETE FROM kudesnik.products
		WHERE product_id = ANY($1);
	`

	cmdTag, err := r.pool.Exec(ctx, query, IDs)
	if err != nil {
		return fmt.Errorf("exec delete query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf(
			"products with IDs=%v: %w",
			IDs,
			core_errors.ErrNotFound,
		)
	}

	return nil
}
