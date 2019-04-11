package yourUserEntity

import (
	"github.com/gorilla/mux"
	"lfkk/be/myHttp"
	"net/http"
	//		"lfkk/be/yourUser"
	"strings"
)

func (c *UserEntityController) DeleteUserEntity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	myHttp.MyWriteHeaderSet(w, r, myHttp.APPJASON_UTF_8)

	resp := c.UserEntityRepository.DeleteUserEntityDB(username)

	if strings.Contains(resp, "404") {
		w.WriteHeader(http.StatusNotFound)

	} else if strings.Contains(resp, "500") {
		w.WriteHeader(http.StatusInternalServerError)
		myHttp.MyWriteResponses(w, r, 507)

	} else if strings.Contains(resp, "200") {
		//				uC = *yourUser.UserController
		//			  uC.UserRepository.DeleteUserDB(username)
		w.WriteHeader(http.StatusOK)
		myHttp.MyWriteResponses(w, r, 240)
		return
	}
}
