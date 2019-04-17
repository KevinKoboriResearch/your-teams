package Interface

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

//__ INSERT __________________________________________________________________//
func InsertDB(collection string, entity interface{}) (err error) {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\nInsertDB\n")

	err = OpenSession(collection).Insert(entity)
	return
}

//__ UPDATE SINGLE ___________________________________________________________//
func UpdateSingleDB(collection string, id string, idValue string, position string, value interface{}) (err error) {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\nUpdateSingleDB\n")

	if id == "_id" {
		err = OpenSession(collection).Update(bson.M{id: bson.ObjectIdHex(idValue)}, bson.M{"$set": bson.M{position: value}})
		return
	}
	err = OpenSession(collection).Update(bson.M{id: idValue}, bson.M{"$set": bson.M{position: value}})
	return
}

//__ UPDATE __________________________________________________________________//
func UpdateDB(collection string, id string, idValue string, entity interface{}) (err error) {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\nUpdateDB\n")

	change := mgo.Change{
		Update:    bson.M{"$set": entity},
		ReturnNew: true,
	}
	if id == "_id" {
		_, err = OpenSession(collection).Find(bson.M{id: bson.ObjectIdHex(idValue)}).Apply(change, &entity)
		return
	}
	_, err = OpenSession(collection).Find(bson.M{id: idValue}).Apply(change, &entity)
	return
}

//__ UPDATE __________________________________________________________________//
func UpdateUnitDB(collection string, entity interface{}, idValues ...string) (err error) {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\nUpdateUnitDB\n")

	change := mgo.Change{
		Update:    bson.M{"$set": entity},
		ReturnNew: true,
	}
	if idValues[0] == "_id" {
		_, err = OpenSession(collection).Find(bson.M{idValues[0]: bson.ObjectIdHex(idValues[1]), idValues[2]: bson.ObjectIdHex(idValues[3])}).Apply(change, &entity)
		return
	}
	_, err = OpenSession(collection).Find(bson.M{idValues[0]: idValues[1], idValues[2]: idValues[3]}).Apply(change, &entity)
	return
}

//__ FIND ____________________________________________________________________//
func FindDB(collection string, id string, idValue string, entity interface{}) (err error) {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\nFindDB\n")

	if id == "_id" {
		err = OpenSession(collection).Find(bson.M{id: bson.ObjectIdHex(idValue)}).One(entity)
		return
	}
	err = OpenSession(collection).Find(bson.M{id: idValue}).One(entity)
	return
}

//__ FIND ____________________________________________________________________//
func FindUnitDB(collection string, entity interface{}, idValues ...string) (err error) {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\nFindUnitDB\n")

	if idValues[0] == "_id" {
		err = OpenSession(collection).Find(bson.M{idValues[0]: bson.ObjectIdHex(idValues[1])}).One(entity)
		return
	}
	err = OpenSession(collection).Find(bson.M{idValues[0]: idValues[1], idValues[2]: idValues[3]}).One(entity)
	return
}

//__ FIND ALL WHILE __________________________________________________________//
func FindAllWhileDB(collection string, position string, value interface{}, entity interface{}) (err error) {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\nFindAllWhileDB\n")

	err = OpenSession(collection).Find(bson.M{position: value}).All(entity)
	return
}

//__ FIND ALL ENABLED WHILE __________________________________________________//
func FindAllEnabledWhileDB(collection string, position string, value interface{}, entity interface{}) (err error) {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\nFindAllEnabledWhileDB\n")

	err = OpenSession(collection).Find(bson.M{position: value, "enable": true}).All(entity)
	return
}

//__ FIND ALL ENABLED ________________________________________________________//
func FindAllEnabledDB(collection string, entity interface{}) (err error) {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\nFindAllEnabledDB\n")

	err = OpenSession(collection).Find(bson.M{"enable": true}).All(entity)
	return
}

//__ FIND ALL ________________________________________________________________//
func FindAllDB(collection string, entity interface{}) (err error) {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\n...\n")

	err = OpenSession(collection).Find(nil).All(entity)
	return
}

//__ DELETE __________________________________________________________________//
func DeleteDB(collection string, id string, idValue string, entity interface{}) (err error) {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\n...\n")

	if id == "_id" {
		err = OpenSession(collection).Remove(bson.M{id: bson.ObjectIdHex(idValue)})
		return
	}
	err = OpenSession(collection).Remove(bson.M{id: idValue})
	return
}

//__ DELETE __________________________________________________________________//
func DeleteUnitDB(collection string, entity interface{}, idValues ...string) (err error) {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\n...\n")

	if idValues[0] == "_id" {
		err = OpenSession(collection).Remove(bson.M{idValues[0]: bson.ObjectIdHex(idValues[1]), idValues[2]: bson.ObjectIdHex(idValues[3])})
		return
	}
	err = OpenSession(collection).Remove(bson.M{idValues[0]: idValues[1], idValues[2]: idValues[3]})
	return
}

//__ DELETE ALL COLLECTION __________________________________________________________________//
func DeleteAllCollection(collection string) (err error) {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\n...\n")

	err = OpenSession(collection).Remove(nil)
	return
}
