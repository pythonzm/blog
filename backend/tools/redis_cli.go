package tools

import (
	"backend/models"
	"backend/utils"
)

type Options struct {
	Timeout bool // 是否设置过期时间
}

var defaultOptions = Options{
	Timeout: false,
}

type Option func(*Options)

func newOptions(opts ...Option) Options {
	// 初始化默认值
	opt := defaultOptions

	for _, o := range opts {
		o(&opt) // 依次调用opts函数列表中的函数，为服务选项（opt变量）赋值
	}

	return opt
}

func SetTimeout(timeout bool) Option {
	return func(o *Options) {
		o.Timeout = timeout
	}
}

func SetKey(key string, value interface{}, opts ...Option) error {
	conn := models.RedisPool.Get()
	defer func() {
		if e := conn.Close(); e != nil {
			return
		}
	}()

	options := newOptions(opts...)
	if options.Timeout {
		_, err := conn.Do("SET", key, value, "EX", utils.RedisInfo.CacheTime)
		return err
	}
	_, err := conn.Do("SET", key, value)
	return err
}

func GetKey(key string) (data interface{}, err error) {
	conn := models.RedisPool.Get()
	defer func() {
		if e := conn.Close(); e != nil {
			return
		}
	}()

	data, err = conn.Do("GET", key)
	return
}

func DelKey(key string) error {
	conn := models.RedisPool.Get()
	defer func() {
		if e := conn.Close(); e != nil {
			return
		}
	}()

	_, err := conn.Do("DEL", key)
	return err
}

func INCRKey(key string) error {
	conn := models.RedisPool.Get()
	defer func() {
		if e := conn.Close(); e != nil {
			return
		}
	}()

	_, err := conn.Do("INCR", key)
	return err
}

func ZADDKey(rankKey, member string, score int) error {
	conn := models.RedisPool.Get()
	defer func() {
		if e := conn.Close(); e != nil {
			return
		}
	}()
	_, err := conn.Do("ZADD", rankKey, score, member)
	return err
}

func ZREVRANGE(rankKey string) (data interface{}, err error) {
	conn := models.RedisPool.Get()
	defer func() {
		if e := conn.Close(); e != nil {
			return
		}
	}()
	data, err = conn.Do("ZREVRANGE", rankKey, 0, 9, "WITHSCORES")
	return
}
