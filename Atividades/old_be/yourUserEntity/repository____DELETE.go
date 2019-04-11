package yourUserEntity

import (
	"lfkk/be/myDB"
	"log"
)

func (r UserEntityRepository) DeleteUserEntityDB(username string) string {

	if &username == nil {
		return "404"
	}

	err := session.DB(DBNAME).C(DOCNAME).Remove(myDB.GetKeyQuery(KEY, username))

	if err != nil {
		log.Println(err)
		return "500"
	}

	return "200"
}
