package users_postgres_repository

import core_repository_postgres "github.com/qury53/GC_learn/sso/internal/core/repository/postgres"

type UserRepository struct {
	pool *core_repository_postgres.Pool
}

func NewUserRepository(pool *core_repository_postgres.Pool) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}
