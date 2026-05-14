package products_repository_postges

import core_postgres_pool "github.com/nikitavaulin/kudesnik/internal/core/repository/postgres/pool"

type ProductsRepositoryPostgres struct {
	pool core_postgres_pool.Pool
}

func NewProductsRepositoryPostgres(pool core_postgres_pool.Pool) *ProductsRepositoryPostgres {
	return &ProductsRepositoryPostgres{
		pool: pool,
	}
}
