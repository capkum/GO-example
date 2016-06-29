package main

import (
  "fmt"
  "gopkg.in/mgo.v2"
  _ "reflect"
)

// mongo db setting
type MongoConf struct {
  dbpath string
}

// mongodb model
type Mst_area struct {
    areaId int
    chara string
    icon string
    name string
    open int
    rank int
    sx int
    sy int
    tex string
    timestamp int64
    worldId string
    xpos string
    ypos string
}

// mongodb connection initialize
func (mgo *MongoConf) InitialMgo()  {
  mgo.dbpath = "mongodb://localhost"
}

// mongo connection
func GetSession()  *mgo.Session {
  mgodb := new(MongoConf)
  mgodb.InitialMgo()

  session, err := mgo.Dial(mgodb.dbpath)
  if err != nil {
    fmt.Println(err.Error())
  }
  return session
}

func main() {
  fmt.Println("==== [mongodb connect] ====")
  conn := GetSession()
  c := conn.DB("CHRODB2").C("mst_place")
  fmt.Println(c)
}
