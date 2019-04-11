package yourUser

import ("fmt")

func (r UserEntityRepository) GetUserEntitiesDB() UserEntities {
		c := session.DB(DBNAME).C(DOCNAME)
		results := UserEntities{}
		err := c.Find(nil).All(&results)

		if err != nil {
			fmt.Println("Failed to write results:", err)
		}

		return results
}
