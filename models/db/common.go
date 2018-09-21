package db

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"myblog/models/model"
	"reflect"
)

var mgoModelMap map[string]ObjectEntry

func init() {
	mgoModelMap = make(map[string]ObjectEntry)
	RegistModel(&ObjectEntry{CollectionName: MONGO_COLL_USER, Object: model.User{}})
	RegistModel(&ObjectEntry{CollectionName:MONGO_COLL_ARTICLE, Object: model.Article{}})
	RegistModel(&ObjectEntry{CollectionName:MONGO_COLL_NOTICE,Object:model.Notice{}})
}

type MongoDBAble interface {
	GetOneObjectFromMongo(dbName string, query []bson.M)
	InsertOneObjectFromMongo(dbName string, object interface{})
	UpdateObjectFromMongo(dbName string, selector bson.M, update bson.M)
	DeleteOneObjectByIdFromMongo(dbName string, id bson.ObjectId)
	GetObjectsFromMongo(dbName string, query []bson.M)
	InsertObjectsFromMongo(dbName string, objects []interface{})
	DeleteObjectsFromMongo(dbName string, query bson.M)
}

type RedisAble interface {
	InsertOneObjectFromRedis(key string, data interface{})
	GetOneObjectFromRedis(key string)
	SetExpert(key string, expertTime int)
	DelKey(key string)

	HSet(key, id string, data interface{})
	HGet(key, id string)

	LIndex(key string, index int)
	LPush(key string, data interface{})
	LRange(key string, start, offset int) int
	LSet(key string, index int, data []byte)
	LTrim(key string, start, offset int)
	LLen(key string) int
	FlushDB()
}

type ObjectEntry struct {
	CollectionName string
	Object   interface{}
	Result   interface{}
}

func RegistModel(entry *ObjectEntry) {
	if entry == nil {
		log.Panicf("[init] regist mongo model failed, Error: mongo model is nill")
	}
	if _, isExist := mgoModelMap[entry.CollectionName]; isExist {
		log.Panicf("[init] mongo model is exist, typename is: %s", entry.CollectionName)
	}
	mgoModelMap[entry.CollectionName] = *entry
}

func (self *ObjectEntry) GetOneObjectFromMongo(dbName string, query []bson.M) {
	mgoSession := MongoConn.GetMgoSession()
	defer MongoConn.CloseMgoSeeion(mgoSession)

	query = append(query, bson.M{"$limit": 1})
	object := CreateObjectByTypeName(self.CollectionName)
	err := mgoSession.DB(dbName).C(self.CollectionName).Pipe(query).One(object)
	if err != nil {
		if err != mgo.ErrNotFound {
			log.Panic("[mongo] get one object failed, Error:", err)
		}
		self.Result = nil
		return
	}
	self.Result = object
}

func (self *ObjectEntry) InsertOneObjectFromMongo(dbName string, object interface{}) {
	mgoSession := MongoConn.GetMgoSession()
	defer MongoConn.CloseMgoSeeion(mgoSession)

	if err := mgoSession.DB(dbName).C(self.CollectionName).Insert(object); err != nil {
		log.Panicf("[mongo] insert one object failed, Error:%+v", err)
	} else {
		if user, ok := object.(*model.User); ok {
			self.Result = user.ID
		} else if user, ok := object.(model.User); ok {
			self.Result = user.ID
		}
	}
}

func (self *ObjectEntry) UpdateObjectFromMongo(dbName string, selector bson.M, update bson.M) {
	mgoSession := MongoConn.GetMgoSession()
	defer MongoConn.CloseMgoSeeion(mgoSession)

	if err := mgoSession.DB(dbName).C(self.CollectionName).Update(selector, update); err != nil {
		log.Panicf("[mongo] update one object failed, Error:%+v", err)
	}
}

func (self *ObjectEntry) DeleteOneObjectByIdFromMongo(dbName string, id bson.ObjectId) {
	mgoSession := MongoConn.GetMgoSession()
	defer MongoConn.CloseMgoSeeion(mgoSession)

	if err := mgoSession.DB(dbName).C(self.CollectionName).RemoveId(id); err != nil {
		log.Panicf("[mongo] remove one object failed, Error:%+v", err)
	}
}

