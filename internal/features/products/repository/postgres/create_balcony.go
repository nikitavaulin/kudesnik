package products_repository_postges

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

// CreateBalcony создает новый балкон
func (r *ProductsRepositoryPostgres) CreateBalcony(
	ctx context.Context,
	balcony domain.Balcony,
) (domain.Balcony, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return domain.Balcony{}, fmt.Errorf("CreateBalcony (begin tx): %w", err)
	}
	defer tx.Rollback(ctx)

	// 1. Вставка в products
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
		uuid.New(), balcony.ProductName, balcony.Price,
		balcony.Description, balcony.IsVisible,
		balcony.CategoryCode, balcony.ProducerID,
	).Scan(&productID, &version)

	if err != nil {
		return domain.Balcony{}, fmt.Errorf("CreateBalcony (insert product): %w", err)
	}

	// 2. Вставка в balconies
	balconyQuery := `
		INSERT INTO kudesnik.balconies (
			balcony_id, purpose, material
		) VALUES ($1, $2, $3);
	`

	_, err = tx.Exec(
		ctx, balconyQuery,
		productID, balcony.Purpose, balcony.Material,
	)

	if err != nil {
		return domain.Balcony{}, fmt.Errorf("CreateBalcony (insert balcony): %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return domain.Balcony{}, fmt.Errorf("CreateBalcony (commit): %w", err)
	}

	balcony.ID = productID
	balcony.Version = version

	return balcony, nil
}
