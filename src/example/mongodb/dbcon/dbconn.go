package dbcon

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// MongoConf config structure
type MongoConf struct {
	dbPath string
	dbName string
}

// MongoInitial mongodb structure initialize
func MongoInitial() *MongoConf {
	conf := MongoConf{}
	conf.dbPath = "mongodb://localhost"
	conf.dbName = "go_rest_tutorial"
	return &conf
}

// MgoConn mongodb connection
func MgoConn() *mgo.Session {
	session, err := mgo.Dial(MongoInitial().dbPath)
	if err != nil {
		fmt.Println(err.Error())
	}
	return session
}

// DbConn database connect
func DbConn() (*mgo.Session, *mgo.Database) {
	dbcon := MgoConn()
	con := dbcon.DB(MongoInitial().dbName)

	return dbcon, con
}
