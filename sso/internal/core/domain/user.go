package domain

import "time"

type VerificationStatus int32

const (
	StatusUnknown             VerificationStatus = iota // 0 — обязательный нулевой вариант
	StatusActive                                        // 1
	StatusPendingVerification                           // 2
	StatusBlocked                                       // 3
)

// доменная сущность пользователя
type User struct {
	ID        int64
	Version   int64
	Username  string
	Password  string
	Email     string
	State     VerificationStatus
	CreatedAt time.Time
	UpdatedAt *time.Time
}

// CanAuthorize проверяет, имеет ли право пользователь войти в систему
func (u *User) CanAuthorize() bool {
	return u.State == StatusActive
}

// IsPending проверяет, ожидает ли пользователь подтверждения SMTP
func (u *User) IsPending() bool {
	return u.State == StatusPendingVerification
}

// NewUserUninitialized создает  пользователя, у которого
// поля ID, Version, UpdatedAt не заданы
//
// Returns: пользователя

func NewUserUninitialized(
	username string,
	password string,
	email string,
) User {
	t := time.Now()
	return User{
		ID:        UserIDUnititialized,
		Version:   UserVersionUnititialized,
		Username:  username,
		Password:  password,
		Email:     email,
		CreatedAt: t,
	}
}
