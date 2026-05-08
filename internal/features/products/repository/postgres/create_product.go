package products_repository_postges

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (r *ProductsRepositoryPostgres) CreateProduct(ctx context.Context, product domain.BaseProduct) (domain.BaseProduct, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		INSERT INTO kudesnik.products (
			product_id, product_name, price, description, is_visible, category_id, producer_id
		) VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING
			product_id, version, product_name, price, description, is_visible, category_id, producer_id;
	`

	row := r.pool.QueryRow(
		ctx, query,
		uuid.New(), product.ProductName, product.Price,
		product.Description, product.IsVisible,
		product.CategoryID, product.ProducerID,
	)

	var model ProductModel
	err := row.Scan(
		&model.ID, &model.Version,
		&model.ProductName, &model.Price, &model.Description,
		&model.IsVisible, &model.CategoryID, &model.ProducerID,
	)
	if err != nil {
		return domain.BaseProduct{}, fmt.Errorf("CreateProduct in Repo: %w", err)
	}

	product = productDomainFromModel(model)
	return product, nil
}
