package auth

type AuthResponse interface {
	Auth() AuthResponse
	GetAccessToken() string
	GetID() string
	GetInstanceURL() string
	GetScope() string
	GetTokenType() string
}
