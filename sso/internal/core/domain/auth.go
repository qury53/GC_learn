package domain

import "errors"

// Provider определяет тип OAuth провайдера
type Provider string

const (
	ProviderGoogle Provider = "google"
	ProviderGithub Provider = "github"
	ProviderYandex Provider = "yandex"
)

// CodeChallengeMethod определяет метод хеширования PKCE
type CodeChallengeMethod string

const (
	PKCEMethodS256 CodeChallengeMethod = "S256"
)

var ErrUnsupportedProvider = errors.New("unsupported oauth provider")
