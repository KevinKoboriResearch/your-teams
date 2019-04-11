package yourUser

import ("encoding/json"
	"lfkk/be/myHttp"
		"github.com/gorilla/mux"
		"net/http")

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	ue := c.UserRepository.GetUserDB(username)

	data, _ := json.Marshal(ue)

	myHttp.MyWriteHeaderSet(w, r, myHttp.APPJASON_UTF_8)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
