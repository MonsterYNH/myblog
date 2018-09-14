package mongodb

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
	if err := MongoConn.InitMongo("localhost"); err != nil {
		log.Panic("[init] mongoDB init failed", err)
	}
	beego.Debug("[init] mongoDB init success")
}

func (self *MongoConnection) InitMongo(url string) error {
	dbSession, err := mgo.Dial(url)
	if err != nil {
		return err
	}
	config := GetDBConfig()
	self.mgoSession = dbSession
	// 对Session进行必要的设置
	self.mgoSession.SetSocketTimeout(config.MongoSocketTimeOut)
	self.mgoSession.SetSyncTimeout(config.MongoSyncTimeOut)
	self.mgoSession.SetPoolLimit(config.MongoPoolLimit)
	self.mgoSession.SetMode(mgo.Strong, true)
	return nil
}

// 获取mgoSession
func (mongoConn *MongoConnection) GetMgoSession() *mgo.Session {
	return mongoConn.mgoSession.Copy()
}

// 关闭mgoSession
func (mongoConn *MongoConnection) CloseMgoSeeion(session *mgo.Session) {
	defer func() {
		if err := recover(); err != nil {
			//loger.Errorf("关闭mgoSession连接出现错误, Error:", err)
			log.Panic("关闭mgoSession连接出现错误, Error:", err)
		}
	}()
	session.Close()
}
