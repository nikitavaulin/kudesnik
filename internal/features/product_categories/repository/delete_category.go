package product_categories_repository

import (
	"context"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func (r *ProductCategoriesRepository) DeleteProductCategory(ctx context.Context, categoryCode domain.ProductCategoryCode) error {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		DELETE FROM kudesnik.product_categories
		WHERE product_category_code = $1;
	`

	cmdTag, err := r.pool.Exec(ctx, query, categoryCode)
	if err != nil {
		return fmt.Errorf("exec delete query: %w", err)
	}
	if cmdTag.RowsAffected() == 0 {
		return fmt.Errorf(
			"category with code=%s: %w",
			categoryCode,
			core_errors.ErrNotFound,
		)
	}

	return nil
}
