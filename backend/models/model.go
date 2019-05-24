package models

import (
	"backend/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	var err error
	db, err = sqlx.Open(utils.DBInfo.Mode, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		utils.DBInfo.User,
		utils.DBInfo.Password,
		utils.DBInfo.Host,
		utils.DBInfo.Port,
		utils.DBInfo.DBName,
	))
	if err != nil {
		panic(fmt.Sprintf("数据库连接失败，err: %s", err))
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
}
