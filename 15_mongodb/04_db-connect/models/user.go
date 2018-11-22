package models

import (
	"gopkg.in/mgo.v2/bson"
)

// User ...
type User struct {
	Name   string `json:"name" bson:"name"`
	Gender string `json:"gender" bson:"gender"`

	// tag json 用處在於告訴 json decode 至 struct 時，json 對應 struct 之名稱以及資料型別
	// bson 則是告訴 struct 與 mongodb 資料中對應關係
	Age int           `json:"age,string" bson:"age"`
	ID  bson.ObjectId `json:"id" bson:"_id"`
}
