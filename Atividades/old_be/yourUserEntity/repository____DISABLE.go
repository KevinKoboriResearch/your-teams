package yourUserEntity

import (
	"gopkg.in/mgo.v2/bson"
	"lfkk/be/myDB"
	"log"
)

func (r UserEntityRepository) DisableUserEntityDB(uev UserEntityVerify) (bool, error) {
	ue := UserEntity{}
	if err := session.DB(DBNAME).C(DOCNAME).Find(bson.M{"username": uev.Username, "password": uev.Password}).One(&ue); err != nil {
		log.Print("[ERROR] failed to login: ", err)
		return false, err
	}

	if ue.Enable == true {
	ue.Enable = false
		if err := session.DB(DBNAME).C(DOCNAME).Update(myDB.GetKeyQuery(KEY, ue.Username), ue); err != nil {
			log.Print("[ERROR] failed to update DISABLE: ", err)
		}
	}

	return true, nil
}
