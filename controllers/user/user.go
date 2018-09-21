package user

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"myblog/models/db"
	"myblog/models/logic"
	"myblog/models/logic/user"
	"myblog/models/model/httparg"
)

type UserController struct {
	beego.Controller
}

var logicUser = user.LogicUser{}

func (self *UserController) Login() {
	var loginArg httparg.LoginArg
	var reply httparg.Reply
	reply.Action = "login"
	if err := json.Unmarshal(self.Ctx.Input.RequestBody, &loginArg); err != nil {
		beego.Debug("[login] user login failed, Error:", err)
		reply.Ret = logic.JSON_ERROR
		return
	}
	beego.Debug("[login] user login, account:", loginArg.Account, "password: ", loginArg.Password, "phone: ", loginArg.Phone)
	reply.Data, reply.Ret = logicUser.Login(loginArg)

	self.Data["json"] = reply
	self.ServeJSON()
}

func (self *UserController) Logout() {
	var logoutArg httparg.LogoutArg
	var reply httparg.Reply
	reply.Action = "logout"
	if err := json.Unmarshal(self.Ctx.Input.RequestBody, &logoutArg); err != nil {
		beego.Debug("[logout] user logout failed, Error:", err)
		reply.Ret = logic.JSON_ERROR
		return
	}
	userEntry := db.GetObjectEntryByTypeName(db.MONGO_COLL_USER)
	userEntry.DelKey(logoutArg.Token)
	reply.Ret = 0
	self.Data["json"] = reply
	self.ServeJSON()
}


