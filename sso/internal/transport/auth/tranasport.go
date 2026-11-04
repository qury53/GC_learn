package auth_transport

import (
	"context"
	ssov1 "shared/pkg/proto/sso/v1"

	"github.com/qury53/GC_learn/sso/internal/core/domain"
)

type AuthTransport struct {
	ssov1.UnimplementedAuthServiceServer
	service AuthService
}

type AuthService interface {
	Register(ctx context.Context, user domain.User) (int64, domain.VerificationStatus, error)
	Login(ctx context.Context, username string, password string) (string, int, domain.VerificationStatus)
	Refresh(ctx context.Context, refreshToken string) (string, error)
	VerifyEmail(ctx context.Context, email string, code string) string
	GetConnectURL(
		ctx context.Context,
		provider domain.Provider,
		redirectURI string,
		challenge string,
	) (url string, state string, err error)
	ExchangeCallback(
		ctx context.Context,
		provider domain.Provider,
		code string,
		state string,
		codeVerifier string,
	)
}

type SMTPEmailSender interface {
	Send(ctx context.Context)
}
