package yourUser

import ("encoding/json"
	  "lfkk/be/myHttp"
		"net/http")

func (c *UserEntityController) GetUserEntities(w http.ResponseWriter, r *http.Request) {
	ues := c.UserEntityRepository.GetUserEntitiesDB()
	data, _ := json.Marshal(ues)

  myHttp.MyWriteHeaderSet(w, r, myHttp.APPJASON_UTF_8)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
