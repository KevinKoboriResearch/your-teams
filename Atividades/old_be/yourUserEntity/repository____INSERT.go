package yourUserEntity

import (
	"gopkg.in/mgo.v2/bson"
	"log"
)

func (r UserEntityRepository) InsertUserEntityDB(userentity UserEntity) (bool, error) {
	ue := UserEntity{}
	email := EmailVerify{}
	email.Email = userentity.Email
	err := session.DB(DBNAME).C(DOCNAME).Find(bson.M{"email": email.Email}).One(&ue)

	if err != nil {
		err = session.DB(DBNAME).C(DOCNAME).Insert(userentity)
		if err != nil {
			log.Println(err)
			return false, err
		}
		return false, err
	}

	return true, err
}
