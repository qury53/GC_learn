package domain

import "time"

type AuthResponse struct {
	AccessToken string
	ExpiredAt   time.Time
}
