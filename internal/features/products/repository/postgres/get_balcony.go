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

// GetBalcony получает балкон по ID
func (r *ProductsRepositoryPostgres) GetBalcony(
	ctx context.Context,
	id uuid.UUID,
) (domain.Balcony, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		SELECT 
			b.balcony_id,
			b.purpose,
			b.material,
			p.product_name,
			p.price,
			p.description,
			p.category_code,
			p.is_visible,
			p.version,
			p.producer_id
		FROM kudesnik.balconies b
		INNER JOIN kudesnik.products p ON b.balcony_id = p.product_id
		WHERE b.balcony_id = $1
	`

	row := r.pool.QueryRow(ctx, query, id)

	var balcony domain.Balcony

	err := row.Scan(
		&balcony.ID,
		&balcony.Purpose,
		&balcony.Material,
		&balcony.ProductName,
		&balcony.Price,
		&balcony.Description,
		&balcony.CategoryCode,
		&balcony.IsVisible,
		&balcony.Version,
		&balcony.ProducerID,
	)

	if err != nil {
		if errors.Is(err, core_postgres_pool.ErrNoRows) {
			return domain.Balcony{}, fmt.Errorf("GetBalcony: %v: %w", err, core_errors.ErrNotFound)
		}
		return domain.Balcony{}, fmt.Errorf("GetBalcony: %w", err)
	}

	return balcony, nil
}
