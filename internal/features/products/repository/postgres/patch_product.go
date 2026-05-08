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
	product domain.BaseProduct,
) (domain.BaseProduct, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		UPDATE kudesnik.products
		SET 
			product_name = $2,
			price = $3,
			description = $4,
			is_visible = $5,
			category_id = $6,
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
			category_id,
			producer_id;
	`

	row := r.pool.QueryRow(
		ctx, query,
		id,
		product.ProductName, product.Price, product.Description,
		product.IsVisible, product.CategoryID, product.ProducerID,
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
		&productModel.CategoryID,
		&productModel.ProducerID,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.BaseProduct{}, fmt.Errorf(
				"product with ID=%v concurrently accessed: %v: %w",
				id,
				err,
				core_errors.ErrConflict,
			)
		}
		return domain.BaseProduct{}, fmt.Errorf("scan product error: %v", err)
	}

	patchedProduct := productDomainFromModel(productModel)
	return patchedProduct, nil
}
