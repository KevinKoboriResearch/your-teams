package myHttp

import (
	"net/http"
)

func MyWriteHeaderSet(w http.ResponseWriter, r *http.Request, typeHeader string) {
	w.Header().Set("Content-Type", typeHeader)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	return
}
