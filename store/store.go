package store

import (
	"labix.org/v2/mgo"
	"log"
	"os"
)

func init() {
	if os.Getenv("mongourl") == "" {
		os.Setenv("mongourl", "localhost/fp")
		log.Println("needed to set url!!!")
	}
}

//connect to database
func connect() *mgo.Session {
	session, err := mgo.Dial(os.Getenv("mongourl"))
	if err != nil {
		panic(err)
	}
	return session
}
