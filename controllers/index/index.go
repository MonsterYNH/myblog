package index

import (
	"github.com/astaxie/beego"
	"myblog/models/logic/article"
	"myblog/models/logic/notice"
)

type IndexController struct {
	beego.Controller
}

func (self *IndexController) Index() {
	// 顶置文章
	top := article.GetTop(2)
	// 最新文章
	newArticle := article.GetNewArticle(5)
	// 公告
	notices := notice.GetNotice(-1)

	self.Data["notice"] = notices
	self.Data["newArticle"] = newArticle
	self.Data["top"] = top
	self.TplName = "index.html"
}
