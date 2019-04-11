package yourUser

import ("lfkk/be/myJson"
	"lfkk/be/myHttp"
	"lfkk/be/myValidator"
	"net/http")

func (c *UserController) InsertUser(w http.ResponseWriter, r *http.Request) {
	var ue UserEntity
	myHttp.MyWriteHeaderSet(w, r, myHttp.APPJASON_UTF_8)

	if 	err := myJson.DecodeJson(r.Body, &ue); err != nil {
		myHttp.MyWriteResponses(w, r, 478)
		return
	}

	myValidator.InsertUserValidator()

	if 	err := myValidator.GetVarValidate().Struct(ue); err != nil {
		myHttp.MyWriteResponses(w, r, ErrorInsertUserEntity(w, r, ue, err))
		return
	}

	if err := c.UserRepository.InsertUserDB(ue); err != nil {
		myHttp.MyWriteResponses(w, r, ErrorInsertUserEntity(w, r, ue, err))
		return
	}

	myHttp.MyWriteResponses(w, r, 220)
	return
}