func (self *ObjectEntry) GetObjectsFromMongo(dbName string, query []bson.M) {
	mgoSession := MongoConn.GetMgoSession()
	defer MongoConn.CloseMgoSeeion(mgoSession)

	objects := CreateObjectsByTypeName(self.CollectionName)
	if err := mgoSession.DB(dbName).C(self.CollectionName).Pipe(query).All(objects); err != nil {
		log.Panicf("[mongo] get objects failed, Error:%+v", err)
	}
	self.Result = objects
}

func (self *ObjectEntry) InsertObjectsFromMongo(dbName string, objects []interface{}) {
	mgoSession := MongoConn.GetMgoSession()
	defer MongoConn.CloseMgoSeeion(mgoSession)

	mgoSession.DB(dbName).C(self.CollectionName).Bulk().Insert(objects...)
}

func (self *ObjectEntry) DeleteObjectsFromMongo(dbName string, query bson.M) {
	mgoSession := MongoConn.GetMgoSession()
	defer MongoConn.CloseMgoSeeion(mgoSession)

	if info, err := mgoSession.DB(dbName).C(self.CollectionName).RemoveAll(query); err != nil {
		log.Panicf("[mongo] remove all objects by query failed, Error:%+v", err)
	} else {
		self.Result = info
	}
}

func (self *ObjectEntry) InsertOneObjectFromRedis(key string, data interface{}) {
	redisConn := RedisConn.RedisPool.Get()
	defer redisConn.Close()

	if reflect.TypeOf(data).Kind() == reflect.Array {
		log.Panic("[redis] unexcept array parameter")
	}
	bytes, _ := json.Marshal(data)
	if _, err := redisConn.Do("set", key, bytes); err != nil {
		log.Panicf("[redis] set object failed, Error:%+v", err)
	}
}

func (self *ObjectEntry) GetOneObjectFromRedis(key string) {
	redisConn := RedisConn.RedisPool.Get()
	defer redisConn.Close()

	data, err := redis.Bytes(redisConn.Do("get", key))
	if err != nil {
		log.Panicf("[redis] get object failed, Error:%+v", err)
	}
	object := CreateObjectByTypeName(self.CollectionName)
	if err := json.Unmarshal(data, object); err != nil {
		log.Panicf("[redis] unmarshal object failed, Error:%+v", err)
	}
	self.Result = object
}

func (self *ObjectEntry) SetExpert(key string, expertTime int) {
	redisConn := RedisConn.RedisPool.Get()
	defer redisConn.Close()

	if info, err := redisConn.Do("expire", key, expertTime); err != nil {
		log.Panicf("[redis] set expire time failed, key: %s, Error:%+v", key, err)
	} else {
		self.Result = info
	}
}

func (self *ObjectEntry) DelKey(key string) {
	redisConn := RedisConn.RedisPool.Get()
	defer redisConn.Close()

	if info, err := redisConn.Do("del", key); err != nil {
		log.Panicf("[redis] del redis key failed, key: %s, Error:%+v", key, err)
	} else {
		self.Result = info
	}
}

func (self *ObjectEntry) HSet(key, id string, data interface{}) {
	redisConn := RedisConn.RedisPool.Get()
	defer redisConn.Close()

	if _, err := redisConn.Do("hset", key, id, data); err != nil {
		log.Panicf("[redis] hset key failed, key:%s, id:%s, Error:%+v", key, id, err)
	}
}

func (self *ObjectEntry) HGet(key, id string) {
	redisConn := RedisConn.RedisPool.Get()
	defer redisConn.Close()

	data, err := redis.Bytes(redisConn.Do("hget", key, id))
	if err != nil {
		if err != redis.ErrNil {
			log.Panicf("[redis] hget key failed, key:%s, id:%s, Error:%+v", key, id, err)
		}
	}
	self.Result = data
}

func (self *ObjectEntry) LIndex(key string, index int) {
	redisConn := RedisConn.RedisPool.Get()
	defer redisConn.Close()

	data, err := redis.Bytes(redisConn.Do("lindex", key, index))
	if err != nil {
		log.Panicf("[redis] lindex key failed, key:%s, index:%d, Error:%+v", key, index, err)
	}
	objectEntry := CreateObjectByTypeName(self.CollectionName)
	if err := json.Unmarshal(data, objectEntry); err != nil {
		log.Panicf("[redis] unmarshal data failed, key:%s, typeName:%s, index:%d, Error:%+v", key, self.CollectionName, index, err)
	}
	self.Result = objectEntry
}

