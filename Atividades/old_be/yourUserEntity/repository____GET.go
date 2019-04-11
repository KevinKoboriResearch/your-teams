package yourUserEntity

import (
	"fmt"
	"lfkk/be/myDB"
)

func (r UserEntityRepository) GetUserEntityDB(username string) UserEntities {
	c := session.DB(DBNAME).C(DOCNAME)
	results := UserEntities{}
	err := c.Find(myDB.GetKeyQuery(KEY, username)).All(&results)

	if err != nil {
		fmt.Println("Failed to write results:", err)
	}

	return results
}
