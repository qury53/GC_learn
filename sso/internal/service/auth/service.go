package auth_service

import (
	"context"

	"github.com/qury53/GC_learn/sso/internal/core/domain"
)

type AuthService struct {
	repository  AuthRepository
	smtp        Smtp
	cacheClient CacheClient
}

type AuthRepository interface {
	SaveUser(ctx context.Context, user domain.User) (domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (domain.User, error)
}

type CacheClient interface {
	Get(ctx context.Context, key string, dst any) (bool, error)
	Set(ctx context.Context, key string, value any) error
	Del(ctx context.Context, key string) error
	GenerateCacheKey(args ...string) string
}

type Smtp interface {
	Send(
		ctx context.Context,
		toEmail string,
		code string,
	) error
}
