package ai_provider_gc_models

import (
	"time"

	"github.com/qury53/GC_learn/agent/internal/domain"
)

type AuthResponse struct {
	AccessToken string
	ExpiredAt   int64
}

func AuthModelToDomain(model AuthResponse) domain.AuthResponse {
	return domain.AuthResponse{
		AccessToken: model.AccessToken,
		ExpiredAt:   time.UnixMicro(model.ExpiredAt),
	}
}
