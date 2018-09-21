package subscribe

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"myblog/models/db"
	"myblog/models/model"
)

type SubArticle struct {}

func (self *SubArticle) Init() {
	objectEntry := db.GetObjectEntryByTypeName(db.MONGO_COLL_ARTICLE)
	query := []bson.M {
		bson.M{"$match":bson.M{"status":1}},
		bson.M{"$sort":bson.M{"issue":-1}},
	}
	objectEntry.GetObjectsFromMongo(db.MONGO_DB, query)
	if articles, ok := objectEntry.Result.(*[]*model.Article); ok {
		for index, article := range *articles {
			objectEntry.LPush("articles", article)
			objectEntry.HSet("article_id", article.ID.Hex(), index)
		}
		beego.Info("[redis init] synchronization article success, length: ", len(*articles))
	}

	objectEntry = db.GetObjectEntryByTypeName(db.MONGO_COLL_ARTICLE)
	query = []bson.M {
		bson.M{"$match":bson.M{"is_top":true,"status":1}},
		bson.M{"$sort":bson.M{"issue":-1}},
	}
	objectEntry.GetObjectsFromMongo(db.MONGO_DB, query)
	if tops, ok := objectEntry.Result.(*[]*model.Article); ok {
		for _, top := range *tops {
			objectEntry.LPush("top", top)
		}
		beego.Info("[redis init] synchronization top success, length: ", len(*tops))
	}

}

func (self *SubArticle) Update(arg SubArg) {
	self.Init()
}
