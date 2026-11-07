package auth_service

import (
	"context"
	"fmt"

	"github.com/qury53/GC_learn/sso/internal/core/domain"
)

func (s *AuthService) Register(
	ctx context.Context,
	user domain.User,
) (int64, domain.VerificationStatus, error) {
	user, err := s.repository.SaveUser(ctx, user)
	if err != nil {
		return 0, domain.StatusUnknown, fmt.Errorf(
			"save user to db: %w",
			err,
		)
	}

	verificationCode, err := s.GenerateVerificationCode()

	if err != nil {
		return 0, domain.StatusUnknown, fmt.Errorf(
			"generate verification code: %w",
			err,
		)
	}

	if err := s.smtp.Send(
		ctx,
		user.Email,
		verificationCode,
	); err != nil {
		return 0, domain.StatusUnknown, fmt.Errorf(
			"send verification code: %w",
			err,
		)
	}

	return user.ID, domain.StatusPendingVerification, nil
}
