package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Notice struct {
	ID bson.ObjectId `json:"id" bson:"_id"`
	Content string `json:"content" bson:"content"`
	Issue int64 `json:"issue" bson:"issue"`
	IsImportant bool `json:"important" bson:"important"`
}

func CreateNotice(content string, isImportant bool) *Notice {
	return & Notice{
		ID:bson.NewObjectId(),
		Content:content,
		Issue:time.Now().Unix(),
		IsImportant:isImportant,
	}
}
