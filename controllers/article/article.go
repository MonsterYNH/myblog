package article

import (
	"github.com/astaxie/beego"
	"myblog/models/logic/article"
	"strconv"
)

type ArticleController struct {
	beego.Controller
}

func (self *ArticleController) GetArticles() {
	// 获取分页数据
	page, _ := self.GetInt("page", 1)
	size, _ := self.GetInt("size", 4)
	data, all := article.GetArticleByPage(page, size)
	self.Data["list"] = data
	allPage := all / size
	if all % size > 0 {
		allPage++
	}
	// 获取分页参数
	urlArray := make([]string, 0)
	for i := 1; i <= allPage; i++ {
		urlArray = append(urlArray, "article?page="+strconv.Itoa(i)+"&size="+strconv.Itoa(size))
	}
	self.Data["urlArray"] = urlArray
	self.Data["all"] = allPage
	self.TplName = "article.html"
}
