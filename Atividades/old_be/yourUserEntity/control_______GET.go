package yourUserEntity

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"lfkk/be/myHttp"
	"net/http"
)

func (c *UserEntityController) GetUserEntity(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	username := vars["username"]

	ue := c.UserEntityRepository.GetUserEntityDB(username)

	data, _ := json.Marshal(ue)

	myHttp.MyWriteHeaderSet(w, r, myHttp.APPJASON_UTF_8)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
