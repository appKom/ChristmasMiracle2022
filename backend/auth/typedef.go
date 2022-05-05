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
	User JWTClaimUser
	jwt.StandardClaims
}

type JWTClaimUser struct {
	ID       uint
	Email    string
	Username string
	Points   int
}
