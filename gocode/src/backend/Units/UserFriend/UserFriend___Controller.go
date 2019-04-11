package UserFriend

import (
	"backend/HyperText"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

//____________________________ INSERT ________________________________________//
func (c *UserFriendController) AddFriend(w http.ResponseWriter, r *http.Request) {

	err := HyperText.BodyValidate(r, &uf)

	if err != nil {
		HyperText.HttpTrueResponse(w, http.StatusOK, []byte(HyperText.CustomResponses["wrong-body"]))
		return
	}

	uf.User2 = strings.ToLower(mux.Vars(r)["username"])

	num := c.UserFriendRepository.InsertUserFriend(uf)

	if num == 0 {
		HyperText.HttpTrueResponse(w, http.StatusOK, []byte(HyperText.CustomResponses["error-database"]))
		return
	}
	if num == 1 {
		HyperText.HttpTrueResponse(w, http.StatusOK, []byte(HyperText.CustomResponses["friend-request"]))
		return
	}

	HyperText.HttpTrueResponse(w, http.StatusOK, []byte(HyperText.CustomResponses["friend-accpect"]))
	return
}

//____________________________ DELETE ________________________________________//
func (c *UserFriendController) Delete(w http.ResponseWriter, r *http.Request) {

	err := HyperText.BodyValidate(r, &uf)

	if err != nil {
		HyperText.HttpTrueResponse(w, http.StatusOK, []byte(HyperText.CustomResponses["wrong-body"]))
		return
	}

	uf.User2 = strings.ToLower(mux.Vars(r)["username"])

	err = c.UserFriendRepository.DeleteUserFriend(uf)

	if err != nil {
		HyperText.HttpTrueResponse(w, http.StatusOK, []byte(HyperText.CustomResponses["not-found-entity"]))
		return
	}

	HyperText.HttpTrueResponse(w, http.StatusOK, []byte(HyperText.CustomResponses["success-delete"]))
	return
}
