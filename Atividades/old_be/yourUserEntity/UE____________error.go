package yourUserEntity

import "net/http"

func ErrorInsertUserEntity(w http.ResponseWriter, r *http.Request, username string, err error) int {
	usernameValidation := `Key: 'UserEntity.Username' Error:Field validation for 'Username' failed on the 'username-req' tag`
	emailValidation := `Key: 'UserEntity.Email' Error:Field validation for 'Email' failed on the 'email' tag`
	passwordValidation := `Key: 'UserEntity.Password' Error:Field validation for 'Password' failed on the 'password-req' tag`
	dupKeyDB := `E11000 duplicate key error collection: ` + DBNAME + `.` + DOCNAME + ` index: username_1 dup key: { : "` + username + `" }`
	httpValue := 475

	if err.Error() == emailValidation {
		httpValue = 477
	} else if err.Error() == passwordValidation {
		httpValue = 474
	} else if err.Error() == usernameValidation {
		httpValue = 473
	} else if err.Error() == usernameValidation+"\n"+emailValidation {
		httpValue = 471
	} else if err.Error() == usernameValidation+"\n"+passwordValidation {
		httpValue = 472
	} else if err.Error() == emailValidation+"\n"+passwordValidation {
		httpValue = 476
	} else if err.Error() == dupKeyDB {
		httpValue = 407
	}
	return httpValue
}
