package db

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2"
	"log"
)

var MongoConn MongoConnection

type MongoConnection struct {
	mgoSession *mgo.Session
}

func init() {
	MongoConn.InitMongo()
	beego.Debug("[init] mongoDB init success")
}

func (self *MongoConnection) InitMongo() {
	config := GetDBConfig()
	dbSession, err := mgo.Dial(config.MongoUrl)
	if err != nil {
		log.Panicf("[init] mongoDB connect failed, mongo url:%s", config.MongoUrl)
	}
	self.mgoSession = dbSession
	// 对Session进行必要的设置
	self.mgoSession.SetSocketTimeout(config.MongoSocketTimeOut)
	self.mgoSession.SetSyncTimeout(config.MongoSyncTimeOut)
	self.mgoSession.SetPoolLimit(config.MongoPoolLimit)
	self.mgoSession.SetMode(mgo.Strong, true)
}

// 获取mgoSession
func (mongoConn *MongoConnection) GetMgoSession() *mgo.Session {
	return mongoConn.mgoSession.Copy()
}

// 关闭mgoSession
func (mongoConn *MongoConnection) CloseMgoSeeion(session *mgo.Session) {
	defer func() {
		if err := recover(); err != nil {
			log.Panic("关闭mgoSession连接出现错误, Error:", err)
		}
	}()
	session.Close()
}
