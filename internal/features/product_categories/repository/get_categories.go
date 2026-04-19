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

	rows, err := r.pool.Query(ctx, query, *limit, *offset)
	if err != nil {
		return []domain.ProductCategory{}, fmt.Errorf("failed to get categories from pool: %w", err)
	}
	defer rows.Close()

	var categoriesModels []ProductCategoriesModel
	if err := rows.Scan(&categoriesModels); err != nil {
		return []domain.ProductCategory{}, fmt.Errorf("failed to scan categories from rows: %w", err)
	}

	return nil, nil
}

// func domainsFromModels(categoriesModels []ProductCategoriesModel) []domain.ProductCategory {
// 	categories := make([]domain.ProductCategory, len(categoriesModels))
// 	for i, model := range categoriesModels {
// 		categories[i] = *domain.NewProductCategory(model)
// 	}
// }
