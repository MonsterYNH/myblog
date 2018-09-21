package logic

import (
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"log"
	"myblog/models/db"
	"strconv"
)

func GetDataByPage(page, size int, typeName, redisKey string) (interface{}, int) {
	objectEntry := db.GetObjectEntryByTypeName(typeName)
	pageData := db.CreateObjectsByTypeName(typeName)
	num := objectEntry.LRange(redisKey, size*(page-1), page*size-1)
	pageData = db.AppendDatasToObjectInterface(pageData, objectEntry.Result)
	if num < size {
		query := []bson.M {
			bson.M{"$match":bson.M{"status":1}},
			bson.M{"$sort":bson.M{"issue":-1}},
			bson.M{"$skip":size*(page-1) + num},
			bson.M{"$limit":size - num},
		}
		objectEntry.GetObjectsFromMongo(db.MONGO_DB, query)
		pageData = db.AppendDatasToObjectInterface(pageData, objectEntry.Result)
	}
	all, err := strconv.Atoi(strconv.FormatInt(objectEntry.LLen(redisKey), 10))
	if err != nil {
		log.Panicf("[convert] convert num error, Error:%+v", err)
	}
	return pageData, all
}

func GetDataByIdFormList(id bson.ObjectId, typeName, hkey, lkey string) interface{} {
	objectEntry := db.GetObjectEntryByTypeName(typeName)
	objectEntry.HGet(hkey, id.Hex())
	var index int
	if objectEntry.Result.([]byte) != nil {
		if err := json.Unmarshal(objectEntry.Result.([]byte), &index); err != nil {
			log.Panicf("[redis] get data from redis list failed, key:%s, id:%s, Error:%+v", hkey, id.Hex(), err)
		}
		objectEntry.LIndex(lkey, index)
		return objectEntry.Result
	}
	if id.Valid() {
		objectEntry.GetOneObjectFromMongo(db.MONGO_DB, []bson.M{bson.M{"$match":bson.M{"_id":id, "status":1}}})
		return objectEntry.Result
	}
	return nil
}
