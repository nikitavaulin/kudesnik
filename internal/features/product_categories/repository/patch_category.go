package product_categories_repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_postgres_pool "github.com/nikitavaulin/kudesnik/internal/core/repository/postgres/pool"
)

func (r *ProductCategoriesRepository) PatchProductCategory(
	ctx context.Context,
	categoryCode domain.ProductCategoryCode,
	category domain.ProductCategory,
) (domain.ProductCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		UPDATE kudesnik.product_categories 
		SET
			product_category_name = $1,
			installation_price = $2
		WHERE product_category_code = $3
		RETURNING
			product_category_code,
			product_category_name,
			installation_price;
	`

	row := r.pool.QueryRow(
		ctx, query,
		category.CategoryName, category.InstallationPrice,
		categoryCode,
	)

	var categoryModel ProductCategoriesModel

	err := row.Scan(
		&categoryModel.Code,
		&categoryModel.CategoryName,
		&categoryModel.InstallationPrice,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.ProductCategory{}, fmt.Errorf(
				"category with code=%v: %v: %w",
				categoryCode,
				err,
				core_errors.ErrNotFound,
			)
		}
		return domain.ProductCategory{}, fmt.Errorf(
			"scan category error: %v: %v",
			categoryCode,
			err,
		)
	}

	patchedCategory := domainFromModel(categoryModel)
	return patchedCategory, nil
}
