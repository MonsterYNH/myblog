package subscribe

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"myblog/models/db"
	"myblog/models/model"
)

type SubNotice struct {}

func (self *SubNotice) Init() {
	noticeEntry := db.GetObjectEntryByTypeName(db.MONGO_COLL_NOTICE)
	query := []bson.M {
		bson.M{"$sort":bson.M{"issue":-1}},
	}
	noticeEntry.GetObjectsFromMongo(db.MONGO_DB, query)
	if notices, ok := noticeEntry.Result.(*[]*model.Notice); ok {
		for _, notice := range *notices {
			noticeEntry.LPush(db.REDIS_NOTICE_KEY, notice)
		}
		beego.Debug("[redis init] synchronization notice success, length: ", len(*notices))
	}
}

func (self *SubNotice) Update(arg SubArg) {
	noticeEntry := db.GetObjectEntryByTypeName(db.MONGO_COLL_NOTICE)
	noticeEntry.LTrim(db.MONGO_COLL_NOTICE, 0, -1)
	self.Init()
}
