package product_categories_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_postgres_pool "github.com/nikitavaulin/kudesnik/internal/core/repository/postgres/pool"
)

func (r *ProductCategoriesRepository) GetProductCategory(ctx context.Context, categoryCode domain.ProductCategoryCode) (domain.ProductCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		SELECT * FROM kudesnik.product_categories
		WHERE product_category_code = $1;
	`

	row := r.pool.QueryRow(ctx, query, categoryCode)

	var categoryModel ProductCategoriesModel

	err := row.Scan(
		&categoryModel.Code,
		&categoryModel.CategoryName,
		&categoryModel.InstallationPrice,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.ProductCategory{}, fmt.Errorf(
				"category with code=%s: %w",
				categoryCode,
				core_errors.ErrNotFound,
			)
		}
		return domain.ProductCategory{}, fmt.Errorf(
			"failed to scan category with code=%s: %w",
			categoryCode,
			err,
		)
	}

	category := *domain.NewProductCategory(
		categoryModel.Code,
		categoryModel.CategoryName,
		categoryModel.InstallationPrice,
	)

	return category, nil
}
