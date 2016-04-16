package services

import (
	"fmt"
	"github.com/evcraddock/supp-store/config"
	"gopkg.in/mgo.v2"
)

func GetSession() *mgo.Session {

	configuration := config.LoadFile()
	mongoconfig := configuration.MongoDb.Address + ":" + configuration.MongoDb.Port

	s, err := mgo.Dial("mongodb://" + mongoconfig)

	if err != nil {
		fmt.Printf("error: %v", err)
		panic(err)
	}

	return s
}
