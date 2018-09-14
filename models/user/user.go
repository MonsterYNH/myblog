package user

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID            bson.ObjectId `json:"id" bson:"_id, omitempty"`
	Account       string        `json:"account" bson:"account, omitempty"`
	Password      string        `json:"password" bson:"password, omitempty"`
	Name          string        `json:"name" bson:"name, omitempty"`
	RealName      string        `json:"real" bson:"real_name, omitempty"`
	Phone         string        `json:"phone" bson:"phone, omitempty"`
	Email         string        `json:"email" bson:"email, omitempty"`
	RegistTime    int64         `json:"regist" bson:"regist, omitempty"`
	LastLoginTime int64         `json:"last" bson:"last, omitempty"`
}

func (self *User) GetOneObject(dbName, collName string, query bson.M, objectName string) (interface{}, error) {
	return User{}, nil
}

func (self *User) InsertOneObject(dbName, collName, objectName string, object interface{}) (int, error) {
	return 0, nil
}

func (self *User) UpdateOneObject(dbName, collName, query bson.M, objectName string, object interface{}) error {
	return nil
}

func (self *User) DeleteOneObject(dbName, collName, query bson.M, objectName string, object interface{}) error {
	return nil
}
