package myHttp

import("net/http"
		"strconv"
	"lfkk/be/myUserEntity")

func MyWriteUserEntityData(w http.ResponseWriter, r *http.Request, ue myUserEntity.UserEntity) {
	w.Write([]byte("Username/Artistic Name: " + ue.Username +
		"\nFull Name: " + ue.FirstName + " " + ue.LastName +
		"\nEmail: " + ue.Email +
		"\nPassword: " + ue.Password +
		"\nMain Role: " + ue.MainRole +
		"\nStatus: " + strconv.Itoa(ue.Status) +
		"\nAge: " + strconv.Itoa(ue.Age)))
}
