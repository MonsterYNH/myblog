package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Article struct {
	ID bson.ObjectId `json:"id" bson:"_id"`
	Img string `json:"img" bson:"img"`
	Title string `json:"title" bson:"title"`
	Issue int64 `json:"issue" bson:"issue"`
	Editor string `json:"string" bson:"editor"`
	Read int `json:"read" bson:"read"`
	Content string `json:"content" bson:"content"`
	KeyWord []string `json:"key" bson:"key"`
	IsTop bool `json:"isTop" bson:"is_top"`
	Comment int `json:"comment" bson:"comment"`
	Status int `json:"status" bson:"status"`
}

func CreateArticle(title, editor, img string, content string, keyWords []string, isTop bool) *Article {
	return &Article {
		ID:bson.NewObjectId(),
		Img:img,
		Title:title,
		Issue:time.Now().Unix(),
		Editor:editor,
		Read:0,
		Content:content,
		KeyWord:keyWords,
		IsTop:isTop,
		Status:1,
	}
}
