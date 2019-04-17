package Interface

import (
	"gopkg.in/mgo.v2"
	"log"
)

const (
	DBNAME     = "your_teams"
	MONGO_IP   = "localhost"
	MONGO_PORT = ":27017"
	MONGO_HOST = MONGO_IP + MONGO_PORT
)

var database *mgo.Database

func StartConectionDatabase() (err error) {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nInterface.StartConectionDatabase\n")

	session, err := mgo.Dial(MONGO_HOST)
	database = session.DB(DBNAME)
	return
}

func OpenSession(docname string) *mgo.Collection {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nInterface.OpenSession\n")

	return database.C(docname)
}
