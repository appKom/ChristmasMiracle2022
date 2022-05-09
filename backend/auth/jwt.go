package auth

import (
	"errors"
	"time"

	"github.com/appKom/ChristmasMiracle2022/api"
	"github.com/golang-jwt/jwt"
)

func CreateNewTokenPair(user api.User, jwtSecret string) (api.TokenResponse, error) {
	jwtWrapper := JWTWrapper{
		SecretKey: jwtSecret,
		Issuer:    "appKom",
	}

	var resp api.TokenResponse
	access, err := jwtWrapper.GenerateToken(user)
	refresh, err2 := jwtWrapper.GenerateRefreshToken(user)

	if err != nil || err2 != nil {
		return resp, err
	}

	resp.Access = access
	resp.Refresh = refresh

	return resp, nil
}

func CheckTokenValidity(token string, jwtSecret string) (uint, error) {
	jwtWrapper := JWTWrapper{
		SecretKey: jwtSecret,
		Issuer:    "appKom",
	}

	return jwtWrapper.validateToken(token)
}

func CheckTokenValidityWithClaims(token string, jwtSecret string) (uint, error) {
	jwtWrapper := JWTWrapper{
		SecretKey: jwtSecret,
		Issuer:    "appKom",
	}

	return jwtWrapper.validateTokenWithClaims(token)
}

func (j *JWTWrapper) GenerateToken(user api.User) (string, error) {
	claims := &JWTClaim{
		Sub:    user.ID,
		Name:   user.Username,
		Email:  user.Email,
		Points: user.Points,
		Admin:  user.Admin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (j *JWTWrapper) GenerateRefreshToken(user api.User) (string, error) {
	claims := &JWTRefreshClaim{
		Sub: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshToken, err := token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}

func getClaimsAndValidate(tokenString string, jwtSecret string) (*JWTClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTClaim)

	if !ok {
		err = errors.New("invalid token")
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return nil, err
	}
	return claims, err
}

func (j *JWTWrapper) validateToken(tokenString string) (uint, error) {
	claims, err := getClaimsAndValidate(tokenString, j.SecretKey)
	if err != nil {
		return 0, err
	}
	return claims.Sub, nil
}

func (j *JWTWrapper) validateTokenWithClaims(tokenString string) (uint, error) {
	claims, err := getClaimsAndValidate(tokenString, j.SecretKey)
	if err != nil {
		return 0, err
	}

	if claims.Name == "" || claims.Email == "" {
		return 0, errors.New("invalid token")
	}
	return claims.Sub, nil
}
