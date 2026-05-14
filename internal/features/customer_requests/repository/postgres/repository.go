package customer_requests_repository_postgres

import core_postgres_pool "github.com/nikitavaulin/kudesnik/internal/core/repository/postgres/pool"

type CustomerRequestsRepository struct {
	pool core_postgres_pool.Pool
}

func NewCustomerRequestsRepository(pool core_postgres_pool.Pool) *CustomerRequestsRepository {
	return &CustomerRequestsRepository{
		pool: pool,
	}
}
