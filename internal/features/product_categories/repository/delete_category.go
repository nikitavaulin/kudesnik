package product_categories_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func (r *ProductCategoriesRepository) DeleteProductCategory(ctx context.Context, categoryID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		DELETE FROM kudesnik.product_categories
		WHERE product_category_id = $1;
	`

	cmdTag, err := r.pool.Exec(ctx, query, categoryID)
	if err != nil {
		return fmt.Errorf("exec delete query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf(
			"category with id=%s: %w",
			categoryID.String(),
			core_errors.ErrNotFound,
		)
	}

	return nil
}
