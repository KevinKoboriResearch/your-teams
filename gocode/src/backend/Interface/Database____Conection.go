package Interface

import (
	"gopkg.in/mgo.v2"
)

const (
	DBNAME     = "your_teams"
	MONGO_IP   = "localhost"
	MONGO_PORT = ":27017"
	MONGO_HOST = MONGO_IP + MONGO_PORT
)

var database *mgo.Database

func StartConectionDatabase() (err error) {
	session, err := mgo.Dial(MONGO_HOST)
	database = session.DB(DBNAME)
	return
}

func OpenSession(docname string) *mgo.Collection {
	return database.C(docname)
}