func (self *ObjectEntry) LPush(key string, data interface{}) {
	redisConn := RedisConn.RedisPool.Get()
	defer redisConn.Close()

	bytes, _ := json.Marshal(data)
	if info, err := redisConn.Do("lpush", key, bytes); err != nil {
		log.Panicf("[redis] lpush redis key failed, key: %s, Error:%+v", key, err)
	} else {
		self.Result = info
	}
}

func (self *ObjectEntry) LRange(key string, start, offset int) int {
	redisConn := RedisConn.RedisPool.Get()
	defer redisConn.Close()

	datas, err := redis.ByteSlices(redisConn.Do("lrange", key, start, offset))
	if err != nil {
		log.Panicf("[redis] lrange key failed, key: %s, Error:%+v", key, err)
	}
	objects := CreateObjectsByTypeName(self.CollectionName)
	for _, data := range datas {
		object := CreateObjectByTypeName(self.CollectionName)
		if err := json.Unmarshal(data, object); err != nil {
			log.Panicf("[redis] unmarshal object failed")
		}
		objects = AppendDataToObjectsInterface(objects, object)
	}
	self.Result = objects
	return len(datas)
}

func (self *ObjectEntry) LSet(key string, index int, data []byte) {
	redisConn := RedisConn.RedisPool.Get()
	defer redisConn.Close()

	if _, err := redisConn.Do("lset", key, index, data); err != nil {
		log.Panicf("[redis] lset key failed, key: %s, index: %d, Error:%+v", key, index, err)
	}
}

func (self *ObjectEntry) LTrim(key string, start, offset int) {
	redisConn := RedisConn.RedisPool.Get()
	defer redisConn.Close()

	if _, err := redisConn.Do("ltrim", key, start, offset); err != nil {
		log.Panicf("[redis] ltrim key failed, key:%s, start:%d, offset:%d, Error:%+v", key, start, offset, err)
	}
}

func (self *ObjectEntry) FlushDB() {
	redisConn := RedisConn.RedisPool.Get()
	defer redisConn.Close()

	if _, err := redisConn.Do("flushdb"); err != nil {
		log.Panicf("[redis] flush redis failed, Error:%+v", err)
	}
}

func (self *ObjectEntry) LLen(key string) int64 {
	redisConn := RedisConn.RedisPool.Get()
	defer redisConn.Close()

	num, err := redisConn.Do("llen", key)
	if err != nil {
		log.Panicf("[redis] get list len failed, Error:%+v", err)
	}
	return num.(int64)
}

// 根据类型名获取对象
func GetObjectEntryByTypeName(typeName string) ObjectEntry { // 不能使用*会改变已经初始化好的值
	objectEntry, isExist := mgoModelMap[typeName]
	if !isExist {
		log.Panic("model doesn't regist in mgoModelMap, type name is:", typeName)
	}
	return objectEntry
}

// 根据类型名创建对象
func CreateObjectByTypeName(typeName string) interface{} {
	return reflect.New(reflect.TypeOf(GetObjectEntryByTypeName(typeName).Object)).Interface()
}

// 根据类型名创建对象数组
func CreateObjectsByTypeName(typeName string) interface{} {
	objectsValue := reflect.MakeSlice(reflect.SliceOf(reflect.TypeOf(CreateObjectByTypeName(typeName))), 0, 0)
	return reflect.New(objectsValue.Type()).Interface()
}

// 在slice中增加数据
func AppendDataToObjectsInterface(slice interface{}, data interface{}) interface{} {
	sliceValue := reflect.ValueOf(slice)
	dataValue := reflect.ValueOf(data)

	if sliceValue.Kind() == reflect.Ptr {
		sliceValue = sliceValue.Elem()
	}
	slicePtrVal := reflect.New(reflect.SliceOf(dataValue.Type()))
	slicePtr := reflect.Indirect(slicePtrVal)
	slicePtr.Set(reflect.Append(sliceValue, dataValue))

	return slicePtr.Addr().Interface()
}

func AppendDatasToObjectInterface(slice interface{}, data interface{}	) interface{} {
	sliceValue := reflect.ValueOf(slice)
	dataValue := reflect.ValueOf(data)

	if sliceValue.Kind() == reflect.Ptr {
		sliceValue = sliceValue.Elem()
	}

	for i := 0; i < reflect.Indirect(dataValue).Len(); i++ {
		sliceValue = reflect.Append(sliceValue, reflect.Indirect(dataValue).Index(i))
	}
	return sliceValue.Interface()
}
