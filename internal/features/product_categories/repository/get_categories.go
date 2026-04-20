package product_categories_repository

import (
	"context"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (r *ProductCategoriesRepository) GetProductCategories(ctx context.Context, limit, offset *int) ([]domain.ProductCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		SELECT * FROM kudesnik.product_categories
		LIMIT $1 OFFSET $2;
	`

	rows, err := r.pool.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to get categories from pool: %w", err)
	}
	defer rows.Close()

	var categoriesModels []ProductCategoriesModel
	for rows.Next() {
		var categoryModel ProductCategoriesModel

		err := rows.Scan(
			&categoryModel.ID,
			&categoryModel.CategoryName,
			&categoryModel.InstallationPrice,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan categories from pool: %w", err)
		}

		categoriesModels = append(categoriesModels, categoryModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	categories := domainsFromModels(categoriesModels)

	return categories, nil
}
