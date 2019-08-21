package service

import (
	"backend/utils"
	"crypto/md5"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Auth struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

var jwtSecret = []byte(utils.AppInfo.JwtSecret)

type Claims struct {
	Auth
	jwt.StandardClaims
}

func (a Auth) CheckAuth() bool {
	var auth Auth
	if e := db.Get(&auth, "select * from blog_auth where username=? and password=?", a.Username, md5pass(a.Password)); e != nil {
		return false
	}
	return true
}

func md5pass(password string) string {
	m := md5.New()
	m.Write([]byte(password))
	return hex.EncodeToString(m.Sum([]byte(utils.DBInfo.Salt)))
}

func (a Auth) GenToken() (string, error) {
	claims := Claims{a, jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + utils.AppInfo.TokenTimeout,
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
