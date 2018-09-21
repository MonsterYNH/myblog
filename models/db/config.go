package db

import (
	"github.com/astaxie/beego"
	"time"
)

type ConfigDB struct {
	MongoUrl           string
	MongoSocketTimeOut time.Duration
	MongoSyncTimeOut   time.Duration
	MongoPoolLimit     int

	RedisServer      string
	RedisPassword    string
	RedisMaxIdle     int
	RedisIdleTimeOut time.Duration
}

func GetDBConfig() *ConfigDB {
	return &ConfigDB{
		MongoUrl:           beego.AppConfig.DefaultString("mongoUrl", "39.104.162.243"),
		MongoSocketTimeOut: time.Second * time.Duration(beego.AppConfig.DefaultInt64("mongoSocketTimeOut", 60)),
		MongoSyncTimeOut:   time.Second * time.Duration(beego.AppConfig.DefaultInt64("mongoSyncTimeOut", 60)),
		MongoPoolLimit:     beego.AppConfig.DefaultInt("mongoPoolLimit", 500),
		RedisServer:        beego.AppConfig.DefaultString("redisServer", "39.104.162.243:6379"),
		RedisPassword:      beego.AppConfig.DefaultString("redisPassword", ""),
		RedisMaxIdle:       beego.AppConfig.DefaultInt("redisMaxIdle", 60),
		RedisIdleTimeOut:   time.Second * time.Duration(beego.AppConfig.DefaultInt64("redisIdleTimeOut", 60)),
	}
}
