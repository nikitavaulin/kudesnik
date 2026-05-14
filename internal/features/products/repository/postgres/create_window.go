package products_repository_postges

import (
	"context"
	"fmt"

	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (r *ProductsRepositoryPostgres) CreateWindow(ctx context.Context, window domain.Window) (domain.Window, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return domain.Window{}, fmt.Errorf("CreateWindow in Repo (begin tx): %w", err)
	}
	defer tx.Rollback(ctx)

	productID, version, err := r.createProductInTx(ctx, tx, window.ProductBase)
	if err != nil {
		return domain.Window{}, fmt.Errorf("CreateWindow (insert product): %w", err)
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
