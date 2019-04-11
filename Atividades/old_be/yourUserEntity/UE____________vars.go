package yourUserEntity

import (
	"lfkk/be/myDB"
)

var session = myDB.GetSession(DBNAME, DOCNAME, KEY)
