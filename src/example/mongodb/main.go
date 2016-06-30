package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// MstArea master area data structure
type MstArea struct {
	ID     int    `json:"id" bson:"_id"`
	AreaID int    `json:"areaId" bson:"areaId"`
	Name   string `json:"name" bson:"name"`
}

// GetSession connection mongodb
func GetSession() *mgo.Session {
	mongoPath := "mongodb://localhost"
	s, err := mgo.Dial(mongoPath)

	if err != nil {
		fmt.Println(err.Error())
	}
	return s
}

func main() {
	fmt.Println("##########[start]###########")
	s := GetSession()
	defer s.Close()

	conn := s.DB("CHRODB2_KOR").C("mst_area")
	// oid := bson.ObjectIdHex("5774903a807d572c6ea8632b")

	// results := MstArea{}
	var results []MstArea
	err := conn.Find(nil).All(&results)
	// err := conn.FindId(oid).One(&results)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, v := range results {
		fmt.Println(v)
	}

}
