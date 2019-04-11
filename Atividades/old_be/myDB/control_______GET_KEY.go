package myDB

import (
	"gopkg.in/mgo.v2/bson"
)

func GetKeyQuery(keyField string, key string) bson.M {
	return bson.M{keyField: key}
}
