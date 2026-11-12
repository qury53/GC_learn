package auth_service

import (
	"context"
	"errors"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

// Claims данные, которые хранятся в jwt токене, нужны в основном для фронта
type Claims struct {
	UserID int64
	jwt.RegisteredClaims
}

// GetRefreshTokenTTL возвращает срок жизни refresh токена
//
// Returns: time.Duration

func (s *AuthService) GetRefreshTokenTTL() time.Duration {
	return s.RefreshTokenTTL
}

// GenerateAccessToken генерирует access токен
//
// Returns: string, error

func (s *AuthService) GenerateAccessToken(userID int64) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.AccessTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "zapiski",
			Subject:   fmt.Sprintf("%d", userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(s.JWTSecret)
	if err != nil {
		return "", fmt.Errorf(
			"failed to create access token: %w",
			err,
		)
	}

	return tokenStr, nil
}

// GenerateRefreshToken генерирует refresh токен
//
// Returns: string, error

func (s *AuthService) GenerateRefreshToken(userID int64) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.RefreshTokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "zapiski",
			Subject:   fmt.Sprintf("%d", userID),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(s.JWTSecret)
	if err != nil {
		return "", fmt.Errorf(
			"failed to generate refresh token: %w",
			err,
		)
	}

	return tokenStr, nil
}

// ValidateToken валидирует токен и возвращает Claims
//
// Returns: *Claims, error

func (s *AuthService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return s.JWTSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// ValidateRefreshToken валидирует refresh токен
//
// Returns: int, bool, error

func (s *AuthService) ValidateRefreshToken(ctx context.Context, tokenStr string) (int64, bool, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.JWTSecret, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return 0, false, nil
		}

		return 0, false, fmt.Errorf("failed to parse jwt: %w", err)
	}

	if !token.Valid {
		return 0, false, nil
	}

	return claims.UserID, true, nil
}

func (s *AuthService) RotateToken(ctx context.Context, oldJTI string)
