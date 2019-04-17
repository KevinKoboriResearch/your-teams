package UserEntity

import (
	"backend/HyperText"
	"backend/Interface"
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"log"
)

//__ USERNAME USED ___________________________________________________________//
func ValidateUsernameUsed(username validator.FieldLevel) bool {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\nValidateUsernameUsed\n")

	c := Interface.OpenSession(DOCNAME)
	u := strings.ToLower(username.Field().String())
	result := UserEntity{}
	if err := c.Find(bson.M{"username": u}).One(&result); err != nil {
	}
	return result.Username == ""
}

//__ USERNAME EXIST __________________________________________________________//
func ValidateUsernameExist(username validator.FieldLevel) bool {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\n...\n")

	c := Interface.OpenSession(DOCNAME)
	u := strings.ToLower(username.Field().String())
	result := UserEntity{}
	if err := c.Find(bson.M{"username": u}).One(&result); err != nil {
	}
	return result.Username != ""
}

//__ USERNAME LENGTH__________________________________________________________//
func ValidateUsernameLength(username validator.FieldLevel) bool {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\n...\n")

	length := len(username.Field().String())
	return length >= 5
}

//__ PASSWORD LENGTH _________________________________________________________//
func ValidatePasswordLength(password validator.FieldLevel) bool {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\n...\n")

	length := len(password.Field().String())
	return length >= 8
}

func ValidatePasswordMatch(password validator.FieldLevel) bool {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\n...\n")

	c := Interface.OpenSession(DOCNAME)
	u := password.Field().String()
	result := UserEntity{}
	if err := c.Find(bson.M{"username": u}).One(&result); err != nil {
	}
	return result.Username != ""
}

//__ EMAIL USED ______________________________________________________________//
func ValidateEmailUsed(email validator.FieldLevel) bool {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\n...\n")

	c := Interface.OpenSession(DOCNAME)
	e := strings.ToLower(email.Field().String())
	result := UserEntity{}
	if err := c.Find(bson.M{"email": e}).One(&result); err != nil {
	}
	return result.Email == ""
}

//__ UPDATE SINGLE ___________________________________________________________//
func ValidateUpdateSingle(c *UserEntityController, ueus UserEntityUpdateSingle) (err interface{}) {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\n...\n")

	if ueus.Position == "username" {
		var u UserUsername
		u.Username = ueus.Value
		err = HyperText.StructValidate(&u)
		return
	}
	if ueus.Position == "email" {
		var e UserEmail
		e.Email = ueus.Value
		err = HyperText.StructValidate(&e)
		return
	}
	/*	if ueus.Position == "image" {
		var i UserImage
		i.Image = ueus.Value
		err = HyperText.StructValidate(&i)
		return
	}*/
	if ueus.Position == "password" {
		var p UserPassword
		p.Password = ueus.Value
		err = HyperText.StructValidate(&p)
		if err == nil {
			var uev UserEntityVerify
			uev.Username = ueus.Username
			uev.Password = ueus.Password
			err = c.UserEntityRepository.VerifyUserEntity(c, uev)
			return
		}
		return
	}
	return
}

//__ UPDATE __________________________________________________________________//
func ValidateUpdate(ueu UserEntityUpdate) interface{} {
/*________________________________________TESTING FUNCTION________________________________________*/log.Println("\n\n...\n")

	if ueu.Email != "" {
		var e UserEmail
		e.Email = ueu.Email
		err := HyperText.StructValidate(e)
		if err != nil {
			return err
		}
	}
	if ueu.Image != "" {
		var i UserImage
		i.Image = ueu.Image
		err := HyperText.StructValidate(i)
		if err != nil {
			return err
		}
	}
	return nil
}
