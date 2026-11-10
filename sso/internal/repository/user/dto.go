package users_postgres_repository

import (
	"time"

	"github.com/qury53/GC_learn/sso/internal/core/domain"
)

// User модель пользователя для базы данных
type User struct {
	ID        int64
	Version   int64
	Username  string
	Password  string
	Email     string
	State     string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

// dtoToDomain конвертер модели для базы в доменную сущность
//
// Returns: domain.User

func dtoToDomain(dto User) domain.User {
	var state domain.VerificationStatus = domain.StatusUnknown
	switch dto.State {
	case "active":
		state = domain.StatusActive
	case "pending_verification":
		state = domain.StatusPendingVerification
	case "blocked":
		state = domain.StatusBlocked
	}

	return domain.User{
		ID:        dto.ID,
		Version:   dto.Version,
		Username:  dto.Username,
		Password:  dto.Password,
		Email:     dto.Email,
		State:     domain.VerificationStatus(state),
		CreatedAt: dto.CreatedAt,
		UpdatedAt: dto.UpdatedAt,
	}
}
