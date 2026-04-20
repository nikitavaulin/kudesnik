package product_categories_repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func (r *ProductCategoriesRepository) GetProductCategory(ctx context.Context, categoryID uuid.UUID) (domain.ProductCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		SELECT * FROM kudesnik.product_categories
		WHERE product_category_id = $1;
	`

	row := r.pool.QueryRow(ctx, query, categoryID)

	var categoryModel ProductCategoriesModel

	err := row.Scan(
		&categoryModel.ID,
		&categoryModel.CategoryName,
		&categoryModel.InstallationPrice,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ProductCategory{}, fmt.Errorf(
				"category with id=%s: %w",
				categoryID.String(),
				core_errors.ErrNotFound,
			)
		}
		return domain.ProductCategory{}, fmt.Errorf(
			"failed to scan category with id=%s: %w",
			categoryID.String(),
			err,
		)
	}

	category := *domain.NewProductCategory(
		categoryModel.ID,
		categoryModel.CategoryName,
		categoryModel.InstallationPrice,
	)

	return category, nil
}
