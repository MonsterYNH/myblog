package code

import (
	"github.com/astaxie/beego"
	"html/template"
	"myblog/models/logic/code"
	"myblog/models/model/httparg"
)

type CodeController struct {
	beego.Controller
}

func (self *CodeController) GetPhoneCode() {
	reply := httparg.Reply{
		Action:"getPhoneCode",
		Ret:-1,
		Data:"暂未开通",
	}
	self.Data["json"] = reply
	self.ServeJSON()
}

func (self *CodeController) GetPicCode() {
	picMap := code.GenerateCaptchaHandle()
	self.Data["PicId"] = picMap["captchaId"]
	self.Data["PicCode"] = template.URL(picMap["base64"])
	self.TplName = "code.tpl"
}


