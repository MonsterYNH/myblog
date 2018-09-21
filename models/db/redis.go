package db

import (
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
)

var RedisConn RedisConnection

func init() {
	InitRedis()
	beego.Debug("[init] init redis success")
}

type RedisConnection struct {
	RedisPool *redis.Pool
}

func InitRedis() {
	config := GetDBConfig()
	RedisConn.RedisPool = &redis.Pool{
		MaxIdle:     config.RedisMaxIdle,
		IdleTimeout: config.RedisIdleTimeOut,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", config.RedisServer, redis.DialPassword(config.RedisPassword))
		},
	}
}
