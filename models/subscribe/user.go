package subscribe

import (
	"gopkg.in/mgo.v2/bson"
	"myblog/models/db"
	"myblog/models/model"
)

type SubUser struct {}

func (self *SubUser) Init() {}

func (self *SubUser) Update(arg SubArg) {
	userEntry := db.GetObjectEntryByTypeName(db.MONGO_COLL_USER)
	userEntry.GetOneObjectFromRedis(arg.Token.Hex())
	user, ok := userEntry.Result.(*model.User)
	if ok {
		query := []bson.M {
			bson.M{"$match":bson.M{"_id":user.ID, "status":1}},
		}
		userEntry.GetOneObjectFromMongo(db.MONGO_DB, query)
		if user, ok := userEntry.Result.(*model.User); ok {
			userEntry.InsertOneObjectFromRedis(arg.Token.Hex(), user)
		}
	}

}
