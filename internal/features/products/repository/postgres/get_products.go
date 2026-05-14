package products_repository_postges

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func (r *ProductsRepositoryPostgres) GetProducts(ctx context.Context, categoryID *uuid.UUID, limit, offset *int) ([]domain.ProductBase, error) {
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

	var products []domain.ProductBase
	for rows.Next() {
		var product domain.ProductBase

		err := rows.Scan(
			&product.ID, &product.Version,
			&product.ProductName, &product.Price, &product.Description,
			&product.IsVisible, &product.CategoryCode, &product.ProducerID,
			&product.ImageURL, &product.ThumbnailURL,
		)
		if err != nil {
			return nil, fmt.Errorf("GetProducts from repo: %v: %w", err, core_errors.ErrNotFound)
		}

		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("next rows: %w", err)
	}

	return products, nil
}
