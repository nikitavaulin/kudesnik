package products_repository_postges

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
	core_postgres_pool "github.com/nikitavaulin/kudesnik/internal/core/repository/postgres/pool"
)

func (r *ProductsRepositoryPostgres) GetProduct(ctx context.Context, id uuid.UUID) (domain.ProductBase, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		SELECT * FROM kudesnik.products
		WHERE product_id = $1;
	`

	row := r.pool.QueryRow(ctx, query, id)

	var product domain.ProductBase

	err := row.Scan(
		&product.ID, &product.Version,
		&product.ProductName, &product.Price, &product.Description,
		&product.IsVisible, &product.CategoryCode, &product.ProducerID,
		&product.ImageURL, &product.ThumbnailURL,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.ProductBase{}, fmt.Errorf("GetProduct from repo: %v: %w", err, core_errors.ErrNotFound)
		}
		return domain.ProductBase{}, fmt.Errorf("GetProduct from repo: %w", err)
	}

	return product, nil
}

func (r *ProductsRepositoryPostgres) GetProductDetailed(ctx context.Context, id uuid.UUID) (domain.ProductBaseDetailed, error) {
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
		WHERE p.product_id = $1
	`

	var product domain.ProductBaseDetailed

	err := r.pool.QueryRow(ctx, query, id).Scan(
		&product.ID, &product.Version,
		&product.ProductName, &product.Price, &product.Description,
		&product.IsVisible, &product.CategoryCode, &product.ProducerID,
		&product.ImageURL, &product.ThumbnailURL,
		&product.CategoryName, &product.InstallationPrice,
		&product.ProducerCompanyName,
	)

	if err != nil {
		if err == core_postgres_pool.ErrNoRows {
			return domain.ProductBaseDetailed{}, fmt.Errorf("product with id %s not found: %w", id, core_errors.ErrNotFound)
		}
		return domain.ProductBaseDetailed{}, fmt.Errorf("failed to get product detailed: %w", err)
	}

	return product, nil
}

func (r *ProductsRepositoryPostgres) GetProductDetails(ctx context.Context, id uuid.UUID) (domain.ProductDetails, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		SELECT 
			pc.product_category_name,
			pc.installation_price,
			pr.company_name
		FROM kudesnik.products p
		LEFT JOIN kudesnik.product_categories pc ON p.category_code = pc.product_category_code
		LEFT JOIN kudesnik.producers pr ON p.producer_id = pr.producer_id
		WHERE p.product_id = $1
	`

	var product domain.ProductDetails

	err := r.pool.QueryRow(ctx, query, id).Scan(
		&product.CategoryName, &product.InstallationPrice,
		&product.ProducerCompanyName,
	)

	if err != nil {
		if err == core_postgres_pool.ErrNoRows {
			return domain.ProductDetails{}, fmt.Errorf("product with id %s not found: %w", id, core_errors.ErrNotFound)
		}
		return domain.ProductDetails{}, fmt.Errorf("failed to get product details: %w", err)
	}

	return product, nil
}
