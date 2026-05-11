package admin_repository_postgres

import core_postgres_pool "github.com/nikitavaulin/kudesnik/internal/core/repository/postgres/pool"

type AdminRepositoryPostgres struct {
	pool core_postgres_pool.Pool
}

func NewAdminRepositoryPostgres(pool core_postgres_pool.Pool) *AdminRepositoryPostgres {
	return &AdminRepositoryPostgres{
		pool: pool,
	}
}
