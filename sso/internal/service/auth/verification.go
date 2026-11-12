package auth_service

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	core_errors "github.com/qury53/GC_learn/sso/internal/core/errors"
)

// verificationCacheItem структура кода верификации, записывается в кэш, хранит код и дату создания
type verificationCacheItem struct {
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
}

const cooldownDuration = 5 * time.Minute

// GenerateVerificationCode отвечает за генерацию 6-значного криптостойкого кода
func (s *AuthService) GenerateVerificationCode() (string, error) {
	max := big.NewInt(900000)

	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	code := n.Int64() + 100000

	return fmt.Sprintf("%d", code), nil
}

// CacheVerificationCode проверяет кулдаун и сохраняет сгенерированный код в кэш
func (s *AuthService) CacheVerificationCode(ctx context.Context, email string, code string) error {
	cacheKey := s.cacheClient.GenerateCacheKey("verification", "email", email)

	var existingItem verificationCacheItem
	found, err := s.cacheClient.Get(ctx, cacheKey, &existingItem)
	if err != nil {
		return fmt.Errorf("failed to check cache: %w", err)
	}

	if found {
		timePassed := time.Since(existingItem.CreatedAt)
		if timePassed < cooldownDuration {
			timeLeft := cooldownDuration - timePassed

			return fmt.Errorf(
				"verification code cooldown, time left:%v %w",
				timeLeft,
				core_errors.ErrTooManyRequest,
			)
		}
	}

	newItem := verificationCacheItem{
		Code:      code,
		CreatedAt: time.Now(),
	}

	err = s.cacheClient.Set(ctx, cacheKey, newItem)
	if err != nil {
		return fmt.Errorf("failed to save code to cache: %w", err)
	}

	return nil
}
