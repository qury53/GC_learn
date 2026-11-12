package auth_service

import (
	"context"
	"fmt"

	core_encrypt "github.com/qury53/GC_learn/sso/internal/core/encrypt"
	core_errors "github.com/qury53/GC_learn/sso/internal/core/errors"
)

func (s *AuthService) Login(
	ctx context.Context,
	username string,
	password string,
) (string, string, int64, error) {
	user, err := s.repository.GetUserByUsername(ctx, username)
	if err != nil {
		return "", "", 0, fmt.Errorf(
			"service: get user by username: %w",
			err,
		)
	}

	if !user.CanAuthorize() {
		s.log.Debug("login rejected: user not authorized", "username", username, "state", user.State)
		return "", "", 0, core_errors.ErrInvalidCredentials
	}

	if err := core_encrypt.CompareHashAndPassword(
		user.Password,
		password,
	); err != nil {
		return "", "", 0, fmt.Errorf(
			"service: invalid credentials: %w",
			err,
		)
	}

	accessToken, err := s.GenerateAccessToken(user.ID)
	if err != nil {
		return "", "", 0, fmt.Errorf("service: generate access token: %w", err)
	}

	refreshToken, err := s.GenerateRefreshToken(user.ID)
	if err != nil {
		return "", "", 0, fmt.Errorf("service: generate refresh token: %w", err)
	}

	return accessToken, refreshToken, user.ID, nil
}
