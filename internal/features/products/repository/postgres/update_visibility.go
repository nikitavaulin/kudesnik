package products_repository_postges

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func (r *ProductsRepositoryPostgres) UpdateProductVisability(ctx context.Context, id uuid.UUID, isVisible bool) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		UPDATE kudesnik.products
		SET is_visible = $1
		WHERE product_id = $2;
	`

	cmdTag, err := r.pool.Exec(ctx, query, isVisible, id)
	if err != nil {
		return fmt.Errorf("exec update visibility query: %w", err)
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

func (r *ProductsRepositoryPostgres) UpdateProductsVisability(ctx context.Context, IDs []uuid.UUID, isVisible bool) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		UPDATE kudesnik.products
		SET is_visible = $1
		WHERE product_id = ANY($2);
	`

	cmdTag, err := r.pool.Exec(ctx, query, isVisible, IDs)
	if err != nil {
		return fmt.Errorf("exec update visibility query: %w", err)
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
