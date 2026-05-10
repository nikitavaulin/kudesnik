package products_repository_postges

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

// PatchBalcony обновляет балкон
func (r *ProductsRepositoryPostgres) PatchBalcony(
	ctx context.Context,
	id uuid.UUID,
	balcony domain.Balcony,
) (domain.Balcony, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return domain.Balcony{}, fmt.Errorf("PatchBalcony (begin tx): %w", err)
	}
	defer tx.Rollback(ctx)

	updatedProduct, err := r.updateProductBaseInTx(ctx, tx, id, balcony.ProductBase, balcony.Version)
	if err != nil {
		return domain.Balcony{}, fmt.Errorf("PatchBalcony (update product): %w", err)
	}

	updateBalconyQuery := `
		UPDATE kudesnik.balconies
		SET 
			purpose = $2,
			material = $3
		WHERE balcony_id = $1
		RETURNING 
			purpose,
			material;
	`

	var balconySpecific struct {
		Purpose  string
		Material string
	}

	err = tx.QueryRow(
		ctx, updateBalconyQuery,
		id,
		balcony.Purpose,
		balcony.Material,
	).Scan(
		&balconySpecific.Purpose,
		&balconySpecific.Material,
	)

	if err != nil {
		return domain.Balcony{}, fmt.Errorf("PatchBalcony (update balconies): %w", err)
	}

	if err = tx.Commit(ctx); err != nil {
		return domain.Balcony{}, fmt.Errorf("PatchBalcony (commit): %w", err)
	}

	patchedBalcony := domain.Balcony{
		ProductBase: updatedProduct,
		Purpose:     balconySpecific.Purpose,
		Material:    balconySpecific.Material,
	}

	return patchedBalcony, nil
}
