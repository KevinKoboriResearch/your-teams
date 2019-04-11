package myValidator

import (
	"gopkg.in/go-playground/validator.v9"
)

func UserEntityValidator() {
	validate = validator.New()
	validate.RegisterValidation("username-req", ValidateUserEntityUsername)
	validate.RegisterValidation("password-req", ValidateUserEntityPassword)
}

func ValidateUserEntityUsername(username validator.FieldLevel) bool {
	length := len(username.Field().String())
	return length >= 5
}

func ValidateUserEntityPassword(password validator.FieldLevel) bool {
	length := len(password.Field().String())
	return length >= 6
}
