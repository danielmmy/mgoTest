package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Log struct {
	Id       bson.ObjectId       `bson:"_id"`
	Timstamp bson.MongoTimestamp `bson:"timestamp,omitempty"`
	Level    string              `bson:"level,omitempty"`
	Message  string              `bson:"message,omitempty"`
}

func main() {
	session, err1 := mgo.Dial("mongodb://localhost")
	if err1 != nil {
		panic(err1)
	}

	col := session.DB("dock-gateway").C("winston-log")

	count, err2 := col.Count()
	if err2 != nil {
		panic(err2)
	}
	fmt.Println("Document count:", count)

	var results []Log
	err3 := col.Find(nil).All(&results)
	if err3 != nil {
		panic(err3)
	}
	fmt.Println("Documents:", results[0])
}
