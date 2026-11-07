package auth_service

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func (s *AuthService) GenerateVerificationCode() (string, error) {
	max := big.NewInt(900000)

	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return "", err
	}

	code := n.Int64() + 100000

	return fmt.Sprintf("%d", code), nil
}
