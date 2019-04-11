package myValidator

import "gopkg.in/go-playground/validator.v9"

func GetVarValidate() *validator.Validate {
	val := validate
	return val
}
