package ai_provider_gc

import (
	"net"
	"net/http"
	"sync"
	"time"
)

// GCProvier - провайдер GigaChat API
type GCProvider struct {
	client      *http.Client
	apiKey      string
	mu          sync.RWMutex
	accessToken string
	expiresAt   time.Time
	scope       string
}

func NewGCProvider(config *Config) *GCProvider {
	return &GCProvider{
		client: &http.Client{
			Transport: &http.Transport{
				DialContext: (&net.Dialer{
					Timeout: 5 * time.Second,
				}).DialContext,
				TLSHandshakeTimeout:   5 * time.Second,
				ResponseHeaderTimeout: 90 * time.Second,
			},
		},
		apiKey:      config.APIKey,
		mu:          sync.RWMutex{},
		accessToken: "",
		expiresAt:   time.Now(),
		scope:       config.Scope,
	}
}
