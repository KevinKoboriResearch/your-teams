package HyperText

import (
	"encoding/json"
	"io"
	"log"
)

func DecodeJson(body io.Reader, entity interface{}) (err error) {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\n...\n")

	d := json.NewDecoder(body)
	err = d.Decode(entity)
	if err != nil {
		log.Print("[ERROR] wrong JSON")
		return
	}
	return
}
