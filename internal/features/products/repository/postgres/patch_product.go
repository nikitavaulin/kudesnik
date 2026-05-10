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
			producer_id;
	`

	row := r.pool.QueryRow(
		ctx, query,
		id,
		product.ProductName, product.Price, product.Description,
		product.IsVisible, product.CategoryCode, product.ProducerID,
		product.Version,
	)

	var productModel ProductModel

	err := row.Scan(
		&productModel.ID,
		&productModel.Version,
		&productModel.ProductName,
		&productModel.Price,
		&productModel.Description,
		&productModel.IsVisible,
		&productModel.CategoryCode,
		&productModel.ProducerID,
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

	patchedProduct := productDomainFromModel(productModel)
	return patchedProduct, nil
}

func (r *ProductsRepositoryPostgres) PatchWindow(
	ctx context.Context,
	id uuid.UUID,
	window domain.Window,
) (domain.Window, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return domain.Window{}, fmt.Errorf("PatchWindow in Repo (begin tx): %w", err)
	}
	defer tx.Rollback(ctx)

	updateProductsQuery := `
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
			producer_id;
	`

	var productBase domain.ProductBase
	err = tx.QueryRow(
		ctx, updateProductsQuery,
		id,
		window.ProductName,
		window.Price,
		window.Description,
		window.IsVisible,
		window.CategoryCode,
		window.ProducerID,
		window.Version,
	).Scan(
		&productBase.ID,
		&productBase.Version,
		&productBase.ProductName,
		&productBase.Price,
		&productBase.Description,
		&productBase.IsVisible,
		&productBase.CategoryCode,
		&productBase.ProducerID,
	)

	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Window{}, fmt.Errorf(
				"window with ID=%v concurrently accessed: %w",
				id,
				core_errors.ErrConflict,
			)
		}
		return domain.Window{}, fmt.Errorf("PatchWindow in Repo (update products): %w", err)
	}

	updateWindowsQuery := `
		UPDATE kudesnik.windows
		SET 
			purpose = $2,
			width = $3,
			height = $4,
			material = $5
		WHERE window_id = $1
		RETURNING 
			purpose,
			width,
			height,
			material;
	`

	var windowSpecific struct {
		Purpose  string
		Width    int
		Height   int
		Material string
	}
	err = tx.QueryRow(
		ctx, updateWindowsQuery,
		id,
		window.Purpose,
		window.Width,
		window.Height,
		window.Material,
	).Scan(
		&windowSpecific.Purpose,
		&windowSpecific.Width,
		&windowSpecific.Height,
		&windowSpecific.Material,
	)
	if err != nil {
		return domain.Window{}, fmt.Errorf("PatchWindow in Repo (update windows): %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return domain.Window{}, fmt.Errorf("PatchWindow in Repo (commit): %w", err)
	}

	// Формируем результат из полученных данных
	patchedWindow := domain.Window{
		ProductBase: productBase,
		Purpose:     windowSpecific.Purpose,
		Width:       windowSpecific.Width,
		Height:      windowSpecific.Height,
		Material:    windowSpecific.Material,
	}

	return patchedWindow, nil
}
