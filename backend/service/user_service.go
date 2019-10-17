package service

import (
	"backend/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type User struct {
	ID           int    `json:"id" db:"id" form:"id"`
	Username     string `json:"username" db:"username" form:"username"`
	Password     string `json:"password" db:"password" form:"password"`
	Introduction string `json:"introduction" db:"introduction" form:"introduction"`
	Avatar       string `json:"avatar" db:"avatar" form:"avatar"`
	Nickname     string `json:"nickname" db:"nickname" form:"nickname"`
	About        string `json:"about" db:"about" form:"about"`
}

var jwtSecret = []byte(utils.AppInfo.JwtSecret)

type CustomClaims struct {
	User
	jwt.StandardClaims
}

func (u User) CheckAuth() bool {
	var count int
	if e := db.Get(&count, "select count(1) from blog_user where username=? and password=?", u.Username, utils.EncodeMD5(u.Password)); e != nil {
		utils.WriteErrorLog(e.Error())
		return false
	}
	return count > 0
}

func (u User) GenToken() (string, error) {
	claims := CustomClaims{u, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Second * time.Duration(utils.AppInfo.TokenTimeout)).Unix(),
		Id:        fmt.Sprintf("%v", time.Now().UnixNano()),
	}}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, e := tokenClaims.SignedString(jwtSecret)
	return token, e
}

func ParseToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if token != nil {
		if _, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			return nil
		} else {
			return err
		}
	}

	return err
}

func GetUser() (User, error) {
	var user User
	e := db.Get(&user, "select avatar,introduction,nickname from blog_user")
	return user, e
}

func (u User) EditUser() error {
	_, e := db.Exec("update blog_user set introduction=?,avatar=?,nickname=?", u.Introduction, u.Avatar, u.Nickname)
	return e
}

func (u User) ResetPassword() error {
	_, e := db.Exec("update blog_user set password=?", utils.EncodeMD5(u.Password))
	return e
}

func (u User) EditAbout() error {
	_, e := db.Exec("update blog_user set about=?", u.About)
	return e
}

func GetAbout() (string, error) {
	var a string
	e := db.Get(&a, "select about from blog_user")
	return a, e
}
