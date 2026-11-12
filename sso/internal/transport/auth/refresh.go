package auth_transport

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"shared/pkg/logger"
	ssov1 "shared/pkg/proto/sso/v1"
	"strings"
	"time"

	core_transport_grpc_response "github.com/qury53/GC_learn/sso/internal/core/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// TODO: на
func (t *AuthTransport) Refresh(
	ctx context.Context,
	req *ssov1.RefreshRequest,
) (*ssov1.RefreshResponse, error) {
	refreshToken := extractRefreshFromCookie(ctx)
	fromCookie := refreshToken != ""
	if !fromCookie {
		refreshToken = req.GetRefreshToken()
	}

	if refreshToken == "" {
		return nil, status.Error(codes.Unauthenticated, "refresh token required")
	}

	refresh, access, err := t.service.Refresh(ctx, refreshToken)
	if err != nil {
		t.log.Warn("refresh failed", slog.Any("error", err))
		clearRefreshCookie(ctx, t.log)
		return nil, status.Error(
			core_transport_grpc_response.RespondCode(err),
			"failed to refresh token",
		)
	}

	resp := &ssov1.RefreshResponse{AccessToken: access}
	if !fromCookie {
		resp.RefreshToken = refresh
		return resp, nil
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

	return resp, nil
}

// extractRefreshFromCookie достает refresh токен из cookie
// если куки нет, то возвращается пустая строка
// Returns: string

func extractRefreshFromCookie(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	var rawCookies []string
	rawCookies = append(rawCookies, md.Get("grpcgateway-cookie")...)
	rawCookies = append(rawCookies, md.Get("cookie")...)

	if len(rawCookies) == 0 {
		return ""
	}

	req := &http.Request{Header: http.Header{}}
	for _, line := range rawCookies {
		req.Header.Add("Cookie", strings.TrimSpace(line))
	}

	if cookie, err := req.Cookie("refresh_token"); err == nil {
		return cookie.Value
	}
	return ""
}

// clearRefreshCookie очищает refresh_token куки если истек срок жизни
//
// Returns: error

func clearRefreshCookie(ctx context.Context, log *logger.Logger) {
	cookie := "refresh_token" + "=; Path=/; Max-Age=0; HttpOnly; Secure; SameSite=Lax"
	if err := grpc.SendHeader(ctx, metadata.Pairs("set-cookie", cookie)); err != nil {
		log.Warn("failed to clear refresh cookie", slog.Any("error", err))
	}
}
