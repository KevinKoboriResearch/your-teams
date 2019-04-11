package myHttp

import (
	"net/http"
)

func MyWriteResponses(w http.ResponseWriter, r *http.Request, i int) {
	w.Write([]byte(responses[i]))
}
