package subscribe

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"log"
	"myblog/models/db"
	"strings"
)

func init() {
	subscribe = Subscribe{
		Commond:    make(chan SubArg),
		CommondMap: make(map[string]SubscribeAble),
	}
	initEntry := db.GetObjectEntryByTypeName(db.MONGO_COLL_USER)
	initEntry.FlushDB()
	Add(db.MONGO_COLL_ARTICLE, &SubArticle{})
	Add(db.MONGO_COLL_USER, &SubUser{})
	Add(db.MONGO_COLL_NOTICE, &SubNotice{})
	go subscribe.Listen()
}

var subscribe Subscribe

type SubscribeAble interface {
	// 初始化操作
	Init()
	// 更新对应的操作
	Update(arg SubArg)
}

type Subscribe struct {
	Commond    chan SubArg
	CommondMap map[string]SubscribeAble
}

type SubArg struct {
	Option string
	Token bson.ObjectId
	IDs    []bson.ObjectId
}

func (self *Subscribe) Listen() {
	for {
		select {
		case commond := <-self.Commond:
			if objectEntry, ok := self.CommondMap[commond.Option]; ok {
				objectEntry.Update(commond)
			} else if strings.EqualFold(commond.Option, "stop") {
				beego.Debug("[subscribe] stop subscribe listen")
				return
			} else {
				beego.Error("[subscribe] update object failed, Error: option is not exist")
			}
		}
	}
}

func SendArg(arg SubArg) {
	subscribe.Commond <- arg
}

func Add(objectName string, able SubscribeAble) {
	if _, ok := subscribe.CommondMap[objectName]; ok {
		log.Panicf("[subscribe] object is exist, object name: %s", objectName)
	}
	able.Init()
	subscribe.CommondMap[objectName] = able
}
