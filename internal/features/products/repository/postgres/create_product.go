package products_repository_postges

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
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

func (r *ProductsRepositoryPostgres) CreateWindow(ctx context.Context, window domain.Window) (domain.Window, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return domain.Window{}, fmt.Errorf("CreateWindow in Repo (begin tx): %w", err)
	}
	defer tx.Rollback(ctx)

	productQuery := `
		INSERT INTO kudesnik.products (
			product_id, product_name, price, description, is_visible, category_code, producer_id
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING product_id, version;
	`

	var productID uuid.UUID
	var version int

	err = tx.QueryRow(
		ctx, productQuery,
		uuid.New(), window.ProductName, window.Price,
		window.Description, window.IsVisible,
		window.CategoryCode, window.ProducerID,
	).Scan(&productID, &version)

	if err != nil {
		return domain.Window{}, fmt.Errorf("CreateWindow in Repo (insert product): %w", err)
	}

	windowQuery := `
		INSERT INTO kudesnik.windows (
			window_id, purpose, width, height, material
		) VALUES ($1, $2, $3, $4, $5);
	`

	_, err = tx.Exec(
		ctx, windowQuery,
		productID, window.Purpose, window.Width,
		window.Height, window.Material,
	)

	if err != nil {
		return domain.Window{}, fmt.Errorf("CreateWindow in Repo (insert window): %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return domain.Window{}, fmt.Errorf("CreateWindow in Repo (commit): %w", err)
	}

	window.ID = productID
	window.Version = version

	return window, nil
}
