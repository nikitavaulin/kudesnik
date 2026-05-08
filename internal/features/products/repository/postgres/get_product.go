package products_repository_postges

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
	core_errors "github.com/nikitavaulin/kudesnik/internal/core/errors"
)

func (r *ProductsRepositoryPostgres) GetProduct(ctx context.Context, id uuid.UUID) (domain.BaseProduct, error) {
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
		&productModel.IsVisible, &productModel.CategoryID, &productModel.ProducerID,
	)
	if err != nil {
		return domain.BaseProduct{}, fmt.Errorf("GetProduct from repo:%v: %w", err, core_errors.ErrNotFound)
	}

	product := productDomainFromModel(productModel)
	return product, nil
}
