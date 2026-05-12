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

func (r *ProductsRepositoryPostgres) GetWindow(ctx context.Context, id uuid.UUID) (domain.Window, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
        SELECT 
            w.window_id,
            w.purpose,
            w.width,
            w.height,
            w.material,
            p.product_name,
            p.price,
			p.description,
			p.category_code,
            p.is_visible,
            p.version
        FROM kudesnik.windows w
        INNER JOIN kudesnik.products p ON w.window_id = p.product_id
        WHERE w.window_id = $1
    `

	row := r.pool.QueryRow(ctx, query, id)

	var window domain.Window

	err := row.Scan(
		&window.ID,
		&window.Purpose,
		&window.Width,
		&window.Height,
		&window.Material,
		&window.ProductName,
		&window.Price,
		&window.Description,
		&window.CategoryCode,
		&window.IsVisible,
		&window.Version,
	)
	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Window{}, fmt.Errorf("Get window from repo: %v: %w", err, core_errors.ErrNotFound)
		}
		return domain.Window{}, fmt.Errorf("Get window from repo: %w", err)
	}

	return window, nil
}
