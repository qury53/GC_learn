package auth_service

import (
	"context"
	"shared/pkg/logger"
	"time"

	"github.com/qury53/GC_learn/sso/internal/core/domain"
)

type AuthService struct {
	log             *logger.Logger
	repository      UserRepository
	smtp            Smtp
	cacheClient     CacheClient
	JWTSecret       []byte
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

type UserRepository interface {
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

func NewAuthService(
	log *logger.Logger,
	config *Config,
	repository UserRepository,
	smtp Smtp,
	cacheClient CacheClient,
) *AuthService {
	return &AuthService{
		log:             log,
		repository:      repository,
		smtp:            smtp,
		cacheClient:     cacheClient,
		JWTSecret:       []byte(config.JWTSecret),
		AccessTokenTTL:  config.AccessTokenTTL,
		RefreshTokenTTL: config.RefreshTokenTTL,
	}
}
