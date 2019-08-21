package models

import (
	"backend/utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func init() {
	DB = sqlx.MustConnect(utils.DBInfo.Mode, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		utils.DBInfo.User,
		utils.DBInfo.Password,
		utils.DBInfo.Host,
		utils.DBInfo.Port,
		utils.DBInfo.DBName,
	))

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(10)
}
