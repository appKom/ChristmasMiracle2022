package auth

import (
	"github.com/golang-jwt/jwt"
)

type JWTWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type JWTClaim struct {
	Sub    uint   `json:"sub"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Points int    `json:"points"`
	Admin  bool   `json:"admin"`
	jwt.StandardClaims
}

type JWTRefreshClaim struct {
	Sub uint `json:"sub"`
	jwt.StandardClaims
}
