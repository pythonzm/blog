package utils

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte(AppInfo.JwtSecret)

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenToken(username, password string) (string, error) {
	claims := Claims{username, password, jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + AppInfo.TokenTimeout,
	}}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, e := tokenClaims.SignedString(jwtSecret)
	return token, e
}

func ParseToken(token string) (*Claims, error) {
	claims, e := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return jwtSecret, nil
	})

	if claims != nil {
		if c, ok := claims.Claims.(*Claims); ok && claims.Valid {
			return c, nil
		}
	}
	return nil, e
}
