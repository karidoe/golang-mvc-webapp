package db

import (
	"fmt"
	. "golang-mvc-webapp/config"
	"gopkg.in/mgo.v2"
	"log"
	"src/gopkg.in/mgo.v2"
)

type Mongodb struct {
	session *mgo.Session
}

var (
	session *mgo.Session
	err error
)

func init() {
	fmt.Println("Connecting DB....")

	dsn := fmt.Sprintf(
		"mongodb://%s:%s@mongo/%s", 
		Getenv("APP_MONGO_USERNAME"),
		Getenv("APP_MONGO_PASSWORD"),
		Getenv("APP_MONGO_DATABASE"),
	)
	
	session, err = mgo.Dial(dsn)

	if (err != nil) {
		log.Fatalf("Cannot Connect Mongodb %v", err)
	}
}

func GetMongodb() *Mongodb {
	return &Mongodb{session.Copy()}
}

func (c *Mongodb) GetSession() *mgo.Session {
	return c.session
}

func (c *Mongodb) CloseSession() {
	c.session.Close()
}