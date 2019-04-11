package yourUserEntity

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"lfkk/be/myHttp"
	"log"
	"net/http"
)

func (c *UserEntityController) UpdateUserEntity(w http.ResponseWriter, r *http.Request) {
	var ue UserEntity
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))

	if err != nil {
		log.Fatalln("Error UpdateUser", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error UpdateUser", err)
	}

	if err := json.Unmarshal(body, &ue); err != nil {
		myHttp.MyWriteHeaderSet(w, r, myHttp.APPJASON_UTF_8)
		w.WriteHeader(422)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error UpdateUser unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	success := c.UserEntityRepository.UpdateUserEntityDB(ue)
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	myHttp.MyWriteHeaderSet(w, r, myHttp.APPJASON_UTF_8)
	w.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(ue)
	w.Write(data)
	return
}
