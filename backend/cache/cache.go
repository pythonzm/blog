package cache

import (
	"backend/models"
	"backend/utils"
)

func SetKey(key string, value interface{}) error {
	conn := models.RedisPool.Get()
	defer func() {
		if e := conn.Close(); e != nil {
			return
		}
	}()

	_, err := conn.Do("SET", key, value, "EX", utils.RedisInfo.CacheTime)
	return err
}

func GetKey(key string) (data interface{}, err error) {
	conn := models.RedisPool.Get()
	defer func() {
		if e := conn.Close(); e != nil {
			return
		}
	}()

	reply, err := conn.Do("GET", key)
	if err != nil {
		return nil, err
	}
	return reply, nil
}