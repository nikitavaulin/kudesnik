package product_categories_repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/nikitavaulin/kudesnik/internal/core/domain"
)

func (r *ProductCategoriesRepository) CreateProductCategory(
	ctx context.Context,
	category domain.ProductCategory,
) (domain.ProductCategory, error) {
	ctx, cancel := context.WithTimeout(ctx, r.pool.OperationTime())
	defer cancel()

	query := `
		INSERT INTO kudesnik.product_categories 
		(product_category_id, product_category_name, installation_price)
		VALUES ($1, $2, $3)
		RETURNING product_category_id, product_category_name, installation_price;
	`

	row := r.pool.QueryRow(
		ctx, query,
		uuid.New(), category.CategoryName, category.InstallationPrice,
	)

	var model ProductCategoriesModel
	err := row.Scan(&model.ID, &model.CategoryName, &model.InstallationPrice)
	if err != nil {
		return domain.ProductCategory{}, fmt.Errorf("CreateProductCategoryRepo: %w", err)
	}

	category = *domain.NewProductCategory(
		model.ID, model.CategoryName, model.InstallationPrice,
	)

	return category, nil
}
