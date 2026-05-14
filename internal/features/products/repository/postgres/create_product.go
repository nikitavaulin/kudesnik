package products_repository_postges

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_postgres_pool "github.com/nikitavaulin/kudesnik/internal/core/repository/postgres/pool"
)

func (r *ProductsRepositoryPostgres) CreateProduct(ctx context.Context, product domain.ProductBase) (domain.ProductBase, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		INSERT INTO kudesnik.products (
			product_id, product_name, price, description, is_visible, category_code, producer_id
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING
			product_id, version, product_name, price, description, is_visible, category_code, producer_id;
	`

	row := r.pool.QueryRow(
		ctx, query,
		uuid.New(), product.ProductName, product.Price,
		product.Description, product.IsVisible,
		product.CategoryCode, product.ProducerID,
	)

	var model ProductModel
	err := row.Scan(
		&model.ID, &model.Version,
		&model.ProductName, &model.Price, &model.Description,
		&model.IsVisible, &model.CategoryCode, &model.ProducerID,
	)
	if err != nil {
		return domain.ProductBase{}, fmt.Errorf("CreateProduct in Repo: %w", err)
	}

	product = productDomainFromModel(model)
	return product, nil
}

func (r *ProductsRepositoryPostgres) createProductInTx(ctx context.Context, tx core_postgres_pool.Tx, product domain.ProductBase) (uuid.UUID, int, error) {
	productQuery := `
		INSERT INTO kudesnik.products (
			product_id, product_name, price, description, is_visible, category_code, producer_id
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING product_id, version;
	`

	var productID uuid.UUID
	var version int

	err := tx.QueryRow(
		ctx, productQuery,
		uuid.New(), product.ProductName, product.Price,
		product.Description, product.IsVisible,
		product.CategoryCode, product.ProducerID,
	).Scan(&productID, &version)

	return productID, version, err
}
