package products_repository_postges

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func (r *ProductsRepositoryPostgres) GetProducts(ctx context.Context, categoryID *uuid.UUID, limit, offset *int) ([]domain.BaseProduct, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	var query string
	var args []any

	if categoryID != nil {
		query = `
			SELECT * FROM kudesnik.products
			WHERE category_id = $1
			LIMIT $2 OFFSET $3;
		`
		args = []any{categoryID, limit, offset}
	} else {
		query = `
			SELECT * FROM kudesnik.products
			LIMIT $1 OFFSET $2;
		`
		args = []any{limit, offset}
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get products from pool: %w", err)
	}
	defer rows.Close()

	var productsModels []ProductModel
	for rows.Next() {
		var productModel ProductModel

		err := rows.Scan(
			&productModel.ID, &productModel.Version,
			&productModel.ProductName, &productModel.Price, &productModel.Description,
			&productModel.IsVisible, &productModel.CategoryID, &productModel.ProducerID,
		)
		if err != nil {
			return nil, fmt.Errorf("GetProducts from repo: %v: %w", err, core_errors.ErrNotFound)
		}

		productsModels = append(productsModels, productModel)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	products := productsDomainFromModels(productsModels...)
	return products, nil
}
