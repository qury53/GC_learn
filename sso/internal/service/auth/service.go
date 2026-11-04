package auth_service

import (
	"context"

	"github.com/qury53/GC_learn/sso/internal/core/domain"
)

type AuthService struct {
}

type AuthRepository interface {
	Register(ctx context.Context, user domain.User) (domain.User, error)
	GetUserByUsername(ctx context.Context, username string) (domain.User, error)
}
