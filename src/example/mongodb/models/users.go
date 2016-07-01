package models

import "gopkg.in/mgo.v2/bson"

// User define users mongodb collection
type User struct {
	ID     bson.ObjectId `json:"id" bson:"_id"`
	Age    int           `json:"age" bson:"age"`
	Name   string        `json:"name" bson:"name"`
	Gender string        `json:"gender" bson:"gender"`
}
