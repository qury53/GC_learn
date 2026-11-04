package auth_transport

import (
	"context"
	ssov1 "shared/pkg/proto/sso/v1"

	"github.com/qury53/GC_learn/sso/internal/core/domain"
	core_transport_grpc_response "github.com/qury53/GC_learn/sso/internal/core/transport/grpc"
	"google.golang.org/grpc/status"
)

func (t *AuthTransport) Register(
	ctx context.Context,
	req *ssov1.RegisterRequest,
) (*ssov1.RegisterResponse, error) {
	user := domain.NewUserUninitialized(
		req.Username,
		req.Password,
		req.Email,
	)

	user_id, verificationStatus, err := t.service.Register(ctx, user)
	if err != nil {
		return &ssov1.RegisterResponse{}, status.Error(
			core_transport_grpc_response.RespondCode(err),
			"failed to register",
		)
	}

	return &ssov1.RegisterResponse{
		UserId: user_id,
		Status: ssov1.AuthStatus(verificationStatus),
	}, nil
}
