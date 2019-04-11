package myDB

import (
	"gopkg.in/mgo.v2"
)

func myEnsureIndex(s *mgo.Session, DBNAME string, DOCNAME string, KEY string) {
	session := s.Copy()
	defer session.Close()

	c := session.DB(DBNAME).C(DOCNAME)

	index := mgo.Index{
		Key:        []string{KEY},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err := c.EnsureIndex(index)

	if err != nil {
		panic(err)
	}
}
