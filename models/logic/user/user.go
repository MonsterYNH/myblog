package user

import (
	"gopkg.in/mgo.v2/bson"
	"myblog/models/db"
	"myblog/models/logic"
	"myblog/models/model"
	"myblog/models/model/httparg"
	"strings"
)

type LogicUser struct {}

var userEntry = db.GetObjectEntryByTypeName(db.MONGO_COLL_USER)

func (self *LogicUser) Login(arg httparg.LoginArg) (reply *httparg.LoginReply, errCode int) {
	// 检查参数
	if len(arg.Account) == 0 || strings.EqualFold(arg.Account, "") {
		return nil, logic.ACCOUNT_OR_PASSWORD_ERROR
	}
	// 验证图片验证码

	// 验证手机验证码

	// 验证手机号


	user := model.CreateUser(arg.Account, arg.Password, arg.Phone)
	userEntry.InsertOneObjectFromMongo(db.MONGO_DB, user)
	token := bson.NewObjectId()
	userEntry.InsertOneObjectFromRedis(token.Hex(), user)
	userEntry.SetExpert(token.Hex(), 3600)
	reply = &httparg.LoginReply{
		Token:token.Hex(),
	}
	return
}

func (self *LogicUser) Logout(token string) {
	userEntry.DelKey(token)
}
