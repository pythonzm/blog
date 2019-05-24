package models

import (
	"backend/utils"
	"crypto/md5"
	"encoding/hex"
)

type Auth struct {
	ID       int    `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	if e := db.Get(&auth, "select * from blog_auth where username=? and password=?", username, md5pass(password)); e != nil {
		return false
	}
	return true
}

func md5pass(password string) string {
	m := md5.New()
	m.Write([]byte(password))
	return hex.EncodeToString(m.Sum([]byte(utils.DBInfo.Salt)))
}
