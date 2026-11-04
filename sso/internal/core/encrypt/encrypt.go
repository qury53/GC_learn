package core_encrypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// GeneratePasswordHash генерит хэш пароля
//
// Returns: хэш пароля
func GeneratePasswordHash(val string) (string, error) {
	hashedVal, err := bcrypt.GenerateFromPassword([]byte(val), 8)
	if err != nil {
		return "", fmt.Errorf("hash error")
	}

	return string(hashedVal), nil
}

// CompareHashAndPassword сравнивает хэш с паролем
//
// Returns: ошибку wrong password если пароль не совпадает
func CompareHashAndPassword(hash string, val string) error {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(hash),
		[]byte(val),
	); err != nil {
		return fmt.Errorf("wrong password")
	}

	return nil
}
