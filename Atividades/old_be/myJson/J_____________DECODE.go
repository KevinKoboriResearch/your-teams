package myJson

import (
	"encoding/json"
	"io"
	"log"
)

func DecodeJson(body io.Reader, entity interface{}) (err error) {

	dbody := json.NewDecoder(body)
	err = dbody.Decode(entity)

	if err != nil {
		log.Print("[ERROR] wrong JSON")
		return
	}

	return
}
