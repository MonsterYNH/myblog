package routers

import (
	"github.com/astaxie/beego"
	"myblog/controllers/about"
	"myblog/controllers/article"
	"myblog/controllers/code"
	"myblog/controllers/index"
	"myblog/controllers/user"
)

func init() {
	beego.Router("/v1/index", &index.IndexController{}, "post,get:Index")
	beego.Router("/v1/about", &about.AboutController{}, "post,get:About")
	beego.Router("/v1/picCode", &code.CodeController{}, "post,get:GetPicCode")
	beego.Router("/v1/phoneCode", &code.CodeController{}, "post:GetPhoneCode")
	beego.Router("/v1/login", &user.UserController{}, "post,get:Login")
	beego.Router("/v1/logout", &user.UserController{}, "post,get:Logout")
	beego.Router("/v1/article", &article.ArticleController{}, "post,get:GetArticles")
}
