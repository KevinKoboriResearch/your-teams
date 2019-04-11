package yourUserEntity

import (
	"lfkk/be/myDB"
	"log"
)

func (r UserEntityRepository) UpdateUserEntityDB(userentity UserEntity) bool {
	err := session.DB(DBNAME).C(DOCNAME).Update(myDB.GetKeyQuery(KEY, userentity.Username), userentity)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
