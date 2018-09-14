package mongodb

import (
	"gopkg.in/mgo.v2/bson"
	"mixotc.com/dailyread/pkg/libs/define"
	"myblog/models/user"
	"reflect"
	"test/coinreadss/vender/mixotc.com/logrus-logger"
)

var mgoModelMap map[string]MgoObjectEntry

func init() {
	mgoModelMap = make(map[string]MgoObjectEntry)
	RegistMgoModel(define.MONGO_COLLECTION_USER, &user.User{})
}

type MongoDBAble interface {
	GetOneObject(dbName, collName string, query bson.M, objectName string) (interface{}, error)
	InsertOneObject(dbName, collName, objectName string, object interface{}) (int, error)
	UpdateOneObject(dbName, collName, query bson.M, objectName string, object interface{}) error
	DeleteOneObject(dbName, collName, query bson.M, objectName string, object interface{}) error
}

// 结构体和mongo数据库文档的映射
type MgoObjectEntry struct {
	TypeName string
	Object   MongoDBAble
}

func RegistMgoModel(typeName string, object MongoDBAble) {
	if _, isExist := mgoModelMap[typeName]; isExist {
		log.Panicf("[init] model is exist, typename is: %s", typeName)
	}
	if object == nil {
		log.Panicf("[init] regist model failed, Error: object is nill")
	}
	mgoModelMap[typeName] = MgoObjectEntry{
		TypeName: typeName,
		Object:   object,
	}
}

func GetObjectEntryByTypeName(typeName string) MgoObjectEntry { // 不能使用*会改变已经初始化好的值
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
	objectsType := reflect.PtrTo(reflect.TypeOf(CreateObjectByTypeName(typeName)).Elem())
	objectsValue := reflect.MakeSlice(reflect.SliceOf(objectsType), 0, 0)
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
