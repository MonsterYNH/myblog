package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	ID            bson.ObjectId `json:"id" bson:"_id"`
	Account       string        `json:"account" bson:"account"`
	Password      string        `json:"password" bson:"password"`
	Name          string        `json:"name" bson:"name"`
	RealName      string        `json:"real" bson:"real_name"`
	Phone         string        `json:"phone" bson:"phone"`
	Email         string        `json:"email" bson:"email"`
	RegistTime    int64         `json:"regist" bson:"regist"`
	LastLoginTime int64         `json:"last" bson:"last"`
	Status        int           `json:"status" bson:"status"`
}

func CreateUser(account, passWord, phone string) *User {
	return &User{
		ID:         bson.NewObjectId(),
		Account:    account,
		Password:   passWord,
		Phone:      phone,
		RegistTime: time.Now().Unix(),
		Status:     1,
	}
}
