package auth

import (
	"errors"
	"time"

	"github.com/appKom/ChristmasMiracle2022/api"
	"github.com/golang-jwt/jwt"
)

func CreateNewToken(user api.User, jwtSecret string) (string, error) {
	jwtWrapper := JWTWrapper{
		SecretKey:       jwtSecret,
		Issuer:          "appKom",
		ExpirationHours: 24,
	}

	return jwtWrapper.GenerateToken(user)
}

func CheckTokenValidity(token string, jwtSecret string) (uint, error) {
	jwtWrapper := JWTWrapper{
		SecretKey: jwtSecret,
		Issuer:    "appKom",
	}

	return jwtWrapper.ValidateToken(token)
}

func (j *JWTWrapper) GenerateToken(user api.User) (string, error) {
	var jwtUser JWTClaimUser
	jwtUser.ID = user.ID
	jwtUser.Email = user.Email
	jwtUser.Username = user.Username
	jwtUser.Points = user.Points

	claims := &JWTClaim{
		User: jwtUser,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
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

func (j *JWTWrapper) ValidateToken(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.SecretKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*JWTClaim)

	if !ok {
		err = errors.New("invalid token")
		return 0, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return 0, err
	}

	return claims.User.ID, nil
}
