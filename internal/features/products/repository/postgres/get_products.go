package products_repository_postges

import (
	"context"
	"fmt"
	"strings"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func (r *ProductsRepositoryPostgres) GetProducts(
	ctx context.Context,
	category *domain.ProductCategoryCode,
	order *string,
	limit, offset *int,
) ([]domain.ProductBaseDetailed, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		SELECT 
			p.product_id,
			p.version,
			p.product_name,
			p.price,
			p.description,
			p.is_visible,
			p.category_code,
			p.producer_id,
			p.image_url,
			p.thumbnail_url,
			pc.product_category_name,
			pc.installation_price,
			pr.company_name
		FROM kudesnik.products p
		LEFT JOIN kudesnik.product_categories pc ON p.category_code = pc.product_category_code
		LEFT JOIN kudesnik.producers pr ON p.producer_id = pr.producer_id
	`

	var args []any
	argCounter := 1
	conditions := []string{}

	if category != nil && *category != "" {
		conditions = append(conditions, fmt.Sprintf("p.category_code = $%d", argCounter))
		args = append(args, *category)
		argCounter++
	}

	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	if order != nil && *order != "" {
		if *order == domain.DescendingOrder {
			query += " ORDER BY p.price DESC"
		} else {
			query += " ORDER BY p.price ASC"
		}
	} else {
		query += " ORDER BY p.price ASC"
	}

	if limit != nil && *limit > 0 {
		query += fmt.Sprintf(" LIMIT $%d", argCounter)
		args = append(args, *limit)
		argCounter++
	} else {
		query += fmt.Sprintf(" LIMIT $%d", argCounter)
		args = append(args, 100)
		argCounter++
	}

	if offset != nil && *offset > 0 {
		query += fmt.Sprintf(" OFFSET $%d", argCounter)
		args = append(args, *offset)
		argCounter++
	}

	rows, err := r.pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get products from pool: %w", err)
	}
	defer rows.Close()

	var products []domain.ProductBaseDetailed
	for rows.Next() {
		var product domain.ProductBaseDetailed

		err := rows.Scan(
			&product.ID, &product.Version,
			&product.ProductName, &product.Price, &product.Description,
			&product.IsVisible, &product.CategoryCode, &product.ProducerID,
			&product.ImageURL, &product.ThumbnailURL,
			&product.CategoryName, &product.InstallationPrice,
			&product.ProducerCompanyName,
		)
		if err != nil {
			return nil, fmt.Errorf("GetProducts from repo: %v: %w", err, core_errors.ErrNotFound)
		}

		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return []domain.ProductBaseDetailed{}, fmt.Errorf("next rows: %w", err)
	}

	return products, nil
}
