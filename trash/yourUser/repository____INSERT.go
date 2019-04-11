package yourUser

import ("log")

func (r UserRepository) InsertUserDB(user User) (err error){
	err = session.DB(DBNAME).C(DOCNAME).Insert(user)

	if err != nil {
		log.Println(err)
		return
	}

	return
}
