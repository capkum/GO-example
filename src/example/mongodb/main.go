package main

import (
	"example/mongodb/dbcon"
	"example/mongodb/models"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	// mongodb session
	dbcon, conn := dbcon.DbConn()
	// mongodb close connection
	defer dbcon.Close()
	// create user
	// CreateUser(conn)
	// user list
	GetUsers(conn)
	//  remove user
	// DelUser(conn, "5775d541807d57128e2e2a8e")
	// update user
	// UpdateUser(conn)
}

// DelUser remove user data
func DelUser(s *mgo.Database, id string) {
	oid := bson.ObjectIdHex(id)
	err := s.C("users").RemoveId(oid)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("삭제 완료")
}

// UpdateUser update user data
func UpdateUser(s *mgo.Database) {
	sampleData := bson.M{
		"name": "Seoungjin1",
		"age":  100,
	}

	sid := bson.ObjectIdHex("5775e602807d57188f616f76")
	condition := bson.M{"_id": sid}
	change := bson.M{"$set": sampleData}
	err := s.C("users").Update(condition, change)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("수정 완료")
}

// CreateUser insert new user data
func CreateUser(s *mgo.Database) {
	sampleData := models.User{
		Name:   "Seoungjin Kim",
		Age:    43,
		Gender: "male",
		ID:     bson.NewObjectId(),
	}
	err := s.C("users").Insert(sampleData)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("입력완료")
}

// GetUsers get user list
func GetUsers(s *mgo.Database) {
	var result []models.User
	err := s.C("users").Find(nil).All(&result)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, v := range result {
		fmt.Println(v.Name, v.Age, v.Gender)
	}

}
