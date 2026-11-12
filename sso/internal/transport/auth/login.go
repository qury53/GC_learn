package auth_transport

import (
	"context"
	"fmt"
	"log/slog"
	ssov1 "shared/pkg/proto/sso/v1"
	"time"
	"unicode/utf8"

	core_errors "github.com/qury53/GC_learn/sso/internal/core/errors"
	core_transport_grpc_response "github.com/qury53/GC_learn/sso/internal/core/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (t *AuthTransport) Login(
	ctx context.Context,
	req *ssov1.LoginRequest,
) (*ssov1.LoginResponse, error) {
	if err := validateLogin(req.Username, req.Password); err != nil {
		t.log.Debug("login validation error", slog.String("error", err.Error()))
		return nil, status.Error(
			core_transport_grpc_response.RespondCode(err),
			"login validation error",
		)
	}

	access, refresh, userID, err := t.service.Login(ctx, req.Username, req.Password)
	if err != nil {
		t.log.Debug(
			"login failed",
			slog.String("username", req.Username),
			slog.Any("error", err),
		)

		return nil, status.Error(
			core_transport_grpc_response.RespondCode(err),
			"service login error",
		)
	}

	cookieValue := fmt.Sprintf(
		"refresh_token=%s; Path=/; Max-Age=%d; HttpOnly; Secure; SameSite=Lax",
		refresh,
		int(t.service.GetRefreshTokenTTL()/time.Second),
	)

	md := metadata.Pairs("set-cookie", cookieValue)

	if err := grpc.SendHeader(ctx, md); err != nil {
		return nil, status.Errorf(codes.Internal, "failed to set cookie: %v", err)
	}

	return &ssov1.LoginResponse{
		AccessToken: access,
		UserId:      userID,
	}, nil
}

func validateLogin(username string, password string) error {
	usernameLen := utf8.RuneCountInString(username)
	if usernameLen < 3 || usernameLen > 32 {
		return fmt.Errorf("username length has to be between 3 and 32: %w", core_errors.ErrInvalidArgument)
	}

	passwordLen := utf8.RuneCountInString(password)
	if passwordLen < 6 || passwordLen > 72 {
		return fmt.Errorf("password length has to be between 6 and 72: %w", core_errors.ErrInvalidArgument)
	}

	return nil
}
