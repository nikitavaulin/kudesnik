package product_categories_repository

import core_postgres_pool "github.com/nikitavaulin/kudesnik/internal/core/repository/postgres/pool"

type ProductCategoriesRepository struct {
	pool core_postgres_pool.Pool
}

func NewProductCategoriesRepository(pool core_postgres_pool.Pool) *ProductCategoriesRepository {
	return &ProductCategoriesRepository{
		pool: pool,
	}
}
