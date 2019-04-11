package myDB

import (
	"gopkg.in/mgo.v2"
	"log"
)

func GetSession(DBNAME string, DOCNAME string, KEY string) *mgo.Session {
	session, err := mgo.Dial(MONGO_HOST)
	if err != nil {
		log.Println("[ERROR] failed to establish connection to Database session: ", err)
	}

	session.SetMode(mgo.Monotonic, true)
	myEnsureIndex(session, DBNAME, DOCNAME, KEY)

	return session
}
