package service

import (
	"backend/utils"
	"errors"
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
}

var jwtSecret = []byte(utils.AppInfo.JwtSecret)
var UserInfo = map[string]User{}

type CustomClaims struct {
	User
	jwt.StandardClaims
}

func (u User) CheckAuth() bool {
	var count int
	if e := db.Get(&count, "select count(1) from blog_user where username=? and password=?", u.Username, utils.EncodeMD5(u.Password)); e != nil {
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
	if e == nil {
		user, _ := u.getUser()
		UserInfo[token] = user
	}
	return token, e
}

func ParseToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if _, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return nil
	} else {
		return err
	}

}

func GetUserInfo(t string) (User, error) {
	value, ok := UserInfo[t]
	if ok {
		return value, nil
	}
	return User{}, errors.New("获取用户信息失败，token不存在或已过期")
}

func (u User) getUser() (User, error) {
	var user User
	e := db.Get(&user, "select * from blog_user where username=?", u.Username)
	return user, e
}

func Logout() {
	UserInfo = map[string]User{}
}

func (u User) EditUser() error {
	_, e := db.Exec("update blog_user set introduction=?,avatar=?,nickname=? where username=?", u.Introduction, u.Avatar, u.Nickname, u.Username)
	return e
}

func (u User) ResetPassword() error {
	_, e := db.Exec("update blog_user set password=? where username=?", utils.EncodeMD5(u.Password), u.Username)
	return e
}