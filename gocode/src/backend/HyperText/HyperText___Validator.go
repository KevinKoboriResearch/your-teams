package HyperText

import (
	"gopkg.in/go-playground/validator.v9"
	"log"
	"net/http"
	"bytes"
	//"encoding/json"
)

var Validate *validator.Validate

func StartValidator() {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nHyperText.StartValidator()\n")

	Validate = validator.New()
}

func BodyValidateJson(j *bytes.Reader, entity interface{}) interface{} {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nnHyperText.BodyValidate()\n")

	if err := DecodeJson(j, entity); err != nil {
		log.Println("[ERROR] Can't decode json: ", err)
		return CustomResponses["wrong-json"]
	}
	if err := StructValidate(entity); err != nil {
		log.Println("[ERROR] Can't validate struct: ", err)
		return err
	}
	return nil
}

func BodyValidate(r *http.Request, entity interface{}) interface{} {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nnHyperText.BodyValidate()\n")

	if err := DecodeJson(r.Body, entity); err != nil {
		log.Println("[ERROR] Can't decode json: ", err)
		return CustomResponses["wrong-json"]
	}
	if err := StructValidate(entity); err != nil {
		log.Println("[ERROR] Can't validate struct: ", err)
		return err
	}
	return nil
}

func StructValidate(entity interface{}) interface{} {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\n...\n")

	err := Validate.Struct(entity)
	if err != nil {
		log.Println("[ERROR] Can't validate struct: ", err)
		var fields []string
		for _, value := range err.(validator.ValidationErrors) {
			fields = append(fields, value.Tag()+": "+value.Field())
		}
		return fields[len(fields)-1]
	}
	return nil
}
