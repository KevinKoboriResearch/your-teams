package yourUserEntity

import (
	"lfkk/be/myHttp"
	"lfkk/be/myJson"
	"lfkk/be/myValidator"
	"net/http"
)

func (c *UserEntityController) InsertUserEntity(w http.ResponseWriter, r *http.Request) {
	var ue UserEntity
	myHttp.MyWriteHeaderSet(w, r, myHttp.APPJASON_UTF_8)

	if err := myJson.DecodeJson(r.Body, &ue); err != nil {
		myHttp.MyWriteResponses(w, r, 478)
		return
	}

	myValidator.UserEntityValidator()

	if err := myValidator.GetVarValidate().Struct(ue); err != nil {
		myHttp.MyWriteResponses(w, r, ErrorInsertUserEntity(w, r, ue.Username, err))
		return
	}

	boolean, err := c.UserEntityRepository.InsertUserEntityDB(ue)

	if err != nil {
		myHttp.MyWriteResponses(w, r, ErrorInsertUserEntity(w, r, ue.Username, err))
		return
	} else if boolean == true {
		myHttp.MyWriteResponses(w, r, 470)
		return
	}

	myHttp.MyWriteResponses(w, r, 220)
	return
}
