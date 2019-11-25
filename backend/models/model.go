package models

import (
	"backend/utils"
	"fmt"
	"github.com/elastic/go-elasticsearch/v8"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	"time"
)

var DB *sqlx.DB
var RedisPool *redis.Pool
var ESClient *elasticsearch.Client

func init() {
	DB = sqlx.MustConnect(utils.DBInfo.Mode, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
		utils.DBInfo.User,
		utils.DBInfo.Password,
		utils.DBInfo.Host,
		utils.DBInfo.Port,
		utils.DBInfo.DBName,
	))

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(10)

	redisAddr := fmt.Sprintf("%s:%s", utils.RedisInfo.Host, utils.RedisInfo.Port)
	RedisPool = &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", redisAddr, redis.DialPassword(utils.RedisInfo.Password), redis.DialDatabase(utils.RedisInfo.DB), redis.DialConnectTimeout(time.Second*5))
		},
		MaxIdle:     20,
		MaxActive:   1000,
		IdleTimeout: time.Second * 100,
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	esAddr := fmt.Sprintf("http://%s:%s", utils.ESInfo.Host, utils.ESInfo.Port)
	ESClient, _ = elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{
			esAddr,
		},
	})
}

