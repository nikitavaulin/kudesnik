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

	var productModel ProductModel

	err := row.Scan(
		&productModel.ID, &productModel.Version,
		&productModel.ProductName, &productModel.Price, &productModel.Description,
		&productModel.IsVisible, &productModel.CategoryCode, &productModel.ProducerID,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.ProductBase{}, fmt.Errorf("GetProduct from repo: %v: %w", err, core_errors.ErrNotFound)
		}
		return domain.ProductBase{}, fmt.Errorf("GetProduct from repo: %w", err)
	}

	product := productDomainFromModel(productModel)
	return product, nil
}

func (r *ProductsRepositoryPostgres) GetProductDetailed(ctx context.Context, id uuid.UUID) (domain.ProductBaseDetailed, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		SELECT 
			p.product_name as product_name,
			p.price,
			c.product_category_name as category_name,
			pc.company_name as producer_company_name
		FROM kudesnik.products p
		LEFT JOIN kudesnik.product_categories c ON p.category_code = c.product_category_code
		LEFT JOIN kudesnik.producers pc ON p.producer_id = pc.producer_id
		WHERE p.product_id = $1
	`

	var product domain.ProductBaseDetailed

	err := r.pool.QueryRow(ctx, query, id).Scan(
		&product.ProductName,
		&product.Price,
		&product.CategoryName,
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
