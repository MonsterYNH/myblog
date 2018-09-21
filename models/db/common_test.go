package db

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"myblog/models/model"
	"reflect"
	"testing"
)

func TestObjectEntry_GetOneObjectFromMongo(t *testing.T) {
	objectEntry := GetObjectEntryByTypeName(MONGO_COLL_USER)
	query := []bson.M {
		bson.M{"$match":bson.M{"status":1}},
	}
	objectEntry.GetOneObjectFromMongo(MONGO_DB, query)
	fmt.Println(reflect.TypeOf(objectEntry.Result), objectEntry.Result)
}

func TestObjectEntry_InsertOneObjectFromMongo(t *testing.T) {
	objectEntry := GetObjectEntryByTypeName(MONGO_COLL_ARTICLE)
	objectEntry.InsertOneObjectFromMongo(MONGO_DB, model.CreateArticle("test", "test", "test", "test", []string{"test"}, true))
	objectEntry.InsertOneObjectFromMongo(MONGO_DB, model.CreateArticle("test1", "test1", "test1", "test1", []string{"test1"}, false))
	objectEntry.InsertOneObjectFromMongo(MONGO_DB, model.CreateArticle("test2", "test2", "test2", "test2", []string{"test2"}, true))
	objectEntry.InsertOneObjectFromMongo(MONGO_DB, model.CreateArticle("test3", "test3", "test3", "test3", []string{"test3"}, true))
	objectEntry.InsertOneObjectFromMongo(MONGO_DB, model.CreateArticle("test4", "test4", "test4", "test4", []string{"test4"}, false))
	objectEntry.InsertOneObjectFromMongo(MONGO_DB, model.CreateArticle("test5", "test5", "test5", "test5", []string{"test5"}, true))
}

func TestObjectEntry_InsertNoticeFromMongo(t *testing.T) {
	objectEntry := GetObjectEntryByTypeName(MONGO_COLL_NOTICE)
	objectEntry.InsertOneObjectFromMongo(MONGO_DB, model.CreateNotice("test", true))
	objectEntry.InsertOneObjectFromMongo(MONGO_DB, model.CreateNotice("test1", true))
	objectEntry.InsertOneObjectFromMongo(MONGO_DB, model.CreateNotice("test2", true))
}

func TestObjectEntry_GetObjectsFromMongo(t *testing.T) {
	objectEntry := GetObjectEntryByTypeName(MONGO_COLL_USER)
	query := []bson.M {
		bson.M{"$match":bson.M{"status":1}},
	}
	objectEntry.GetObjectsFromMongo(MONGO_DB, query)
	fmt.Println(reflect.TypeOf(objectEntry.Result), objectEntry.Result)
}

func TestObjectEntry_InsertOneObjectFromRedis(t *testing.T) {
	objectEntry := GetObjectEntryByTypeName(MONGO_COLL_USER)
	objectEntry.InsertOneObjectFromRedis(REDIS_USER_KEY, model.CreateUser("test", "test", "test"))
}

func TestObjectEntry_GetOneObjectFromRedis(t *testing.T) {
	objectEntry := GetObjectEntryByTypeName(MONGO_COLL_USER)
	objectEntry.GetOneObjectFromRedis(REDIS_USER_KEY)
	fmt.Println(reflect.TypeOf(objectEntry), objectEntry.Result)
}

func TestObjectEntry_LPush(t *testing.T) {
	objectEntry := GetObjectEntryByTypeName(MONGO_COLL_USER)
	objectEntry.LPush("user_list", model.CreateUser("test", "test", "test"))
}

func TestObjectEntry_LRange(t *testing.T) {
	objectEntry := GetObjectEntryByTypeName(MONGO_COLL_ARTICLE)
	objectEntry.LRange("articles", 0, -1)
	if users, ok := objectEntry.Result.(*[]*model.User); ok {
		for _, user := range *users {
			fmt.Println(*user)
		}
		fmt.Println(reflect.TypeOf(objectEntry.Result), *users)
	}
}

func TestObjectEntry_HSet(t *testing.T) {
	objectEntry := GetObjectEntryByTypeName(MONGO_COLL_ARTICLE)
	objectEntry.HSet("article_id", bson.NewObjectId().Hex(), 1)
}

func TestObjectEntry_HGet(t *testing.T) {
	objectEntry := GetObjectEntryByTypeName(MONGO_COLL_ARTICLE)
	objectEntry.HGet("article_id", "5ba0bf6e6fe3ce1be0bf10112")
	var index int
	if objectEntry.Result.([]byte) == nil {
		t.Fatal("lalalalala")
	}
	if err := json.Unmarshal(objectEntry.Result.([]byte), &index); err != nil {
		t.Fatal(err)
	}
	fmt.Println(index)
}

func TestObjectEntry_LIndex(t *testing.T) {
	objectEntry := GetObjectEntryByTypeName(MONGO_COLL_ARTICLE)
	objectEntry.LIndex("articles", 0)
	fmt.Println(objectEntry.Result)
}
