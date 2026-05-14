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

func (r *ProductsRepositoryPostgres) PatchProduct(
	ctx context.Context,
	id uuid.UUID,
	product domain.ProductBase,
) (domain.ProductBase, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		UPDATE kudesnik.products
		SET 
			product_name = $2,
			price = $3,
			description = $4,
			is_visible = $5,
			category_code = $6,
			producer_id = $7,
			image_url = $8,
			thumbnail_url = $9,
			version = version + 1
		WHERE product_id = $1 AND version = $10
		RETURNING 
			product_id,
			version,
			product_name,
			price,
			description,
			is_visible,
			category_code,
			producer_id,
			image_url,
			thumbnail_url;
	`

	row := r.pool.QueryRow(
		ctx, query,
		id,
		product.ProductName, product.Price, product.Description,
		product.IsVisible, product.CategoryCode, product.ProducerID,
		product.ImageURL, product.ThumbnailURL,
		product.Version,
	)

	var patchedProduct domain.ProductBase

	err := row.Scan(
		&patchedProduct.ID,
		&patchedProduct.Version,
		&patchedProduct.ProductName,
		&patchedProduct.Price,
		&patchedProduct.Description,
		&patchedProduct.IsVisible,
		&patchedProduct.CategoryCode,
		&patchedProduct.ProducerID,
		&patchedProduct.ImageURL,
		&patchedProduct.ThumbnailURL,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.ProductBase{}, fmt.Errorf(
				"product with ID=%v concurrently accessed: %v: %w",
				id,
				err,
				core_errors.ErrConflict,
			)
		}
		return domain.ProductBase{}, fmt.Errorf("scan product error: %v", err)
	}

	return patchedProduct, nil
}

// Общая функция для обновления базовых полей продукта
func (r *ProductsRepositoryPostgres) updateProductBaseInTx(
	ctx context.Context,
	tx core_postgres_pool.Tx,
	id uuid.UUID,
	product domain.ProductBase,
	currentVersion int,
) (domain.ProductBase, error) {
	query := `
		UPDATE kudesnik.products
		SET 
			product_name = $2,
			price = $3,
			description = $4,
			is_visible = $5,
			category_code = $6,
			producer_id = $7,
			version = version + 1
		WHERE product_id = $1 AND version = $8
		RETURNING 
			product_id,
			version,
			product_name,
			price,
			description,
			is_visible,
			category_code,
			producer_id,
			image_url,
			thumbnail_url;
	`

	var updatedProduct domain.ProductBase
	err := tx.QueryRow(
		ctx, query,
		id,
		product.ProductName,
		product.Price,
		product.Description,
		product.IsVisible,
		product.CategoryCode,
		product.ProducerID,
		currentVersion,
	).Scan(
		&updatedProduct.ID,
		&updatedProduct.Version,
		&updatedProduct.ProductName,
		&updatedProduct.Price,
		&updatedProduct.Description,
		&updatedProduct.IsVisible,
		&updatedProduct.CategoryCode,
		&updatedProduct.ProducerID,
		&updatedProduct.ImageURL,
		&updatedProduct.ThumbnailURL,
	)

	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.ProductBase{}, fmt.Errorf(
				"product with ID=%v concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}
		return domain.ProductBase{}, fmt.Errorf("update product base: %w", err)
	}

	return updatedProduct, nil
}
