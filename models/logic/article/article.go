package article

import (
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"myblog/models/db"
	"myblog/models/logic"
	"myblog/models/model"
)

func GetTop(num int) ([]model.Article) {
	var articleEntry = db.GetObjectEntryByTypeName(db.MONGO_COLL_ARTICLE)
	articleEntry.Result = nil
	articleEntry.LRange("top", 0, num)
	if objects, ok := articleEntry.Result.(*[]*model.Article); ok {
		articles := make([]model.Article, 0)
		for _, article := range *objects {
			articles = append(articles, *article)
		}
		return articles
	}
	return nil
}

func GetNewArticle(num int) []model.Article {
	var articleEntry = db.GetObjectEntryByTypeName(db.MONGO_COLL_ARTICLE)
	articleEntry.Result = nil
	articleEntry.LRange("articles", 0, num)
	if objects, ok := articleEntry.Result.(*[]*model.Article); ok {
		articles := make([]model.Article, 0)
		for _, article := range *objects {
			articles = append(articles, *article)
		}
		return articles
	}
	return nil
}

func GetArticleByPage(page, size int) ([]model.Article, int) {
	beego.Debug("[article] list page:", page, "size: ", size)
	datas, num :=  logic.GetDataByPage(page, size, db.MONGO_COLL_ARTICLE, "articles")
	articles := make([]model.Article, 0)
	if objects, ok := datas.([]*model.Article); ok {
		for _, article := range objects {
			articles = append(articles, *article)
		}
		return articles, num
	}
	return nil, 0
}

func GetArticleInfo(articleId bson.ObjectId) *model.Article {
	data := logic.GetDataByIdFormList(articleId, db.MONGO_COLL_ARTICLE, "article_id", "articles")
	if article, ok := data.(*model.Article); ok {
		return article
	}
	return nil
}

func GetArticleHot(num int) []model.Article {

}
